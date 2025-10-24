package core

import (
	"muse/internal/config"
	"muse/internal/display"
	"muse/internal/fetch"
	"muse/internal/ipc"
	"muse/internal/lyric"
	"muse/internal/shared"
	"os"
	"os/signal"
	"syscall"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	godbus "github.com/godbus/dbus"
)

type App struct {
	config *config.Config

	conn   *godbus.Conn
	player string

	// Current state
	currentTrack *shared.Track
	Lyrics       *[]shared.Lyric
	lastLine     *shared.Lyric

	// Control channels
	stopChan        chan struct{}
	trackChangeChan chan *shared.Track

	// Display modes
	tuiProgram  *tea.Program
	fyneDisplay *display.WindowDisplay

	// Notifications
	notifier *ipc.Notifier
}

func NewApp(cfg *config.Config) *App {
	return &App{
		config:          cfg,
		stopChan:        make(chan struct{}),
		trackChangeChan: make(chan *shared.Track),
	}
}

func (a *App) Start() error {

	// Connect to dbus
	conn, err := ipc.Connect()
	if err != nil {
		return err
	}
	a.conn = conn

	// Set up notifications
	a.notifier = ipc.NewNotifier(a.conn, a.config.General.ProgramName)

	// Find prefered player
	player, err := ipc.FindActivePlayer(conn, a.config.Player.Preferred)
	if err != nil {
		return err
	}
	a.player = player
	shared.Debug("Found player: %s\n", player)

	// Create tui mode
	switch a.config.General.DisplayMode {
	case "tui":
		model := display.NewTUI(a.config)
		a.tuiProgram = tea.NewProgram(model, tea.WithAltScreen())

		// Run in goroutine so Start() can continue
		go func() {
			a.tuiProgram.Run()
			a.Stop()
		}()

	case "window":
		a.fyneDisplay = display.NewWindow(a.config)
		a.fyneDisplay.Start()
	}

	track, err := ipc.GetTrackInfo(a.conn, a.player)
	if err == nil {
		shared.Debug("Initial track: %s - %s\n", track.Artist, track.Title)
		a.handleTrackChange(track)
	}

	// Start watching for track changes
	go a.watchTrackChanges()

	// Start syncing lyrics
	go a.syncLoop()

	// Exit
	a.waitForShutdown()

	return nil
}

func (a *App) syncLoop() {
	ticker := time.NewTicker(time.Duration(a.config.Player.PollInterval) * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-a.stopChan:
			return
		case <-ticker.C:
			if a.Lyrics == nil || len(*a.Lyrics) == 0 {
				shared.Debug("No lyrics available")
				continue
			}
			position, err := ipc.GetPlayerPosition(a.conn, a.player)
			if err != nil {
				shared.Debug("Failed to get position: %v\n", err)
				continue
			}

			var lyric shared.Lyric
			lyric.Lyric = "…"

			// Apply position offset
			position -= int(a.config.Player.PositionOffset) * 1_000
			shared.Debug("Position: %d ms\n", position)

			idx := GetCurrentLine(*a.Lyrics, position)
			shared.Debug("Current line index: %d\n", idx)

			// Only display if line changed
			if idx < len(*a.Lyrics) && len(*a.Lyrics) > 0 {
				lyric = (*a.Lyrics)[idx]
				a.displayLine(&lyric)
			}
		}
	}
}

func (a *App) watchTrackChanges() {
	ipc.WatchTrackChanges(a.conn, a.player, func(track *shared.Track) {
		a.handleTrackChange(track)
	})
}

func (a *App) handleTrackChange(track *shared.Track) {
	shared.Debug("Track changed: %s - %s\n", track.Artist, track.Title)

	a.currentTrack = track
	a.lastLine = nil
	a.Lyrics = nil

	rawLyrics, err := fetch.FetchLyrics(track)
	if err != nil {
		if err == shared.ErrNoLyricsFound {
			ipc.Notify(a.notifier, "No Lyrics Found", "No lyrics could be found for this song.")
		}

		shared.Debug("Failed to fetch lyrics: %v\n", err)
		a.Lyrics = nil
		return
	}

	shared.Debug("Fetched lyrics, length: %d\n", len(rawLyrics))

	lyrics, err := lyric.ParseLrc(rawLyrics)
	if err != nil {
		shared.Debug("Failed to parse lyrics: %v\n", err)
		a.Lyrics = nil
		return
	}
	shared.Debug("Parsed %d lyric lines\n", len(*lyrics))

	a.Lyrics = lyrics
}

func (a *App) displayLine(lyric *shared.Lyric) {
	switch a.config.General.DisplayMode {
	case "simple":
		display.Minimal(lyric.Lyric)
	case "tui":
		if a.tuiProgram != nil {
			a.tuiProgram.Send(display.TextUpdateMsg(*lyric))
		}
	case "window":
		a.fyneDisplay.Send(lyric.Lyric)
	default:
		display.Minimal(lyric.Lyric)
	}
}

func (a *App) waitForShutdown() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan

	a.Stop()
}

func (a *App) Stop() {
	close(a.stopChan)
	if a.conn != nil {
		a.conn.Close()
	}
	os.Exit(0)
}

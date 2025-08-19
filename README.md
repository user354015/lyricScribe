## What?
A simple Go script that downloads, checks and displays the synced lyrics for the song currently playing using MPRIS.

[>Showcase](Showcase.mp4)

## Why?
Nobody else was gonna do it properly, so screw it, I'm doing it myself.

## How?
1. Clone this repo ```git clone https://github.com/user354015/lyricScribe```
2. Move to the src directory ```cd lyricScribe/src```
3. Build the project ```go build .```

Or download the binary from the [releases](https://github.com/user354015/lyricScribe/releases)


### Floating lyrics on Hyprland
<details>
<summary>If you're on hyprland and want to replicate a setup similar to mine you can use these window rules along with a terminal of your choice</summary>

``` bash
windowrule = size 1100 160, class:lyricscribe
windowrule = move onscreen 410 60, class:lyricscribe
# windowrule = move onscreen 320 1100, class:lyricscribe
windowrule = pin, class:lyricscribe
windowrule = float, class:lyricscribe
windowrule = nofocus, class:lyricscribe
windowrule = opacity 0.8, class:lyricscribe
windowrule = noshadow, class:lyricscribe
windowrule = noblur, class:lyricscribe
```
and a minimal foot (any other terminal emulator works too) config
``` ini
font=IosevkatermSlab Nerd Font:size=30
app-id = "lyricscribe"

[colors]
alpha = 0
background = 000000
foreground = ff3b30
```
then bind launching a terminal to a shortcut:
```pgrep -f "foot.*lyricscribe" > /dev/null && pkill -f "foot.*lyricscribe" || foot -c ~/.config/foot/display.ini -e sh -c "~/.local/bin/lyricscribe"```
</details>

## Future of the project
This is my first real project and I plan to take good care of it. I have [many plans](/roadmap.md) for the future and will do my best to maintain and improve this for as long as possible. That being said, I am just some guy doing this on my (limited) free time and cannot make any promises.

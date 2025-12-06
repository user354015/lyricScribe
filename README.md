## What?
A simple Go script that downloads, checks and displays the synced lyrics for the song currently playing using MPRIS.

https://github.com/user-attachments/assets/7905037f-5c62-43a2-86e6-96e41b18af12

## Why?
Nobody else was gonna do it properly, so screw it, I'm doing it myself.

## How?
1. Clone this repo
```bash
git clone https://github.com/user354015/muse
```
2. Move to the cloned directory
```bash
cd muse
```
3. Build the project and make it executable
```bash
go build .
chmod +x muse
```
4. Move the default config to the correct location
```bash
mkdir -p ~/.config/muse
cp config.toml ~/.config/muse/
```

Or download the binary from the [releases](https://github.com/user354015/muse/releases)


## Configuration
Information about the config location and options can be found [here](/CONFIGURING.md)

### Floating lyrics on Hyprland
<details>
<summary>Click here if you're on hyprland and want to replicate a setup similar to mine. </summary>

You can use these window rules along with a terminal of your choice

``` bash
windowrule = size 1100 200, class:^(muse)$
windowrule = move onscreen 410 0, class:^(muse)$
# windowrule = move onscreen 410 160, class:^(muse)$
# windowrule = move onscreen 410 1000, class:^(muse)$
windowrule = pin, class:^(muse)$
windowrule = float, class:^(muse)$
windowrule = nofocus, class:^(muse)$
windowrule = opacity 0.9, class:^(muse)$
windowrule = noshadow, class:^(muse)$
windowrule = noblur, class:^(muse)$
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


## Credits
This project would have been impossible without the existance of [LRCLIB.net](https://www.LRCLIB.net). Please go donate to them if you are able to!


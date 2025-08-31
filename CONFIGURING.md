## List of options
This is a list of all available options. The configuration file is written in `TOML` format and must be located in `$XDG_CONFIG_HOME/lyrics/config.toml` (usually `~/.config/lyrics/config.toml`). Each of the following sub-sections should also be a section in the config.

Each subsection below corresponds to a section in the config file.

An example config file can be found [here](/config.toml)

---

### General (`[general]`)  

| Name                | Description                                                                                               | Accepted Values     | Default       |
| ------------------- | --------------------------------------------------------------------------------------------------------- | ------------------- | ------------- |
| `program_name`      | The name used to identify the program (e.g., in notifications).                                           | Any `string`        | `LyricScribe` |
| `check_for_updates` | Notifies when program version is different from latest version on GitHub.                                 | `true`, `false`     | `true`        |
| `program_mode`      | How lyrics should be displayed. `display` is for normal use, `debug` is for scripting/debugging.          | `display`, `debug`  | `display`     |
| `logging`           | Controls whether issues should be shown in the program output.                                            | `display`, `silent` | `silent`      |

### Search (`[search]`)  

| Name    | Description                                                                                                                                                                                                          | Accepted Values           | Default |
|---------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------|---------|
| `depth` | Controls how the program searches for lyrics. `match` requires an exact match (artist, album, song name, duration). `search` uses only artist and song name. `both` uses `match` first, with `search` as a fallback. | `match`, `search`, `both` | `both`  |

### Player (`[player]`)

| Name              | Description                                                                            | Accepted Values | Default |
| ----------------- | -------------------------------------------------------------------------------------- | --------------- | ------- |
| `player`          | The media player to connect to. Any string contained in the player’s name suffices.    | Any `string`    | `"mpv"` |
| `position_offset` | Time offset (in seconds) applied to the current playback position                      | Any `float`     | `-0.52` |
| `step`            | The time interval (in seconds) used for polling/stepping through playback.             | Any `float`     | `0.3`   |
| `silence_timeout` | The maximum silence duration (in seconds) before lyrics are considered paused/stopped. | Any `float`     | `3`     |

### Display (`[display]`)
| Name         | Description                                                              | Accepted values     | Default   |
| ------------ | ------------------------------------------------------------------------ | ------------------- | --------- |
| `foreground` | Text color used in display mode.                                         | Any hex color value | `#ffffff` |
| `background` | Background color used in display mode. Leave empty to disable completely | Any hex color value | `#000000` |

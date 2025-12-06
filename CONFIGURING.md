## List of options
This is a list of all available options. The configuration file is written in `TOML` format and must be located in `$XDG_CONFIG_HOME/lyrics/config.toml` (usually `~/.config/lyrics/config.toml`). Each of the following sub-sections should also be a section in the config file.

Each subsection below corresponds to a section in the config file.

An example config file can be found [here](/config.toml)

---

### General `[general]`

| Name                | Description                                                                                                         | Accepted Values           | Default  |
| ------------------- | ------------------------------------------------------------------------------------------------------------------- | ------------------------- | -------- |
| `program_name`      | The name used to identify the program (e.g., in notifications).                                                     | Any `string`              | `muse`   |
| `check_for_updates` | Notifies when program version is different from latest version on GitHub.                                           | `true`, `false`           | `true`   |
| `display_mode`      | How lyrics should be displayed. `window` or `tui` are intended for normal use, `simple` is for scripting/debugging. | `window`, `tui`, `simple` | `simple` |

### Player `[player]`

| Name              | Description                                                                                                                   | Accepted Values  | Default |
| ----------------- | ----------------------------------------------------------------------------------------------------------------------------- | ---------------- | ------- |
| `player`          | The media player to connect to. Any string contained in the player's name suffices.                                           | Any `string`     | `"mpv"` |
| `position_offset` | Time offset (in seconds) applied to the current playback position.                                                            | Any `int`, in ms | `-0.52` |
| `step`            | The time interval (in seconds) used for polling/stepping through playback.                                                    | Any `int`, in ms | `0.3`   |
| `poll_interval`   | A cooldown period between song changes to prevent spamming the API. The default value does not interrupt the user experience. | Any `int`, in ms | `3000`  |
| `silence_timeout` | The maximum silence duration (in seconds) before lyrics are considered paused/stopped.                                        | Any `int`, in ms | `3`     |

### Display `[display]`

| Name         | Description                                     | Accepted values     | Default   |
| ------------ | ----------------------------------------------- | ------------------- | --------- |
| `foreground` | Text color used in display mode.                | Any hex color value | `#ffffff` |
| `background` | Background color used in display mode.          | Any hex color value | `#000000` |

### Display `[display]` / *Exclusive to `window` mode*

| Name            | Description                                                                  | Accepted values              | Default |
| --------------- | ---------------------------------------------------------------------------- | ---------------------------- | ------- |
| `font`          | Typeface used when rendering lyrics. Leave empty for fallback.               | Any `string`, use full paths | `""`    |
| `font_size`     | Font size for the lyrics displayed.                                          | Any `int`, in px             | `32`    |
| `window_x`      | Window's spawn position. May not work depending on the compositor.           | Any `int`, in px             | `410`   |
| `window_y`      | Window's spawn position. May not work depending on the compositor.           | Any `int`, in px             | `0`     |
| `window_width`  | Window's spawn dimensions. May not work depending on the compositor.         | Any `int`, in px             | `1100`  |
| `window_height` | Window's spawn dimensions. May not work depending on the compositor.         | Any `int`, in px             | `250`   |

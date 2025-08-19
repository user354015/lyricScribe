## Planned features (in no particular order):

- [x] A better configuration experience (such as ability to use flags and a dedicated config file in `~/.config`).
- [ ] Support for specifically Spotify's (and other platforms) api
- [ ] A proper window without the need for a terminal
- [ ] A client-server setup allowing for more flexibility
    - [ ] Ability for multiple clients to hook up to a server instead of fetching duplicate information each.
    - [ ] A third output mode, returning only the current lyric, most suitable for scripting.
    - [ ] Locally caching lyrics up to a certain size for speed and to not abuse `lrclib.net`.
- [x] Performance optimizations
    - [x] Switch to using dbus signals instead of fetching information every cycle

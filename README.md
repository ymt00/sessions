# Bemenu sessions launcher

Depends on bemenu

## Usage in i3status-rs config file

```
[[block]]
block = "custom"
command = "echo Sessions"
interval = "once"
[[block.click]]
button = "left"
cmd = "/path/to/sessions_binary /path/to/sessions.json"
```
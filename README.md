# hG
## Intent 
- `hG` is a TUI for filtering/running commands returned from `history` with vim bindings
 
## Build
```sh
make build
```

## Usage
### Define alias
```sh
alias hG='function f() { history | grep "$*" | hgbin; unset -f f; }; f'
```
### Select from all commands
```sh
hG
```
### Filter commands
```sh
hG <search term>
```

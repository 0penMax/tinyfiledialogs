## tinyfiledialogs for Go

Pure Go file dialog helpers with the same public API as this package's earlier tinyfiledialogs bindings.

The package no longer depends on the bundled C implementation or on `cgo`. It uses platform-native dialog commands from Go.

### Installation

```
go get -v -u github.com/0penMax/tinyfiledialogs
```

Supported exported dialogs:

- `OpenFileDialog`
- `SaveFileDialog`
- `SelectFolderDialog`

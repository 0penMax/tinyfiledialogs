package tinyfiledialogs

import "strings"

// OpenFileDialog opens a file selection dialog.
//
// title - the title of the dialog.
//
// defaultPathOrFile - the default path or file to open. For directories, use a trailing slash.
//
// filePatterns is a list of file patterns to filter the files.
//
// singleFilterDescription is the description of the filter.
//
// allowMultipleSelection is a boolean to allow multiple file selection.
//
// Returns the selected file(s). If multiple files are selected, they are separated by a pipe character "|".
func OpenFileDialog(title string, defaultPathOrFile string, filePatterns []string, singleFilterDescription string, allowMultipleSelection bool) (path string, ok bool) {
	return openFileDialog(title, defaultPathOrFile, filePatterns, singleFilterDescription, allowMultipleSelection)
}

// SaveFileDialog opens a file save dialog.
//
// title - the title of the dialog.
//
// defaultPathOrFile - the default path or file to open. For directories, use a trailing slash.
//
// filePatterns is a list of file patterns to filter the files.
//
// singleFilterDescription is the description of the filter.
func SaveFileDialog(title string, defaultPathOrFile string, filePatterns []string, singleFilterDescription string) (path string, ok bool) {
	return saveFileDialog(title, defaultPathOrFile, filePatterns, singleFilterDescription)
}

// SelectFolderDialog opens a folder selection dialog.
//
// title - the title of the dialog.
//
// defaultPath - the default path to open.
func SelectFolderDialog(title string, defaultPath string) (path string, ok bool) {
	return selectFolderDialog(title, defaultPath)
}

func joinSelection(values []string) string {
	return strings.Join(values, "|")
}

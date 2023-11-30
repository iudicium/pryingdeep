package fsutils

import (
	"os"
	"path/filepath"
)

// / Similar to the touch command on *nix, where the file
// or directory will be created if it does not already exist.
// Returns the absolute path.
// The optional second boolean argument will force
// the method to treat the path as a file instead of a directory
// (useful when the filename has not extension).
// An optional 3rd boolean argument will force the method
// to treat the path as a directory even if a file extension is present.
//
// For example:
// `fsutil.Touch("./path/to/archive.old", false, true)`
//
// Normally, any file path with an extension is determined
// to be a file. However; the second argument (`false`)
// instructs the command to **not** force a file. The third
// argument (`true`) instructs the command to **treat the path
// like a directory**.
func Touch(path string, flags ...interface{}) string {
	abs := Abs(path)

	if !Exists(path) {
		forceFile := false
		forceDir := false

		if len(flags) > 0 {
			for i, flag := range flags {
				if i == 0 {
					forceFile = flag.(bool)
				} else if i == 1 {
					forceDir = flag.(bool)
				}
			}
		}

		ext := filepath.Ext(abs)

		if !forceDir && (forceFile || len(ext) > 0) {
			Mkdirp(filepath.Dir(abs))

			file, err := os.Create(abs)
			if err != nil {
				panic(err)
			}

			file.Close()
		} else {
			Mkdirp(abs)
		}
	}

	return abs
}

// Returns the fully resolved path, even if the
// path does not exist.
//
// ```
// fsutil.Abs("./does/not/exist")
// ```
// If the code above was run within `/home/user`, the
// result would be `/home/user/does/not/exist`.
func Abs(path string) string {
	abs, _ := filepath.Abs(path)
	return abs
}

// Mkdirp is the equivalent of [mkdir -p](https://en.wikipedia.org/wiki/Mkdir)
// It will generate the full directory path if it does not already
// exist.
func Mkdirp(path string) string {
	path = Abs(path)
	os.MkdirAll(path, os.ModePerm)
	return path
}

// Exists is a helper method to quickly
// determine whether a directory or file exists.
func Exists(path string) bool {
	if len(Abs(path)) == 0 {
		return false
	}

	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}

	return true

}

// IsDirectory determines whether the specified path
// represents a directory.
func IsDirectory(path string) bool {
	if !Exists(path) {
		return false
	}

	stat, err := os.Stat(path)
	if err != nil {
		return false
	}

	return stat.IsDir()
}

func WriteTextFile(path string, content string, args ...interface{}) error {
	path = Touch(path, true)
	perm := os.ModePerm

	if len(args) > 0 {
		perm = args[0].(os.FileMode)
	}

	return os.WriteFile(path, []byte(content), perm)
}

func ReadTextFile(path string) (string, error) {
	data, err := os.ReadFile(Abs(path))
	if err != nil {
		return "", err
	}

	return string(data), nil
}

package fs

import (
	"os"
	"os/user"
	"path/filepath"

	"github.com/dwarukira/findcare/pkg/rnd"
)

const IgnoreFile = ".ppignore"
const HiddenPath = ".findcare"
const PathSeparator = string(filepath.Separator)
const Home = "~"
const HomePath = Home + PathSeparator

// FileExists returns true if file exists and is not a directory.
func FileExists(fileName string) bool {
	if fileName == "" {
		return false
	}

	info, err := os.Stat(fileName)

	return err == nil && !info.IsDir()
}

// Abs returns the full path of a file or directory, "~" is replaced with home.
func Abs(name string) string {
	if name == "" {
		return ""
	}

	if len(name) > 2 && name[:2] == HomePath {
		if usr, err := user.Current(); err == nil {
			name = filepath.Join(usr.HomeDir, name[2:])
		}
	}

	result, err := filepath.Abs(name)

	if err != nil {
		panic(err)
	}

	return result
}

// PathExists tests if a path exists, and is a directory or symlink.
func PathExists(path string) bool {
	if path == "" {
		return false
	}

	info, err := os.Stat(path)

	if err != nil {
		return false
	}

	m := info.Mode()

	return m&os.ModeDir != 0 || m&os.ModeSymlink != 0
}

// PathWritable tests if a path exists and is writable.
func PathWritable(path string) bool {
	if !PathExists(path) {
		return false
	}

	tmpName := filepath.Join(path, "."+rnd.Token(8))

	if f, err := os.Create(tmpName); err != nil {
		return false
	} else if err := f.Close(); err != nil {
		return false
	} else if err := os.Remove(tmpName); err != nil {
		return false
	}

	return true
}

// Overwrite overwrites the file with data. Creates file if not present.
func Overwrite(fileName string, data []byte) bool {
	f, err := os.Create(fileName)
	if err != nil {
		return false
	}

	_, err = f.Write(data)
	return err == nil
}

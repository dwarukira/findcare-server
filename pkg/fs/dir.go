package fs

var OriginalPaths = []string{
	"~/.findcare",
}

var AssetPaths = []string{
	"assets",
}

func FindDir(dirs []string) string {
	for _, dir := range dirs {
		absDir := Abs(dir)
		if PathExists(absDir) {
			return absDir
		}
	}

	return ""
}

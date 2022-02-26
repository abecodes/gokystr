package gokystr

import (
	"io/fs"
	"regexp"
)

var extReg = regexp.MustCompile(`/^\.[^.^\/^\\]+$/gm`)

func validateExt(ext string) bool {
	return extReg.MatchString(ext)
}

func validatePath(dir string) bool {
	return fs.ValidPath(dir)
}

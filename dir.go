package gokystr

import (
	"os"
	"path"
)

func readFiles(dir, ext string, handle func(p string)) error {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		// just interested in files on root lvl

		if !entry.IsDir() && path.Ext(entry.Name()) == ext {
			handle(path.Join(dir, entry.Name()))
		}
	}

	return nil
}

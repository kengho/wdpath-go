package wdpath

import (
	"fmt"
	"os"
	"path/filepath"
	"github.com/kengho/errors"
)

func WdPath(path string) (string) {
	wd, err := filepath.Abs(filepath.Dir(os.Args[0]))
	errors.HandleErr(err)
	wdPath := fmt.Sprintf("%s\\%s", wd, path)
	return wdPath
}
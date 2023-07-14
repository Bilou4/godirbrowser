package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func checkPath(path string) (string, error) {
	c := filepath.Clean(path)

	r, err := filepath.EvalSymlinks(c)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return c, errors.New("invalid path specified")
		} else {
			return c, errors.New("unsafe path specified")
		}
	}

	err = inTrustedRoot(r, rootDir)
	if err != nil {
		fmt.Println("Error " + err.Error())
		if errors.Is(err, os.ErrNotExist) {
			return r, errors.New("invalid path specified")
		} else {
			return r, errors.New("unsafe path specified")
		}
	} else {
		return r, nil
	}
}
func inTrustedRoot(path string, trustedRoot string) error {
	for path != "/" {
		path = filepath.Dir(path)
		if path == trustedRoot {
			return nil
		}
	}
	return errors.New("path is outside of trusted root")
}

func byteCountBinary(b int64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %ciB", float64(b)/float64(div), "KMGTPE"[exp])
}

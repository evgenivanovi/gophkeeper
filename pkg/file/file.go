package file

import (
	"errors"
	"os"
	"path/filepath"
)

/* __________________________________________________ */

func Exists(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}

func Create(path string) error {

	if Exists(path) {
		return nil
	}

	dir := filepath.Dir(path)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	return nil

}

func CreateForce(path string) error {

	if Exists(path) {
		err := Remove(path)
		if err != nil {
			return err
		}
	}

	dir := filepath.Dir(path)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	return nil

}

func Remove(path string) error {
	return os.Remove(path)
}

/* __________________________________________________ */

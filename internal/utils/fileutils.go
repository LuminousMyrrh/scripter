package utils

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func IsDirExist(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return info.IsDir(), nil
}

func copyFile(src string, dest string) error {
	srcFileStat, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !srcFileStat.Mode().IsRegular() {
		return fmt.Errorf("File %s is not a regular file.", srcFileStat.Name())
	}

	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destination.Close()

	nBytes, err := io.Copy(destination, source)
	if err != nil {
		return err 
	}
	if nBytes == 0 {
		return fmt.Errorf("No bytes were copied.")
	}

	return nil
}

func CopyDir(dirPath string, destPath string) error {
	dir, err := os.Open(dirPath)
	if err != nil {
		return err
	}
	dirFiles, err := dir.ReadDir(0)
	if err != nil {
		return err
	}
	for _, entity := range dirFiles {
		entityPath := dirPath + "/" + entity.Name()
		if entity.Type().IsDir() {
			if !strings.HasPrefix(entity.Name(), ".") {
				newDir := destPath + "/" + entity.Name()
				os.Mkdir(newDir, 0755)
				CopyDir(entityPath, newDir)
			}
		} else {
			err := copyFile(entityPath, destPath + "/" + entity.Name())
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func CopyTemplate(template string, dest string) error {
	err := CopyDir(template, dest)
	if err != nil {
		return err
	}
	return nil
}

package utils

import (
	"fmt"
	"io"
	"os"
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

func CopyFiles(dirPath string, destPath string) error {
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
		fmt.Println(entityPath)
		fmt.Println(destPath)
		if entity.Type().IsDir() {
			CopyFiles(entityPath, destPath)
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
	err := CopyFiles(template, dest)
	if err != nil {
		return err
	}
	return nil
}


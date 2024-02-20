package main

import (
	"os"
)

type Filesystem interface {
	readFile(filepath string) (string, error)
	writeFile(filepath string, content string) error
	fileExists(filepath string) (bool, error)
}

type RealFilesystem struct{}

func (fs RealFilesystem) readFile(filepath string) (string, error) {
	content, err := os.ReadFile(filepath)

	if err != nil {
		return "", err
	}

	return string(content), nil
}

func (fs RealFilesystem) writeFile(filepath string, content string) error {
	err := os.WriteFile(filepath, []byte(content), 0644)

	if err != nil {
		return err
	}

	return nil
}

func (fs RealFilesystem) fileExists(filepath string) (bool, error) {
	_, err := os.Stat(filepath)

	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

type MockFilesystem struct {
	files map[string]string
}

func (fs MockFilesystem) readFile(filepath string) (string, error) {
	content, ok := fs.files[filepath]

	if !ok {
		return "", os.ErrNotExist
	}

	return content, nil
}

func (fs MockFilesystem) writeFile(filepath string, content string) error {
	fs.files[filepath] = content
	return nil
}

func (fs MockFilesystem) fileExists(filepath string) (bool, error) {
	_, ok := fs.files[filepath]

	if !ok {
		return false, nil
	}

	return true, nil
}

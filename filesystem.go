package main

import (
	"os"
)

type Filesystem interface {
	ReadFile(filepath string) (string, error)
	WriteFile(filepath string, content string) error
	FileExists(filepath string) (bool, error)
}

type RealFilesystem struct{}

func (fs RealFilesystem) ReadFile(filepath string) (string, error) {
	content, err := os.ReadFile(filepath)

	if err != nil {
		return "", err
	}

	return string(content), nil
}

func (fs RealFilesystem) WriteFile(filepath string, content string) error {
	err := os.WriteFile(filepath, []byte(content), 0644)

	if err != nil {
		return err
	}

	return nil
}

func (fs RealFilesystem) FileExists(filepath string) (bool, error) {
	_, err := os.Stat(filepath)

	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

type MockFilesystem struct {
	files map[string]string
}

func (fs MockFilesystem) ReadFile(filepath string) (string, error) {
	content, ok := fs.files[filepath]

	if !ok {
		return "", os.ErrNotExist
	}

	return content, nil
}

func (fs MockFilesystem) WriteFile(filepath string, content string) error {
	fs.files[filepath] = content
	return nil
}

func (fs MockFilesystem) FileExists(filepath string) (bool, error) {
	_, ok := fs.files[filepath]

	if !ok {
		return false, nil
	}

	return true, nil
}

func MakeMockFilesystem(files map[string]string) MockFilesystem {
	return MockFilesystem{files: files}
}

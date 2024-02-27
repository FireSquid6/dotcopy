package main

import (
	"log"
	"os"
	"path"

	"gopkg.in/yaml.v2"
)

type Dotfile struct {
	TemplateText     string // raw text of the template file
	SlotText         string // raw text contained in the slotfile
	TemplateFilepath string // path to the template file. Absolute path.
	SlotFilepath     string // path to the slot file. Absolute path.
	CompiledFilepath string // path to the compiled file. Absolute path.
}

type DotcopyYaml struct {
	Template string `yaml:"template"`
	Slotfile string `yaml:"slotfile"`
	Location string `yaml:"location"`
}

type LocalConfig struct {
	RootFilepath     string `yaml:"root_filepath"`
	MachineDirectory string `yaml:"machine_directory"`
}

func ParseDotfiles(fs Filesystem, localConfig LocalConfig) ([]Dotfile, error) {
	yamlData, err := unmarshalYaml(fs, path.Join(localConfig.RootFilepath, "dotcopy.yaml"))
	if err != nil {
		return nil, err
	}

	dotfiles := []Dotfile{}

	for _, yamlObj := range yamlData {
		dotfile := Dotfile{
			TemplateFilepath: path.Join(localConfig.RootFilepath, yamlObj.Template),
			SlotFilepath:     path.Join(localConfig.RootFilepath, localConfig.MachineDirectory, yamlObj.Slotfile),
			CompiledFilepath: yamlObj.Location,
			TemplateText:     "",
			SlotText:         "",
		}

		templateText, err := fs.ReadFile(dotfile.TemplateFilepath)
		if err == nil {
			dotfile.TemplateText = templateText
		}

		slotText, err := fs.ReadFile(dotfile.SlotFilepath)
		if err == nil {
			dotfile.SlotText = slotText
		}

		dotfiles = append(dotfiles, dotfile)

	}

	return dotfiles, nil
}

func unmarshalYaml(fs Filesystem, filepath string) ([]DotcopyYaml, error) {
	rootFileText, err := fs.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	yamlData := []DotcopyYaml{}
	err = yaml.Unmarshal([]byte(rootFileText), &yamlData)

	if err != nil {
		return nil, err
	}

	return yamlData, nil
}

func ParseLocalConfig(fs Filesystem, filepath string) (LocalConfig, error) {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	localConfig := LocalConfig{
		RootFilepath:     path.Join(dirname, "dotfiles"),
		MachineDirectory: hostname,
	}
	localConfigText, err := fs.ReadFile(filepath)
	if err != nil {
		return localConfig, err
	}

	err = yaml.Unmarshal([]byte(localConfigText), &localConfig)
	if err != nil {
		return localConfig, err
	}

	return localConfig, nil
}

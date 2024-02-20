package main

import (
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

func ParseDotfiles(fs Filesystem, rootFilepath string) ([]Dotfile, error) {
	// read the root file as an array of dotcopy yaml objects
	yamlData, err := unmarshalYaml(fs, rootFilepath)

	if err != nil {
		return nil, err
	}

	dotfiles := []Dotfile{}

	for _, yamlObj := range yamlData {
		dotfile := Dotfile{
			TemplateFilepath: yamlObj.Template,
			SlotFilepath:     yamlObj.Slotfile,
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

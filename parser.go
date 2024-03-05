package main

import (
	"log"
	"os"
	"path"
	"strings"

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

type GlobalVar struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
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

		if strings.HasPrefix(dotfile.CompiledFilepath, "~/") {
			dotfile.CompiledFilepath = path.Join(fs.Homedir(), dotfile.CompiledFilepath[1:])
		}

		templateText, err := fs.ReadFile(dotfile.TemplateFilepath)
		if err == nil {
			dotfile.TemplateText = templateText
		}

		// if the slotfile is empty, we just don't want to bother
		// if there is no slotfile, we just copy the dotfile as is
		if yamlObj.Slotfile != "" {
			slotText, err := fs.ReadFile(dotfile.SlotFilepath)
			if err == nil {
				dotfile.SlotText = slotText
			}
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

func ParseLocalConfig(fs Filesystem) (LocalConfig, error) {
	dirname := fs.Homedir()

	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	filepath := path.Join(dirname, ".config/dotcopy/localconfig.yaml")

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

func ParseGlobalVars(fs Filesystem, localconfig LocalConfig) ([]GlobalVar, error) {
	filepath := path.Join(localconfig.RootFilepath, "vars.yaml")

	globalVars := []GlobalVar{}
	globalVarText, err := fs.ReadFile(filepath)
	if err != nil {
		return globalVars, err
	}

	err = yaml.Unmarshal([]byte(globalVarText), &globalVars)

	return globalVars, nil
}

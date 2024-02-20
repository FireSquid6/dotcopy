package main

type Dotfile struct {
	templateText     string // raw text of the template file
	slotText         string // raw text contained in the slotfile
	templateFilepath string // path to the template file. Absolute path.
	slotFilepath     string // path to the slot file. Absolute path.
	compiledFilepath string // path to the compiled file. Absolute path.
}

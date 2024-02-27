package main

import (
	"reflect"
	"testing"
)

func TestParseDotfiles(t *testing.T) {
	// create a new mock filesystem
	fs := MakeMockFilesystem(map[string]string{
		"/dotfiles/dotcopy.yaml": `---
- template: template.txt
  slotfile: slotfile.txt
  location: /path/to/somewhere/compiledfile.txt`,
		"/dotfiles/template.txt":         "template text",
		"/dotfiles/machine/slotfile.txt": "slot text",
	})

	// parse the dotfiles
	dotfiles, err := ParseDotfiles(fs, LocalConfig{
		RootFilepath:     "/dotfiles",
		MachineDirectory: "machine",
	})
	// check for errors
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	// result should be a single dotfile
	if len(dotfiles) != 1 {
		t.Errorf("expected 1 dotfile, got %d", len(dotfiles))
	}
	// check the dotfile's fields
	expected := Dotfile{
		TemplateText:     "template text",
		SlotText:         "slot text",
		TemplateFilepath: "/dotfiles/template.txt",
		SlotFilepath:     "/dotfiles/machine/slotfile.txt",
		CompiledFilepath: "/path/to/somewhere/compiledfile.txt",
	}

	if !reflect.DeepEqual(dotfiles[0], expected) {
		t.Errorf("expected %v, got %v", expected, dotfiles[0])
	}
}

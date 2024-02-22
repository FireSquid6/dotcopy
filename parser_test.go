package main

import (
	"reflect"
	"testing"
)

func TestParseDotfiles(t *testing.T) {
	// create a new mock filesystem
	fs := MakeMockFilesystem(map[string]string{
		"/path/to/rootfile.yaml": `---
- template: /path/to/template.txt
  slotfile: /path/to/slotfile.txt
  location: /path/to/compiledfile.txt`,
		"/path/to/template.txt": "template text",
		"/path/to/slotfile.txt": "slot text",
	})

	// parse the dotfiles
	dotfiles, err := ParseDotfiles(fs, "/path/to/rootfile.yaml")

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
		TemplateFilepath: "/path/to/template.txt",
		SlotFilepath:     "/path/to/slotfile.txt",
		CompiledFilepath: "/path/to/compiledfile.txt",
	}

	if !reflect.DeepEqual(dotfiles[0], expected) {
		t.Errorf("expected %v, got %v", expected, dotfiles[0])
	}

}

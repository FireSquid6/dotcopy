package main

import (
	"reflect"
	"testing"
)

func TestParseDotfiles(t *testing.T) {
	fs := MakeMockFilesystem(map[string]string{
		"/dotfiles/dotcopy.yaml": `---
- template: template.txt
  slotfile: slotfile.txt
  location: /path/to/somewhere/compiledfile.txt`,
		"/dotfiles/template.txt":         "template text",
		"/dotfiles/machine/slotfile.txt": "slot text",
	})

	dotfiles, err := ParseDotfiles(fs, LocalConfig{
		RootFilepath:     "/dotfiles",
		MachineDirectory: "machine",
	})
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if len(dotfiles) != 1 {
		t.Errorf("expected 1 dotfile, got %d", len(dotfiles))
	}
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

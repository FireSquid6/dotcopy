package main

import (
	"testing"
)

func TestCompileDotfile(t *testing.T) {
	dotfile := Dotfile{
		TemplateText: "Hello, {{name}}!",
		SlotText:     "--- {{name}}\nWorld\n---",
	}

	compiledText, _ := CompileDotfile(dotfile, []GlobalVar{})
	expected := "Hello, World!"
	if compiledText != expected {
		t.Errorf("expected %s, got %s", expected, compiledText)
	}
}

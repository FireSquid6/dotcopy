package main

import (
	"strings"
)

func CompileDotfile(dotfile Dotfile) (string, string) {
	compiledText := dotfile.TemplateText
	// iterate through the lines of the SlotText using a FSM and generate compiledText
	// return compiledText, dotfile.CompiledFilepath
	slots := getSlots(dotfile.SlotText)

	for _, slot := range slots {
		if strings.Contains(compiledText, slot.slotText) {
			compiledText = strings.ReplaceAll(compiledText, slot.slotText, slot.insert)
		}
	}

	return compiledText, dotfile.CompiledFilepath
}

type Slot struct {
	insert   string // the text to insert into the template
	slotText string // the slot text to look for
}

func getSlots(slotfile string) []Slot {
	slots := []Slot{}
	state := "outside" // whether we are inside or outside of a slot
	currentSlot := Slot{}

	lines := strings.Split(slotfile, "\n")

	for _, line := range lines {
		if state == "inside" {
			if strings.HasPrefix(line, "---") {
				currentSlot.insert = strings.TrimSuffix(currentSlot.insert, "\n")
				slots = append(slots, currentSlot)
				state = "outside"
			} else {
				currentSlot.insert += line + "\n"
			}

		} else {
			if strings.HasPrefix(line, "---") {
				currentSlot.slotText = strings.Split(line, " ")[1]
				state = "inside"
			}
		}

	}

	return slots
}

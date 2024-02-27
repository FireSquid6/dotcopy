package main

import (
	"fmt"
	"log"
	"os"

	"github.com/radovskyb/watcher"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "dotcopy",
		Usage: "A dotfile compiler",
		Action: func(c *cli.Context) error {
			fmt.Println("Hey there! Try dotcopy help for more info.")
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:  "build",
				Usage: "Compiles and copies your dotfiles to their appropriate location",
				Action: func(c *cli.Context) error {
					errors := Dotcopy()
					if errors != "" {
						fmt.Println(errors)
					} else {
						fmt.Println("\nDotfiles compiled successfully!")
					}

					return nil
				},
			},
			{
				Name:  "watch",
				Usage: "Starts the file watcher. Also probably managed by systemd.",
				Action: func(c *cli.Context) error {
					Watch()
					return nil
				},
			},
			{
				Name:  "init",
				Usage: "Initializes a basic localconfig",
				Action: func(c *cli.Context) error {
					fmt.Println("Not implemented yet")
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

// the core of dotcopy
// compiles and copies your dotfiles to their appropriate location
func Dotcopy() string {
	fs := MakeRealFilesystem()

	localConfig, err := ParseLocalConfig(fs)
	if err != nil {
		return "Error parsing localconfig. Does ~/.config/dotcopy/localconfig.yaml exist?"
	}

	dotfiles, err := ParseDotfiles(fs, localConfig)
	if err != nil {
		log.Println(err)
		return "Error parsing dotfiles"
	}

	for _, dotfile := range dotfiles {
		text, filepath := CompileDotfile(dotfile)

		err := fs.WriteFile(filepath, text)
		if err != nil {
			log.Println(err)
			return "Error writing file"
		}

		slotfileLog := dotfile.SlotFilepath
		if slotfileLog != "" {
			slotfileLog = "No slotfile"
		}

		fmt.Println("Compiled:", dotfile.TemplateFilepath, "+", slotfileLog, "-->", filepath)
	}

	return ""
}

const (
	CHANGE_NOTIFICATION  = "Dotcopy detected a change in your dotfiles. Recompiling..."
	FAIL_NOTIFICATION    = "Dotcopy failed to compile your dotfiles. Run `dotcopy build` for more info."
	SUCCESS_NOTIFICATION = "Dotcopy successfully compiled your dotfiles. Make sure to reload."
)

func Watch() error {
	fs := MakeRealFilesystem()
	localconfig, err := ParseLocalConfig(fs)
	if err != nil {
		return err
	}

	log.Println("Watching", localconfig.RootFilepath)

	// send notification using notify-send on update
	w := watcher.New()
	if err := w.AddRecursive(localconfig.RootFilepath); err != nil {
		log.Fatalln(err)
	}

	go func() {
		for {
			select {
			case event := <-w.Event:
				fmt.Println(event) // Print the event's info.
			case err := <-w.Error:
				log.Fatalln(err)
			case <-w.Closed:
				return
			}
		}
	}()

	return nil
}

// TODO:
// make a cli with commands:
// - build - runs Dotcopy() and gives diagnostic notification
// - watch - starts the file watcher

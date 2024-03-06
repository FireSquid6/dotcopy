package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name: "dotcopy",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "silent",
				Usage:   "Silence all output to stdout. Use `-s -d` to silence all output.",
				Aliases: []string{"s"},
				Value:   false,
			},
			&cli.BoolFlag{
				Name:    "disable-notifications",
				Usage:   "Disable system notifications",
				Aliases: []string{"d", "n"},
				Value:   false,
			},
		},
		Usage: "A dotfile compiler",
		Action: func(c *cli.Context) error {
			logger := MakeRealLogger(c.Bool("disable-notifications"), c.Bool("silent"))
			output := Dotcopy(logger)

			if output == "" {
				logger.SuccessfulBuild()
			} else {
				logger.Error(output)
			}

			return nil
		},
		Commands: []*cli.Command{
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
func Dotcopy(logger Logger) string {
	fs := MakeRealFilesystem()

	localConfig, err := ParseLocalConfig(fs)
	if err != nil {
		return "Error parsing localconfig. Does ~/.config/dotcopy/localconfig.yaml exist?"
	}

	dotfiles, err := ParseDotfiles(fs, localConfig)
	if err != nil {
		return "Error parsing dotfiles"
	}

	globalVars, err := ParseGlobalVars(fs, localConfig)
	if err != nil {
		log.Println(err)
	}

	for _, dotfile := range dotfiles {
		text, filepath := CompileDotfile(dotfile, globalVars)

		err := fs.WriteFile(filepath, text)
		if err != nil {
			return "Error writing file"
		}

		slotfileLog := dotfile.SlotFilepath
		if slotfileLog != "" {
			slotfileLog = "No slotfile"
		}

		logger.Info(strings.Join([]string{"Compiled:", dotfile.TemplateFilepath, "+", slotfileLog, "-->", filepath}, " "))
	}

	return ""
}

//
// func watch(paths ...string) {
// 	if len(paths) < 1 {
// 		os.Exit(1)
// 	}
//
// 	// Create a new watcher.
// 	w, err := fsnotify.NewWatcher()
// 	if err != nil {
// 		os.Exit(1)
// 	}
// 	defer w.Close()
//
// 	// Start listening for events.
// 	go watchLoop(w)
//
// 	// Add all paths from the commandline.
// 	for _, p := range paths {
// 		err = w.Add(p)
// 		if err != nil {
// 			os.Exit(1)
// 		}
// 	}
//
// 	fmt.Println("Watching", paths)
// 	<-make(chan struct{}) // Block forever
// }
//
// func watchLoop(w *fsnotify.Watcher) {
// 	i := 0
// 	for {
// 		select {
// 		// Read from Errors.
// 		case err, ok := <-w.Errors:
// 			if !ok { // Channel was closed (i.e. Watcher.Close() was called).
// 				return
// 			}
// 			fmt.Println("error:", err)
// 		// Read from Events.
// 		case e, ok := <-w.Events:
// 			if !ok { // Channel was closed (i.e. Watcher.Close() was called).
// 				return
// 			}
//
// 			// Just print the event nicely aligned, and keep track how many
// 			// events we've seen.
// 			notify.Notify("Dotcopy", "Dotcopy: Change detected", "rebuilding dotfiles...", "")
// 			Dotcopy()
// 			notify.Notify("Dotcopy", "Dotcopy: Dotfiles Built", "Remember to reload your apps! (i3, polybar, etc.)", "")
// 			i++
// 			fmt.Println(i, e)
// 		}
// 	}
// }

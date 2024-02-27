package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "dotcopy",
		Usage: "A dotfile compiler",
		Action: func(c *cli.Context) error {
			fmt.Println("Hey there! Try dotcopy --help for more info.")
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
				Usage: "Starts the file watcher",
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

		err := fs.WriteFile(text, filepath)
		if err != nil {
			log.Println(err)
			return "Error writing file"
		}

		fmt.Println("Wrote ", filepath)
	}

	return ""
}

// TODO:
// make a cli with commands:
// - build - runs Dotcopy() and gives diagnostic notification
// - watch - starts the file watcher

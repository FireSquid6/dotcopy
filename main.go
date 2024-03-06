package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

const VERSION = "0.3.0"

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
		Usage: "Builds your dotfiles. See https://dotcopy.firesquid.co",
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
			{
				Name:  "version",
				Usage: "Prints the version of dotcopy",
				Action: func(c *cli.Context) error {
					fmt.Println("v" + VERSION)
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

package main

import (
	"github.com/fsnotify/fsnotify"
	"log"
)

func main() {
	// I'm like 90% sure this doesn't work right
	log.Println("Starting watcher")
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				if event.Has(fsnotify.Write) {
					log.Println("modified file:", event.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add("/tmp")
	if err != nil {
		log.Fatal(err)
	}

	<-make(chan struct{})
}

func Compile() {
	// filesystem := RealFilesystem{}
	// figure out where the dotfiles are located
	// parse the dotfiles using ParseDotfiles
	// for each dotfile, compile it using CompileDotfile
	// for each compiled dotfile, write it to the output directory
}

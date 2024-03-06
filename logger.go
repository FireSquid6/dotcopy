package main

import (
	"log"

	"github.com/martinlindhe/notify"
)

type Logger interface {
	Info(msg string)
	Error(err string)
	SuccessfulBuild()
	Notification(title string, msg string)
}

type RealLogger struct {
	useNotification bool
	useStdout       bool
}

func (l RealLogger) Info(msg string) {
	if l.useStdout {
		log.Println(msg)
	}
}

func (l RealLogger) Error(err string) {
	if l.useStdout {
		log.Println(err)
	}

	if l.useNotification {
		notify.Notify("Dotcopy", "Error", err, "")
	}
}

func (l RealLogger) SuccessfulBuild() {
	if l.useStdout {
		log.Println("Dotfiles compiled successfully!")
	}

	if l.useNotification {
		notify.Notify("Dotcopy", "Dotfiles compiled successfully!", "", "")
	}
}

func (l RealLogger) Notification(title string, msg string) {
	if l.useNotification {
		// send notification
		notify.Notify("Dotcopy", title, msg, "")
	}
}

func MakeRealLogger(useNotification bool, useStdout bool) Logger {
	return RealLogger{useNotification, useStdout}
}

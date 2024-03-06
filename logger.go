package main

import (
	"github.com/martinlindhe/notify"
	"log"
)

type Logger interface {
	Info(msg string)
	Error(err error)
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
func (l RealLogger) Error(err error) {
	if l.useStdout {
		log.Println(err)
	}

	if l.useNotification {
		notify.Notify("Dotcopy", "Error", err.Error(), "")
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

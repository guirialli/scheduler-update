package main

import (
	"time"

	"github.com/guirialli/scheduler-update/internals/scheduler/apt"
	"github.com/guirialli/scheduler-update/internals/scheduler/flatpak"
	"github.com/guirialli/scheduler-update/internals/whoami"
)

func main() {
	err := whoami.NewWhoami().IsRoot()
	if err != nil {
		panic(err)
	}

	apt := apt.NewApt()
	flatpak := flatpak.NewFlatpak()
	for true {
		go apt.Upgrade()
		go flatpak.Update()
		time.Sleep(1 * time.Hour)
	}
}

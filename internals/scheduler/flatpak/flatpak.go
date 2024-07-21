package flatpak

import (
	"log"
	"os/exec"
)

type flatpak struct{}

func NewFlatpak() *flatpak {
	return &flatpak{}
}

func (f *flatpak) Update() error {
	cmd := exec.Command("flatpak", "update", "-y")
	err := cmd.Run()
	if err != nil {
		log.Println(err.Error())
		return err
	}

	log.Println("Flatpak update exec with sucess")
	return nil
}

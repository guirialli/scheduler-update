package apt

import (
	"log"
	"os/exec"
)

type apt struct{}

func NewApt() *apt {
	return &apt{}
}
func (a *apt) Update() error {
	cmd := exec.Command("apt", "update")
	err := cmd.Run()
	if err != nil {
		log.Println(err)
		return err
	}

	log.Println("Apt update exec with sucess")
	return nil
}

func (a *apt) Upgrade() error {
	err := a.Update()
	if err != nil {
		log.Println(err)
		return err
	}

	cmd := exec.Command("apt", "dist-upgrade", "-y")
	err = cmd.Run()
	if err != nil {
		log.Println(err)
		return err
	}

	log.Println("Apt upgrade exec with sucess")
	return nil
}

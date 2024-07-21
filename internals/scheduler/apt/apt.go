package apt

import (
	"log"
	"os/exec"
)

type Apt struct{}

func NewApt() *Apt {
	return &Apt{}
}
func (a *Apt) Update() error {
	cmd := exec.Command("apt", "update")
	err := cmd.Run()
	if err != nil {
		log.Println(err)
		return err
	}

	log.Println("Apt update exec with sucess")
	return nil
}

func (a *Apt) Upgrade() error {
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

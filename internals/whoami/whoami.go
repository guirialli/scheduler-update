package whoami

import (
	"errors"
	"os/exec"
)

type whoami struct{}

func NewWhoami() *whoami {
	return &whoami{}
}

func (w *whoami) IsRoot() error {
	out, err := exec.Command("whoami").Output()

	if err != nil {
		return err
	}
	if string(out) != "root\n" {
		return errors.New("run this program with root")
	}
	return nil
}

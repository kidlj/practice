package main

import (
	"errors"
	"os"

	"github.com/kidlj/log"
	"github.com/kidlj/log/handlers/cli"
)

func main() {
	h := cli.New(os.Stderr)
	log.SetHandler(h)

	log.WithFields(log.Fields{
		"name": "mellon",
		"age":  30,
	}).Info("logged")

	log.WithField("sex", "male").Infof("This is a %s message", "debug")
	log.WithField("sex", "male").Debugf("This is a %s message", "debug")

	log.WithError(errors.New("This is a error")).Errorf("good night, %s.", "Jian Li")
}

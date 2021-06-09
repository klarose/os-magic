package main

import (
	"os"
	"runtime"

	"github.com/klarose/os-magic/pkg/oslayer"
	"github.com/sirupsen/logrus"
)

func main() {
	osIf := oslayer.Get(runtime.GOOS)
	if osIf == nil {
		logrus.Errorf("OS '%v' not supported", runtime.GOOS)
		os.Exit(1)
	}

	err := osIf.DoThing()

	if err != nil {
		logrus.Errorf("Failed to do thing: %v", err.Error())
		os.Exit(1)
	}
}

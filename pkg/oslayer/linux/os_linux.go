package linux

import (
	"github.com/klarose/os-magic/pkg/oslayer/osinterface"
	"github.com/sirupsen/logrus"
)

// LinuxInterface is the windows implementation of OSInterface
type linuxInterface struct {
}

func init() {
	print("Linux init")

	linuxIf := linuxInterface{}

	osinterface.Register(&linuxIf)
}

// OS returns linux
func (w *linuxInterface) OS() string {
	return "linux"
}

// DoThing does the linux thing
func (w *linuxInterface) DoThing() error {
	localDoIt()
	return nil
}

// Not exported, yet doesn't make linters angry
func localDoIt() {
	logrus.Print("Running on linux")
}

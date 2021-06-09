package windows

import (
	"github.com/klarose/os-magic/pkg/oslayer/osinterface"
	"github.com/sirupsen/logrus"
)

// WindowsInterface is the windows implementation of OSInterface
type windowsInterface struct {
}

func init() {

	windowsIf := windowsInterface{}

	osinterface.Register(&windowsIf)

}

// OS returns windows
func (w *windowsInterface) OS() string {
	return "windows"
}

// DoThing does the windows thing
func (w *windowsInterface) DoThing() error {
	windowsLocalDoIt()
	return nil
}

// Not exported, yet doesn't make linters angry
func windowsLocalDoIt() {
	logrus.Print("Running on windows")
}

package osinterface

import (
	"errors"
)

// OsInterface provides a per-os abstraction
type OsInterface interface {
	OS() string
	DoThing() error
}

var runningImplementations map[string]OsInterface

func init() {
	runningImplementations = make(map[string]OsInterface)
}

// Register registers an interface which may be later looked up using Get
func Register(perOS OsInterface) error {
	if perOS == nil {
		return errors.New("perOS cannot be nil")
	}

	runningImplementations[perOS.OS()] = perOS
	return nil
}

func Get(os string) OsInterface {
	perOS, ok := runningImplementations[os]

	if !ok {
		return nil
	}

	return perOS
}

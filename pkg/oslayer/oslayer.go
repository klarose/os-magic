package oslayer

import (
	_ "github.com/klarose/os-magic/pkg/oslayer/imports"
	"github.com/klarose/os-magic/pkg/oslayer/osinterface"
)

func Get(osName string) osinterface.OsInterface {
	return osinterface.Get(osName)
}

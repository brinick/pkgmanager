package pkginstaller

import (
	"fmt"

	"github.com/brinick/pkgmanager/ayum"
)

type Opts interface{ fmt.Stringer }

func choose(name string) func(opts Opts, logpath string) PkgInstaller {
	switch name {
	case "ayum":
		return ayum.New
	case "dnf":
		panic("dnf installer not yet implemented")
	default:
		panic(fmt.Sprintf("%s: unknown package manager, must be ayum or dnf"))
	}
}

type PkgInstaller interface {
}

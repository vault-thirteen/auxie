package dependency

import (
	"runtime/debug"
	"testing"

	"github.com/vault-thirteen/auxie/tester"
)

func Test_New(t *testing.T) {
	aTest := tester.New(t)
	var err error
	var dep *Dependency

	// 1.
	dep, err = New(nil)
	aTest.MustBeAnError(err)

	// 2.
	dep, err = New(&debug.Module{Path: "golang.org/x/sys", Version: "v0.0.1"})
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(dep.Version(), "v0.0.1")

	// 3.
	dep, err = New(&debug.Module{Path: "golang.org/x/sys", Version: "x"})
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(dep.Version(), "x")

	// 4.
	dep, err = New(&debug.Module{Path: "garbage", Version: ""})
	aTest.MustBeAnError(err)

}
func Test_parseModulePath(t *testing.T) {
	aTest := tester.New(t)
	var err error

	// 1.
	x := Dependency{fullPath: "garbage"}
	err = x.parseModulePath()
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(x.Name(), "")

	// 2.
	x = Dependency{fullPath: "github.com/kr/pretty"}
	err = x.parseModulePath()
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(x.Name(), "pretty")

	// 3.
	x = Dependency{fullPath: "github.com/vault-thirteen/Simpel-Chat-Server/src"}
	err = x.parseModulePath()
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(x.Name(), "Simpel-Chat-Server")

	// 4.
	x = Dependency{fullPath: "golang.org/x/sys"}
	err = x.parseModulePath()
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(x.Name(), "x/sys")

	// 5.
	x = Dependency{fullPath: "github.com/is-broken"}
	err = x.parseModulePath()
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(x.Name(), "")

	// 6.
	x = Dependency{fullPath: "github.com/acc/repo/v2"}
	err = x.parseModulePath()
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(x.Name(), "repo/v2")

	// 7.
	x = Dependency{fullPath: "example.org/something"}
	err = x.parseModulePath()
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(x.Name(), "example.org/something")

	// 8.
	var xx *Dependency
	xx = nil
	err = xx.parseModulePath() // Yes, Golang allows to call methods from NULL ! What a disaster !
	aTest.MustBeAnError(err)
}
func Test_Name(t *testing.T) {
	aTest := tester.New(t)

	x := Dependency{name: "name"}
	aTest.MustBeEqual(x.Name(), x.name)
}
func Test_Version(t *testing.T) {
	aTest := tester.New(t)

	x := Dependency{version: "version"}
	aTest.MustBeEqual(x.Version(), x.version)
}

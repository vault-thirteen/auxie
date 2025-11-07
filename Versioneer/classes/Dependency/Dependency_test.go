package dependency

import (
	"testing"

	"github.com/vault-thirteen/auxie/tester"
)

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

func Test_isStringAVersionPostfix(t *testing.T) {
	aTest := tester.New(t)

	aTest.MustBeEqual(isStringAVersionPostfix(""), false)
	aTest.MustBeEqual(isStringAVersionPostfix("x"), false)
	aTest.MustBeEqual(isStringAVersionPostfix("v"), false)
	aTest.MustBeEqual(isStringAVersionPostfix("qwerty"), false)
	aTest.MustBeEqual(isStringAVersionPostfix("vzero"), false)
	aTest.MustBeEqual(isStringAVersionPostfix("v0"), false)
	aTest.MustBeEqual(isStringAVersionPostfix("v1"), true)
	aTest.MustBeEqual(isStringAVersionPostfix("v123"), true)
	aTest.MustBeEqual(isStringAVersionPostfix("v1.2"), false)
	aTest.MustBeEqual(isStringAVersionPostfix("v1.k"), false)
	aTest.MustBeEqual(isStringAVersionPostfix("v100m"), false)
}

func Test_getLastTwoPathParts(t *testing.T) {
	aTest := tester.New(t)

	aTest.MustBeEqual(getLastTwoPathParts(""), "")
	aTest.MustBeEqual(getLastTwoPathParts("github.com/someone/something/v2"), "something/v2")
}

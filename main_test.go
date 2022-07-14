package main

import (
	"fmt"
	"os"
	"testing"
)

// FIXME: https://stackoverflow.com/questions/49193480/golang-flag-redefined

func TestMain(m *testing.M) {
	fmt.Println("TEST SETUP")
	code := m.Run()
	os.Exit(code)
}

func TestPackageWithDefaultFlags(t *testing.T) {
	main()
}

func TestPackageWithDebugFlag(t *testing.T) {
	os.Args = append(os.Args, "-debug")
	main()
}

// +build mage

package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"runtime"
)

// Main task
var Default = Build
func Build() {
	color.Yellow("~*~ Welcome To Corp Bot ~*~\n")

	// Build info
	printHeader("BUILD INFO")
	printOperation(fmt.Sprintf("Running Go version: %s", runtime.Version()))

	mg.Deps(Bin.Compile)
}

// Simple utility functions
func printHeader(text string) {
	color.Cyan("\n~~ %s ~~\n", text)
}

func printOperation(text string) {
	color.Magenta("ⓘ %s\n", text)
}

// Commands involving the compilation of the project
type Bin mg.Namespace

func (Bin) Compile() {
	mg.SerialDeps(Bin.Modules, Test.All, Bin.Out)
	color.Green("\n✓ All Done\n\n")
}

func (Bin) Tidy() error {
	printOperation("Cleaning modules")
	err := sh.RunV("go", "mod", "tidy", "-v")
	return err
}

func (Bin) Vendor() error {
	printOperation("Creating vendor files")
	err := sh.RunV("go", "mod", "vendor")
	return err
}

func (Bin) Modules() {
	printHeader("MODULES")
	mg.SerialDeps(Bin.Tidy, Bin.Vendor)
}

func (Bin) Out() error {
	printHeader("CREATE BIN")
	printOperation("Running go build")
	err := sh.RunV(
		"go", "build",
		"-o", "./bin/serve",
		"./cmd/serve/serve.go",
	)
	return err
}

// Commands involving running automated tests for the project
type Test mg.Namespace

// Runs gotestsum without -short flag
func (Test) All() error {
	printHeader("PKG TESTS (ALL)")
	err := sh.RunV("gotestsum", "./pkg/...")
	return err
}

// Runs gotestsum with -short flag
func (Test) Unit() error {
	printHeader("PKG TESTS (UNIT)")
	err := sh.RunV("gotestsum", "./pkg/...", "--", "-short")
	return err
}

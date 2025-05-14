// A generated module for Fyne functions
//
// This module has been generated via dagger init and serves as a reference to
// basic module structure as you get started with Dagger.
//
// Two functions have been pre-created. You can modify, delete, or add to them,
// as needed. They demonstrate usage of arguments and return types using simple
// echo and grep commands. The functions can be called from the dagger CLI or
// from one of the SDKs.
//
// The first line in this comment block is a short description line and the
// rest is a long description with more detail on the module's purpose or usage,
// if appropriate. All modules should have a short description.

package main

import "dagger/fyne/internal/dagger"

type Fyne struct{}

func (f *Fyne) WithFyne(ctr *dagger.Container) *dagger.Container {
	return ctr.
		// fyne deps
		WithExec([]string{"apt-get", "install", "-y", "gcc", "libgl1-mesa-dev", "xorg-dev", "libxkbcommon-dev"}).
		// TODO: use the new tools repository for v2.6+
		WithExec([]string{"git", "clone", "https://github.com/dolanor/fyne", "/src/fyne"}).
		WithWorkdir("/src/fyne").
		WithExec([]string{"go", "install", "./cmd/fyne"})
}

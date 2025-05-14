// A generated module for Go functions
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

import (
	"dagger/go/internal/dagger"
	"fmt"
)

type Go struct{}

type GoVersion string

const (
// GoVersion1_24_3 GoVersion = "1.24.3"
)

func goImageName(version string) string {
	return fmt.Sprintf("golang:%s", version)
}

func (g *Go) GoImage(version string) *dagger.Container {
	goCache := dag.CacheVolume("gobuildcache")
	goModCache := dag.CacheVolume("gomodcache")

	imageName := goImageName(version)

	return dag.Container().
		From(imageName).
		WithMountedCache("/root/.cache/go-build", goCache).
		WithMountedCache("/go/pkg/mod", goModCache)
}

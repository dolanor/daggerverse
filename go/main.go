// Package go allows to build and test Go source code.
package main

import (
	"context"
	"dagger/go/internal/dagger"
	"fmt"
	"path/filepath"
	"strings"
)

type Go struct {
	ctr *dagger.Container
}

// GoVersion constraints which Go version can be used.
type GoVersion string

const (
	GoVersion1_25_6 GoVersion = "v1_25_6"
	GoVersion1_24_3 GoVersion = "v1_24_3"
	GoVersion1_23_8 GoVersion = "v1_23_8"
)

func goImageName(version GoVersion) string {
	// we remove the initial v
	version = version[1:]
	// we replace the _ by . so it will actually match
	// a docker image version instead of pleasing graphql limited enum
	// format.
	v := strings.Replace(string(version), "_", ".", -1)
	return fmt.Sprintf("golang:%s", v)
}

// Container gives a Go container based on the docker image "golang"
// with the given version as the tag.
func (g *Go) Container(
	// version of the "golang" image (eg, "v1_24_3")
	version GoVersion,
) *dagger.Container {
	if g.ctr != nil {
		return g.ctr
	}

	goCache := dag.CacheVolume("gobuildcache")
	goModCache := dag.CacheVolume("gomodcache")

	imageName := goImageName(version)

	g.ctr = dag.Container().
		From(imageName).
		WithMountedCache("/root/.cache/go-build", goCache).
		WithMountedCache("/go/pkg/mod", goModCache)

	return g.ctr
}

// Test tests the container
func (g *Go) Test(
	ctx context.Context,

	source *dagger.Directory,
) bool {
	dirName, err := source.Name(ctx)
	if err != nil {
		return false
	}

	testPath := filepath.Join("/usr/src", dirName)

	exitCode, err := g.ctr.
		WithDirectory(testPath, source).
		WithWorkdir(testPath).
		WithExec([]string{"go", "test"}).
		ExitCode(ctx)
	if err != nil {
		return false
	}

	if exitCode != 0 {
		return false
	}
	return true
}

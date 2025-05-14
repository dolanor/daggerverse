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

func (f *Fyne) BuildAPK(
	// source is the root of source we're using to build the app.
	source *dagger.Directory,

	// App ID is the identifier that is use for an app. In general, it is based on
	// reverse name notation (eg. com.fynelabs.nomad).
	appID string,

	// Target Platform represents which os/arch we are building for.
	// (eg. android/arm64, android/arm)
	targetPlatform string,

	// app main dir is where fyne will look for the base of the app to build into an
	// APK, relative to the source directory.
	appMainDir string,

	// APK Path represents the path where fyne has built it.
	// normally, it should be in the source dir, but the name of the apk could vary.
	// We need to set this so we know where we need to look to be able to export the file.
	apkPath string,
) *dagger.File {
	apk := dag.Go().Container("1.24.3").
		With(dag.Android().WithAndroid).
		With(f.WithFyne).
		WithDirectory("/src", source).
		WithWorkdir("/src").
		//Terminal().
		WithExec([]string{
			"fyne",
			"package",
			"--appID",
			appID,
			"--target",
			targetPlatform,
			"--sourceDir",
			appMainDir,
		}).
		File(apkPath)

	return apk
}

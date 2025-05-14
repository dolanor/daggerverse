// A generated module for Android functions
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
	"dagger/android/internal/dagger"
	"path/filepath"
)

type Android struct{}

func (a *Android) WithAndroid(ctr *dagger.Container) *dagger.Container {
	aptArchives := dag.CacheVolume("aptarchives")

	androidMount := "/mnt"
	androidRoot := "/android"
	androidSDKHome := filepath.Join(androidMount, androidRoot, "sdk")
	androidNDKHome := filepath.Join(androidMount, androidRoot, "ndk")

	return ctr.
		WithMountedCache("/var/cache/apt/archives", aptArchives).

		// android sdk deps
		WithExec([]string{"dpkg", "--add-architecture", "i386"}).
		WithExec([]string{"apt-get", "update"}).
		WithExec([]string{"apt-get", "install", "-y", "libc6:i386", "libncurses5:i386", "libstdc++6:i386", "lib32z1", "libbz2-1.0:i386"}).
		WithWorkdir("/tmp/").
		// WithExec([]string{"curl", "-O", "-L", "https://dl.google.com/android/repository/android-ndk-r27c-linux.zip"}).
		// WithExec([]string{"apt-get", "install", "-y", "unzip"}).
		// WithExec([]string{"unzip", "android-ndk-r27c-linux.zip"}).
		// WithExec([]string{"mv", "android-ndk-r27c", "/android/ndk"}).

		WithDirectory(androidMount, androidSDK()).
		//WithExec([]string{"cp", "-r", androidMount, androidRoot}).
		WithEnvVariable("ANDROID_HOME", androidSDKHome).
		WithEnvVariable("ANDROID_SDK_ROOT", androidSDKHome).
		WithEnvVariable("ANDROID_NDK_HOME", androidNDKHome)
	// WithExec([]string{"cp", "-r", "/mnt/android-sdk", "/android/sdk"}).
	// https://dl.google.com/android/repository/platform-tools-latest-linux.zip
}

func androidSDK() *dagger.Directory {
	return dag.Container().
		From("reg.txg.re/android-sdk:27").
		Rootfs()
	// Using Directory here instead of Rootfs() will bust the cache, somehow
	// related: https://github.com/dagger/dagger/issues/3705
	//Directory("/android")
}

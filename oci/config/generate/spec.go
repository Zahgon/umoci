// SPDX-License-Identifier: Apache-2.0
/*
 * umoci: Umoci Modifies Open Containers' Images
 * Copyright (C) 2016-2025 SUSE LLC
 * Copyright (C) 2026 Aleksa Sarai <cyphar@cyphar.com>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Package generate provides an API for modifying the OCI image configuration
// object in a slightly less manual way than constructing structs and doing nil
// checks manually. It is analogous to the runtime-tools generate package, and
// is a properly working version of the image-tools generate package.
package generate

import (
	"time"

	"github.com/opencontainers/go-digest"
	ispec "github.com/opencontainers/image-spec/specs-go/v1"
)

// FIXME: Because we are not a part of upstream, we have to add some tests that
//        ensure that this set of getters and setters is complete. This should
//        be possible through some reflection.

// Generator allows you to generate a mutable OCI image-spec configuration
// which can be written to a file (and its digest computed). It is the
// recommended way of handling modification and generation of image-spec
// configuration blobs.
type Generator struct {
	image ispec.Image
}

// init makes sure everything has a "proper" zero value.
func (g *Generator) init() {
	if g.image.Config.ExposedPorts == nil {
		g.ClearConfigExposedPorts()
	}
	if g.image.Config.Env == nil {
		g.ClearConfigEnv()
	}
	if g.image.Config.Entrypoint == nil {
		g.ClearConfigEntrypoint()
	}
	if g.image.Config.Cmd == nil {
		g.ClearConfigCmd()
	}
	if g.image.Config.Volumes == nil {
		g.ClearConfigVolumes()
	}
	if g.image.Config.Labels == nil {
		g.ClearConfigLabels()
	}
	if g.image.RootFS.DiffIDs == nil {
		g.ClearRootfsDiffIDs()
	}
	if g.image.History == nil {
		g.ClearHistory()
	}
}

// New creates a new Generator with the initial template set to a default. It is
// not recommended to leave any of the options as their default values (they
// may change in the future without warning and may be invalid images).
func New() *Generator {
	_ = "STUB: not implemented"
	// FIXME: Come up with some sane default.
	return nil
}

// NewFromImage generates a new generator with the initial template being the
// given ispec.Image.
func NewFromImage(image ispec.Image) (*Generator, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Image returns a copy of the current state of the generated image.
func (g *Generator) Image() ispec.Image {
	_ = "STUB: not implemented"

	// SetConfigUser sets the username or UID which the process in the container should run as.
	return *new(ispec.Image)
}

func (g *Generator) SetConfigUser(user string) { _ = "STUB: not implemented"; return }

// ConfigUser returns the username or UID which the process in the container should run as.
func (g *Generator) ConfigUser() string { _ = "STUB: not implemented"; return "" }

// ClearConfigExposedPorts clears the set of ports to expose from a container running this image.
func (g *Generator) ClearConfigExposedPorts() { _ = "STUB: not implemented"; return }

// AddConfigExposedPort adds a port the set of ports to expose from a container running this image.
func (g *Generator) AddConfigExposedPort(port string) { _ = "STUB: not implemented"; return }

// RemoveConfigExposedPort removes a port the set of ports to expose from a container running this image.
func (g *Generator) RemoveConfigExposedPort(port string) { _ = "STUB: not implemented"; return }

// ConfigExposedPorts returns the set of ports to expose from a container running this image.
func (g *Generator) ConfigExposedPorts() []string { _ = "STUB: not implemented"; return nil }

// ClearConfigEnv clears the list of environment variables to be used in a container.
func (g *Generator) ClearConfigEnv() { _ = "STUB: not implemented"; return }

// AddConfigEnv appends to the list of environment variables to be used in a container.
func (g *Generator) AddConfigEnv(name, value string) {
	_ = "STUB: not implemented"
	// If the key already exists in the environment set, we replace it.
	// This ensures we don't run into POSIX undefined territory.
	return
}

// ConfigEnv returns the list of environment variables to be used in a container.
func (g *Generator) ConfigEnv() []string {
	_ = "STUB: not implemented"
	// We have to make a copy to preserve the privacy of g.image.Config.
	return nil
}

// ClearConfigEntrypoint clears the list of arguments to use as the command to execute when the container starts.
func (g *Generator) ClearConfigEntrypoint() { _ = "STUB: not implemented"; return }

// SetConfigEntrypoint sets the list of arguments to use as the command to execute when the container starts.
func (g *Generator) SetConfigEntrypoint(entrypoint []string) { _ = "STUB: not implemented"; return }

// ConfigEntrypoint returns the list of arguments to use as the command to execute when the container starts.
func (g *Generator) ConfigEntrypoint() []string {
	_ = "STUB: not implemented"
	// We have to make a copy to preserve the privacy of g.image.Config.
	return nil
}

// ClearConfigCmd clears the list of default arguments to the entrypoint of the container.
func (g *Generator) ClearConfigCmd() { _ = "STUB: not implemented"; return }

// SetConfigCmd sets the list of default arguments to the entrypoint of the container.
func (g *Generator) SetConfigCmd(cmd []string) { _ = "STUB: not implemented"; return }

// ConfigCmd returns the list of default arguments to the entrypoint of the container.
func (g *Generator) ConfigCmd() []string {
	_ = "STUB: not implemented"
	// We have to make a copy to preserve the privacy of g.image.Config.
	return nil
}

// ClearConfigVolumes clears the set of directories which should be created as data volumes in a container running this image.
func (g *Generator) ClearConfigVolumes() { _ = "STUB: not implemented"; return }

// AddConfigVolume adds a volume to the set of directories which should be created as data volumes in a container running this image.
func (g *Generator) AddConfigVolume(volume string) { _ = "STUB: not implemented"; return }

// RemoveConfigVolume removes a volume from the set of directories which should be created as data volumes in a container running this image.
func (g *Generator) RemoveConfigVolume(volume string) { _ = "STUB: not implemented"; return }

// ConfigVolumes returns the set of directories which should be created as data volumes in a container running this image.
func (g *Generator) ConfigVolumes() []string { _ = "STUB: not implemented"; return nil }

// ClearConfigLabels clears the set of arbitrary metadata for the container.
func (g *Generator) ClearConfigLabels() { _ = "STUB: not implemented"; return }

// AddConfigLabel adds a label to the set of arbitrary metadata for the container.
func (g *Generator) AddConfigLabel(label, value string) { _ = "STUB: not implemented"; return }

// RemoveConfigLabel removes a label from the set of arbitrary metadata for the container.
func (g *Generator) RemoveConfigLabel(label string) { _ = "STUB: not implemented"; return }

// ConfigLabels returns the set of arbitrary metadata for the container.
func (g *Generator) ConfigLabels() map[string]string {
	_ = "STUB: not implemented"
	// We have to make a copy to preserve the privacy of g.image.Config.
	return nil
}

// SetConfigWorkingDir sets the current working directory of the entrypoint process in the container.
func (g *Generator) SetConfigWorkingDir(workingDir string) { _ = "STUB: not implemented"; return }

// ConfigWorkingDir returns the current working directory of the entrypoint process in the container.
func (g *Generator) ConfigWorkingDir() string { _ = "STUB: not implemented"; return "" }

// SetConfigStopSignal sets the system call signal that will be sent to the container to exit.
func (g *Generator) SetConfigStopSignal(stopSignal string) { _ = "STUB: not implemented"; return }

// ConfigStopSignal returns the system call signal that will be sent to the container to exit.
func (g *Generator) ConfigStopSignal() string { _ = "STUB: not implemented"; return "" }

// SetRootfsType sets the type of the rootfs.
func (g *Generator) SetRootfsType(rootfsType string) { _ = "STUB: not implemented"; return }

// RootfsType returns the type of the rootfs.
func (g *Generator) RootfsType() string { _ = "STUB: not implemented"; return "" }

// ClearRootfsDiffIDs clears the array of layer content hashes (DiffIDs), in order from bottom-most to top-most.
func (g *Generator) ClearRootfsDiffIDs() { _ = "STUB: not implemented"; return }

// AddRootfsDiffID appends to the array of layer content hashes (DiffIDs), in order from bottom-most to top-most.
func (g *Generator) AddRootfsDiffID(diffid digest.Digest) { _ = "STUB: not implemented"; return }

// RootfsDiffIDs returns the the array of layer content hashes (DiffIDs), in order from bottom-most to top-most.
func (g *Generator) RootfsDiffIDs() []digest.Digest {
	_ = "STUB: not implemented"
	// We have to make a copy to preserve the privacy of g.image.RootFS.
	return nil
}

// ClearHistory clears the history of each layer.
func (g *Generator) ClearHistory() { _ = "STUB: not implemented"; return }

// AddHistory appends to the history of the layers.
func (g *Generator) AddHistory(history ispec.History) { _ = "STUB: not implemented"; return }

// History returns the history of each layer.
func (g *Generator) History() []ispec.History {
	_ = "STUB: not implemented"
	// We have to make a copy to preserve the privacy of g.image.History.
	return nil
}

// ISO8601 represents the format of an ISO-8601 time string, which is identical
// to Go's RFC3339 specification.
const ISO8601 = time.RFC3339Nano

// SetCreated sets the combined date and time at which the image was created.
func (g *Generator) SetCreated(created time.Time) { _ = "STUB: not implemented"; return }

// Created gets the combined date and time at which the image was created.
func (g *Generator) Created() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// TODO: Maybe we should be returning pointers?

// SetAuthor sets the name and/or email address of the person or entity which created and is responsible for maintaining the image.
func (g *Generator) SetAuthor(author string) { _ = "STUB: not implemented"; return }

// Author returns the name and/or email address of the person or entity which created and is responsible for maintaining the image.
func (g *Generator) Author() string { _ = "STUB: not implemented"; return "" }

// SetPlatformOS sets the name of the operating system which the image is built to run on.
func (g *Generator) SetPlatformOS(os string) {
	_ = "STUB: not implemented"

	// PlatformOS returns the name of the operating system which the image is built to run on.
	return
}

func (g *Generator) PlatformOS() string {
	_ = "STUB: not implemented"

	// SetPlatformArchitecture is the CPU architecture which the binaries in this image are built to run on.
	return ""
}

func (g *Generator) SetPlatformArchitecture(arch string) { _ = "STUB: not implemented"; return }

// PlatformArchitecture returns the CPU architecture which the binaries in this image are built to run on.
func (g *Generator) PlatformArchitecture() string { _ = "STUB: not implemented"; return "" }

// SetPlatformVariant is the CPU architecture variant which the binaries in this image are built to run on.
func (g *Generator) SetPlatformVariant(variant string) { _ = "STUB: not implemented"; return }

// PlatformVariant returns the CPU architecture variant which the binaries in this image are built to run on.
func (g *Generator) PlatformVariant() string { _ = "STUB: not implemented"; return "" }

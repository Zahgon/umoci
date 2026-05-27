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

package convert

import (
	ispec "github.com/opencontainers/image-spec/specs-go/v1"
	rspec "github.com/opencontainers/runtime-spec/specs-go"
)

// Annotations described by the OCI image-spec document (these represent fields
// in an image configuration that do not have a native representation in the
// runtime-spec).
const (
	platformOsAnnotation      = "org.opencontainers.image.os"
	platformArchAnnotation    = "org.opencontainers.image.architecture"
	platformVariantAnnotation = "org.opencontainers.image.variant"
	authorAnnotation          = "org.opencontainers.image.author"
	createdAnnotation         = "org.opencontainers.image.created"
	stopSignalAnnotation      = "org.opencontainers.image.stopSignal"
	exposedPortsAnnotation    = "org.opencontainers.image.exposedPorts"
)

// ToRuntimeSpec converts the given OCI image configuration to a runtime
// configuration appropriate for use, which is templated on the default
// configuration specified by the OCI runtime-tools. It is equivalent to
// MutateRuntimeSpec("runtime-tools/generate".New(), image).Spec().
func ToRuntimeSpec(rootfs string, image ispec.Image) (rspec.Spec, error) {
	_ = "STUB: not implemented"
	return *new(rspec.Spec), nil
}

// parseEnv splits a given environment variable (of the form name=value) into
// (name, value). An error is returned if there is no "=" in the line or if the
// name is empty.
func parseEnv(env string) (string, string, error) { _ = "STUB: not implemented"; return "", "", nil }

// appendEnv takes a (name, value) pair and inserts it into the given
// environment list (overwriting an existing environment if already set).
func appendEnv(env *[]string, name, value string, clobber bool) { _ = "STUB: not implemented"; return }

// allocateNilStruct recursively enumerates all pointers in the given type and
// replaces them with the zero-value of their associated type. It's a shame
// that this is necessary.
//
// TODO: Switch to doing this recursively with reflect.
func allocateNilStruct(spec *rspec.Spec) { _ = "STUB: not implemented"; return }

// MutateRuntimeSpec mutates a given runtime configuration with the image
// configuration provided in accordance with the image specification's
// conversion mechanism (for more information, see
// <https://github.com/opencontainers/image-spec/blob/main/conversion.md>).
func MutateRuntimeSpec(spec *rspec.Spec, rootfs string, image ispec.Image) error {
	_ = "STUB: not implemented"
	return nil
}

// Default config to our rspec version if none was specified.

// Make sure that the previous version of the spec is compatible with us.
// We cannot operate on specifications that are newer than us (because we
// might drop fields that the user finds important).

// Set verbatim fields

// Set the "annotation fields".

// Config.Labels need to be applied after the auto-applied labels.

// Set parsed fields
// Get the *actual* uid and gid of the user. If the image doesn't contain
// an /etc/passwd or /etc/group file then GetExecUserPath will just do a
// numerical parsing.

// We only log an error if were not given a rootfs, and we set execUser
// to the "default" (root:root).

// XXX: This is _fine_ but might cause some issues in the future.

// Remove all seccomp rules.

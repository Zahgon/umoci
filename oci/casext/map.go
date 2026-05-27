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

package casext

import (
	"reflect"

	ispec "github.com/opencontainers/image-spec/specs-go/v1"
)

// Used by walkState.mark() to determine which struct members are descriptors to
// recurse into them. We aren't interested in struct members which are not
// either a slice of ispec.Descriptor or ispec.Descriptor themselves.
var descriptorType = reflect.TypeFor[ispec.Descriptor]()

// DescriptorMapFunc is a function that is used to provide a mapping between
// different descriptor values with MapDescriptors. It will not be called
// concurrently, and will only be called once for each recursively resolved
// element.
type DescriptorMapFunc func(ispec.Descriptor) ispec.Descriptor

// isDescriptor returns whether the given T is a ispec.Descriptor.
func isDescriptor(T reflect.Type) bool { _ = "STUB: not implemented"; return false }

func mapDescriptors(V reflect.Value, mapFunc DescriptorMapFunc) error {
	_ = "STUB: not implemented"
	// We can ignore this value.
	return nil
}

// First check that V isn't actually a ispec.Descriptor, if it is then
// we're done.

//nolint:forcetypeassert // already checked with reflection in isDescriptor

// We only need to do any assignment if the two are not equal.

// P is a ptr to V (or just V if it's already a pointer).

// This is a programmer error.

// Recurse into all the types.
//nolint:exhaustive // no need to handle other types explicitly

// Just deref the pointer/interface.

// Iterate over each element.

// We are only ever going to be interested in registered types.

// We can now actually iterate through a struct to find all descriptors.

// FIXME: Should we log something here? While this will be hit normally
//        (namely when we hit an io.ReadCloser) this seems a bit
//        careless.

// Unreachable.

// MapDescriptors applies the given function once for every instance of
// ispec.Descriptor found in the given type, and replaces it with the returned
// value (which may be the same). This is done through the reflection API in
// Go, which means that hidden attributes may be inaccessible.
// DescriptorMapFunc will only be executed once for every ispec.Descriptor
// found.
func MapDescriptors(i any, mapFunc DescriptorMapFunc) error { _ = "STUB: not implemented"; return nil }

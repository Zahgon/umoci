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
	"context"
	"errors"

	"github.com/opencontainers/go-digest"
	ispec "github.com/opencontainers/image-spec/specs-go/v1"
)

// childDescriptors is a wrapper around MapDescriptors which just creates a
// slice of all of the arguments, and doesn't modify them.
func childDescriptors(i any) []ispec.Descriptor { _ = "STUB: not implemented"; return nil }

// If we got an error, this is a bug in MapDescriptors proper.

// walkState stores state information about the recursion into a given
// descriptor tree.
type walkState struct {
	// engine is the CAS engine we are operating on.
	engine Engine

	// walkFunc is the WalkFunc provided by the user.
	walkFunc WalkFunc
}

// DescriptorPath is used to describe the path of descriptors (from a top-level
// index) that were traversed when resolving a particular reference name. The
// purpose of this is to allow libraries like github.com/opencontainers/umoci/mutate
// to handle generic manifest updates given an arbitrary descriptor walk. Users
// of ResolveReference that don't care about the descriptor path can just use
// .Descriptor.
type DescriptorPath struct {
	// Walk is the set of descriptors walked to reach Descriptor (inclusive).
	// The order is the same as the order of the walk, with the target being
	// the last entry and the entrypoint from index.json being the first.
	Walk []ispec.Descriptor `json:"descriptor_walk"`
}

// Root returns the first step in the DescriptorPath, which is the point where
// the walk started. This is just shorthand for DescriptorPath.Walk[0]. Root
// will *panic* if DescriptorPath is invalid.
func (d DescriptorPath) Root() ispec.Descriptor {
	_ = "STUB: not implemented"
	return *new(ispec.Descriptor)
}

// Descriptor returns the final step in the DescriptorPath, which is the target
// descriptor being referenced by DescriptorPath. This is just shorthand for
// accessing the last entry of DescriptorPath.Walk. Descriptor will *panic* if
// DescriptorPath is invalid.
func (d DescriptorPath) Descriptor() ispec.Descriptor {
	_ = "STUB: not implemented"
	return *new(ispec.Descriptor)
}

// ErrSkipDescriptor is a special error returned by WalkFunc which will cause
// Walk to not recurse into the descriptor currently being evaluated by
// WalkFunc. This interface is roughly equivalent to filepath.SkipDir.
var ErrSkipDescriptor = errors.New("[internal] do not recurse into descriptor")

// WalkFunc is the type of function passed to Walk. It will be a called on each
// descriptor encountered, recursively -- which may involve the function being
// called on the same descriptor multiple times (though because an OCI image is
// a Merkle tree there will never be any loops). If an error is returned by
// WalkFunc, the recursion will halt and the error will bubble up to the
// caller.
//
// TODO: Also provide Blob to WalkFunc so that callers don't need to load blobs
//
//	more than once. This is quite important for remote CAS implementations.
type WalkFunc func(descriptorPath DescriptorPath) error

func (ws *walkState) recurse(ctx context.Context, descriptorPath DescriptorPath) (Err error) {
	_ = "STUB: not implemented"
	return nil
}

// Run walkFunc.

// Get blob to recurse into.

// Since FromDescriptor gives us a full VerifiedReadCloser (meaning that
// Close is expensive if we don't read any bytes), we should only try to
// recurse into this thing if we actually can parse it.

// Recurse into the blob now.

// Ignore cases where the descriptor points to an object we don't know
// how to parse.

// Recurse into children.

// Walk preforms a depth-first walk from a given root descriptor, using the
// provided CAS engine to fetch all other necessary descriptors. If an error is
// returned by the provided WalkFunc, walking is terminated and the error is
// returned to the caller.
func (e Engine) Walk(ctx context.Context, root ispec.Descriptor, walkFunc WalkFunc) error {
	_ = "STUB: not implemented"
	return nil
}

// reachable returns the set of digests which can be reached using a descriptor
// path from the provided root descriptor. The returned slice will *not*
// contain any duplicate digest.Digest entries.
//
// Please note that without descriptors, a digest is not particularly meaninful
// (OCI blobs are not self-descriptive). This method primarily exists for GC()
// and any use outside of GC() should be carefully considered (you probably
// want to use Walk directly).
func (e Engine) reachable(ctx context.Context, root ispec.Descriptor) ([]digest.Digest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Don't traverse further if we've already seen this digest.

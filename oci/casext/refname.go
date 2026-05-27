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
	"regexp"

	ispec "github.com/opencontainers/image-spec/specs-go/v1"
)

// refnameRegex is a regex that only matches reference names that are valid
// according to the OCI specification. See IsValidReferenceName for the EBNF.
var refnameRegex = regexp.MustCompile(`^([A-Za-z0-9]+(([-._:@+]|--)[A-Za-z0-9]+)*)(/([A-Za-z0-9]+(([-._:@+]|--)[A-Za-z0-9]+)*))*$`)

// IsValidReferenceName returns whether the provided annotation value for
// "org.opencontainers.image.ref.name" is actually valid according to the
// OCI specification. This only matches against the MUST requirement, not the
// SHOULD requirement. The EBNF defined in the specification is:
//
//	refname   ::= component ("/" component)*
//	component ::= alphanum (separator alphanum)*
//	alphanum  ::= [A-Za-z0-9]+
//	separator ::= [-._:@+] | "--"
func IsValidReferenceName(refname string) bool { _ = "STUB: not implemented"; return false }

// ResolveReference will attempt to resolve all possible descriptor paths to
// Manifests (or any unknown blobs) that match a particular reference name (if
// descriptors are stored in non-standard blobs, Resolve will be unable to find
// them but will return the top-most unknown descriptor).
// ResolveReference assumes that "reference name" refers to the value of the
// "org.opencontainers.image.ref.name" descriptor annotation. It is recommended
// that if the returned slice of descriptors is greater than zero that the user
// be consulted to resolve the conflict (due to ambiguity in resolution paths).
//
// TODO: How are we meant to implement other restrictions such as the
//
//	architecture and feature flags? The API will need to change.
func (e Engine) ResolveReference(ctx context.Context, refname string) ([]DescriptorPath, error) {
	_ = "STUB: not implemented"
	// XXX: It should be possible to override this somehow, in case we are
	//      dealing with an image that abuses the image specification in some
	//      way.
	return nil, nil
}

// Set of root links that match the given refname.

// We only consider the case where AnnotationRefName is defined on the
// top-level of the index tree. While this isn't codified in the spec (at
// the time of writing -- 1.0.0-rc5) there are some discussions to add this
// restriction in 1.0.0-rc6.

// XXX: What should we do if refname == "".

// The resolved set of descriptors.

// Find all manifests or other blobs that are reachable from the given
// descriptor.

// If the media-type should be treated as a "target media-type" for
// reference resolution, we stop resolution here and add it to the
// set of resolved paths.

// XXX: Should the *Reference set of interfaces support DescriptorPath? While
//      it might seem like it doesn't make sense, a DescriptorPath entirely
//      removes ambiguity with regards to which root needs to be operated on.
//      If a user has that information we should provide them a way to use it.

// UpdateReference replaces an existing entry for refname with the given
// descriptor. If there are multiple descriptors that match the refname they
// are all replaced with the given descriptor.
func (e Engine) UpdateReference(ctx context.Context, refname string, descriptor ispec.Descriptor) error {
	_ = "STUB: not implemented"
	// XXX: It should be possible to override this somehow, in case we are
	//      dealing with an image that abuses the image specification in some
	//      way.
	return nil
}

// Get index to modify.

// TODO: Handle refname = "".

// Warn users if the operation is going to remove more than one references.

// Append the descriptor.

// Commit to image.

// DeleteReference removes all entries in the index that match the given
// refname.
func (e Engine) DeleteReference(ctx context.Context, refname string) error {
	_ = "STUB: not implemented"
	// XXX: It should be possible to override this somehow, in case we are
	//      dealing with an image that abuses the image specification in some
	//      way.
	return nil
}

// Get index to modify.

// TODO: Handle refname = "".

// Warn users if the operation is going to remove more than one references.

// Commit to image.

// ListReferences returns all of the ref.name entries that are specified in the
// top-level index. Note that the list may contain duplicates, due to the
// nature of references in the image-spec.
func (e Engine) ListReferences(ctx context.Context) ([]string, error) {
	_ = "STUB: not implemented"
	// Get index.
	return nil, nil
}

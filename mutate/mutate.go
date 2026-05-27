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

// Package mutate implements various functionality to allow for the
// modification of container images in a much higher-level fashion than
// available from github.com/opencontainers/umoci/oci/cas. In particular, this library
// should be viewed as a wrapper around github.com/opencontainers/umoci/oci/cas that
// provides many convenience functions.
package mutate

import (
	"context"
	"io"
	"time"

	"github.com/opencontainers/go-digest"
	ispec "github.com/opencontainers/image-spec/specs-go/v1"

	"github.com/opencontainers/umoci/oci/cas"
	"github.com/opencontainers/umoci/oci/casext"
)

// UmociUncompressedBlobSizeAnnotation is an umoci-specific annotation to
// provide information in descriptors to compressed blobs about the size of the
// underlying uncompressed blob for users that need that information. Note that
// this annotation value should be treated as a hint -- an attacker could
// create an image that has a dummy UmociUncompressedBlobSizeAnnotation value
// for a zip-bomb blob.
const UmociUncompressedBlobSizeAnnotation = "ci.umo.uncompressed_blob_size"

func configPtr(c ispec.Image) *ispec.Image         { _ = "STUB: not implemented"; return nil }
func manifestPtr(m ispec.Manifest) *ispec.Manifest { _ = "STUB: not implemented"; return nil }
func timePtr(t time.Time) *time.Time               { _ = "STUB: not implemented"; return nil }

// XXX: Currently this package is very entangled in modifying of a given
//      Manifest and their associated Config + Layers. While this works fine,
//      really mutate/ should be a far more generic library that allows you to
//      apply a delta for a particular OCI structure and then regenerate the
//      necessary blobs. Something like changing annotations for intermediate
//      manifests is not really possible at the moment, and it's not clear how
//      to make it possible without forcing users to interactively make all the
//      necessary changes.

// Mutator is a wrapper around a cas.Engine instance, and is used to mutate a
// given image (described by a manifest) in a high-level fashion. It handles
// creating all necessary blobs and modfying other blobs. In order for changes
// to be committed you must call .Commit().
//
// TODO: Implement manifest list support.
type Mutator struct {
	// These are the arguments we got in New().
	engine casext.Engine
	source casext.DescriptorPath

	// Cached values of the configuration and manifest.
	manifest *ispec.Manifest
	config   *ispec.Image
}

// Meta is a wrapper around the "safe" fields in ispec.Image, which can be
// modified by users and have no effect on a Mutator or the validity of an
// image.
type Meta struct {
	// Created defines an ISO-8601 formatted combined date and time at which
	// the image was created.
	Created time.Time `json:"created,omitzero"`

	// Author defines the name and/or email address of the person or entity
	// which created and is responsible for maintaining the image.
	Author string `json:"author,omitzero"`

	// Architecture is the CPU architecture which the binaries in this image
	// are built to run on.
	Architecture string `json:"architecture"`

	// Variant is the variant of the CPU architecture which the binaries in
	// this image are built to run on.
	Variant string `json:"variant"`

	// OS is the name of the operating system which the image is built to run
	// on.
	OS string `json:"os"`

	// TODO: Should we embed ispec.Platform?
}

// cache ensures that the cached versions of the related configurations have
// been loaded. Calling this function more than once will do nothing, unless
// you've explicitly cleared the cache.
func (m *Mutator) cache(ctx context.Context) (Err error) {
	_ = "STUB: not implemented"
	// We need the manifest
	return nil
}

// Should _never_ be reached.

// Make a copy of the manifest.

// Should _never_ be reached.

// Make a copy of the config and configDescriptor.

// New creates a new Mutator for the given descriptor (which _must_ have a
// MediaType of ispec.MediaTypeImageManifest.
func New(engine cas.Engine, src casext.DescriptorPath) (*Mutator, error) {
	_ = "STUB: not implemented"
	// We currently only support changing a given manifest through a walk.
	return nil, nil
}

// Config returns the current (cached) image configuration, which should be
// used as the source for any modifications of the configuration using
// Set.
func (m *Mutator) Config(ctx context.Context) (ispec.Image, error) {
	_ = "STUB: not implemented"
	return *new(ispec.Image), nil
}

// Manifest returns the current (cached) image manifest. This is what will be
// appended to when any additional Add() calls are made, and what will be
// Commit()ed if no further changes are made.
func (m *Mutator) Manifest(ctx context.Context) (ispec.Manifest, error) {
	_ = "STUB: not implemented"
	return *new(ispec.Manifest), nil
}

// Meta returns the current (cached) image metadata, which should be used as
// the source for any modifications of the configuration using Set.
func (m *Mutator) Meta(ctx context.Context) (Meta, error) {
	_ = "STUB: not implemented"
	return *new(Meta), nil
}

// Annotations returns the set of annotations in the current manifest. This
// does not include the annotations set in ispec.ImageConfig.Labels. This
// should be used as the source for any modifications of the annotations using
// Set.
func (m *Mutator) Annotations(ctx context.Context) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Set sets the image configuration and metadata to the given values. The
// provided ispec.History entry is appended to the image's history and should
// correspond to what operations were made to the configuration.
func (m *Mutator) Set(ctx context.Context, config ispec.ImageConfig, meta Meta, annotations map[string]string, history *ispec.History) error {
	_ = "STUB: not implemented"
	return nil
}

// Ensure the mediatype is correct.

// Set annotations.

// Set configuration.

// Set metadata.

// Append history.

func (m *Mutator) appendToConfig(history *ispec.History, layerDiffID digest.Digest) {
	_ = "STUB: not implemented"
	return
}

// Append history.

// Some tools get confused if there are layers with no history entry.
// Especially if you have later layers have history entries (which will
// result in the history entries not matching up and everyone getting
// quite confused).

// PickDefaultCompressAlgorithm returns the best option for the compression
// algorithm for new layers. The main preference is to use re-use whatever the
// most recent layer's compression algorithm is (for those we support). As a
// final fallback, we use blobcompress.Default.
func (m *Mutator) PickDefaultCompressAlgorithm(ctx context.Context) (Compressor, error) {
	_ = "STUB: not implemented"
	return *new(Compressor), nil
}

// Don't generate an uncompressed layer even if the previous one is --
// there is no reason to automatically generate uncompressed blobs.

// No supported, non-plain algorithm found. Just use the default.

// Add adds a layer to the image, by reading the layer changeset blob from the
// provided reader. The stream must not be compressed, as it is used to
// generate the DiffIDs for the image metatadata. The provided history entry is
// appended to the image's history and should correspond to what operations
// were made to the configuration.
func (m *Mutator) Add(ctx context.Context, mediaType string, r io.Reader, history *ispec.History, compressor Compressor, annotations map[string]string) (_ ispec.Descriptor, Err error) {
	_ = "STUB: not implemented"
	return *new(ispec.Descriptor), nil
}

// Add DiffID to configuration.

// Build the descriptor.

// Append to layers.

// AddExisting adds a blob that already exists to the layer, using the user
// specified DiffID. It currently checks that the layer exists, but does not
// validate the DiffID.
func (m *Mutator) AddExisting(ctx context.Context, desc ispec.Descriptor, history *ispec.History, diffID digest.Digest) error {
	_ = "STUB: not implemented"
	return nil
}

// Commit writes all of the temporary changes made to the configuration,
// metadata and manifest to the engine. It then returns a new manifest
// descriptor (which can be used in place of the source descriptor provided to
// New).
func (m *Mutator) Commit(ctx context.Context) (_ casext.DescriptorPath, Err error) {
	_ = "STUB: not implemented"
	return *new(casext.DescriptorPath), nil
}

// We first have to commit the configuration blob.

// Now commit the manifest.

// We now have to create a new DescriptorPath that replaces the one we were
// given. Note that we have to walk *up* the path rather than down it
// because we have to replace each blob in order to replace its references.

// Replace the end of the path.

// Walk up the path, mutating the parent reference of each descriptor.

// Get the blob of the parent.

// Replace all references to the child blob with the new one.

// In principle you should never be in a situation where two
// descriptors reference the same data with different
// media-types. This lead to CVE-2021-41190.

// Replace the digest+size with the new blob.

// Copy the embedded data for the new descriptor (if any).

// Do not touch any other bits in case the same blob is being
// referenced with different annotations or platform
// configurations.

// Re-commit the blob.
// TODO: This won't handle foreign blobs correctly, we need to make it
//       possible to write a modified blob through the blob API.

// Update the key parts of the descriptor.

// Clear the embedded data (if present).
// TODO: Auto-embed data if it is reasonably small.

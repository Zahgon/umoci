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

	ispec "github.com/opencontainers/image-spec/specs-go/v1"
)

// Blob represents a "parsed" blob in an OCI image's blob store. MediaType
// offers a type-safe way of checking what the type of Data is.
type Blob struct {
	// Descriptor is the {mediatype,digest,length} 3-tuple. Note that this
	// isn't updated if the Data is modified.
	Descriptor ispec.Descriptor

	// RawData is the un-parsed blob data taken from the OCI image's blob
	// store. If the data is unparseable, this value will be nil.
	RawData []byte

	// Data is the "parsed" blob taken from the OCI image's blob store, and is
	// typed according to the media type. The default mappings from MIME =>
	// type is as follows (more can be registered using RegisterParser).
	//
	// ispec.MediaTypeDescriptor => ispec.Descriptor
	// ispec.MediaTypeImageManifest => ispec.Manifest
	// ispec.MediaTypeImageIndex => ispec.Index
	// ispec.MediaTypeImageLayer => io.ReadCloser
	// ispec.MediaTypeImageLayerGzip => io.ReadCloser
	// ispec.MediaTypeImageLayerNonDistributable => io.ReadCloser
	// ispec.MediaTypeImageLayerNonDistributableGzip => io.ReadCloser
	// ispec.MediaTypeImageConfig => ispec.Image
	// ispec.MediaTypeEmptyJSON => struct{}
	// unknown => io.ReadCloser
	Data any
}

// Close cleans up all of the resources for the opened blob.
func (b *Blob) Close() error { _ = "STUB: not implemented"; return nil }

// FromDescriptor parses the blob referenced by the given descriptor.
func (e Engine) FromDescriptor(ctx context.Context, descriptor ispec.Descriptor) (blob *Blob, Err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// copy if present

// TODO: Should we short-cut this for descriptors with embedded data?

// Include all trailing data.

// Sanity check.

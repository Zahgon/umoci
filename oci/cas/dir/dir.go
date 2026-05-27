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

// Package dir implements the basic local directory-backed layout backend for
// the OCI content-addressible store.
package dir

import (
	"context"
	"io"
	"os"

	"github.com/opencontainers/go-digest"
	ispec "github.com/opencontainers/image-spec/specs-go/v1"

	"github.com/opencontainers/umoci/oci/cas"
)

const (
	// ImageLayoutVersion is the version of the image layout we support. This
	// value is *not* the same as imagespec.Version, and the meaning of this
	// field is still under discussion in the spec. For now we'll just hardcode
	// the value and hope for the best.
	ImageLayoutVersion = ispec.ImageLayoutVersion // "1.0.0"

	// blobDirectory is the directory inside an OCI image that contains blobs.
	blobDirectory = ispec.ImageBlobsDir // "blobs"

	// indexFile is the file inside an OCI image that contains the top-level
	// index.
	indexFile = ispec.ImageIndexFile // "index.json"

	// layoutFile is the file in side an OCI image the indicates what version
	// of the OCI spec the image is.
	layoutFile = ispec.ImageLayoutFile // "oci-layout"
)

// blobPath returns the path to a blob given its digest, relative to the root
// of the OCI image. The digest must be of the form algorithm:hex.
func blobPath(digest digest.Digest) (string, error) { _ = "STUB: not implemented"; return "", nil }

type uidgid struct {
	uid, gid int
}

type dirEngine struct {
	path     string
	temp     string
	tempFile *os.File
	owner    *uidgid
}

func (e *dirEngine) ensureTempDir() error { _ = "STUB: not implemented"; return nil }

// We get an advisory lock to ensure that GC() won't delete our
// temporary directory here. Once we get the lock we know it won't do
// anything until we unlock it or exit.

// chown changes the ownership of the provided file to match the owner of the
// cas directory itself (in order to avoid a root "umoci repack" creating
// inaccessible files for the original user).
func (e *dirEngine) fchown(file *os.File) error { _ = "STUB: not implemented"; return nil }

// skip chown if not running as root

// verify ensures that the image is valid.
func (e *dirEngine) validate() error { _ = "STUB: not implemented"; return nil }

// XXX: Currently the meaning of this field is not adequately defined by
//      the spec, nor is the "official" value determined by the spec.

// Check that "blobs" and "index.json" exist in the image.
// FIXME: We also should check that blobs *only* contains a cas.BlobAlgorithm
//        directory (with no subdirectories) and that refs *only* contains
//        files (optionally also making sure they're all JSON descriptors).

// PutBlob adds a new blob to the image. This is idempotent; a nil error
// means that "the content is stored at DIGEST" without implying "because
// of this PutBlob() call".
func (e *dirEngine) PutBlob(_ context.Context, reader io.Reader) (_ digest.Digest, _ int64, Err error) {
	_ = "STUB: not implemented"
	return *new(digest.Digest), 0, nil
}

// We copy this into a temporary file because we need to get the blob hash,
// but also to avoid half-writing an invalid blob.

// Get the digest.

// Move the blob to its correct path.

// GetBlob returns a reader for retrieving a blob from the image, which the
// caller must Close(). Returns ErrNotExist if the digest is not found.
//
// This function will return a VerifiedReadCloser, meaning that you must call
// Close() and check the error returned from Close() in order to ensure that
// the hash of the blob is verified.
//
// Please note that calling Close() on the returned blob will read the entire
// from disk and hash it (even if you didn't read any bytes before calling
// Close), so if you wish to only check if a blob exists you should use
// StatBlob() instead.
func (e *dirEngine) GetBlob(_ context.Context, digest digest.Digest) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

// Assume the file size is the blob size. This is almost certainly true
// in general, and if an attacker is modifying the blobs underneath us
// then snapshotting the size makes sure we don't read endlessly.

// StatBlob returns whether the specified blob exists in the image. Returns
// false if the blob doesn't exist, true if it does, or an error if any error
// occurred.
//
// NOTE: In future this API may change to return additional information.
func (e *dirEngine) StatBlob(_ context.Context, digest digest.Digest) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// PutIndex sets the index of the OCI image to the given index, replacing the
// previously existing index. This operation is atomic; any readers attempting
// to access the OCI image while it is being modified will only ever see the
// new or old index.
func (e *dirEngine) PutIndex(_ context.Context, index ispec.Index) (Err error) {
	_ = "STUB: not implemented"
	return nil
}

// Make sure the index has the mediatype field set.

// We copy this into a temporary index to ensure the atomicity of this
// operation.

// Encode the index.

// Move the blob to its correct path.

// GetIndex returns the index of the OCI image. Return ErrNotExist if the
// digest is not found. If the image doesn't have an index, ErrInvalid is
// returned (a valid OCI image MUST have an image index).
//
// It is not recommended that users of cas.Engine use this interface directly,
// due to the complication of properly handling references as well as correctly
// handling nested indexes. casext.Engine provides a wrapper for cas.Engine
// that implements various reference resolution functions that should work for
// most users.
func (e *dirEngine) GetIndex(_ context.Context) (ispec.Index, error) {
	_ = "STUB: not implemented"
	return *new(ispec.Index), nil
}

// DeleteBlob removes a blob from the image. This is idempotent; a nil
// error means "the content is not in the store" without implying "because
// of this DeleteBlob() call".
func (e *dirEngine) DeleteBlob(_ context.Context, digest digest.Digest) error {
	_ = "STUB: not implemented"
	return nil
}

// ListBlobs returns the set of blob digests stored in the image.
func (e *dirEngine) ListBlobs(_ context.Context) ([]digest.Digest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Skip the actual directory.

// XXX: Do we need to handle multiple-directory-deep cases?

// Clean executes a garbage collection of any non-blob garbage in the store
// (this includes temporary files and directories not reachable from the CAS
// interface). This MUST NOT remove any blobs or references in the store.
func (e *dirEngine) Clean(ctx context.Context) error {
	_ = "STUB: not implemented"
	// Remove every .umoci directory that isn't flocked.
	return nil
}

func (e *dirEngine) cleanPath(_ context.Context, path string) (Err error) {
	_ = "STUB: not implemented"
	return nil
}

// If we fail to get a flock(2) then it's probably already locked,
// so we shouldn't touch it.

// somebody else beat us to it

// Close releases all references held by the e. Subsequent operations may
// fail.
func (e *dirEngine) Close() error { _ = "STUB: not implemented"; return nil }

// Open opens a new reference to the directory-backed OCI image referenced by
// the provided path.
func Open(path string) (cas.Engine, error) { _ = "STUB: not implemented"; return *new(cas.Engine), nil }

// Create creates a new OCI image layout at the given path. If the path already
// exists, os.ErrExist is returned. However, all of the parent components of
// the path will be created if necessary.
func Create(path string) (Err error) {
	_ = "STUB: not implemented"
	// We need to fail if path already exists, but we first create all of the
	// parent paths.
	return nil
}

// Create the necessary directories and "oci-layout" file.

// FIXME: This is hardcoded at the moment.

// Everything is now set up.

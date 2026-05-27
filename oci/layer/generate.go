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

package layer

import (
	"io"

	"github.com/vbatts/go-mtree"
)

// inodeDeltas is a wrapper around []mtree.InodeDelta that allows for sorting
// the set of deltas by the pathname.
type inodeDeltas []mtree.InodeDelta

func (ids inodeDeltas) Len() int           { _ = "STUB: not implemented"; return 0 }
func (ids inodeDeltas) Less(i, j int) bool { _ = "STUB: not implemented"; return false }
func (ids inodeDeltas) Swap(i, j int)      { _ = "STUB: not implemented"; return }

// GenerateLayer creates a new OCI diff layer based on the mtree diff provided.
// All of the mtree.Modified and mtree.Extra blobs are read relative to the
// provided path (which should be the rootfs of the layer that was diffed). The
// returned reader is for the *raw* tar data, it is the caller's responsibility
// to gzip it.
func GenerateLayer(path string, deltas []mtree.InodeDelta, opt *RepackOptions) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

// Close with the returned error.

// We can't just dump all of the file contents into a tar file. We need
// to emulate a proper tar generator. Luckily there aren't that many
// things to emulate (and we can do them all in tar.go).

// Sort the delta paths.
// FIXME: We need to add whiteouts first, otherwise we might end up
//        doing something silly like deleting a file which we actually
//        meant to modify.

// XXX: It's possible that if we unlink a hardlink, we're going to
//      AddFile() for no reason. Maybe we should drop nlink= from
//      the set of keywords we care about?

// For opaque whiteout directories we need to
// output an entry for the directory itself so that
// the ownership and modes set on the directory are
// included in the archive.

// We should never see these delta types because they are not
// generated for regular mtree.Compare.

//nolint:errcheck // errors are handled in defer func

// GenerateInsertLayer generates a completely new layer from root to be
// inserted into the image at target. If root is an empty string then the
// target will be removed via a whiteout. If opaque is true then the target
// directory will also have an opaque whiteout applied (clearing any files
// inside the directory), followed by the contents of the root.
func GenerateInsertLayer(root, target string, opaque bool, opt *RepackOptions) io.ReadCloser {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser)
}

// Continue on to add the new root contents...

// For opaque whiteout directories we need to
// output an entry for the directory itself so that
// the ownership and modes set on the directory are
// included in the archive.

//nolint:errcheck // errors are handled in defer func

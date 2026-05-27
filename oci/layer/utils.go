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
	"archive/tar"
)

// mapHeader maps a tar.Header generated from the filesystem so that it
// describes the inode as it would be observed by a container process. In
// particular this involves apply an ID mapping from the host filesystem to the
// container mappings. Returns an error if it's not possible to map the given
// UID.
func mapHeader(hdr *tar.Header, mapOptions MapOptions) error { _ = "STUB: not implemented"; return nil }

// It only makes sense to do un-mapping if we're not rootless. If we're
// rootless then all of the files will be owned by us anyway.

// We have special handling for the "user.rootlesscontainers" xattr. If
// we're rootless then we override the owner of the file we're currently
// parsing (and then remove the xattr). If we're not rootless then the user
// is doing something strange, so we log a warning but just ignore the
// xattr otherwise.
//
// TODO: We should probably add a flag to opt-out of this (though I'm not
//       sure why anyone would intentionally use this incorrectly).
//nolint:staticcheck,revive // SA1019: Xattrs is deprecated but PAXRecords is more annoying
// noop

// If the payload isn't uint32(-1) we apply it. The xattr includes the
// *in-container* owner so we don't want to map it.

// Drop the xattr since it's just a marker for us and shouldn't be in
// layers. This is technically out-of-spec, but so is
// "user.rootlesscontainers".
//nolint:staticcheck // SA1019: Xattrs is deprecated but PAXRecords is more annoying

// unmapHeader maps a tar.Header from a tar layer stream so that it describes
// the inode as it would be exist on the host filesystem. In particular this
// involves applying an ID mapping from the container filesystem to the host
// mappings. Returns an error if it's not possible to map the given UID.
func unmapHeader(hdr *tar.Header, mapOptions MapOptions) error {
	_ = "STUB: not implemented"
	// To avoid nil references.
	//nolint:staticcheck // SA1019: Xattrs is deprecated but PAXRecords is more annoying
	return nil
}

// If there is already a "user.rootlesscontainers" we give a warning in
// both rootless and root cases -- but in rootless we explicitly delete the
// entry because we might replace it.
//nolint:staticcheck // SA1019: Xattrs is deprecated but PAXRecords is more annoying

//nolint:staticcheck // SA1019: Xattrs is deprecated but PAXRecords is more annoying

// In rootless mode there are a few things we need to do. We need to map
// all of the files in the layer to have an owner of (0, 0) because we
// cannot lchown(2) anything -- and then if the owner was non-root we have
// to create a "user.rootlesscontainers" xattr for it.

// Fill the rootlesscontainers payload with the original (uid, gid). If
// either is 0, we replace it with uint32(-1). Technically we could
// just leave it as 0 (since that is what the source of truth told us
// the owner was), but this would result in a massive increase in
// xattrs with no real benefit.

// Don't add the xattr if the owner isn't just (0, 0) because that's a
// waste of space.

// While the payload is almost certainly not UTF-8, Go strings can
// actually be arbitrary bytes (in case you didn't know this and
// were confused like me when this worked). See
// <https://blog.golang.org/strings> for more detail.
//nolint:staticcheck // SA1019: Xattrs is deprecated but PAXRecords is more annoying

// CleanPath makes a path safe for use with filepath.Join. This is done by not
// only cleaning the path, but also (if the path is relative) adding a leading
// '/' and cleaning it (then removing the leading '/'). This ensures that a
// path resulting from prepending another path will always resolve to lexically
// be a subdirectory of the prefixed path. This is all done lexically, so paths
// that include symlinks won't be safe as a result of using CleanPath.
//
// This function comes from runC (libcontainer/utils/utils.go).
func CleanPath(path string) string {
	_ = "STUB: not implemented"
	// Deal with empty strings nicely.
	return ""
}

// Ensure that all paths are cleaned (especially problematic ones like
// "/../../../../../" which can cause lots of issues).

// If the path isn't absolute, we need to do more processing to fix paths
// such as "../../../../<etc>/some/path". We also shouldn't convert absolute
// paths to relative ones.

// This can't fail, as (by definition) all paths are relative to root.

// Clean the path again for good measure.

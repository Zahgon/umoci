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
	"errors"
	"io"
	"os"

	"github.com/moby/sys/userns"

	"github.com/opencontainers/umoci/internal/pathtrie"
	"github.com/opencontainers/umoci/pkg/fseval"
)

// inUserNamespace is a cached return value of userns.RunningInUserNS(). We
// compute this once globally rather than for each unpack. It won't change (we
// would hope) after we check it the first time.
var inUserNamespace = userns.RunningInUserNS()

// TarExtractor represents a tar file to be extracted.
type TarExtractor struct {
	// onDiskFormat indicates what kind of rootfs this TarExtractor is going to
	// extract into. [OverlayfsRootfs] will cause whiteouts to be extracted as
	// overlayfs-style whiteouts and some xattrs will be modified. See
	// [OnDiskFormat] for more information.
	onDiskFmt OnDiskFormat

	// partialRootless indicates whether "partial rootless" tricks should be
	// applied in our extraction. Rootless and userns execution have some
	// similar tricks necessary, but not all rootless tricks should be applied
	// when running in a userns -- hence the term "partial rootless" tricks.
	partialRootless bool

	// fsEval is an fseval.FsEval used for extraction.
	fsEval fseval.FsEval

	// upperPaths are paths that have either been extracted in the execution of
	// this TarExtractor or are ancestors of paths extracted. The purpose of
	// having this stored in-memory is to be able to handle opaque whiteouts as
	// well as some other possible ordering issues with malformed archives (the
	// downside of this approach is that it takes up memory -- we could switch
	// to a trie if necessary). These paths are relative to the tar root but
	// are fully symlink-expanded so no need to worry about that line noise.
	upperPaths map[string]struct{}

	// upperWhiteouts is a trie that represents the subset of upperPaths that
	// are whiteout files. This is needed for overlayfs translation because
	// opaque whiteout directories in overlayfs do not play well with regular
	// whiteouts.
	//
	// This is stored separately to upperPaths because this is only needed for
	// overlayfs mode (a niche usecase), and it is far more efficient for us to
	// only walk through whiteout entries (as it's the only thing that matters
	// for upperWhiteouts) as there should be very few of them in most images.
	upperWhiteouts *pathtrie.PathTrie[overlayWhiteoutType]

	// enotsupWarned is a flag set when we encounter the first ENOTSUP error
	// dealing with xattrs. This is used to ensure extraction to a destination
	// file system that does not support xattrs raises a single warning, rather
	// than a warning for every file, which can amount to 1000s of messages that
	// scroll a terminal, and may obscure other more important warnings.
	enotsupWarned bool

	// keepDirlinks is the corresponding flag from the UnpackOptions
	// supplied when this TarExtractor was constructed.
	keepDirlinks bool
}

// NewTarExtractor creates a new TarExtractor.
func NewTarExtractor(opt *UnpackOptions) *TarExtractor { _ = "STUB: not implemented"; return nil }

// We only need the whiteout trie for overlayfs extraction.

// restoreMetadata applies the state described in tar.Header to the filesystem
// at the given path. No sanity checking is done of the tar.Header's pathname
// or other information. In addition, no mapping is done of the header.
func (te *TarExtractor) restoreMetadata(path string, hdr *tar.Header) error {
	_ = "STUB: not implemented"
	// Some of the tar.Header fields don't match the OS API.
	return nil
}

// Get the _actual_ file info to figure out if the path is a symlink.

// Apply the owner. If we are rootless then "user.rootlesscontainers" has
// already been set up by unmapHeader, so nothing to do here.

// NOTE: This is not done through fsEval.

// We cannot apply hdr.Mode to symlinks, because symlinks don't have a mode
// of their own (they're special in that way). We have to apply this after
// we've applied the owner because setuid bits are cleared when changing
// owner (in rootless we don't care because we're always the owner).

// Apply access and modified time. Note that some archives won't fill the
// atime and mtime fields, so we have to set them to a more sane value.
// Otherwise Linux will start screaming at us, and nobody wants that.

// XXX: Should we instead default to atime if it's non-zero?

// Default to the mtime.

// Apply xattrs. In order to make sure that we *only* have the xattr set we
// want, we first clear the set of xattrs from the file then apply the ones
// set in the tar.Header.

//nolint:staticcheck // SA1019: Xattrs is deprecated but PAXRecords is more annoying

// Some xattrs need to be skipped for sanity reasons, such as
// security.selinux, because they are very much host-specific and
// extracting them from layers would be a really bad idea. Also, other
// xattrs may need to be remapped (such as {user,trusted}.overlay.*
// xattrs when in overlayfs mode) to have correct values.

// Avoid outputting a warning if a must-skip xattr already has
// the expected value we wanted.
//
// TODO: Maybe we should still emit a warning even in this case
//       because now that the directory has its xattrs cleared
//       with ToTar, we should probably warn if images have
//       xattrs that only happen to be applied correctly now.

// In rootless mode, some xattrs will fail (security.capability).
// This is _fine_ as long as we're not running as root (in which
// case we shouldn't be ignoring xattrs that we were told to set).
//
// TODO: We should translate all security.capability capabilities
//       into v3 capabilities, which allow us to write them as
//       unprivileged users (we also would need to translate them
//       back when creating archives).

// We cannot do much if we get an ENOTSUP -- this usually means
// that extended attributes are simply unsupported by the
// underlying filesystem (such as AUFS or NFS).

// applyMetadata applies the state described in tar.Header to the filesystem at
// the given path, using the state of the TarExtractor to remap information
// within the header. This should only be used with headers from a tar layer
// (not from the filesystem). No sanity checking is done of the tar.Header's
// pathname or other information.
func (te *TarExtractor) applyMetadata(path string, hdr *tar.Header) error {
	_ = "STUB: not implemented"
	// Modify the header.
	return nil
}

// Restore it on the filesystme.

// isDirlink returns whether the given path is a link to a directory (or a
// dirlink in rsync(1) parlance) which is used by --keep-dirlink to see whether
// we should extract through the link or clobber the link with a directory (in
// the case where we see a directory to extract and a symlink already exists
// there).
func (te *TarExtractor) isDirlink(root, path string) (bool, error) {
	_ = "STUB: not implemented"
	// Make sure it exists and is a symlink.
	return false, nil
}

// Technically a string.TrimPrefix would also work...

// It should be noted that SecureJoin will evaluate all symlinks in the
// path, so we don't need to loop over it or anything like that. It'll just
// be done for us (in UnpackEntry only the dirname(3) is evaluated but here
// we evaluate the whole thing).

// We hit a symlink loop -- which is fine but that means that this
// cannot be considered a dirlink.

// ENOENT or similar just means that it's a broken symlink, which
// means we have to overwrite it (but it's an allowed case).

func (te *TarExtractor) ociWhiteout(_ DirRootfs, root, dir, file string) error {
	_ = "STUB: not implemented"
	return nil

	// We have to be quite careful here. While the most intuitive way of
	// handling whiteouts would be to just RemoveAll without prejudice, We
	// have to be careful here. If there is a whiteout entry for a file
	// *after* a normal entry (in the same layer) then the whiteout must
	// not remove the new entry. We handle this by keeping track of
	// whichpaths have been touched by this layer's extraction (these form
	// the "upperdir"). We also have to handle cases where a directory has
	// been marked for deletion, but a child has been extracted in this
	// layer.
}

// If the root doesn't exist we've got nothing to do.
// XXX: We currently cannot error out if a layer asks us to remove a
//      non-existent path with this implementation (because we don't
//      know if it was implicitly removed by another whiteout). In
//      future we could add lowerPaths that would help track whether
//      another whiteout caused the removal to "fail" or if the path
//      was actually missing -- which would allow us to actually error
//      out here if the layer is invalid).

// Need to use securejoin.IsNotExist to handle ENOTDIR.

// Walk over the path to remove it. We remove a given path as soon as
// it isn't present in upperPaths (which includes ancestors of paths
// we've extracted so we only need to look up the one path). Otherwise
// we iterate over any children and try again. The only difference
// between opaque whiteouts and regular whiteouts is that we don't
// delete the directory itself with opaque whiteouts.

// If we are passed an error, bail unless it's ENOENT.

// If something was deleted outside of our knowledge it's not
// the end of the world. In principle this shouldn't happen
// though, so we log it for posterity.

// Get the relative form of subpath to root to match
// te.upperPaths.

// Remove the path only if it hasn't been touched.

// Opaque whiteouts don't remove the directory itself, so skip
// the top-level directory.

// Purge the path. We skip anything underneath (if it's a
// directory) since we just purged it -- and we don't want to
// hit ENOENT during iteration for no good reason.

func (te *TarExtractor) overlayfsWhiteout(onDiskFmt OverlayfsRootfs, root, dir, file string) error {
	_ = "STUB: not implemented"
	// Unlike standard dir whiteouts, we need to ensure that the path we are
	// whiting out exists, because this layer is applied to lower layers where
	// the target path might exist. As with UnpackEntry, we expect the tar
	// archive itself to contain information about the directory (and since we
	// are extracting overlayfs we can't really be sure of the underlying
	// directory's ownership and modes either).
	//
	// TODO: Same TODO as UnpackEntry regarding consistency.
	return nil
}

// In the case of opaque whiteouts we need to make sure the target path
// is definitely a directory, so if it's a non-directory clear it.

// For opaque whiteouts, we just need to set the overlayfs xattr for
// directory. Any files already there were added in this layer (since
// OverlayfsRootfs is used to generate each layer in separate
// directories) and so shouldn't be removed anyway.

// Overlayfs has strange behaviour if we extract a whiteout into an
// opaque directory (namely readdir will report the whiteouts as being
// there while all other syscalls will fail to operate on them). The
// solution is to delete any pre-existing whiteouts we have when we hit
// an opaque whiteout, and the creation of any subsequent regular
// whiteouts inside an opaque whiteout should be skipped.

// Skip any opaque whiteouts, because stacking them is fine and we
// cannot just RemoveAll them anyway (we would need to clear the
// xattr).

// Clear the subpath.

// For regular whiteouts, just remove any pre-existing inode and
// replace it with a whiteout inode.

// If the path is inside an opaque whiteout, don't bother creating a
// whiteout inode. We still need to RemoveAll first to make sure
// anything there gets removed though.

// mkdirAll is like te.fsEval.MkdirAll except it handles cases of (arguably
// invalid) tar archives where a path component of the target path is a
// non-directory and so standard os.MkdirAll would error out with ENOTDIR. In
// such cases the problematic component will be removed and replaced with
// MkdirAll of the remaining components.
func (te *TarExtractor) mkdirAll(root, subpath string, mode os.FileMode) error {
	_ = "STUB: not implemented"
	// Fast path -- just try MkdirAll.
	return nil
}

// Convert the path to a in-root path.

// Look for the first parent component that exists and can be resolved
// (which is presumably whatever is giving us the ENOTDIR).

// TODO: Should we check to see if it is actually not a directory?

// Clear the problematic parent component and retry MkdirAll.

// If we are in overlayfs mode, then it is possible that what happened is
// that the offending non-directory parent component was a regular
// whiteout, and then a later entry (in the same layer) added a path
// underneath the deleted directory. The correct behaviour in this case is
// to replace the whiteout with an opaque directory whiteout (this matches
// the upstream overlayfs behaviour).

// TODO: If there is an orphaned subpath in upperWhiteouts that we have
//       deleted now, what should we do? Is that even possible?

// Make the directory an opaque whiteout as if there were a real
// opaque tar entry for consistency.

var errInvalidWhiteout = errors.New("invalid whiteout")

// UnpackEntry extracts the given tar.Header to the provided root, ensuring
// that the layer state is consistent with the layer state that produced the
// tar archive being iterated over. This does handle whiteouts, so a tar.Header
// that represents a whiteout will result in the path being removed.
func (te *TarExtractor) UnpackEntry(root string, hdr *tar.Header, r io.Reader) (Err error) {
	_ = "STUB: not implemented"
	// Make the paths safe.
	return nil
}

// Get directory and filename, but we have to safely get the directory
// component of the path. SecureJoinVFS will evaluate the path itself,
// which we don't want (we're clever enough to handle the actual path being
// a symlink).

// If we got an entry for the root, then unsafeDir is the full path.

// If we're being asked to change the root type, bail because they may
// change it to a symlink which we could inadvertently follow.

// Before we do anything, get the state of dir. Because we might be adding
// or removing files, our parent directory might be modified in the
// process. As a result, we want to be able to restore the old state
// (because we only apply state that we find in the archive we're iterating
// over). We can safely ignore an error here, because a non-existent
// directory will be fixed by later archive entries.

// FIXME: This is really stupid.

// More faking to trick restoreMetadata to actually restore the directory.

// os.Lstat doesn't get the list of xattrs by default. We need to fill
// this explicitly. Note that while Go's "archive/tar" takes strings,
// in Go strings can be arbitrary byte sequences so this doesn't
// restrict the possible values.
// TODO: Move this to a separate function so we can share it with
//       tar_generate.go.

//nolint:staticcheck // SA1019: Xattrs is deprecated but PAXRecords is more annoying

// Because restoreMetadata will re-apply these xattrs
// (potentially remapping them if we have specialXattrs
// filters) we need to map their names to match what we would
// get from an actual archive.
//
// However, since these are xattrs on the underlying filesystem
// we don't need to provide any user warnings.

// If the xattr should be ignored we can safely skip it
// here because MaskedOnDisk will also stop them from
// being cleared. However, just to be safe we should
// verify that this is actually true (otherwise you'll
// end up with silently wrong extractions).

// TODO: Find a nicer setup that doesn't require
// this fatal error.

//nolint:staticcheck // SA1019: Xattrs is deprecated but PAXRecords is more annoying

// Ensure that after everything we correctly re-apply the old metadata.
// We don't map this header because we're restoring files that already
// existed on the filesystem, not from a tar layer.

// Currently the spec doesn't specify what the hdr.Typeflag of whiteout
// files is meant to be. We specifically only produce regular files
// ('\x00') but it could be possible that someone produces a different
// Typeflag, expecting that the path is the only thing that matters in a
// whiteout entry.

// special value to indicate opaque whiteout

// Get information about the path. This has to be done after we've dealt
// with whiteouts because it turns out that lstat(2) will return EPERM if
// you try to stat a whiteout on AUFS.

// File doesn't exist, just switch fi to the file header.

// Attempt to create the parent directory of the path we're unpacking.
// We do a MkdirAll here because even though you need to have a tar entry
// for every component of a new path, applyMetadata will correct any
// inconsistencies.
// FIXME: We have to make this consistent, since if the tar archive doesn't
//        have entries for some of these components we won't be able to
//        verify that we have consistent results during unpacking.

// We remove whatever existed at the old path to clobber it so that
// creating a new path will not break. The only exception is if the path is
// a directory in both the layer and the current filesystem, in which case
// we don't delete it for obvious reasons. In all other cases we clobber.
//
// Note that this will cause hard-links in the "lower" layer to not be able
// to point to "upper" layer inodes even if the extracted type is the same
// as the old one, however it is not clear whether this is something a user
// would expect anyway. In addition, this will incorrectly deal with a
// TarLink that is present before the "upper" entry in the layer but the
// "lower" file still exists (so the hard-link would point to the old
// inode). It's not clear if such an archive is actually valid though.

// If we are in --keep-dirlinks mode and the existing fs object is a
// symlink to a directory (with the pending object is a directory), we
// don't remove the symlink (and instead allow subsequent objects to be
// just written through the symlink into the directory). This is a very
// specific usecase where layers that were generated independently from
// each other (on different base filesystems) end up with weird things
// like /lib64 being a symlink only sometimes but you never want to
// delete libraries (not just the ones that were under the "real"
// directory).
//
// TODO: This code should also handle a pending symlink entry where the
//       existing object is a directory. I'm not sure how we could
//       disambiguate this from a symlink-to-a-file but I imagine that
//       this is something that would also be useful in the same vein
//       as --keep-dirlinks (which currently only prevents clobbering
//       in the opposite case).

//nolint:staticcheck // QF1001: this form is easier to understand

// Now create or otherwise modify the state of the path. Right now, either
// the type of path matches hdr or the path doesn't exist. Note that we
// don't care about umasks or the initial mode here, since applyMetadata
// will fix all of that for us.

// regular file
//nolint:staticcheck // SA1019: TypeRegA is deprecated but for compatibility we need to support it
// Create a new file, then just copy the data.

// We need to make sure that we copy all of the bytes.

// Force close here so that we don't affect the metadata.

// directory

// Attempt to create the directory. We do a MkdirAll here because even
// though you need to have a tar entry for every component of a new
// path, applyMetadata will correct any inconsistencies.

// hard link, symbolic link

// Hardlinks and symlinks act differently when it comes to the scoping.
// In both cases, we have to just unlink and then re-link the given
// path. But the function used and the argument are slightly different.

// Because hardlinks are inode-based we need to scope the link to
// the rootfs using SecureJoinVFS. As before, we need to be careful
// that we don't resolve the last part of the link path (in case
// the user actually wanted to hardlink to a symlink).

// Link the new one.

// FIXME: Currently this can break if tar hardlink entries occur
//        before we hit the entry those hardlinks link to. I have a
//        feeling that such archives are invalid, but the correct
//        way of handling this is to delay link creation until the
//        very end. Unfortunately this won't work with symlinks
//        (which can link to directories).

// character device node, block device node

// In rootless mode we have no choice but to fake this, since mknod(2)
// doesn't work as an unprivileged user here.
//
// TODO: We need to add the concept of a fake block device in
//       "user.rootlesscontainers", because this workaround suffers
//       from the obvious issue that if the file is touched (even the
//       metadata) then it will be incorrectly copied into the layer.
//       This would break distribution images fairly badly.

// Otherwise the handling is the same as a FIFO.

// fifo node

// We have to remove and then create the device. In the FIFO case we
// could choose not to do so, but we do it anyway just to be on the
// safe side.

// Create the node.

// We should never hit any other headers (Go abstracts them away from us),
// and we can't handle any custom Tar extensions. So just error out.

// Apply the metadata, which will apply any mappings necessary. We don't
// apply metadata for hardlinks, because hardlinks don't have any separate
// metadata from their link (and the tar headers might not be filled).

// Everything is done -- the path now exists. Add it (and all its
// ancestors) to the set of upper paths. We first have to figure out the
// proper path corresponding to hdr.Name though.

// Really shouldn't happen because of the guarantees of SecureJoinVFS.

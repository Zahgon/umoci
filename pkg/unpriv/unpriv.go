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

// Package unpriv provides rootless emulation of CAP_DAC_READ_SEARCH without
// the need for rootless user namespaces. This is necessary in general because
// it turns out that a lot of distributions have a rootfs with `chmod 000`
// directories that rely on root having CAP_DAC_READ_SEARCH to be normally
// accessible.
//
// Note that the implementation of CAP_DAC_READ_SEARCH requires write access to
// any normally-inaccessible components of paths.
//
// Users should use fseval.FsEval instead to allow programs to switch between
// fseval.Rootless and fseval.Default based on whether the program is
// privileged or not.
package unpriv

import (
	"os"
	"path/filepath"
	"time"

	"golang.org/x/sys/unix"
)

// fiRestore restores the state given by an os.FileInfo instance at the given
// path by ensuring that an Lstat(path) will return as-close-to the same
// os.FileInfo.
func fiRestore(path string, fi os.FileInfo) error {
	_ = "STUB: not implemented"
	// archive/tar handles the OS-specific syscall stuff required to get atime
	// and mtime information for a file.
	return nil
}

// Apply the relevant information from the FileInfo.

// splitpath splits the given path into each of the path components.
func splitpath(path string) []string { _ = "STUB: not implemented"; return nil }

// WrapFunc is a function that can be passed to Wrap. It takes a path (and
// presumably operates on it -- since Wrap only ensures that the path given is
// resolvable) and returns some form of error.
type WrapFunc func(path string) error

// Wrap will wrap a given function, and call it in a context where all of the
// parent directories in the given path argument are such that the path can be
// resolved (you may need to make your own changes to the path to make it
// readable). Note that the provided function may be called several times, and
// if the error returned is such that !os.IsPermission(err), then no trickery
// will be performed. If fn returns an error, so will this function. All of the
// trickery is reverted when this function returns (which is when fn returns).
func Wrap(path string, fn WrapFunc) (Err error) {
	_ = "STUB: not implemented"
	// FIXME: Should we be calling fn() here first?
	return nil
}

// We need to chown all of the path components we don't have execute rights
// to. Specifically these are the path components which are parents of path
// components we cannot stat. However, we must make sure to not touch the
// path itself.

// We've hit the first element we can chown.

// This is a legitimate error.

// Chown from the top down.

// Add +rwx permissions to directories. If we have the access to change
// the mode at all then we are the user owner (not just a group owner).

// Everything is wrapped. Return from this nightmare.

// Open is a wrapper around os.Open which has been wrapped with unpriv.Wrap to
// make it possible to open paths even if you do not currently have read
// permission. Note that the returned file handle references a path that you do
// not have read access to (since all changes are reverted when this function
// returns), so attempts to do Readdir() or similar functions that require
// doing lstat(2) may fail.
func Open(path string) (*os.File, error) { _ = "STUB: not implemented"; return nil, nil }

// Get information so we can revert it.

// Add +r permissions to the file.

// Open the damn thing.

// Create is a wrapper around os.Create which has been wrapped with unpriv.Wrap
// to make it possible to create paths even if you do not currently have read
// permission. Note that the returned file handle references a path that you do
// not have read access to (since all changes are reverted when this function
// returns).
func Create(path string) (*os.File, error) { _ = "STUB: not implemented"; return nil, nil }

// Readdir is a wrapper around (*os.File).Readdir which has been wrapper with
// unpriv.Wrap to make it possible to get []os.FileInfo for the set of children
// of the provided directory path. The interface for this is quite different to
// (*os.File).Readdir because we have to have a proper filesystem path in order
// to get the set of child FileInfos (because all of the child paths need to be
// resolveable).
func Readdir(path string) ([]os.FileInfo, error) { _ = "STUB: not implemented"; return nil, nil }

// Get information so we can revert it.

// Add +rx permissions to the file.

// Open the damn thing.

// Get the set of dirents.

// Lstat is a wrapper around os.Lstat which has been wrapped with unpriv.Wrap
// to make it possible to get os.FileInfo about a path even if you do not
// currently have the required mode bits set to resolve the path. Note that you
// may not have resolve access after this function returns because all of the
// trickery is reverted by unpriv.Wrap.
func Lstat(path string) (os.FileInfo, error) {
	_ = "STUB: not implemented"
	return *new(os.FileInfo), nil
}

// Fairly simple.

// Lstatx is like Lstat but uses unix.Lstat and returns unix.Stat_t instead.
func Lstatx(path string) (unix.Stat_t, error) {
	_ = "STUB: not implemented"
	return *new(unix.Stat_t), nil
}

// Readlink is a wrapper around os.Readlink which has been wrapped with
// unpriv.Wrap to make it possible to get the target of a symlink even if you
// do not currently have the required mode bits set to resolve the path. Note
// that you may not have resolve access after this function returns because all
// of this trickery is reverted by unpriv.Wrap.
func Readlink(path string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// Fairly simple.

// Symlink is a wrapper around os.Symlink which has been wrapped with
// unpriv.Wrap to make it possible to create a symlink even if you do not
// currently have the required access bits to create the symlink. Note that you
// may not have resolve access after this function returns because all of the
// trickery is reverted by unpriv.Wrap.
func Symlink(target, linkname string) error { _ = "STUB: not implemented"; return nil }

// Link is a wrapper around unix.Link(..., 0) which has been wrapped with
// unpriv.Wrap to make it possible to create a hard link even if you do not
// currently have the required access bits to create the hard link. Note that
// you may not have resolve access after this function returns because all of
// the trickery is reverted by unpriv.Wrap.
func Link(target, linkname string) error { _ = "STUB: not implemented"; return nil }

// We have to double-wrap this, because you need search access to the
// linkname. This is safe because any common ancestors will be reverted
// in reverse call stack order.

// We need to explicitly pass 0 as a flag because POSIX allows the
// default behaviour of link(2) when it comes to target being a
// symlink to be implementation-defined. Only linkat(2) allows us
// to guarantee the right behaviour.
//  <https://pubs.opengroup.org/onlinepubs/9699919799/functions/link.html>

// Chmod is a wrapper around os.Chmod which has been wrapped with unpriv.Wrap
// to make it possible to change the permission bits of a path even if you do
// not currently have the required access bits to access the path.
func Chmod(path string, mode os.FileMode) error { _ = "STUB: not implemented"; return nil }

// Chtimes is a wrapper around os.Chtimes which has been wrapped with
// unpriv.Wrap to make it possible to change the modified times of a path even
// if you do not currently have the required access bits to access the path.
func Chtimes(path string, atime, mtime time.Time) error { _ = "STUB: not implemented"; return nil }

// Lutimes is a wrapper around system.Lutimes which has been wrapped with
// unpriv.Wrap to make it possible to change the modified times of a path even
// if you do no currently have the required access bits to access the path.
func Lutimes(path string, atime, mtime time.Time) error { _ = "STUB: not implemented"; return nil }

// Remove is a wrapper around os.Remove which has been wrapped with unpriv.Wrap
// to make it possible to remove a path even if you do not currently have the
// required access bits to modify or resolve the path.
func Remove(path string) error { _ = "STUB: not implemented"; return nil }

// foreachSubpath executes WrapFunc for each child of the given path (not
// including the path itself). If path is not a directory, then WrapFunc will
// not be called and no error will be returned. This should be called within a
// context where path has already been made resolveable, however the . If WrapFunc returns an
// error, the first error is returned and iteration is halted.
func foreachSubpath(path string, wrapFn WrapFunc) (Err error) {
	_ = "STUB: not implemented"
	// Is the path a directory?
	return nil
}

// Open the directory.

// We need to change the mode to Readdirnames. We don't need to worry about
// permissions because we're already in a context with filepath.Dir(path)
// is at least a+rx. However, because we are calling wrapFn we need to
// restore the original mode immediately.

// Make iteration order consistent.

// Call on all the sub-directories. We run it in a Wrap context to ensure
// that the path we pass is resolveable when executed.

// RemoveAll is similar to os.RemoveAll but with all of the internal functions
// wrapped with unpriv.Wrap to make it possible to remove a path (even if it
// has child paths) even if you do not currently have enough access bits.
func RemoveAll(path string) error { _ = "STUB: not implemented"; return nil }

// If remove works, we're done.

// Is this a directory?

// Use securejoin's IsNotExist to handle ENOTDIR sanely.

// Return error from remove if it's not a directory.

// We must have hit a race, but we don't care.

// Remove the directory. This should now work.

// Mkdir is a wrapper around os.Mkdir which has been wrapped with unpriv.Wrap
// to make it possible to remove a path even if you do not currently have the
// required access bits to modify or resolve the path.
func Mkdir(path string, perm os.FileMode) error { _ = "STUB: not implemented"; return nil }

// MkdirAll is similar to os.MkdirAll but in order to implement it properly all
// of the internal functions were wrapped with unpriv.Wrap to make it possible
// to create a path even if you do not currently have enough access bits.
func MkdirAll(path string, perm os.FileMode) error { _ = "STUB: not implemented"; return nil }

// Check whether the path already exists.

// Create parent.

// Parent exists, now we can create the path.

// Handle "foo/.".

// Mknod is a wrapper around unix.Mknod which has been wrapped with unpriv.Wrap
// to make it possible to remove a path even if you do not currently have the
// required access bits to modify or resolve the path.
func Mknod(path string, mode os.FileMode, dev uint64) error { _ = "STUB: not implemented"; return nil }

// Llistxattr is a wrapper around system.Llistxattr which has been wrapped with
// unpriv.Wrap to make it possible to remove a path even if you do not
// currently have the required access bits to resolve the path.
func Llistxattr(path string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

// Lremovexattr is a wrapper around system.Lremovexattr which has been wrapped
// with unpriv.Wrap to make it possible to remove a path even if you do not
// currently have the required access bits to resolve the path.
func Lremovexattr(path, name string) error { _ = "STUB: not implemented"; return nil }

// Lsetxattr is a wrapper around system.Lsetxattr which has been wrapped
// with unpriv.Wrap to make it possible to set a path even if you do not
// currently have the required access bits to resolve the path.
func Lsetxattr(path, name string, value []byte, flags int) error {
	_ = "STUB: not implemented"
	return nil
}

// Lgetxattr is a wrapper around system.Lgetxattr which has been wrapped
// with unpriv.Wrap to make it possible to get a path even if you do not
// currently have the required access bits to resolve the path.
func Lgetxattr(path, name string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Lclearxattrs is a wrapper around system.Lclearxattrs which has been wrapped
// with unpriv.Wrap to make it possible to get a path even if you do not
// currently have the required access bits to resolve the path.
func Lclearxattrs(path string, skipFn func(xattrName string) bool) error {
	_ = "STUB: not implemented"
	return nil
}

// walk is the inner implementation of Walk.
func walk(path string, info os.FileInfo, walkFn filepath.WalkFunc) error {
	_ = "STUB: not implemented"
	// Always run walkFn first. If we're not a directory there's no children to
	// iterate over and so we bail even if there wasn't an error.
	return nil
}

// Now just execute walkFn over each subpath.
// TODO: We should handle the Readdirnames failing case that stdlib does.

// If it doesn't exist, just pass it directly to walkFn.

// To match stdlib, SkipDir assumes a non-existent path is a
// directory and so SkipDir just skips that path.

// If this entry is a directory then SkipDir will just skip
// this entry and continue walking the current directory, but
// otherwise we need to skip the whole directory. This matches
// the stdlib behaviour.
//nolint:staticcheck // QF1001: this form is easier to understand

// Walk is a reimplementation of filepath.Walk, wrapping all of the relevant
// function calls with Wrap, allowing you to walk over a tree even in the face
// of multiple nested cases where paths are not normally accessible. The
// os.FileInfo passed to walkFn is the "pristine" version (as opposed to the
// currently-on-disk version that may have been temporarily modified by Wrap).
func Walk(root string, walkFn filepath.WalkFunc) error { _ = "STUB: not implemented"; return nil }

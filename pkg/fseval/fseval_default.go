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

package fseval

import (
	"os"
	"path/filepath"
	"time"

	"github.com/vbatts/go-mtree"
	"golang.org/x/sys/unix"
)

// Default is the "identity" form of FsEval. In particular, it does not do any
// trickery and calls directly to the relevant os.* functions (and does not
// wrap KeywordFunc). This should be used by default, because there are no
// weird side-effects.
var Default FsEval = osFsEval(0)

// osFsEval is a hack to be able to make DefaultFsEval a const.
type osFsEval int

// Open is equivalent to os.Open.
func (fs osFsEval) Open(path string) (*os.File, error) {
	_ = "STUB: not implemented"
	return nil,

		// Create is equivalent to os.Create.
		nil
}

func (fs osFsEval) Create(path string) (*os.File, error) {
	_ = "STUB: not implemented"
	return nil,

		// Readdir is equivalent to os.Readdir.
		nil
}

func (fs osFsEval) Readdir(path string) (_ []os.FileInfo, Err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Lstat is equivalent to os.Lstat.
func (fs osFsEval) Lstat(path string) (os.FileInfo, error) {
	_ = "STUB: not implemented"
	return *

	// Lstatx is equivalent to unix.Lstat.
	new(os.FileInfo), nil
}

func (fs osFsEval) Lstatx(path string) (unix.Stat_t, error) {
	_ = "STUB: not implemented"
	return *new(unix.Stat_t), nil
}

// Readlink is equivalent to os.Readlink.
func (fs osFsEval) Readlink(path string) (string, error) {
	_ = "STUB: not implemented"
	return "",

		// Symlink is equivalent to os.Symlink.
		nil
}

func (fs osFsEval) Symlink(target, linkname string) error { _ = "STUB: not implemented"; return nil }

// Link is equivalent to unix.Link(..., ~AT_SYMLINK_FOLLOW).
func (fs osFsEval) Link(target, linkname string) error {
	_ = "STUB: not implemented"
	// We need to explicitly pass 0 as a flag because POSIX allows the default
	// behaviour of link(2) when it comes to target being a symlink to be
	// implementation-defined. Only linkat(2) allows us to guarantee the right
	// behaviour.
	//
	//	<https://pubs.opengroup.org/onlinepubs/9699919799/functions/link.html>
	return nil
}

// Chmod is equivalent to os.Chmod.
func (fs osFsEval) Chmod(path string, mode os.FileMode) error {
	_ = "STUB: not implemented"
	return nil
}

// Lutimes is equivalent to os.Lutimes.
func (fs osFsEval) Lutimes(path string, atime, mtime time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

// RemoveAll is equivalent to os.RemoveAll.
func (fs osFsEval) RemoveAll(path string) error { _ = "STUB: not implemented"; return nil }

// Mknod is equivalent to unix.Mknod.
func (fs osFsEval) Mknod(path string, mode os.FileMode, dev uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// MkdirAll is equivalent to os.MkdirAll.
func (fs osFsEval) MkdirAll(path string, perm os.FileMode) error {
	_ = "STUB: not implemented"
	return nil
}

// Llistxattr is equivalent to system.Llistxattr.
func (fs osFsEval) Llistxattr(path string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Lremovexattr is equivalent to system.Lremovexattr.
}

func (fs osFsEval) Lremovexattr(path, name string) error { _ = "STUB: not implemented"; return nil }

// Lsetxattr is equivalent to system.Lsetxattr.
func (fs osFsEval) Lsetxattr(path, name string, value []byte, flags int) error {
	_ = "STUB: not implemented"
	return nil
}

// Lgetxattr is equivalent to system.Lgetxattr.
func (fs osFsEval) Lgetxattr(path, name string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Lclearxattrs is equivalent to system.Lclearxattrs.
func (fs osFsEval) Lclearxattrs(path string, skipFn func(xattrName string) bool) error {
	_ = "STUB: not implemented"
	return nil
}

// KeywordFunc returns a wrapper around the given mtree.KeywordFunc.
func (fs osFsEval) KeywordFunc(fn mtree.KeywordFunc) mtree.KeywordFunc {
	_ = "STUB: not implemented"

	// Walk is equivalent to filepath.Walk.
	return *new(mtree.KeywordFunc)
}

func (fs osFsEval) Walk(root string, fn filepath.WalkFunc) error {
	_ = "STUB: not implemented"
	return nil
}

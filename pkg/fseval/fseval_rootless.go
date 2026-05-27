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

// Rootless is an FsEval implementation that uses "umoci/pkg/unpriv".*
// functions in order to provide the ability for unprivileged users (those
// without CAP_DAC_OVERRIDE and CAP_DAC_READ_SEARCH) to evaluate parts of a
// filesystem that they own. Note that by necessity this requires modifying the
// filesystem (and thus will not work on read-only filesystems).
var Rootless FsEval = unprivFsEval(0)

// unprivFsEval is a hack to be able to make RootlessFsEval a const.
type unprivFsEval int

// Open is equivalent to unpriv.Open.
func (fs unprivFsEval) Open(path string) (*os.File, error) {
	_ = "STUB: not implemented"
	return nil,

		// Create is equivalent to unpriv.Create.
		nil
}

func (fs unprivFsEval) Create(path string) (*os.File, error) {
	_ = "STUB: not implemented"
	return nil,

		// Readdir is equivalent to unpriv.Readdir.
		nil
}

func (fs unprivFsEval) Readdir(path string) ([]os.FileInfo, error) {
	_ = "STUB: not implemented"
	return nil,

		// Lstat is equivalent to unpriv.Lstat.
		nil
}

func (fs unprivFsEval) Lstat(path string) (os.FileInfo, error) {
	_ = "STUB: not implemented"
	return *new(os.FileInfo), nil
}

func (fs unprivFsEval) Lstatx(path string) (unix.Stat_t, error) {
	_ = "STUB: not implemented"
	return *new(unix.Stat_t), nil
}

// Readlink is equivalent to unpriv.Readlink.
func (fs unprivFsEval) Readlink(path string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil

	// Symlink is equivalent to unpriv.Symlink.
}

func (fs unprivFsEval) Symlink(target, linkname string) error {
	_ = "STUB: not implemented"
	return nil
}

// Link is equivalent to unpriv.Link.
func (fs unprivFsEval) Link(target, linkname string) error { _ = "STUB: not implemented"; return nil }

// Chmod is equivalent to unpriv.Chmod.
func (fs unprivFsEval) Chmod(path string, mode os.FileMode) error {
	_ = "STUB: not implemented"
	return nil
}

// Lutimes is equivalent to unpriv.Lutimes.
func (fs unprivFsEval) Lutimes(path string, atime, mtime time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

// RemoveAll is equivalent to unpriv.RemoveAll.
func (fs unprivFsEval) RemoveAll(path string) error { _ = "STUB: not implemented"; return nil }

// Mknod is equivalent to unpriv.Mknod.
func (fs unprivFsEval) Mknod(path string, mode os.FileMode, dev uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// MkdirAll is equivalent to unpriv.MkdirAll.
func (fs unprivFsEval) MkdirAll(path string, perm os.FileMode) error {
	_ = "STUB: not implemented"
	return nil
}

// Llistxattr is equivalent to unpriv.Llistxattr.
func (fs unprivFsEval) Llistxattr(path string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Lremovexattr is equivalent to unpriv.Lremovexattr.
}

func (fs unprivFsEval) Lremovexattr(path, name string) error { _ = "STUB: not implemented"; return nil }

// Lsetxattr is equivalent to unpriv.Lsetxattr.
func (fs unprivFsEval) Lsetxattr(path, name string, value []byte, flags int) error {
	_ = "STUB: not implemented"
	return nil
}

// Lgetxattr is equivalent to unpriv.Lgetxattr.
func (fs unprivFsEval) Lgetxattr(path, name string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Lclearxattrs is equivalent to unpriv.Lclearxattrs.
func (fs unprivFsEval) Lclearxattrs(path string, skipFn func(xattrName string) bool) error {
	_ = "STUB: not implemented"
	return nil
}

// KeywordFunc returns a wrapper around the given mtree.KeywordFunc.
func (fs unprivFsEval) KeywordFunc(fn mtree.KeywordFunc) mtree.KeywordFunc {
	_ = "STUB: not implemented"
	return *new(mtree.KeywordFunc)
}

// Walk is equivalent to filepath.Walk.
func (fs unprivFsEval) Walk(root string, fn filepath.WalkFunc) error {
	_ = "STUB: not implemented"
	return nil
}

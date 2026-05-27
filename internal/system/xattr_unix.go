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

package system

// Llistxattr is a wrapper around unix.Llistattr, to abstract the NUL-splitting
// and resizing of the returned []string.
func Llistxattr(path string) ([]string, error) {
	_ = "STUB: not implemented"
	//nolint:prealloc // we do pre-allocate later
	return nil, nil
}

// Find the size.

// Could not get the size.

// Get the buffer.

// If we got an ERANGE then we have to resize the buffer because
// someone raced with us getting the list. Don't you just love C
// interfaces.

// Split the buffer.

// "" is not a valid xattr (weirdly you get ERANGE -- not EINVAL -- if
// you try to touch it). So just skip it.

// Lgetxattr is a wrapper around unix.Lgetattr, to abstract the resizing of the
// returned []string.
func Lgetxattr(path, name string) ([]byte, error) {
	_ = "STUB: not implemented"
	//nolint:prealloc // we do pre-allocate later
	return nil, nil
}

// Find the size.

// Could not get the size.

// Get the buffer.

// If we got an ERANGE then we have to resize the buffer because
// someone raced with us getting the list. Don't you just love C
// interfaces.

// Lclearxattrs is a wrapper around Llistxattr and Lremovexattr, which attempts
// to remove all xattrs from a given file.
//
// If skipFn is non-nil and returns true when passed an xattr we planned to
// remove, that xattr is skipped and remains set on the path.
func Lclearxattrs(path string, skipFn func(xattrName string) bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Ignore permission errors, because hitting a permission error
// means that it's a security.* xattr label or something similar.

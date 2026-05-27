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

import (
	"io"
)

// Copy has identical semantics to io.Copy except it will automatically resume
// the copy after it receives an EINTR error.
func Copy(dst io.Writer, src io.Reader) (int64, error) {
	_ = "STUB: not implemented"
	// Make a buffer so io.Copy doesn't make one for each iteration.
	return 0, nil
}

// n is always non-negative

// CopyN has identical semantics to io.CopyN except it will automatically
// resume the copy after it receives an EINTR error.
func CopyN(dst io.Writer, src io.Reader, n int64) (int64, error) {
	_ = "STUB: not implemented"
	// This is based on the stdlib io.CopyN implementation.
	return 0, nil
}

// somewhat confusing io.CopyN semantics

// if the source ends prematurely, io.EOF

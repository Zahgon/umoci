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
	"errors"
	"io"

	ispec "github.com/opencontainers/image-spec/specs-go/v1"
)

var errInvalidDescriptorSize = errors.New("descriptor size must not be negative")

// GetVerifiedBlob returns a VerifiedReadCloser for retrieving a blob from the
// image, which the caller must Close() *and* read-to-EOF (checking the error
// code of both). Returns ErrNotExist if the digest is not found, and
// ErrBlobDigestMismatch on a mismatched blob digest. In addition, the reader
// is limited to the descriptor.Size.
func (e Engine) GetVerifiedBlob(ctx context.Context, descriptor ispec.Descriptor) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	// Negative sizes are not permitted by the spec, and are a DoS vector.
	return *new(io.ReadCloser), nil
}

// The empty blob descriptor only has one valid value so we should validate
// it before allowing it to be opened.

// Embedded data.

// If the digest is small enough to fit in the descriptor, we can
// validate it immediately without deferring to VerifiedReadCloser.

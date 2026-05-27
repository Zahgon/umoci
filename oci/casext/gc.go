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

	"github.com/opencontainers/go-digest"
)

// GCPolicy is a policy function that returns 'true' if a blob can be GC'ed.
type GCPolicy func(ctx context.Context, digest digest.Digest) (bool, error)

// GC will perform a mark-and-sweep garbage collection of the OCI image
// referenced by the given CAS engine. The root set is taken to be the set of
// references stored in the image, and all blobs not reachable by following a
// descriptor path from the root set will be removed.
//
// GC will only call ListBlobs and ListReferences once, and assumes that there
// is no change in the set of references or blobs after calling those
// functions. In other words, it assumes it is the only user of the image that
// is making modifications. Things will not go well if this assumption is
// challenged.
//
// Furthermore, GC policies (zero or more) can also be specified which given a
// blob's digest can indicate whether that blob needs to garbage collected. The
// blob is skipped for garbage collection if a policy returns false.
func (e Engine) GC(ctx context.Context, policies ...GCPolicy) error {
	_ = "STUB: not implemented"
	// Generate the root set of descriptors.
	return nil
}

// Mark from the root sets.

// Sweep all blobs in the white set.

// Digest is in the black set.

// skip this blob for GC

// Finally, tell CAS to GC it.

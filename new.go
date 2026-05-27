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

package umoci

import (
	"time"

	"github.com/opencontainers/umoci/oci/casext"
)

// NewImage creates a new empty image (tag) in the existing layout.
func NewImage(engineExt casext.Engine, tagName string, sourceDateEpoch *time.Time) error {
	_ = "STUB: not implemented"
	// Create a new manifest.
	return nil
}

// Create a new image config.

// Make sure we have no diffids.

// Update config and create a new blob for it.

// Create a new manifest that just points to the config and has an
// empty layer set. FIXME: Implement ManifestList support.

// FIXME: This is hardcoded at the moment.

// Now create a new reference, and either add it to the engine or spew it
// to stdout.

// FIXME: Support manifest lists.

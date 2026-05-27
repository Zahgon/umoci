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

	ispec "github.com/opencontainers/image-spec/specs-go/v1"

	"github.com/opencontainers/umoci/mutate"
	"github.com/opencontainers/umoci/oci/casext"
	"github.com/opencontainers/umoci/pkg/mtreefilter"
)

// Repack repacks a bundle into an image adding a new layer for the changed
// data in the bundle.
//
// If layerCompressor is nil, the compression algorithm is auto-selected.
func Repack(engineExt casext.Engine, tagName, bundlePath string,
	meta Meta,
	history *ispec.History,
	filters []mtreefilter.FilterFunc,
	refreshBundle bool,
	mutator *mutate.Mutator,
	layerCompressor mutate.Compressor, //nolint:staticcheck // SA1019: this interface is defined by us and we keep it for compatibility
	sourceDateEpoch *time.Time,
) (Err error) {
	_ = "STUB: not implemented"
	return nil
}

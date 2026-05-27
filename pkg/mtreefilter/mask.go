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

package mtreefilter

import (
	"github.com/vbatts/go-mtree"
)

// FilterFunc is a function used when filtering deltas with FilterDeltas.
type FilterFunc func(path string) bool

// makeRoot does a very simple job of converting a path to a lexical
// relative-to-root. In mtree we don't deal with any symlink components.
func makeRoot(path string) string { _ = "STUB: not implemented"; return "" }

func maskFilter(maskedPaths map[string]struct{}, includeSelf bool) FilterFunc {
	_ = "STUB: not implemented"
	return *new(FilterFunc)
}

// Convert the path to be cleaned and relative-to-root.

// Check that no ancestor of the path is a masked path.

// MaskFilter is a factory for FilterFuncs that will mask all InodeDelta paths
// that are lexical children of any path in the mask slice. All paths are
// considered to be relative to '/'.
func MaskFilter(masks []string) FilterFunc { _ = "STUB: not implemented"; return *new(FilterFunc) }

// SimplifyFilter is a factory that takes a list of InodeDelta and creates a
// filter to filter out all deletion entries that have a parent which also has
// a deletion entry. This is necessary to both reduce our image sizes and
// remain compatible with Docker's now-incompatible image format (the OCI spec
// doesn't require this behaviour but it's now needed because of course Docker
// won't fix their own bugs).
func SimplifyFilter(deltas []mtree.InodeDelta) FilterFunc {
	_ = "STUB: not implemented"
	return *new(FilterFunc)
}

// FilterDeltas is a helper function to easily filter []mtree.InodeDelta with a
// filter function. Only entries which have `filter(delta.Path()) == true` will
// be included in the returned slice.
func FilterDeltas(deltas []mtree.InodeDelta, filters ...FilterFunc) []mtree.InodeDelta {
	_ = "STUB: not implemented"
	return nil
}

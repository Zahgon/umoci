//go:build gofuzz

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

package mutate

import (
	ispec "github.com/opencontainers/image-spec/specs-go/v1"

	"github.com/opencontainers/umoci/oci/cas"
)

// fuzzSetup() does the necessary setup for the fuzzer, it takes a data
// parameter provided by the fuzzer.
func fuzzSetup(dir string, data []byte) (cas.Engine, ispec.Descriptor, error) {
	_ = "STUB: not implemented"
	return *new(cas.Engine), *new(ispec.Descriptor), nil
}

// Write a tar layer.

// Push the base layer.

// Create a config.

// Create the manifest.

// FuzzMutate implements the fuzzer.
func FuzzMutate(data []byte) int { _ = "STUB: not implemented"; return 0 }

//nolint:errcheck

//nolint:errcheck

// This isn't a valid image, but whatever.

// Add a new layer.

// Cache the data to check it.

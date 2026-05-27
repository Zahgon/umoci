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

package layer

import (
	ispec "github.com/opencontainers/image-spec/specs-go/v1"

	"github.com/opencontainers/umoci/oci/casext"
)

func createRandomFile(dirpath string, filename, filecontents []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func createRandomDir(basedir string, dirname []byte, dirArray []string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func isLetter(input []byte) bool { _ = "STUB: not implemented"; return false }

// FuzzGenerateLayer implements a fuzzer that targets layer.GenerateLayer().
func FuzzGenerateLayer(data []byte) int { _ = "STUB: not implemented"; return 0 }

//nolint:errcheck

//nolint:prealloc

// lower
// upper

// Get initial.

// lower
// upper

// Get post.

//nolint:errcheck

func makeFuzzImage(base641, base642 string) (string, ispec.Manifest, casext.Engine, error) {
	_ = "STUB: not implemented"
	return "", *new(ispec.Manifest), *new(casext.Engine), nil
}

// Create our image.

// Set up the CAS and an image from the above layers.

// Create the config.

// Create the manifest.

// FuzzUnpack implements a fuzzer that targets UnpackManifest().
func FuzzUnpack(data []byte) int {
	_ = "STUB: not implemented"
	// We would like as little log output as possible:
	return 0
}

//nolint:errcheck

//nolint:errcheck

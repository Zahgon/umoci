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

package blobcompress

import (
	"io"
)

// Gzip provides concurrent gzip blobcompression and deblobcompression.
var Gzip Algorithm = gzipAlgo{}

type gzipAlgo struct{}

func (gz gzipAlgo) MediaTypeSuffix() string { _ = "STUB: not implemented"; return "" }

// gzipBlockSize is the block size we use when generating gzip blobs. Changing
// this value could result in different hashes (compared to the old setting)
// for the same inputs, so it must not be changed except in exceptional
// circumstances.
//
// This value was chosen to match the buffer size of containerd/docker because
// it seems Docker will transparently re-compress blobs and a different block
// size will result in different hashes. This is probably an unintentional
// implementation detail that could change in the future, but given that it
// will cause all layer blobs to have different hashes you can hope they would
// notice if they break it.
//
// This also matches the new default for github.com/klauspost/pgzip.
//
// TODO: Make this configurable, with a warning to only change it in
// exceptional circumstances.
const gzipBlockSize = 1 << 20

func (gz gzipAlgo) Compress(reader io.Reader) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

// We don't CloseWithError because we cannot override the Close.

func (gz gzipAlgo) Decompress(reader io.Reader) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

func init() {
	MustRegisterAlgorithm(Gzip)
}

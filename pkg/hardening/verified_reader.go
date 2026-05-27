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

package hardening

import (
	"errors"
	"io"

	"github.com/apex/log"
	"github.com/opencontainers/go-digest"
)

// Exported errors for verification issues that occur during processing within
// VerifiedReadCloser.
var (
	ErrDigestMismatch      = errors.New("verified reader digest mismatch")
	ErrSizeMismatch        = errors.New("verified reader size mismatch")
	ErrInvalidExpectedSize = errors.New("verified reader has invalid expected size")
)

// VerifiedReadCloser is a basic io.ReadCloser which allows for simple
// verification that a stream matches an expected hash. The entire stream is
// hashed while being passed through this reader, and on EOF it will verify
// that the hash matches the expected hash. If not, an error is returned. Note
// that this means you need to read all input to EOF in order to find
// verification errors.
//
// If Reader is a VerifiedReadCloser (with the same ExpectedDigest), all of the
// methods are just piped to the underlying methods (with no verification in
// the upper layer).
type VerifiedReadCloser struct {
	// Reader is the underlying reader.
	Reader io.ReadCloser

	// ExpectedDigest is the expected digest. When the underlying reader
	// returns an EOF, the entire stream's sum will be compared to this hash
	// and an error will be returned if they don't match.
	ExpectedDigest digest.Digest

	// ExpectedSize is the expected amount of data to be read overall. If the
	// underlying reader hasn't returned an EOF by the time this value is
	// exceeded, an error is returned and no further reads will occur.
	ExpectedSize int64

	// digester stores the current state of the stream's hash.
	digester digest.Digester

	// currentSize is the number of bytes that have been read so far.
	currentSize int64
}

func (v *VerifiedReadCloser) init() {
	// Define digester if not already set.
	if v.digester == nil {
		alg := v.ExpectedDigest.Algorithm()
		if !alg.Available() {
			log.Fatalf("verified reader: unsupported hash algorithm %s", alg) //nolint:revive // panic is for extra safety
			panic("verified reader: unreachable section")                     // should never be hit
		}
		v.digester = alg.Digester()
	}
}

func (v *VerifiedReadCloser) isNoop() bool { _ = "STUB: not implemented"; return false }

func (v *VerifiedReadCloser) check() error { _ = "STUB: not implemented"; return nil }

func (v *VerifiedReadCloser) verify(nilErr error) error {
	_ = "STUB: not implemented"
	// Digest mismatch (always takes precedence)?
	return nil
}

// Not enough bytes in the stream.

// We don't read the entire blob, so the message needs to be slightly adjusted.

// Forward the provided error.

// Read is a wrapper around VerifiedReadCloser.Reader, with a digest check on
// EOF.  Make sure that you always check for EOF and read-to-the-end for all
// files.
func (v *VerifiedReadCloser) Read(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Make sure we don't read after v.ExpectedSize has been passed.

// We still have something left to read.

// Piped to the underling read.

// We have either read everything, or just happened to land on a boundary
// (with potentially more things afterwards). So we must check if there is
// anything left by doing a 1-byte read (Go doesn't allow for zero-length
// Read()s to give EOFs).

// We just want to know whether we read something (n>0). Whatever we
// read is irrelevant because if we read something that means the
// reader will fail to verify.

// Are we going to be a noop?

// Make sure we're ready.

// Forward it to the digester.

// hash.Hash guarantees Write() never fails and is never short.

//nolint:revive // panic is for extra safety
// should never be hit

// We have finished reading -- let's verify the state!

// sourceName returns a debugging-friendly string to indicate to the user what
// the source reader is for this verified reader.
func (v *VerifiedReadCloser) sourceName() string { _ = "STUB: not implemented"; return "" }

// TODO: Maybe handle things like io.NopCloser by using reflection?

// Close is a wrapper around VerifiedReadCloser.Reader, but with a digest check
// which will return an error if the underlying Close() didn't.
func (v *VerifiedReadCloser) Close() error { _ = "STUB: not implemented"; return nil }

// Consume any remaining bytes to make sure that we've actually read to the
// end of the stream. VerifiedReadCloser.Read will not read past
// ExpectedSize+1, so we don't need to add a limit here.

// If there's trailing bytes being discarded at this point, that
// indicates whatever you used to generate this blob is adding trailing
// gunk.

// Piped to underlying close.

// Are we going to be a noop?

// Make sure we're ready.

// Verify the state.

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

package mediatype

import (
	"errors"
	"io"
	"sync"

	ispec "github.com/opencontainers/image-spec/specs-go/v1"
)

// ErrNilReader is returned by the parsers in this package when they are called
// with a nil Reader. You may use this error for the same purpose if you wish,
// but it's not required.
var ErrNilReader = errors.New("")

// ParseFunc is a parser that is registered for a given mediatype and called
// to parse a blob if it is encountered. If possible, the blob should be
// represented as a native Go object (with all Descriptors represented as
// ispec.Descriptor objects) -- this will allow umoci to recursively discover
// blob dependencies.
//
// NOTE: Your ParseFunc must be able to accept a nil Reader (the error value is
// not relevant) and must return a struct. This is used during registration in
// order to determine the type of the struct (thus you must return the same
// struct you would return in a non-nil reader scenario). Go doesn't have a way
// for us to enforce this.
type ParseFunc func(io.Reader) (any, error)

var (
	lock sync.RWMutex

	// parsers is a mapping of media-type to parser function.
	parsers = map[string]ParseFunc{}

	// packages is the set of package paths which have been registered.
	packages = map[string]struct{}{}

	// targets is the set of media-types which are treated as "targets" for the
	// purposes of reference resolution (resolution terminates at these targets
	// as well as any un-parseable blob types).
	targets = map[string]struct{}{}
)

// IsRegisteredPackage returns whether a parser which returns a type from the
// given package path was registered. This is only useful to allow restricting
// reflection recursion (as a first-pass to limit how deep reflection goes).
func IsRegisteredPackage(pkgPath string) bool { _ = "STUB: not implemented"; return false }

// GetParser returns the ParseFunc that was previously registered for the given
// media-type with RegisterParser (or nil if the media-type is unknown).
func GetParser(mediaType string) ParseFunc { _ = "STUB: not implemented"; return *new(ParseFunc) }

// RegisterParser registers a new ParseFunc to be used when the given
// media-type is encountered during parsing or recursive walks of blobs. See
// the documentation of ParseFunc for more detail. The returned ParseFunc must
// return a struct.
func RegisterParser(mediaType string, parser ParseFunc) {
	_ = "STUB: not implemented"
	// Get the return type so we know what packages are white-listed for
	// recursion.
	return
}

// Ensure the returned type is actually a struct. Ideally we would detect
// this with generics but this seems to not be possible with Go generics
// (circa Go 1.20).
// programmer bug

// Register the parser and package.

// programmer bug

// IsTarget returns whether the given media-type should be treated as a "target
// media-type" for the purposes of reference resolution. This means that either
// the media-type has been registered as a target (using RegisterTarget) or has
// not been registered as parseable (using RegisterParser).
func IsTarget(mediaType string) bool { _ = "STUB: not implemented"; return false }

// RegisterTarget registers that a given *parseable* media-type (meaning that
// there is a parser already registered using RegisterParser) should be treated
// as a "target" for the purposes of reference resolution. This means that if
// this media-type is encountered during a reference resolution walk, a
// DescriptorPath to *that* blob will be returned and resolution will not
// recurse any deeper. All un-parseable blobs are treated as targets, so this
// is only useful for blobs that have also been given parsers.
func RegisterTarget(mediaType string) { _ = "STUB: not implemented"; return }

// JSONParser is a minimal wrapper around
//
//	var v T
//	return json.NewDecoder(rdr).Decode(&v)
//
// But also handling the rdr == nil case which is needed for RegisterParser to
// correctly detect what type T is. T must be a struct (this is verified by
// RegisterParser).
func JSONParser[T any](rdr io.Reader) (_ any, err error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// must not return a nil interface{}

func indexParser(rdr io.Reader) (any, error) {
	_ = "STUB: not implemented"
	// Construct a fake struct which contains fields that shouldn't exist, to
	// detect images that have maliciously-inserted fields. CVE-2021-41190
	return *new(any), nil
}

// must not return a nil interface{}

func manifestParser(rdr io.Reader) (any, error) {
	_ = "STUB: not implemented"
	// Construct a fake struct which contains fields that shouldn't exist, to
	// detect images that have maliciously-inserted fields. CVE-2021-41190
	return *new(any), nil
}

// must not return a nil interface{}

// emptyJSONParser only parses "application/vnd.oci.empty.v1+json" and
// validates that it is actually "{}".
func emptyJSONParser(rdr io.Reader) (any, error) {
	_ = "STUB: not implemented"

	// must not return a nil interface{}
	return *new(any), nil
}

// The only valid value for this blob.

// Try to read at least one more byte than emptyJSON so if there is some
// trailing data we will error out without needing to read any more.

// Register the core image-spec types.
func init() {
	RegisterParser(ispec.MediaTypeDescriptor, JSONParser[ispec.Descriptor])
	RegisterParser(ispec.MediaTypeImageIndex, indexParser)
	RegisterParser(ispec.MediaTypeImageConfig, JSONParser[ispec.Image])
	RegisterParser(ispec.MediaTypeEmptyJSON, emptyJSONParser)

	RegisterTarget(ispec.MediaTypeImageManifest)
	RegisterParser(ispec.MediaTypeImageManifest, manifestParser)
}

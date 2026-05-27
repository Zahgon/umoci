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
	"context"
	"encoding/json"
	"io"

	"github.com/opencontainers/go-digest"
	ispec "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/urfave/cli"
	"github.com/vbatts/go-mtree"

	"github.com/opencontainers/umoci/oci/casext"
	"github.com/opencontainers/umoci/oci/layer"
)

// FIXME: This should be moved to a library. Too much of this code is in the
//        cmd/... code, but should really be refactored to the point where it
//        can be useful to other people. This is _particularly_ true for the
//        code which repacks images (the changes to the config, manifest and
//        CAS should be made into a library).

// MtreeKeywords is the set of keywords used by umoci for verification and diff
// generation of a bundle. This is based on mtree.DefaultKeywords, but is
// hardcoded here to ensure that vendor changes don't mess things up.
var MtreeKeywords = []mtree.Keyword{
	"size",
	"type",
	"uid",
	"gid",
	"mode",
	"link",
	"nlink",
	"tar_time",
	"sha256digest",
	"xattr",
}

// MetaName is the name of umoci's metadata file that is stored in all
// bundles extracted by umoci.
const MetaName = "umoci.json"

// MetaVersion is the version of Meta supported by this code. The
// value is only bumped for updates which are not backwards compatible.
const MetaVersion = "2"

// Meta represents metadata about how umoci unpacked an image to a bundle
// and other similar information. It is used to keep track of information that
// is required when repacking an image and other similar bundle information.
type Meta struct {
	// Version is the version of umoci used to unpack the bundle. This is used
	// to future-proof the umoci.json information.
	Version string `json:"umoci_version"`

	// From is a copy of the descriptor pointing to the image manifest that was
	// used to unpack the bundle. Essentially it's a resolved form of the
	// --image argument to umoci-unpack(1).
	From casext.DescriptorPath `json:"from_descriptor_path"`

	// MapOptions is the parsed version of --uid-map, --gid-map and --rootless
	// arguments to umoci-unpack(1). While all of these options technically do
	// not need to be the same for corresponding umoci-unpack(1) and
	// umoci-repack(1) calls, changing them is not recommended and so the
	// default should be that they are the same.
	MapOptions layer.MapOptions `json:"map_options"`

	// WhiteoutMode indicates what style of whiteout was written to disk
	// when this filesystem was extracted.
	//
	// Deprecated: This feature was completely broken. See
	// <https://github.com/opencontainers/umoci/issues/574> for more details.
	WhiteoutMode int `json:"whiteout_mode,omitempty"`
}

// WriteTo writes a JSON-serialised version of Meta to the given io.Writer.
func (m Meta) WriteTo(w io.Writer) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

// WriteBundleMeta writes an umoci.json file to the given bundle path.
func WriteBundleMeta(bundle string, meta Meta) (Err error) { _ = "STUB: not implemented"; return nil }

// ReadBundleMeta reads and parses the umoci.json file from a given bundle path.
func ReadBundleMeta(bundle string) (_ Meta, Err error) {
	_ = "STUB: not implemented"
	return *new(Meta), nil
}

// NOTE: This field has been deprecated, as the feature was completely
// broken. See <https://github.com/opencontainers/umoci/issues/574> for
// more details.

// ManifestStat has information about a given OCI manifest.
// TODO: Implement support for manifest lists, this should also be able to
// contain stat information for a list of manifests.
type ManifestStat struct {
	Manifest manifestStat `json:"manifest"`

	// Config stores information about the configuration of a manifest.
	Config *configStat `json:"config,omitzero"`

	// History stores the history information for the manifest.
	History historyStatList `json:"history,omitzero"`
}

// quote is a wrapper around [strconv.Quote] that only returns a quoted string
// if it is actually necessary. The precise flag indicates whether the field
// being quoted needs to provide extra accuracy to the user (in particular,
// regarding whitespace and empty strings).
func quote(s string, precise bool) string { _ = "STUB: not implemented"; return "" }

func pprint(w io.Writer, prefix, key string, values ...string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Make sure "," leads to quoting.

func pprintSlice(w io.Writer, prefix, name string, data []string) error {
	_ = "STUB: not implemented"
	return nil
}

func pprintMap(w io.Writer, prefix, name string, data map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

func pprintSet(w io.Writer, prefix, name string, data map[string]struct{}) error {
	_ = "STUB: not implemented"
	return nil
}

func xxd(w io.Writer, prefix string, r io.Reader) error { _ = "STUB: not implemented"; return nil }

// like xxd

// truncate

// <prefix><offset>:

// <hex1><hex2> <hex3><hex4> ...

// Pad out the hex output for short lines.

// <textual output>

func pprintBytes(w io.Writer, prefix, name string, data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// Do not use pprint, as we do not want our own suffix to get quote()d.

// We want to limit how much data we dump to the screen.

func pprintPlatform(w io.Writer, prefix string, platform ispec.Platform) error {
	_ = "STUB: not implemented"
	return nil
}

// Do not use pprint, as we do not want our own suffix to get quote()d.

// pprintDescriptor pretty-prints an ispec.Descriptor.
func pprintDescriptor(w io.Writer, prefix string, descriptor ispec.Descriptor) error {
	_ = "STUB: not implemented"
	return nil
}

// Format formats a ManifestStat using the default formatting, and writes the
// result to the given writer.
//
// TODO: This should really be implemented in a way that allows for users to
// define their own custom templates for different blocks (meaning that this
// should use text/template rather than using tabwriters manually.
func (ms ManifestStat) Format(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// manifestStat contains information about the image manifest.
type manifestStat struct {
	// Descriptor is the descriptor for the manifest JSON.
	Descriptor ispec.Descriptor `json:"descriptor"`

	// Manifest is the contents of the image manifest.
	Manifest ispec.Manifest `json:"-"`

	// RawData is the raw data stream of the blob, which is output when we
	// provide JSON output (to make sure no information is lost in --json
	// mode).
	RawData json.RawMessage `json:"blob,omitzero"`
}

func pprintManifest(w io.Writer, prefix string, manifest ispec.Manifest) error {
	_ = "STUB: not implemented"
	return nil
}

func (m manifestStat) pprint(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// configStat contains information about the image configuration of this
// manifest.
type configStat struct {
	// Descriptor is the descriptor for the configuration JSON.
	Descriptor ispec.Descriptor `json:"descriptor"`

	// Image is the contents of the configuration.
	Image *ispec.Image `json:"-"`

	// RawData is the raw data stream of the blob, which is output when we
	// provide JSON output (to make sure no information is lost in --json
	// mode).
	RawData json.RawMessage `json:"blob,omitzero"`
}

func pprintImageConfig(w io.Writer, prefix string, config ispec.ImageConfig) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:staticcheck // we need to support this deprecated field

func pprintImage(w io.Writer, prefix string, image ispec.Image) error {
	_ = "STUB: not implemented"
	return nil
}

func (c configStat) pprint(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// historyStat contains information about a single entry in the history of a
// manifest. This is essentially equivalent to a single record from
// docker-history(1).
type historyStat struct {
	// Layer is the descriptor referencing where the layer is stored. If it is
	// nil, then this entry is an empty_layer (and thus doesn't have a backing
	// diff layer).
	Layer *ispec.Descriptor `json:"layer"`

	// DiffID is an additional piece of information to Layer. It stores the
	// DiffID of the given layer corresponding to the history entry. If DiffID
	// is "", then this entry is an empty_layer.
	DiffID digest.Digest `json:"diff_id"`

	// History is embedded in the stat information.
	ispec.History
}

type historyStatList []historyStat

func (hsl historyStatList) pprint(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// TODO: We need to truncate some of the fields.

// Stat computes the ManifestStat for a given manifest blob. The provided
// descriptor must refer to an OCI Manifest.
func Stat(ctx context.Context, engine casext.Engine, manifestDescriptor ispec.Descriptor) (ManifestStat, error) {
	_ = "STUB: not implemented"
	return *new(ManifestStat), nil
}

// We have to get the actual manifest.

// Should _never_ be reached.

// Get the config.

// If the config is a valid config blob, fill the stat information.

// Generate the history of the image. config.History entries are in the
// same order as manifest.Layer, but "empty layer" entries may be
// interspersed so we need to skip over those when associating layers
// to history entries.

// Only fill the other information and increment layerIdx if it's a
// non-empty layer.

// If the config could be parsed successfully (giving us RawData), then
// fill the raw data section for the JSON output and provide a
// descriptor for the pprint output, but don't pretty-print an image
// config object.

// GenerateBundleManifest creates and writes an mtree of the rootfs in the given
// bundle path, using the supplied fsEval method.
func GenerateBundleManifest(mtreeName, bundlePath string, fsEval mtree.FsEval) (Err error) {
	_ = "STUB: not implemented"
	return nil
}

// ParseIdmapOptions sets up the mapping options for Meta, using
// the arguments specified on the command line.
func ParseIdmapOptions(meta *Meta, ctx *cli.Context) error {
	_ = "STUB: not implemented"
	// We need to set mappings if we're in rootless mode.
	return nil
}

// Should _never_ be reached.

// Should _never_ be reached.

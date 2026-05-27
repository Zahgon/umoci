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
	"context"
	"io"

	// Import is necessary for go-digest.
	_ "crypto/sha256"

	ispec "github.com/opencontainers/image-spec/specs-go/v1"

	"github.com/opencontainers/umoci/oci/cas"
	"github.com/opencontainers/umoci/oci/casext/blobcompress"
)

// AfterLayerUnpackCallback is called after each layer is unpacked.
type AfterLayerUnpackCallback func(manifest ispec.Manifest, desc ispec.Descriptor) error

// UnpackLayer unpacks the tar stream representing an OCI layer at the given
// root. It ensures that the state of the root is as close as possible to the
// state used to create the layer. If an error is returned, the state of root
// is undefined (unpacking is not guaranteed to be atomic).
func UnpackLayer(root string, layer io.Reader, opt *UnpackOptions) error {
	_ = "STUB: not implemented"
	return nil
}

// RootfsName is the name of the rootfs directory inside the bundle path when
// generated.
const RootfsName = "rootfs"

func isLayerType(mediaType string) bool { _ = "STUB: not implemented"; return false }

//nolint:staticcheck // we need to support this deprecated media-type

func getLayerCompressAlgorithm(mediaType string) (string, blobcompress.Algorithm, error) {
	_ = "STUB: not implemented"
	return "", *new(blobcompress.Algorithm), nil
}

// UnpackManifest extracts all of the layers in the given manifest, as well as
// generating a runtime bundle and configuration. The rootfs is extracted to
// <bundle>/<layer.RootfsName>.
//
// FIXME: This interface is ugly.
func UnpackManifest(ctx context.Context, engine cas.Engine, bundle string, manifest ispec.Manifest, opt *UnpackOptions) (Err error) {
	_ = "STUB: not implemented"

	// Create the bundle directory. We only error out if config.json or rootfs/
	// already exists, because we cannot be sure that the user intended us to
	// extract over an existing bundle.
	return nil
}

// We change the mode of the bundle directory to 0700. A user can easily
// change this after-the-fact, but we do this explicitly to avoid cases
// where an unprivileged user could recurse into an otherwise unsafe image
// (giving them potential root access through setuid binaries for example).

// It's too late to care about errors.

// Generate a runtime configuration file from ispec.Image.

// UnpackRootfs extracts all of the layers in the given manifest.
// Some verification is done during image extraction.
func UnpackRootfs(ctx context.Context, engine cas.Engine, rootfsPath string, manifest ispec.Manifest, opt *UnpackOptions) (Err error) {
	_ = "STUB: not implemented"

	// TODO: For now, unpacking layers into a bundle with the overlayfs on-disk
	// format is not supported, because we still unpack everything into a
	// single rootfs directory. For more information about outstanding issues,
	// see <https://github.com/opencontainers/umoci/issues/574>.
	return nil
}

// In order to avoid having a broken rootfs in the case of an error, we
// remove the rootfs. In the case of rootless this is particularly
// important (`rm -rf` won't work on most distro rootfs's).

// It's too late to care about errors.

// Make sure that the owner is correct.

// Currently, many different images in the wild don't specify what the
// atime/mtime of the root directory is. This is a huge pain because it
// means that we can't ensure consistent unpacking. In order to get around
// this, we first set the mtime of the root directory to the Unix epoch
// (which is as good of an arbitrary choice as any).

// In order to verify the DiffIDs as we extract layers, we have to get the
// .Config blob first. But we can't extract it (generate the runtime
// config) until after we have the full rootfs generated.

// Should _never_ be reached.

// We can't understand non-layer images.

// Layer extraction.

//nolint:errcheck // in the non-error path this is a double-close we can ignore

// Should _never_ be reached.

// Pick the decompression algorithm based on the media-type.

// We have to extract a compressed version of the above layer. Also
// note that we have to check the DiffID we're extracting (which is
// the sha256 sum of the *uncompressed* layer).

// Different tar implementations can have different levels of redundant
// padding and other similar weird behaviours. While on paper they are
// all entirely valid archives, Go's tar.Reader implementation doesn't
// guarantee that the entire stream will be consumed (which can result
// in the later diff_id check failing because the digester didn't get
// the whole uncompressed stream). Just blindly consume anything left
// in the layer.

// Same goes for compressed layers -- it seems like some gzip
// implementations add trailing NUL bytes, which Go doesn't slurp up.
// Just eat up the rest of the remaining bytes and discard them.
//
// FIXME: We use layerData here because pgzip returns io.EOF from
// WriteTo, which causes havoc with system.Copy. Ideally we would use
// layerRaw. See <https://github.com/klauspost/pgzip/issues/38>.

// UnpackRuntimeJSON converts a given manifest's configuration to a runtime
// configuration and writes it to the given writer. If rootfs is specified, it
// is sourced during the configuration generation (for conversion of
// Config.User and other similar jobs -- which will error out if the user could
// not be parsed). If rootfs is not specified (is an empty string) then all
// conversions that require sourcing the rootfs will be set to their default
// values.
//
// XXX: I don't like this API. It has way too many arguments.
func UnpackRuntimeJSON(ctx context.Context, engine cas.Engine, configFile io.Writer, rootfs string, manifest ispec.Manifest, opt *MapOptions) (Err error) {
	_ = "STUB: not implemented"
	return nil
}

// In order to verify the DiffIDs as we extract layers, we have to get the
// .Config blob first. But we can't extract it (generate the runtime
// config) until after we have the full rootfs generated.

// Should _never_ be reached.

// Add UIDMapping / GIDMapping options.

// Save the config.json.

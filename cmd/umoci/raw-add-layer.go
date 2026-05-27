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

package main

import (
	"errors"

	"github.com/urfave/cli"
)

var rawAddLayerCommand = uxCompress(uxHistory(uxTag(cli.Command{
	Name:  "add-layer",
	Usage: "add a layer archive verbatim to an image",
	ArgsUsage: `--image <image-path>[:<tag>] <new-layer.tar>

Where "<image-path>" is the path to the OCI image, "<tag>" is the name of the
tagged image to modify (if not specified, defaults to "latest"),
"<new-layer.tar>" is the new layer to add (it must be uncompressed).

Note that using your own layer archives may result in strange behaviours (for
instance, you may need to use --keep-dirlink with umoci-unpack(1) in order to
avoid breaking certain entries).

At the moment, umoci-raw-add-layer(1) will only *append* layers to an image and
only supports uncompressed archives.`,

	// unpack reads manifest information.
	Category: "image",

	Action: rawAddLayer,

	Before: func(ctx *cli.Context) error {
		if ctx.NArg() != 1 {
			return errors.New("invalid number of positional arguments: expected <newlayer.tar>")
		}
		if ctx.Args().First() == "" {
			return errors.New("<new-layer.tar> path cannot be empty")
		}
		ctx.App.Metadata["newlayer"] = ctx.Args().First()
		return nil
	},
})))

func rawAddLayer(ctx *cli.Context) (Err error) { _ = "STUB: not implemented"; return nil }

// Overide the from tag by default, otherwise use the one specified.

// Get a reference to the CAS.

// TODO: Handle this more nicely.

// Create the mutator.

// TODO: Verify that the layer is actually uncompressed.

// XXX: Should we append argv to this?

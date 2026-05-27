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

var rawUnpackCommand = uxRemap(cli.Command{
	Name:  "unpack",
	Usage: "unpacks a reference into a rootfs",
	ArgsUsage: `--image <image-path>[:<tag>] <rootfs>

Where "<image-path>" is the path to the OCI image, "<tag>" is the name of the
tagged image to unpack (if not specified, defaults to "latest") and "<rootfs>"
is the destination to unpack the image to.`,

	// unpack reads manifest information.
	Category: "image",

	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "keep-dirlinks",
			Usage: "don't clobber underlying symlinks to directories",
		},
	},

	Action: rawUnpack,

	Before: func(ctx *cli.Context) error {
		if ctx.NArg() != 1 {
			return errors.New("invalid number of positional arguments: expected <rootfs>")
		}
		if ctx.Args().First() == "" {
			return errors.New("rootfs path cannot be empty")
		}
		ctx.App.Metadata["rootfs"] = ctx.Args().First()
		return nil
	},
})

func rawUnpack(ctx *cli.Context) (Err error) { _ = "STUB: not implemented"; return nil }

// Parse map options.
// We need to set mappings if we're in rootless mode.

// Get a reference to the CAS.

// TODO: Handle this more nicely.

// Get the manifest.

// Should _never_ be reached.

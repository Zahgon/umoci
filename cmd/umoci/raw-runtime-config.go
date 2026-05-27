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

var rawConfigCommand = uxRemap(cli.Command{
	Name:    "runtime-config",
	Aliases: []string{"config"},
	Usage:   "generates an OCI runtime configuration for an image",
	ArgsUsage: `--image <image-path>[:<tag>] [--rootfs <rootfs>] <config.json>

Where "<image-path>" is the path to the OCI image, "<tag>" is the name of the
tagged image to unpack (if not specified, defaults to "latest"), "<rootfs>" is
a rootfs to use as a supplementary "source of truth" for certain generation
operations and "<config.json>" is the destination to write the runtime
configuration to.

Note that the results of this may not agree with umoci-unpack(1) because the
--rootfs flag affects how certain properties are interpreted.`,

	// unpack reads manifest information.
	Category: "image",

	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "rootfs",
			Usage: "path to secondary source of truth (root filesystem)",
		},
	},

	Action: rawConfig,

	Before: func(ctx *cli.Context) error {
		if ctx.NArg() != 1 {
			return errors.New("invalid number of positional arguments: expected <config.json>")
		}
		if ctx.Args().First() == "" {
			return errors.New("config.json path cannot be empty")
		}
		ctx.App.Metadata["config"] = ctx.Args().First()
		return nil
	},
})

func rawConfig(ctx *cli.Context) (Err error) { _ = "STUB: not implemented"; return nil }

// Parse and set up the mapping options.

// Get a reference to the CAS.

// TODO: Handle this more nicely.

// Get the manifest.

// Should _never_ be reached.

// Generate the configuration.

// Write out the generated config.

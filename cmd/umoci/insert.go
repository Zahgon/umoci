// SPDX-License-Identifier: Apache-2.0
/*
 * umoci: Umoci Modifies Open Containers' Images
 * Copyright (C) 2016-2025 SUSE LLC
 * Copyright (C) 2018 Cisco Systems
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
	"fmt"

	"github.com/urfave/cli"
)

var insertCommand = uxCompress(uxRemap(uxHistory(uxTag(cli.Command{
	Name:  "insert",
	Usage: "insert content into an OCI image",
	ArgsUsage: `--image <image-path>[:<tag>] [--opaque] <source> <target>
                                  --image <image-path>[:<tag>] [--whiteout] <target>

Where "<image-path>" is the path to the OCI image, and "<tag>" is the name of
the tag that the content wil be inserted into (if not specified, defaults to
"latest").

The path at "<source>" is added to the image with the given "<target>" name.
If "--whiteout" is specified, rather than inserting content into the image, a
removal entry for "<target>" is inserted instead.

If "--opaque" is specified then any paths below "<target>" (assuming it is a
directory) from previous layers will no longer be present. Only the contents
inserted by this command will be visible. This can be used to replace an entire
directory, while the default behaviour merges the old contents with the new.

Note that this command works by creating a new layer, so this should not be
used to remove (or replace) secrets from an already-built image. See
umoci-config(1) and --config.volume for how to achieve this correctly.

Some examples:
	umoci insert --image oci:foo mybinary /usr/bin/mybinary
	umoci insert --image oci:foo myconfigdir /etc/myconfigdir
	umoci insert --image oci:foo --opaque myoptdir /opt
	umoci insert --image oci:foo --whiteout /some/old/dir
`,

	Category: "image",

	Action: insert,

	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "whiteout",
			Usage: "insert a 'removal entry' for the given path",
		},
		cli.BoolFlag{
			Name:  "opaque",
			Usage: "mask any previous entries in the target directory",
		},
	},

	Before: func(ctx *cli.Context) error {
		// This command is quite weird because we need to support two different
		// positional-argument numbers. Awesome.
		numArgs := 2
		if ctx.IsSet("whiteout") {
			numArgs = 1
		}
		if ctx.NArg() != numArgs {
			return fmt.Errorf("invalid number of positional arguments: expected %d", numArgs)
		}
		for idx, args := range ctx.Args() {
			if args == "" {
				return fmt.Errorf("invalid positional argument %d: arguments cannot be empty", idx)
			}
		}

		// Figure out the arguments.
		var sourcePath, targetPath string
		targetPath = ctx.Args()[0]
		if !ctx.IsSet("whiteout") {
			sourcePath = targetPath
			targetPath = ctx.Args()[1]
		}

		ctx.App.Metadata["--source-path"] = sourcePath
		ctx.App.Metadata["--target-path"] = targetPath
		return nil
	},
}))))

func insert(ctx *cli.Context) (Err error) { _ = "STUB: not implemented"; return nil }

// By default we clobber the old tag.

// Get a reference to the CAS.

// TODO: Handle this more nicely.

// Create the mutator.

// Parse and set up the mapping options.

// XXX: Should we append argv to this?

// If set, takes precedence over SOURCE_DATE_EPOCH.

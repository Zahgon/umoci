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

	"github.com/opencontainers/umoci/oci/casext"
)

var tagAddCommand = cli.Command{
	Name:  "tag",
	Usage: "creates a new tag in an OCI image",
	ArgsUsage: `--image <image-path>[:<tag>] <new-tag>

Where "<image-path>" is the path to the OCI image, "<tag>" is the old name of
the tag and "<new-tag>" is the new name of the tag.`,

	// tag modifies an image layout.
	Category: "image",

	Action: tagAdd,

	Before: func(ctx *cli.Context) error {
		if ctx.NArg() != 1 {
			return errors.New("invalid number of positional arguments: expected <new-tag>")
		}
		if ctx.Args().First() == "" {
			return errors.New("new tag cannot be empty")
		}
		if !casext.IsValidReferenceName(ctx.Args().First()) {
			return errors.New("new tag is an invalid reference")
		}
		ctx.App.Metadata["new-tag"] = ctx.Args().First()
		return nil
	},
}

func tagAdd(ctx *cli.Context) (Err error) { _ = "STUB: not implemented"; return nil }

// Get a reference to the CAS.

// Get original descriptor.

// TODO: Handle this more nicely.

// Add it.

var tagRemoveCommand = cli.Command{
	Name:    "remove",
	Aliases: []string{"rm"},
	Usage:   "removes a tag from an OCI image",
	ArgsUsage: `--image <image-path>[:<tag>]


Where "<image-path>" is the path to the OCI image, "<tag>" is the name of the
tag to remove.`,

	// tag modifies an image layout.
	Category: "image",

	Before: func(ctx *cli.Context) error {
		if ctx.NArg() != 0 {
			return errors.New("invalid number of positional arguments: expected none")
		}
		return nil
	},

	Action: tagRemove,
}

func tagRemove(ctx *cli.Context) (Err error) { _ = "STUB: not implemented"; return nil }

// Get a reference to the CAS.

// Remove it.

var tagListCommand = cli.Command{
	Name:    "list",
	Aliases: []string{"ls"},
	Usage:   "lists the set of tags in an OCI layout",
	ArgsUsage: `--layout <image-path>

Where "<image-path>" is the path to the OCI layout.

Gives the full list of tags in an OCI layout, with each tag name on a single
line. See umoci-stat(1) to get more information about each tagged image.`,

	// tag modifies an image layout.
	Category: "layout",

	Before: func(ctx *cli.Context) error {
		if ctx.NArg() != 0 {
			return errors.New("invalid number of positional arguments: expected none")
		}
		return nil
	},

	Action: tagList,
}

func tagList(ctx *cli.Context) (Err error) { _ = "STUB: not implemented"; return nil }

// Get a reference to the CAS.

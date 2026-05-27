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

	ispec "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/urfave/cli"

	"github.com/opencontainers/umoci/mutate"
)

// FIXME: We should also implement a raw mode that just does modifications of
//
//	JSON blobs (allowing this all to be used outside of our build setup).
var configCommand = uxHistory(uxTag(cli.Command{
	Name:  "config",
	Usage: "modifies the image configuration of an OCI image",
	ArgsUsage: `--image <image-path>[:<tag>] [--tag <new-tag>]

Where "<image-path>" is the path to the OCI image, and "<tag>" is the name of
the tagged image from which the config modifications will be based (if not
specified, it defaults to "latest"). "<new-tag>" is the new reference name to
save the new image as, if this is not specified then umoci will replace the old
image.`,

	// config modifies a particular image manifest.
	Category: "image",

	// Verify the metadata.
	Before: func(ctx *cli.Context) error {
		if ctx.NArg() != 0 {
			return errors.New("invalid number of positional arguments: expected none")
		}
		if _, ok := ctx.App.Metadata["--image-path"]; !ok {
			return errors.New("missing mandatory argument: --image")
		}
		if _, ok := ctx.App.Metadata["--image-tag"]; !ok {
			return errors.New("missing mandatory argument: --image")
		}
		return nil
	},

	// Do not re-order arguments.
	//
	// It turns out that urfave/cli incorrectly handles cases like
	// [--config.cmd -c] during argument re-ordering for subcommands, causing
	// us a fair number of issues when users are trying to pass a flag an
	// argument that starts with a dash. Luckily 'umoci config' doesn't take
	// positional arguments, so disabling argument re-ordering has no other
	// real effect.
	//
	// See <https://github.com/urfave/cli/issues/1152> for more details.
	SkipArgReorder: true,

	Flags: []cli.Flag{
		cli.StringFlag{Name: "config.user"},
		cli.StringSliceFlag{Name: "config.exposedports"},
		cli.StringSliceFlag{Name: "config.env"},
		cli.StringSliceFlag{Name: "config.entrypoint"}, // FIXME: This interface is weird.
		cli.StringSliceFlag{Name: "config.cmd"},        // FIXME: This interface is weird.
		cli.StringSliceFlag{Name: "config.volume"},
		cli.StringSliceFlag{Name: "config.label"},
		cli.StringFlag{Name: "config.workingdir"},
		cli.StringFlag{Name: "config.stopsignal"},
		cli.StringFlag{Name: "created"}, // FIXME: Implement TimeFlag.
		cli.StringFlag{Name: "author"},
		cli.StringFlag{Name: "platform.os,os"},
		cli.StringFlag{Name: "platform.arch,architecture"},
		cli.StringFlag{Name: "platform.variant"},
		// TODO: platform.os.{version,features}
		cli.StringSliceFlag{Name: "manifest.annotation"},
		cli.StringSliceFlag{Name: "clear"},
	},

	Action: config,
}))

func toImage(config ispec.ImageConfig, meta mutate.Meta) ispec.Image {
	_ = "STUB: not implemented"
	return *new(ispec.Image)
}

func fromImage(image ispec.Image) (ispec.ImageConfig, mutate.Meta) {
	_ = "STUB: not implemented"
	return *new(ispec.ImageConfig), *new(mutate.Meta)
}

// parseKV splits a given string (of the form name=value) into (name,
// value). An error is returned if there is no "=" in the line or if the
// name is empty.
func parseKV(input string) (string, string, error) { _ = "STUB: not implemented"; return "", "", nil }

func config(ctx *cli.Context) (Err error) { _ = "STUB: not implemented"; return nil }

// By default we clobber the old tag.

// Get a reference to the CAS.

// TODO: Handle this more nicely.

// g.ClearRootfsDiffIDs()

// How do we handle other formats?

// FIXME: This interface is weird.

// FIXME: This interface is weird.

// If set, takes precedence over SOURCE_DATE_EPOCH.

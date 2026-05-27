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
	"github.com/vbatts/go-mtree"

	"github.com/opencontainers/umoci/pkg/fseval"
)

func parseMtreeKeywordArg(arg string) []mtree.Keyword { _ = "STUB: not implemented"; return nil }

func cmdParseMtreeKeywords(ctx *cli.Context, name string) { _ = "STUB: not implemented"; return }

func uxMtreeKeyword(cmd cli.Command) cli.Command {
	_ = "STUB: not implemented"
	return *new(cli.Command)
}

// for compatibility with gomtree

// for compatibility with gomtree

var rawMtreeValidateCommand = uxMtreeKeyword(uxRootless(cli.Command{
	Name:  "mtree-validate",
	Usage: `validate an mtree manifest (akin to "go-mtree validate")`,
	ArgsUsage: `--manifest <manifest.mtree> --path <directory>

Validate "<directory>" against the mtree(8) manifest in "<manifest.mtree>".

This tool is primarily intended for umoci's integration tests (go-mtree is
missing rootless support in its CLI), and should not be relied upon by users.`,

	// This is only really used for our tests.
	Category: "devtools",
	Hidden:   true,

	Flags: []cli.Flag{
		cli.StringFlag{
			Name:     "manifest,f,file", // for compatibility with gomtree
			Required: true,
			Usage:    "mtree manifest to validate",
		},
		cli.StringFlag{
			Name:     "path,p", // for compatibility with gomtree
			Required: true,
			Usage:    "root path the the mtree manifest is relative to",
		},
		// TODO: Do we need --directory-only / -d ?
	},

	Action: rawMtreeValidate,

	Before: func(ctx *cli.Context) error {
		if ctx.String("manifest") == "" {
			return errors.New("--manifest must be a valid path")
		}
		if ctx.String("path") == "" {
			return errors.New("--path must be a valid path")
		}

		fsEval := fseval.Default
		if ctx.Bool("rootless") {
			fsEval = fseval.Rootless
		}
		ctx.App.Metadata["fseval"] = fsEval

		return nil
	},
}))

func rawMtreeValidate(ctx *cli.Context) (Err error) { _ = "STUB: not implemented"; return nil }

//nolint:errcheck

// Figure out the set of keywords to use for the comparison.

// Just use the manifest keywords if none were specified.

// "type" is necessary for any comparison to make sense.

// NOTE: compareKeywords might contain keywords not in the manifest, but
// this is usually okay.

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
	"github.com/urfave/cli"
)

// foreachSubcommand runs the given closure on every command and (recursively)
// every subcommand, allowing you to apply filters to all commands and
// subcommands.
func foreachSubcommand(cmds []cli.Command, fn func(*cli.Command)) {
	_ = "STUB: not implemented"
	return
}

// uxHistory adds the full set of --history.* flags to the given cli.Command as
// well as adding relevant validation logic to the .Before of the command. The
// values will be stored in ctx.Metadata with the keys "--history.author",
// "--history.created", "--history.created_by", "--history.comment", with
// string values. If they are not set the value will be nil.
func uxHistory(cmd cli.Command) cli.Command { _ = "STUB: not implemented"; return *new(cli.Command) }

// --no-history is incompatible with other --history.* options.

// Include any old befores set.

// uxCompress adds the --compress flag to the given cli.Command as well as
// adding relevant validation logic to the .Before of the command. The value
// will be stored in ctx.Metadata["--compress"] as a string (or nil if --tag
// was not specified).
func uxCompress(cmd cli.Command) cli.Command { _ = "STUB: not implemented"; return *new(cli.Command) }

// Verify compression algorithm value.

// Include any old befores set.

// uxTag adds a --tag flag to the given cli.Command as well as adding relevant
// validation logic to the .Before of the command. The value will be stored in
// ctx.Metadata["--tag"] as a string (or nil if --tag was not specified).
func uxTag(cmd cli.Command) cli.Command { _ = "STUB: not implemented"; return *new(cli.Command) }

// Verify tag value.

// Include any old befores set.

// uxImage adds an --image flag to the given cli.Command as well as adding
// relevant validation logic to the .Before of the command. The values (image,
// tag) will be stored in ctx.Metadata["--image-path"] and
// ctx.Metadata["--image-tag"] as strings (both will be nil if --image is not
// specified).
func uxImage(cmd cli.Command) cli.Command { _ = "STUB: not implemented"; return *new(cli.Command) }

// Verify and parse --image.

// Verify directory value.

// Verify tag value.

// uxLayout adds an --layout flag to the given cli.Command as well as adding
// relevant validation logic to the .Before of the command. The value is stored
// in ctx.App.Metadata["--image-path"] as a string (or nil --layout was not set).
func uxLayout(cmd cli.Command) cli.Command { _ = "STUB: not implemented"; return *new(cli.Command) }

// Verify and parse --layout.

// Verify directory value.

func uxRootless(cmd cli.Command) cli.Command { _ = "STUB: not implemented"; return *new(cli.Command) }

func uxRemap(cmd cli.Command) cli.Command { _ = "STUB: not implemented"; return *new(cli.Command) }

// fetchMeta returns the requested metadata from the current [cli.Context] as
// the requested type. fetchMeta will panic if the type parameter T does not
// match the actual type of the metadata entry.
func fetchMeta[T any](ctx *cli.Context, metaName string) (T, bool) {
	_ = "STUB: not implemented"
	return *new(T), false
}

// programmer error

// mustFetchMeta is like fetchMeta except that it will also panic if the
// metadata is not present in the current [cli.Context].
func mustFetchMeta[T any](ctx *cli.Context, metaName string) T {
	_ = "STUB: not implemented"
	return *new(T)
}

// programmer error

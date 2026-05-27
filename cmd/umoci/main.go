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

// Package main is the cli implementation of umoci.
package main

import (
	"os"

	"github.com/apex/log"
	"github.com/urfave/cli"
)

const (
	usage = `umoci modifies Open Container images`

	// Categories used to automatically monkey-patch flags to commands.
	categoryLayout = "layout"
	categoryImage  = "image"
)

func printVersion(c *cli.Context) { _ = "STUB: not implemented"; return }

//nolint:errcheck // errors not relevant

// This string is matched against in some tests (namely the xattr
// masking tests) to detect whether the umoci binary can be used for
// those tests. Make sure to update umoci-is-test-binary if you change
// this string!
//nolint:errcheck // errors not relevant

//nolint:errcheck // errors not relevant
//nolint:errcheck // errors not relevant
//nolint:errcheck // errors not relevant

// Main is the underlying main() implementation. You can call this directly as
// though it were the command-line arguments of the umoci binary (this is
// needed for umoci's integration test hacks you can find in main_test.go).
func Main(args []string) error { _ = "STUB: not implemented"; return nil }

// Should _never_ be reached.

// In order to make the uxXyz wrappers not too cumbersome we automatically
// add them to images with categories set to categoryImage or
// categoryLayout. Monkey patching was never this neat.

// If an error is a permission based error, give a hint to the user
// that --rootless might help. We probably should only be doing this if
// we're an unprivileged user.

func main() {
	if err := Main(os.Args); err != nil {
		log.Fatalf("%v", err)
	}
}

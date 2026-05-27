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

package convert

import (
	"strings"

	"github.com/blang/semver/v4"
	rspec "github.com/opencontainers/runtime-spec/specs-go"
)

// FIXME: We currently use an unreleased version of the runtime-spec and so we
// have to modify the version string because OCI specifications use "-dev" as
// suffix for not-yet-released versions but in such a way that it produces
// incorrect behaviour. This is compounded with the fact that runtime-tools
// cannot handle any version other than the single version they were compiled
// with.
//
// For instance, 1.0.2-dev is the development version after the release of
// 1.0.2, but according to SemVer 1.0.2-dev should be considered older than
// 1.0.2 (it has a pre-release tag) -- the specs should be using 1.0.2+dev.
var curSpecVersion = semver.MustParse(strings.TrimSuffix(rspec.Version, "-dev"))

// Example returns an example spec file, used as a "good sane default".
// XXX: Really we should just use runc's directly.
func Example() rspec.Spec { _ = "STUB: not implemented"; return *new(rspec.Spec) }

// ToRootless converts a specification to a version that works with rootless
// containers. This is done by removing options and other settings that clash
// with unprivileged user namespaces.
func ToRootless(spec *rspec.Spec) error {
	_ = "STUB: not implemented"
	// Remove additional groups.
	return nil
}

// Remove networkns from the spec, as well as userns (for us to add it
// later without duplicates).

// Add userns to the spec.

// Fix up mounts.

// Ignore all mounts that are under /sys.

// Remove all gid= and uid= mappings.

// Add the sysfs mount as an rbind.

// NOTE: "type: bind" is silly here, see opencontainers/runc#2035.

// Add /etc/resolv.conf as an rbind.

// If /etc/resolv.conf doesn't exist (such as inside OBS), just log a
// warning and continue on. In the worst case, you'll just end up with
// a non-networked container.

// If we are using user namespaces, then we must make sure that we don't
// drop any of the CL_UNPRIVILEGED "locked" flags of the source "mount"
// when we bind-mount. The reason for this is that at the point when runc
// sets up the root filesystem, it is already inside a user namespace, and
// thus cannot change any flags that are locked.

// NOTE: "type: bind" is silly here, see opencontainers/runc#2035.

// Remove cgroup settings.

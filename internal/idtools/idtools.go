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

// Package idtools provides helpers for dealing with Linux ID mappings.
package idtools

import (
	rspec "github.com/opencontainers/runtime-spec/specs-go"
)

// ToHost translates a remapped container ID to an unmapped host ID using the
// provided ID mapping. If no mapping is provided, then the mapping is a no-op.
// If there is no mapping for the given ID an error is returned.
func ToHost(contID int, idMap []rspec.LinuxIDMapping) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// ToContainer takes an unmapped host ID and translates it to a remapped
// container ID using the provided ID mapping. If no mapping is provided, then
// the mapping is a no-op. If there is no mapping for the given ID an error is
// returned.
func ToContainer(hostID int, idMap []rspec.LinuxIDMapping) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Helper to return a uint32 from strconv.ParseUint type-safely.
func parseUint32(str string) (uint32, error) { _ = "STUB: not implemented"; return 0, nil }

// ParseMapping takes a mapping string of the form "container:host[:size]" and
// returns the corresponding rspec.LinuxIDMapping. An error is returned if not
// enough fields are provided or are otherwise invalid. The default size is 1.
func ParseMapping(spec string) (rspec.LinuxIDMapping, error) {
	_ = "STUB: not implemented"
	return *new(rspec.LinuxIDMapping), nil
}

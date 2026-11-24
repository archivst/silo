// Copyright (c) 2015-2025 MinIO, Inc.
//
// This file is part of MinIO Object Storage stack
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package cmd

import (
	"github.com/minio/minio/internal/config"
	"github.com/pgsty/silo-pkg/v3/env"
)

// globalIdleDrives is enabled with SILO_IDLE_DRIVES=on. In this mode the
// server never touches the data drives on its own initiative, so that they
// can stay powered down (or spun down) until a client or admin issues an
// operation, e.g. a backup or a restore.
//
// It switches off:
//   - the data scanner (usage, ILM, replication and heal scans)
//   - the periodic drive write/read health probe
//   - automatic drive healing monitoring
//   - periodic .trash and stale multipart upload cleanup
//   - periodic IAM reload from the drives
//
// Reconnecting drives that are offline stays active, so drives that are
// powered up later are picked up without a restart.
var globalIdleDrives = env.Get("SILO_IDLE_DRIVES", config.EnableOff) == config.EnableOn

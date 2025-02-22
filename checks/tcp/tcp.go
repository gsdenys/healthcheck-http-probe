// Copyright 2021 by the contributors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tcp

import (
	"net"
	"time"

	"github.com/gsdenys/healthcheck/checks"
)

// Dial returns a Check that checks TCP connectivity to the provided
// endpoint.
func Dial(addr string, timeout time.Duration) checks.Check {
	return func() error {
		conn, err := net.DialTimeout("tcp", addr, timeout)

		if err != nil {
			return err
		}

		return conn.Close()
	}
}

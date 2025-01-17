// Copyright The OpenTelemetry Authors
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

package exemplar

import (
	"testing"
)

func TestDrop(t *testing.T) {
	t.Run("Int64", ReservoirTest[int64](func(int) (Reservoir[int64], int) {
		return Drop[int64](), 0
	}))

	t.Run("Float64", ReservoirTest[float64](func(int) (Reservoir[float64], int) {
		return Drop[float64](), 0
	}))
}

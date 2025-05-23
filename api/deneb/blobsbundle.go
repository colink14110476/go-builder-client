// Copyright © 2023 Attestant Limited.
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

package deneb

import (
	"fmt"

	"github.com/goccy/go-yaml"
)

// BlobsBundle is the structure used to store the blobs bundle.
type BlobsBundle struct {
	Commitments [][48]byte     `ssz-max:"4096" ssz-size:"?,48"`
	Proofs      [][48]byte     `ssz-max:"4096" ssz-size:"?,48"`
	Blobs       [][131072]byte `ssz-max:"4096" ssz-size:"?,131072"`
}

// String returns a string version of the structure.
func (b *BlobsBundle) String() string {
	data, err := yaml.Marshal(b)
	if err != nil {
		return fmt.Sprintf("ERR: %v", err)
	}

	return string(data)
}

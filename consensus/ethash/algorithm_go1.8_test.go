// Copyright 2017 The go-sferum Authors
// This file is part of the go-sferum library.
//
// The go-sferum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-sferum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-sferum library. If not, see <http://www.gnu.org/licenses/>.

// +build go1.8

package ethash

import "testing"

// Tests whether the dataset size calculator works correctly by cross checking the
// hard coded lookup table with the value generated by it.
func TestSizeCalculations(t *testing.T) {
	var tests []uint64

	// Verify all the cache sizes from the lookup table
	defer func(sizes []uint64) { cacheSizes = sizes }(cacheSizes)
	tests, cacheSizes = cacheSizes, []uint64{}

	for i, test := range tests {
		if size := cacheSize(uint64(i*epochLength) + 1); size != test {
			t.Errorf("cache %d: cache size mismatch: have %d, want %d", i, size, test)
		}
	}
	// Verify all the dataset sizes from the lookup table
	defer func(sizes []uint64) { datasetSizes = sizes }(datasetSizes)
	tests, datasetSizes = datasetSizes, []uint64{}

	for i, test := range tests {
		if size := datasetSize(uint64(i*epochLength) + 1); size != test {
			t.Errorf("dataset %d: dataset size mismatch: have %d, want %d", i, size, test)
		}
	}
}

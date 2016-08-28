// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package goradix

import "testing"

// ----------------------- Benchmarks ------------------------ //

func BenchmarkRemoveNTS(b *testing.B) {
	rx := New(false)
	insertDataBytes(rx, sampleData3)
	sd3 := sampleData3()
	sdLen := len(sd3)
	tn := 0

	for i := 0; i < b.N; i++ {
		if tn == sdLen {
			rx = New(false)
			insertDataBytes(rx, sampleData3)
		}

		rx.RemoveBytes(randomBytes(sd3))

		tn++
	}
}

func BenchmarkRemoveTS(b *testing.B) {
	rx := New(true)
	insertDataBytes(rx, sampleData3)
	sd3 := sampleData3()
	sdLen := len(sd3)
	tn := 0

	for i := 0; i < b.N; i++ {
		if tn == sdLen {
			rx = New(true)
			insertDataBytes(rx, sampleData3)
		}

		rx.RemoveBytes(randomBytes(sd3))

		tn++
	}
}

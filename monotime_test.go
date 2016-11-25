// Copyright (C) 2016  Arista Networks, Inc.
// Use of this source code is governed by the Apache License 2.0
// that can be found in the COPYING file.

// Package monotime provides a fast monotonic clock source.
package monotime

import (
	"testing"
)

func TestNanoseconds(t *testing.T) {
	for i := 0; i < 100; i++ {
		t1 := Nanoseconds()
		t2 := Nanoseconds()
		if t1 >= t2 {
			t.Fatalf("t1=%d should have been strictly less than t2=%d", t1, t2)
		}
	}
}

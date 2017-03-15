// Derived from SciPy's special/c_misc/misc.h
// https://github.com/scipy/scipy/blob/master/scipy/special/c_misc/misc.h

// Copyright ©2017 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package internal

type objectiveFunc func(float64, []float64) float64

const (
	machEp = 1.0 / (1 << 53)
)

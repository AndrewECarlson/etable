// Copyright (c) 2019, The Emergent Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package clust

import "math"

// DistFunc is a clustering distance function that evaluates aggregate distance
// between nodes, given the indexes of leaves in a and b clusters
// which are indexs into an ntot x ntot similarity (distance) matrix smat.
type DistFunc func(aix, bix []int, ntot int, smat []float64) float64

// MinDist is the minimum-distance or single-linkage weighting function for comparing
// two clusters a and b, given by their list of indexes.
// ntot is total number of nodes, and smat is the square similarity matrix [ntot x ntot].
func MinDist(aix, bix []int, ntot int, smat []float64) float64 {
	md := math.MaxFloat64
	for _, ai := range aix {
		for _, bi := range bix {
			d := smat[ai*ntot+bi]
			if d < md {
				md = d
			}
		}
	}
	return md
}

// MaxDist is the maximum-distance or complete-linkage weighting function for comparing
// two clusters a and b, given by their list of indexes.
// ntot is total number of nodes, and smat is the square similarity matrix [ntot x ntot].
func MaxDist(aix, bix []int, ntot int, smat []float64) float64 {
	md := -math.MaxFloat64
	for _, ai := range aix {
		for _, bi := range bix {
			d := smat[ai*ntot+bi]
			if d > md {
				md = d
			}
		}
	}
	return md
}

// AvgDist is the average-distance or average-linkage weighting function for comparing
// two clusters a and b, given by their list of indexes.
// ntot is total number of nodes, and smat is the square similarity matrix [ntot x ntot].
func AvgDist(aix, bix []int, ntot int, smat []float64) float64 {
	md := 0.0
	n := 0
	for _, ai := range aix {
		for _, bi := range bix {
			d := smat[ai*ntot+bi]
			md += d
			n++
		}
	}
	if n > 0 {
		md /= float64(n)
	}
	return md
}

// ContrastDist computes the average between distance - average within distance
// for two clusters a and b, given by their list of indexes.
// avg between is average distance between all items in a & b versus all outside that.
// ntot is total number of nodes, and smat is the square similarity matrix [ntot x ntot].
func ContrastDist(aix, bix []int, ntot int, smat []float64) float64 {
	wd := AvgDist(aix, bix, ntot, smat)
	nab := len(aix) + len(bix)
	abix := append(aix, bix...)
	abmap := make(map[int]struct{}, ntot-nab)
	for _, ix := range abix {
		abmap[ix] = struct{}{}
	}
	oix := make([]int, ntot-nab)
	octr := 0
	for ix := 0; ix < ntot; ix++ {
		if _, has := abmap[ix]; !has {
			oix[octr] = ix
			octr++
		}
	}
	bd := AvgDist(abix, oix, ntot, smat)
	return bd - wd
}

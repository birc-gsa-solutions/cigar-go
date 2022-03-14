// You can create modules at this level and they will be
// interpreted as under module birc.au.dk, so to import
// package `gsa` you need `import "birc.au.dk/gsa"`

package gsa

import "strings"

// Extract the edit operations from a pairwise alignment.
//
//  Args:
//      p: The first row in the pairwise alignment.
//      q: The second row in the pairwise alignment.
//
//  Returns:
//      The two strings without gaps and the list of edit operations
//      as a string.
func GetEdits(p, q string) (gapFreeP, gapFreeQ, edits string) {
	if len(p) != len(q) {
		panic("p and q must have equal length")
	}

	edits_buf := make([]byte, len(p))
	for i := 0; i < len(p); i++ {
		switch {
		case p[i] == '-':
			edits_buf[i] = 'I'
		case q[i] == '-':
			edits_buf[i] = 'D'
		default:
			edits_buf[i] = 'M'
		}
	}

	gapFreeP = strings.ReplaceAll(p, "-", "")
	gapFreeQ = strings.ReplaceAll(q, "-", "")
	return gapFreeP, gapFreeQ, string(edits_buf)
}

//  Get the distance between p and the string that starts at x[i:]
//  using the edits.
//
//  Args:
//      p: The read string we have mapped against x
//      x: The longer string we have mapped against
//      i: The location where we have an approximative match
//      edits: The list of edits to apply, given as a string
//
//  Returns:
//      The distance from p to x[i:?] described by edits
func EditDist(p, x string, i int, edits string) int {
	j, dist := 0, 0
	for k := 0; k < len(edits); k++ {
		switch edits[k] {
		case 'I':
			dist++
			i++
		case 'D':
			dist++
			j++
		case 'M':
			if x[i] != p[j] {
				dist++
			}
			i++
			j++
		}
	}

	return dist
}

// You can create modules at this level and they will be
// interpreted as under module birc.au.dk, so to import
// package `gsa` you need `import "birc.au.dk/gsa"`

package gsa

// Align two sequences from a sequence of edits.
//
//  Args:
//      p: The first sequence to align.
//      q: The second sequence to align
//      edits: The list of edits to apply, given as a string
//
//  Returns:
//      The two rows in the pairwise alignment
func Align(p, q, edits string) (pRow, qRow string) {
	return LocalAlign(p, q, 0, edits)
}

// Align two sequences from a sequence of edits.
//
//  Args:
//      p: The first sequence to align
//      x: The second sequence to align; we only align locally
//      i: Start position of the alignment in x
//      edits: The list of edits to apply, given as a string
//
//  Returns:
//      The two rows in the pairwise alignment
func LocalAlign(p, x string, i int, edits string) (pRow, xRow string) {
	pr := make([]byte, len(edits))
	xr := make([]byte, len(edits))
	j, k := 0, 0
	for l, e := range edits {
		switch e {
		case 'M':
			pr[l] = p[j]
			xr[l] = x[i+k]
			j++
			k++
		case 'D':
			pr[l] = p[j]
			xr[l] = '-'
			j++

		case 'I':
			pr[l] = '-'
			xr[l] = x[i+k]
			k++
		}
	}
	return string(pr), string(xr)
}

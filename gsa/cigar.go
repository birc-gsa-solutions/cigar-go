// You can create modules at this level and they will be
// interpreted as under module birc.au.dk, so to import
// package `gsa` you need `import "birc.au.dk/gsa"`

package gsa

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Expand the compressed CIGAR encoding into the full list of edits.
//
//  Args:
//      cigar: A CIGAR string
//
//  Returns:
//      The edit operations the CIGAR string describes.
func CigarToEdits(cigar string) string {
	r := regexp.MustCompile(`\d+[MID]`)
	edits := []byte{}

	for _, op := range r.FindAllString(cigar, -1) {
		repeats, _ := strconv.Atoi(op[:len(op)-1])
		opcode := op[len(op)-1]

		for i := 0; i < repeats; i++ {
			edits = append(edits, opcode)
		}
	}

	return string(edits)
}

// Encode a sequence of edits as a CIGAR.
//
//  Args:
//      edits: A sequence of edit operations
//
//  Returns:
//      The CIGAR encoding of edits.
func EditsToCigar(edits string) string {
	var (
		res  = []string{}
		i, j int
	)

	for ; i < len(edits); i = j {
		for j = i + 1; j < len(edits) && edits[i] == edits[j]; j++ {
		}

		res = append(res, fmt.Sprintf("%d%s", j-i, string(edits[i])))
	}

	return strings.Join(res, "")
}

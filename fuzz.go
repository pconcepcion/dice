// +build gofuzz

package dice

// Fuzz is used by dvyukov/go-fuzz to detemrine valid inputs
// See https://github.com/dvyukov/go-fuzz and
// https://medium.com/@dgryski/go-fuzz-github-com-arolek-ase-3c74d5a3150c
func Fuzz(data []byte) int {
	de := string(data)
	// Initial focus on sort strings, this should be changed later
	if len(de) > 10 {
		return 0
	}
	sde := SimpleExpression{expressionText: de}
	err := sde.parse()
	//fmt.Printf("sde: %+v", sde)
	if err != nil {
		return 0
	}
	return 1
}

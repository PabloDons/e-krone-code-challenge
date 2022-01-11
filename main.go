package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	fmt.Println(testValidity("456-abc-123-def-3"))

}

/*  Test the validity of a string in the spec
 *
 *	returns boolean flag true if the given string complies with the format, or false if the string does not comply
 *
 *	Estimated time: 20min
 *  Used time: 21min
 */
func testValidity(str string) bool {
	items := strings.Split(str, "-")

	// Assumption of number format creates some caveats:
	// Accepts invalid numbers like 00123
	// Declines base prefixes like 0xaf and any numbers with base > 10
	for i := 0; i < len(items); i += 2 {
		match, _ := regexp.MatchString("[0-9]", items[i])
		if !match {
			return false
		}
	}

	// Only requirement is non-zero size
	// Accepts non-ascii characters
	for i := 1; i < len(items); i += 2 {
		if len(items[i]) == 0 {
			return false
		}
	}

	return true
}

/*  Takes a string in the given spec
 *  Returns the average of the number parts
 *
 *  Estimated time: 10min
 *  Used time: ...
 */
func averageNumber(str string) {

}

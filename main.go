package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(wholeStory("100-abc-200-def-350-g"))
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
 *  Used time: 6min
 */
func averageNumber(str string) float32 {
	items := strings.Split(str, "-")
	count := 0
	sum := 0

	for i := 0; i < len(items); i += 2 {
		count += 1
		num, _ := strconv.Atoi(items[i])
		sum += num
	}

	return float32(sum) / float32(count)
}

/*  Takes a string in the given spec
 *  Returns a text that is composed from all the text words separated by spaces
 *
 *  Estimated time: 5min
 *  Used time: 7min
 */
func wholeStory(str string) string {
	items := strings.Split(str, "-")
	var resList []string

	for i := 1; i < len(items); i += 2 {
		resList = append(resList, items[i])
	}

	return strings.Join(resList, " ")
}

/*  Takes a string in the given spec
 *  Returns some stats about the string
 *
 *	Estimated time: 25min
 *  Used time: ...
 */
func storyStats() (shortest string, longest string, meanLen int, sameLen []string) {

}

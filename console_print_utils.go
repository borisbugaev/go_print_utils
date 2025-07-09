package printutils

import (
	"fmt"
	"strings"
)

const C_UP string = "\x1b[1A"
const L_CLR string = "\x1b[2K"

func clear_lines(count int) {
	for range count {
		fmt.Printf("%s%s", C_UP, L_CLR)
	}
}

func print_lines(line string) int {
	fmt.Printf("%s", line)
	return strings.Count(line, "\n")
}

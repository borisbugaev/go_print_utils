package printutils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const C_UP string = "\x1b[1A"
const L_CLR string = "\x1b[2K"
const ABC_LEN int = 'Z' - 'A'

func Clear_Lines(count int) {
	for range count {
		fmt.Printf("%s%s", C_UP, L_CLR)
	}
}

func Print_Lines(lines string) int {
	fmt.Printf("%s", lines)
	return strings.Count(lines, "\n")
}

func mc_letter(index int) string {
	my_letter := ""
	if index > ABC_LEN {
		letter_mod := index % ABC_LEN
		letter_intdiv := index / ABC_LEN
		my_letter = fmt.Sprintf("%s%s", mc_letter(letter_intdiv), mc_letter(letter_mod))
	} else {
		my_letter = fmt.Sprintf("%c", 'A'+index)
	}
	return my_letter
}

func mc_input_filter(content string) string {
	content = strings.Trim(content, " \n")
	content = strings.ToUpper(content)
	return content
}

func Line_Select_MC(fields []string) string {
	index := 0
	my_fields := ""
	fields_map := map[string]string{}
	for _, field := range fields {
		letter := mc_letter(index)
		my_fields = fmt.Sprintf("%s%s> %s\n", my_fields, letter, field)
		fields_map[letter] = field
		index++
	}
	my_fields = fmt.Sprintf("%s>> ", my_fields)
	line_count := Print_Lines(my_fields)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	response_seq := strings.SplitSeq(mc_input_filter(scanner.Text()), " ")
	response_fields := ""
	for response := range response_seq {
		selected_field, ok := fields_map[response]
		if !ok {
			continue
		}
		response_fields = fmt.Sprintf("%s,%s", response_fields, selected_field)
	}
	response_fields = strings.Trim(response_fields, ",")
	if response_fields == "" {
		response_fields = "\a"
	}
	Clear_Lines(line_count)
	return response_fields
}

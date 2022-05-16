package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func readFromCli() string {
	reader := bufio.NewReader(os.Stdin)

	var inputString string
	for {

		input, readError := reader.ReadString(' ')
		inputWithoutSpace := strings.TrimSpace(input)

		if input == "" {
			continue
		}

		if readError == io.EOF {
			inputString = inputWithoutSpace
			break
		}
	}

	return inputString
}
func convertRuneToInt(r rune) (*int, error) {
	converted, convertedErr := strconv.Atoi(string(r))
	if convertedErr != nil {
		return nil, convertedErr
	}
	return &converted, nil
}
func main() {
	nepaliDigits := [...]rune{'०', '१', '२', '३', '४', '५', '६', '७', '८', '९'}

	input := readFromCli()

	var converted string
	for _, runeValue := range input {
		integer, _ := convertRuneToInt(runeValue)

		converted += string(nepaliDigits[*integer])
	}
	fmt.Println("converted", converted)
}

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	sum := 0
	for {
		input, inputErr := reader.ReadString(' ')
		n := strings.TrimSpace(input)

		if input == "" {
			continue
		}
		intFormat, err := strconv.Atoi(n)
		if err != nil {
			fmt.Println("error!!!!")
		}
		sum += intFormat

		// ctrl+D will act as EOF on linux (my machine)
		if inputErr == io.EOF {
			break
		}
		//if input ==

		//fmt.Println(input)
	}
	fmt.Print("sum ", sum)

}

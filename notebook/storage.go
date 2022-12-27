package notebook

import (
	"bufio"
	"fmt"
	"os"
)

const HEAD string = "DATE;CODE;ACTION;AMOUNT;PRICE;EXTRA"

func ReadLinesFromFile(filepath string) []string {
	var lines []string
	f, err := os.Open(filepath)

	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if HEAD == scanner.Text() {
			continue
		}
		lines = append(lines, scanner.Text())
	}

	return lines
}

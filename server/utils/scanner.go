package utils

import (
	"bufio"
	"os"
	"strconv"
)

var scanner *bufio.Scanner

func init() {
	scanner = bufio.NewScanner(os.Stdin)
}

func Scan() string {
	scanner.Scan()
	return scanner.Text()
}

func ScanNumber() (number int, err error) {
	input := Scan()
	return strconv.Atoi(input)
}

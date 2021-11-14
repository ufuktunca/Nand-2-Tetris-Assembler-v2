package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	createInstruction := NewInstruction()

	file, err := os.Open("./test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if len(strings.Split(scanner.Text(), "//")[0]) > 0 {
			fmt.Println(scanner.Text())
			createInstruction.ConvertAssemblerToMachineCode(scanner.Text())
		}

	}

}

package main

import (
	constants "assembler/Constants"
	"assembler/helpers"
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	instructionWithoutSymbol := []string{}
	numberOfLabel := 0

	f, _ := os.Create("./Add2.hack")

	defer f.Close()
	w := bufio.NewWriter(f)

	createInstruction := NewInstruction()

	file, err := os.Open("./Pong.asm")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	instructions := []string{}

	for scanner.Scan() {
		if len(helpers.RemoveCommentFromInstruction(scanner.Text())) > 0 {
			instructions = append(instructions, helpers.RemoveCommentFromInstruction(scanner.Text()))
		}

	}

	for index, instruction := range instructions {
		if strings.Contains(instruction, "(") && strings.Contains(instruction, ")") {
			err := helpers.FeedSymbolTableWithLabels(strings.Split(strings.Split(instruction, ")")[0], "(")[1], index-numberOfLabel)
			if err != nil {
				panic(err)
			}
			numberOfLabel += 1
			continue
		}
		instructionWithoutSymbol = append(instructionWithoutSymbol, instruction)
	}

	for _, instruction := range instructionWithoutSymbol {
		if strings.Contains(instruction, "@") {
			_, err := strconv.Atoi(strings.Split(instruction, "@")[1])
			if err != nil && constants.SymbolTable[strings.Split(instruction, "@")[1]] == "" {
				err := helpers.FeedSymbolTableWithLabels(strings.Split(instruction, "@")[1], constants.SymbolTableVariableIndex)
				if err != nil {
					panic(err)
				}
				constants.SymbolTableVariableIndex += 1
			}
		}

	}

	for _, instruction := range instructionWithoutSymbol {
		bit16 := createInstruction.ConvertAssemblerToMachineCode(instruction)
		if len(bit16) > 0 {
			_, err := w.WriteString(bit16 + "\n")
			if err != nil {
				panic(err)
			}

			w.Flush()
		}
	}

}

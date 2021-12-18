package helpers

import (
	constants "assembler/Constants"
	"strconv"
	"strings"
)

func CompleteAddressTo16(binaryAddress string) string {
	for i := len(binaryAddress); i < 16; i++ {
		binaryAddress = "0" + binaryAddress
	}

	return binaryAddress
}

func DecideABit(assembly string) string {
	if strings.Contains(assembly, "M") {
		return "1"
	}
	return "0"
}

func FeedSymbolTableWithLabels(instruction string, index int) error {
	if constants.SymbolTable[instruction] == "" {
		constants.SymbolTable[instruction] = strconv.Itoa(index)
	}
	return nil
}

func RemoveCommentFromInstruction(instruction string) string {
	return strings.TrimSpace(strings.Split(instruction, "//")[0])
}

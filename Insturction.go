package main

import (
	"strconv"
	"strings"
)

type CreateInstructıon struct {
}

func NewInstruction() *CreateInstructıon {
	return &CreateInstructıon{}
}

func (c *CreateInstructıon) ConvertAssemblerToMachineCode(assembly string) string {
	if strings.Contains(assembly, "@") {
		instruction, err := c.CreateAInstruction(assembly)

		if err != nil {
			return ""
		}

		return instruction
	}

	return ""
}

func (c *CreateInstructıon) CreateAInstruction(assembly string) (string, error) {
	address, err := strconv.Atoi(strings.Split(assembly, "@")[1])

	if err != nil {
		return "", nil
	}
	binaryAddress := strconv.FormatInt(int64(address), 2)

	return helpers.CompleteAddressTo16(binaryAddress), nil
}

package main

import (
	constants "assembler/Constants"
	"assembler/helpers"
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

	if strings.Contains(assembly, "=") {
		return c.CreateCInstruction(assembly)
	}

	if strings.Contains(assembly, ";") {
		return c.CreateJump(assembly)
	}

	return ""
}

func (c *CreateInstructıon) CreateAInstruction(assembly string) (string, error) {
	var address string
	var addressVariable int
	var err error

	addressVariable, err = strconv.Atoi(strings.Split(assembly, "@")[1])

	if err != nil {
		address = constants.SymbolTable[strings.Split(assembly, "@")[1]]
		addressVariable, err = strconv.Atoi(address)
		if err != nil {
			return "", nil
		}
	}

	if err != nil {
		return "", nil
	}
	binaryAddress := strconv.FormatInt(int64(addressVariable), 2)

	return helpers.CompleteAddressTo16(binaryAddress), nil
}

func (c *CreateInstructıon) CreateCInstruction(assembly string) string {

	splittedAssembly := strings.Split(assembly, "=")

	return "111" + helpers.DecideABit(splittedAssembly[1]) + constants.C[splittedAssembly[1]] + constants.D[splittedAssembly[0]] + "000"
}

func (c *CreateInstructıon) CreateJump(assembly string) string {
	splittedAssembly := strings.Split(assembly, ";")

	return "111" + helpers.DecideABit(splittedAssembly[0]) + constants.C[splittedAssembly[0]] + "000" + constants.J[splittedAssembly[1]]
}

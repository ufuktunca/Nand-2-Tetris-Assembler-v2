package main

import (
	"testing"
)

func Test_ConvertAssemblerToMachineCode(t *testing.T) {
	createInstruction := NewInstruction()

	result := createInstruction.ConvertAssemblerToMachineCode("@2")

	if result != "0000000000000010" {
		t.Errorf("Assembly doesnt convert properly")
	}
}

func Test_CreateAInstruction(t *testing.T) {
	createInstruction := NewInstruction()

	result, err := createInstruction.CreateAInstruction("@10")
	if result != "0000000000001010" && err == nil {
		t.Errorf("Assembly didn't create correct instruction for 10")
	}

	result, err = createInstruction.CreateAInstruction("@27")
	if result != "0000000000011011" && err == nil {
		t.Errorf("Assembly didn't create correct instruction for 27")
	}

	result, err = createInstruction.CreateAInstruction("@81")
	if result != "0000000001010001" && err == nil {
		t.Errorf("Assembly didn't create correct instruction for 27")
	}
}

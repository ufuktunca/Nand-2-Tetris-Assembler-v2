package main

import (
	constants "assembler/Constants"
	"assembler/helpers"
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

	result, err = createInstruction.CreateAInstruction("@")
	if err != nil {
		t.Errorf("Assembly didn't create correct instruction because there is no address")
	}
}

func Test_CreateCInstruction(t *testing.T) {
	createInstruction := NewInstruction()

	result := createInstruction.CreateCInstruction("D=A+1")
	if result != "1110110111010000" {
		t.Errorf("Assembly didn't create correct instruction for D=A+1")
	}

	result = createInstruction.CreateCInstruction("M=M+1")
	if result != "1111110111001000" {
		t.Errorf("Assembly didn't create correct instruction for M=M+1")
	}

	result = createInstruction.CreateCInstruction("A=D")
	if result != "1110001100100000" {
		t.Errorf("Assembly didn't create correct instruction for A=D")
	}
}

func Test_CreateJump(t *testing.T) {
	createInstruction := NewInstruction()

	result := createInstruction.CreateJump("0;JMP")
	if result != "1110101010000111" {
		t.Errorf("Assembly didn't create correct instruction for 0;JMP " + result)
	}

	result = createInstruction.CreateJump("D;JGT")
	if result != "1110001100000001" {
		t.Errorf("Assembly didn't create correct instruction for D;JGT " + result)
	}

	result = createInstruction.CreateJump("D;JLE")
	if result != "1110001100000110" {
		t.Errorf("Assembly didn't create correct instruction for D;JLE " + result)
	}

}

func Test_FeedSymbolTableWithLabels(t *testing.T) {
	err := helpers.FeedSymbolTableWithLabels("Test", 22)
	if err != nil || constants.SymbolTable["Test"] != "22" {
		t.Errorf("Error occur1")
	}

	err = helpers.FeedSymbolTableWithLabels("LOOP", 45)
	if err != nil || constants.SymbolTable["LOOP"] != "45" {
		t.Errorf("Error occur2")
	}

}

package constants

const INSTRUCTION_TYPE = 15

const COMP_1 = 11
const COMP_2 = 10
const COMP_3 = 9
const COMP_4 = 8
const COMP_5 = 7
const COMP_6 = 6

const DEST_1 = 5
const DEST_2 = 4
const DEST_3 = 3

const JUMP_1 = 2
const JUMP_2 = 1
const JUMP_3 = 0

var C = map[string]string{
	"0":   "101010",
	"1":   "111111",
	"-1":  "111010",
	"D":   "001100",
	"A":   "110000",
	"!D":  "001101",
	"!A":  "110001",
	"-D":  "001111",
	"-A":  "110011",
	"D+1": "011111",
	"A+1": "110111",
	"D-1": "001110",
	"A-1": "110010",
	"D+A": "000010",
	"D-A": "010011",
	"A-D": "000111",
	"D&A": "000000",
	"D|A": "010101",
	"M":   "110000",
	"!M":  "110001",
	"-M":  "110011",
	"M+1": "110111",
	"M-1": "110010",
	"D+M": "000010",
	"D-M": "010011",
	"M-D": "000111",
	"D&M": "000000",
	"D|M": "010101",
}

var D = map[string]string{
	"M":   "001",
	"D":   "010",
	"MD":  "011",
	"A":   "100",
	"AM":  "101",
	"AD":  "110",
	"AMD": "111",
	"0":   "000",
}

var J = map[string]string{
	"JGT": "001",
	"JEQ": "010",
	"JGE": "011",
	"JLT": "100",
	"JNE": "101",
	"JLE": "110",
	"JMP": "111",
}

var SymbolTable = map[string]string{
	"R0":     "0",
	"R1":     "1",
	"R2":     "2",
	"R3":     "3",
	"R4":     "4",
	"R5":     "5",
	"R6":     "6",
	"R7":     "7",
	"R8":     "8",
	"R9":     "9",
	"R10":    "10",
	"R11":    "11",
	"R12":    "12",
	"R13":    "13",
	"R14":    "14",
	"R15":    "15",
	"SCREEN": "16384",
	"KBD":    "24576",
	"SP":     "0",
	"LCL":    "1",
	"ARG":    "2",
	"THIS":   "3",
	"THAT":   "4",
}

var SymbolTableVariableIndex = 16

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var address = 96

	//Attempts to find and open a file, and will throw and error if a problem occurs.
	file, err := os.Open("imtest1_bin.txt") // For read access.
	if err != nil {
		log.Fatal(err)
	}

	//Reads file by lines and puts each line into an array of string type.
	readFile := bufio.NewScanner(file)
	readFile.Split(bufio.ScanLines)
	var instructions []string

	for readFile.Scan() {
		instructions = append(instructions, readFile.Text())
	}

	//fmt.Println(instructions)
	readFile.Close()

	//Verify the data retrieved is the same as the data in the file.
	for _, line := range instructions {

		var op = line[0:11]
		var opcode string
		var instruct string
		var rm string

		if i, err := strconv.ParseInt(op, 2, 64); err != nil {
			fmt.Println(err)
		} else {
			opcode = strconv.Itoa(int(i))
		}

		fmt.Println(opcode)

		if opcode == "1104" {
			instruct = "AND"
			//fmt.Println(opcode)
		} else if opcode == "1112" {
			instruct = "ADD"
			//fmt.Println(opcode)
		} else if opcode == "1160" || opcode == "1161" {
			instruct = "ADDI"
			//fmt.Println(opcode)
		} else if opcode == "1360" {
			instruct = "ORR"
			//fmt.Println(opcode)
		} else if opcode == "1440" || opcode == "1447" {
			instruct = "CBZ"
			//fmt.Println(opcode)
		} else if opcode == "1448" || opcode == "1455" {
			instruct = "CBNZ"
			//fmt.Println(opcode)
		} else if opcode == "1624" {
			instruct = "SUB"
			//fmt.Println(opcode)
		} else if opcode == "1672" || opcode == "1673" {
			instruct = "SUBI"
			//fmt.Println(opcode)
		} else if opcode == "1684" || opcode == "1687" {
			instruct = "MOVZ"

			//fmt.Println(opcode)
		} else if opcode == "1940" || opcode == "1943" {
			instruct = "MOVK"
			//fmt.Println(opcode)
		} else if opcode == "1690" {
			instruct = "LSR"
			//fmt.Println(opcode)
		} else if opcode == "1691" {
			instruct = "LSL"
			//fmt.Println(opcode)
		} else if opcode == "1984" {
			instruct = "STUR"
			//fmt.Println(opcode)
		} else if opcode == "1986" {
			instruct = "LDUR"
			//fmt.Println(opcode)
		} else if opcode == "1692" {
			instruct = "ASR"
			//fmt.Println(opcode)
		} else if opcode == "1872" {
			instruct = "EOR"
			//fmt.Println(opcode)
		}

		fmt.Println(op + "\t" + strconv.Itoa(address) + "\t" + instruct + "\tR" + rm)

		address = address + 4
	}

}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func argConv(rm string) string {
	var newRm string

	if i, err := strconv.ParseInt(rm, 2, 64); err != nil {
		fmt.Println(err)
	} else {
		newRm = strconv.Itoa(int(i))
	}

	return newRm
}

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
		var strAddress string
		var opcode string
		var instruct string
		var rm string
		var rn string
		var rd string

		if i, err := strconv.ParseInt(op, 2, 64); err != nil {
			fmt.Println(err)
		} else {
			opcode = strconv.Itoa(int(i))
		}

		strAddress = strconv.Itoa(address)

		if opcode == "1104" {
			instruct = "AND"
			rm = argConv(line[11:16])
			rn = argConv(line[22:27])
			rd = argConv(line[27:32])

			var toPrint = fmt.Sprintf("%.11s %.5s %.6s %.5s %.5s", line[0:11], line[11:16], line[16:22], line[22:27], line[27:32])
			fmt.Println(toPrint + "\t" + strAddress + "\t" + instruct + "\tR" + rd + ", R" + rn + ", R" + rm)
		} else if opcode == "1112" {
			instruct = "ADD"
			rm = argConv(line[11:16])
			rn = argConv(line[22:27])
			rd = argConv(line[27:32])

			var toPrint = fmt.Sprintf("%.11s %.5s %.6s %.5s %.5s", line[0:11], line[11:16], line[16:22], line[22:27], line[27:32])
			fmt.Println(toPrint + "\t" + strAddress + "\t" + instruct + "\tR" + rd + ", R" + rn + ", R" + rm)
		} else if opcode == "1160" || opcode == "1161" {
			instruct = "ADDI"
			//fmt.Println(opcode)
		} else if opcode == "1360" {
			instruct = "ORR"
			rm = argConv(line[11:16])
			rn = argConv(line[22:27])
			rd = argConv(line[27:32])

			var toPrint = fmt.Sprintf("%.11s %.5s %.6s %.5s %.5s", line[0:11], line[11:16], line[16:22], line[22:27], line[27:32])
			fmt.Println(toPrint + "\t" + strAddress + "\t" + instruct + "\tR" + rd + ", R" + rn + ", R" + rm)
		} else if opcode == "1440" || opcode == "1447" {
			instruct = "CBZ"
			//fmt.Println(opcode)
		} else if opcode == "1448" || opcode == "1455" {
			instruct = "CBNZ"
			//fmt.Println(opcode)
		} else if opcode == "1624" {
			instruct = "SUB"
			rm = argConv(line[11:16])
			rn = argConv(line[22:27])
			rd = argConv(line[27:32])

			var toPrint = fmt.Sprintf("%.11s %.5s %.6s %.5s %.5s", line[0:11], line[11:16], line[16:22], line[22:27], line[27:32])
			fmt.Println(toPrint + "\t" + strAddress + "\t" + instruct + "\tR" + rd + ", R" + rn + ", R" + rm)
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
			rm = argConv(line[11:16])
			rn = argConv(line[22:27])
			rd = argConv(line[27:32])

			var toPrint = fmt.Sprintf("%.11s %.5s %.6s %.5s %.5s", line[0:11], line[11:16], line[16:22], line[22:27], line[27:32])
			fmt.Println(toPrint + "\t" + strAddress + "\t" + instruct + "\tR" + rd + ", R" + rn + ", #")
		} else if opcode == "1691" {
			instruct = "LSL"
			rm = argConv(line[11:16])
			rn = argConv(line[22:27])
			rd = argConv(line[27:32])

			var toPrint = fmt.Sprintf("%.11s %.5s %.6s %.5s %.5s", line[0:11], line[11:16], line[16:22], line[22:27], line[27:32])
			fmt.Println(toPrint + "\t" + strAddress + "\t" + instruct + "\tR" + rd + ", R" + rn + ", #")
		} else if opcode == "1984" {
			instruct = "STUR"
			//fmt.Println(opcode)
		} else if opcode == "1986" {
			instruct = "LDUR"
			//fmt.Println(opcode)
		} else if opcode == "1692" {
			instruct = "ASR"
			rm = argConv(line[11:16])
			rn = argConv(line[22:27])
			rd = argConv(line[27:32])

			var toPrint = fmt.Sprintf("%.11s %.5s %.6s %.5s %.5s", line[0:11], line[11:16], line[16:22], line[22:27], line[27:32])
			fmt.Println(toPrint + "\t" + strAddress + "\t" + instruct + "\tR" + rd + ", R" + rn + ", R" + rm)
		} else if opcode == "1872" {
			instruct = "EOR"
			rm = argConv(line[11:16])
			rn = argConv(line[22:27])
			rd = argConv(line[27:32])

			var toPrint = fmt.Sprintf("%.11s %.5s %.6s %.5s %.5s", line[0:11], line[11:16], line[16:22], line[22:27], line[27:32])
			fmt.Println(toPrint + "\t" + strAddress + "\t" + instruct + "\tR" + rd + ", R" + rn + ", R" + rm)
		}

		address = address + 4
	}

}

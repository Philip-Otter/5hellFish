// This is the core system for interpreting 5hellFish code
// The 2xdropout 2024

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type subShell struct {
	pearl  string
	pShell *Shell

	nextSubShell *subShell
}

type Shell struct {
	pearl string
	meat  string
	wait  bool `default:"false"`

	isSubShell   bool `default:"false"`
	pParentShell *Shell
	pSubShell    *subShell

	nextShell *Shell
}

func main() {
	pearlMap := make(map[string]string)
	pearlMap["#"] = "COMM"
	pearlMap["~"] = "MAIN"
	pearlMap["~~"] = "FUNC"
	pearlMap["!"] = "DECL"
	pearlMap["!!"] = "GLOBAL"
	pearlMap["^"] = "CLASS"
	pearlMap["<"] = "IMPORT"
	pearlMap["|||"] = "BUFF"
	pearlMap["$"] = "COST"

	pearls := [][]string{{"COMM", "MAIN", "FUNC", "DECL", "GLOBAL", "CLASS", "IMPORT", "BUFF", "COST"}, {"#", "~", "~~", "!", "!!", "^", "<", "|||", "$"}}

	targetFile, err := os.Open("./Test/testCode.5F") // This will be removed at a later date
	if err != nil {
		fmt.Println("Failure:  ", err)
	}

	scanner := bufio.NewScanner(targetFile)
	scanner.Split(bufio.ScanLines)
	var fileLines []string

	for scanner.Scan() {
		fileLines = append(fileLines, scanner.Text())
	}

	for _, line := range fileLines { //In the future the index (currently "_" in this line) could be used for better error handling as it shares a relationship with the line number
		line = strings.TrimLeft(line, " ")
		symbolicPearlRegex, _ := regexp.Compile(`[{].+?([ ])`) // This regex may seem to be overly inclusive, but it will allow us to easily modify our mappings later when the interpreter is compiled from source.
		englishPearlRegex, _ := regexp.Compile(`[{](COMM|MAIN|FUNC|DECL|GLOBAL|CLASS|IMPORT|BUFF|COST)[ ]`)

		symbolicPearl := symbolicPearlRegex.FindString(line)
		symbolicPearl = strings.TrimLeft(symbolicPearl, "{")
		symbolicPearl = strings.TrimRight(symbolicPearl, " ")

		englishPearl := englishPearlRegex.FindString(line)
		englishPearl = strings.TrimLeft(englishPearl, "{")
		englishPearl = strings.TrimRight(englishPearl, " ")

		if englishPearl != "" {
			for _, pearl := range pearls[0] {
				if englishPearl == pearl {
					fmt.Println(englishPearl)
				}
			}
		} else if symbolicPearl != "" {
			for _, pearl := range pearls[1] {
				if symbolicPearl == pearl {
					fmt.Println(pearlMap[symbolicPearl])
				}
			}
		}
	}
}

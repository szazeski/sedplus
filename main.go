package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	// command | sed 's/apple/mango/g'

	sedN := ""
	flag.StringVar(&sedN, "n", "", "sed compatible n value ('-n 1p' gets first line)")
	//sedF := ""
	//flag.StringVar(&sedF, "f", "", "sed compatible f value (file to read sed commands from)")

	find := ""
	flag.StringVar(&find, "find", "", "find")
	findLine := ""
	flag.StringVar(&findLine, "find-line", "", "findLine looks for this string and replaces the entire line")
	replace := ""
	flag.StringVar(&replace, "replace", "", "replace")

	toLowercase := false
	flag.BoolVar(&toLowercase, "lowercase", false, "convert to lowercase")
	toUppercase := false
	flag.BoolVar(&toUppercase, "uppercase", false, "convert to uppercase")

	trim := false
	flag.BoolVar(&trim, "trim", false, "trim whitespace")

	outputAlphanumericOnly := false
	flag.BoolVar(&outputAlphanumericOnly, "alphanumeric", false, "output only alphanumeric characters")
	outputNumericOnly := false
	flag.BoolVar(&outputNumericOnly, "numeric", false, "output only numeric characters")
	outputAlphaOnly := false
	flag.BoolVar(&outputAlphaOnly, "alpha", false, "output only alphabetic characters")

	errorIfNotFound := false
	flag.BoolVar(&errorIfNotFound, "error-if-not-found", false, "error if find is not found")
	linesChanged := 0

	flag.Parse()

	stat, _ := os.Stdin.Stat()

	if (stat.Mode() & os.ModeCharDevice) != 0 {
		// todo check if a file input was set instead of stdin

		showHelpText("")
		os.Exit(0)
	}

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		output := string(scanner.Bytes())

		if toLowercase {
			output = strings.ToLower(output)
		}
		if toUppercase {
			output = strings.ToUpper(output)
		}

		if find != "" {
			if replace != "" {
				if strings.Contains(output, find) {
					linesChanged++
					output = strings.ReplaceAll(output, find, replace)
				}
			} else {
				fmt.Println("missing --replace 'value'")
			}
		}

		if findLine != "" {
			if replace != "" {
				if strings.Contains(output, findLine) {
					linesChanged++
					output = replace
				}
			} else {
				fmt.Println("missing --replace 'value'")
			}
		}

		if outputNumericOnly {
			output = numericOnly(output)
		}

		// todo regex

		if trim {
			output = strings.TrimSpace(output)
		}

		fmt.Println(output)

		if errorIfNotFound && linesChanged == 0 {
			os.Exit(1)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func numericOnly(value string) string {
	regexOutput, _ := regexp.Compile("[^0-9]+")
	value = string(regexOutput.ReplaceAll([]byte(value), []byte("")))
	return value
}

func showHelpText(errorLine string) {
	if errorLine != "" {
		fmt.Println(errorLine)
	}
	fmt.Println("sedplus (an easier to use sed-like tool for Stream EDiting)")
	fmt.Println("        --find 'apple' --replace 'orange' --errorIfNotFound")
	fmt.Println("        --findLine 'apple' --replace 'mango'")
	fmt.Println("        --regex 's/apple/mango/g'")
	fmt.Println("        --alphanumericOnly")
	fmt.Println("        --toLowercase")
	fmt.Println("        --trim")
	fmt.Println("        --removeDoubleQuotes")
	fmt.Println("        --blurTimestamps")
}

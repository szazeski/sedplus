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
	caseInsensitive := false
	flag.BoolVar(&caseInsensitive, "case-insensitive", false, "case insensitive for find")

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

	compactWhitespace := false
	flag.BoolVar(&compactWhitespace, "compact-whitespace", false, "replace multiple spaces with a single space")

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

		if trim {
			output = strings.TrimSpace(output)
		}
		if toLowercase {
			output = strings.ToLower(output)
		}
		if toUppercase {
			output = strings.ToUpper(output)
		}

		if find != "" {
			if replace == "" {
				fmt.Println("missing --replace 'value'")
				os.Exit(1)
			}
			if contains(output, find, caseInsensitive) {
				linesChanged++
				re := regexp.MustCompile(`(?i)` + find)
				output = re.ReplaceAllString(output, replace)
			}
		}

		if findLine != "" {
			if replace == "" {
				fmt.Println("missing --replace 'value'")
				os.Exit(1)
			}
			if contains(output, findLine, caseInsensitive) {
				linesChanged++
				output = replace
			}
		}

		if outputNumericOnly {
			output = numericOnly(output)
		}
		if outputAlphaOnly {
			output = alphaOnly(output)
		}
		if outputAlphanumericOnly {
			output = alphanumericOnly(output)
		}

		if compactWhitespace {
			output = strings.ReplaceAll(output, "  ", " ")
		}

		fmt.Println(output)
	} // end for

	if errorIfNotFound && linesChanged == 0 {
		os.Exit(1)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func contains(input string, find string, caseInsensitive bool) bool {
	if caseInsensitive {
		input = strings.ToLower(input)
		find = strings.ToLower(find)
	}
	return strings.Contains(input, find)
}

func numericOnly(value string) string {
	regexOutput, _ := regexp.Compile("[^0-9]+")
	value = string(regexOutput.ReplaceAll([]byte(value), []byte("")))
	return value
}
func alphaOnly(value string) string {
	regexOutput, _ := regexp.Compile("[^a-zA-Z]+")
	value = string(regexOutput.ReplaceAll([]byte(value), []byte("")))
	return value
}
func alphanumericOnly(value string) string {
	regexOutput, _ := regexp.Compile("[^a-zA-Z0-9]+")
	value = string(regexOutput.ReplaceAll([]byte(value), []byte("")))
	return value
}
func replaceTimestamps(value string) string {
	regexOutput, _ := regexp.Compile("[0-9]{2}:[0-9]{2}:[0-9]{2}[ ]+")
	value = string(regexOutput.ReplaceAll([]byte(value), []byte("")))
	return value
}

func showHelpText(errorLine string) {
	if errorLine != "" {
		fmt.Println(errorLine)
	}
	fmt.Println("sedplus (an easier to use sed-like tool for Stream EDiting)")
	fmt.Println("        --find 'apple' --replace 'orange' --errorIfNotFound")
	fmt.Println("        --find-line 'apple' --replace 'mango'")
	fmt.Println("        --regex 's/apple/mango/g'")
	fmt.Println("        --alpha")
	fmt.Println("        --numeric")
	fmt.Println("        --lowercase")
	fmt.Println("        --uppercase")
	fmt.Println("        --trim")
	fmt.Println("        --compact-whitespace")
}

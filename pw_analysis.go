package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"text/tabwriter"
)

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func checkVeryWeak(word string, dictionary string) string {

	// Open the file
	f, err := os.Open(dictionary)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		pwd := scanner.Text()
		r := strings.Compare(word, pwd)

		if r == 0 {
			return "matched"
		}

	}
	return ""
}

func passwordAnalysis(pwdList string, dicList string) {

	// Initiate Regexes
	lowCase, err := regexp.Compile(`[a-z]`)                                  // Checks for lower case substring
	upCase, err := regexp.Compile(`[A-Z]`)                                   // Checks for upper case substring
	numbers, err := regexp.Compile(`[0-9]`)                                  // Checks for digit substring
	symbols, err := regexp.Compile(`[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?~]`) // Checks for symbol substring
	whitespace, err := regexp.Compile(`\S+`)                                 // Checks for whitespace substring
	var strength = ""
	smallerPwdSize := 1000
	biggerPwdSize := 0
	var averagePwd = []int{}
	nExcelent := 0
	nVStrong := 0
	nStrong := 0
	nVWeak := 0
	nWeak := 0

	fmt.Println("[+] Password Analysis Software\n")
	fmt.Println("Starting password analysis of: " + pwdList)
	fmt.Printf("Using %s as dictionary of comparison.\n\n", dicList)

	// Open the file
	f, err := os.Open(pwdList)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	// Read line by line
	scanner := bufio.NewScanner(f)

	writer := tabwriter.NewWriter(os.Stdout, 40, 8, 1, '\t', tabwriter.AlignRight)
	fmt.Fprintln(writer, "\nPASSWORD\tLENGTH\tSTRENGTH\n")

	for scanner.Scan() {
		pwd := scanner.Text()
		pwdSize := len(scanner.Text())

		averagePwd = append(averagePwd, pwdSize)

		if smallerPwdSize > pwdSize {
			smallerPwdSize = pwdSize
		}

		if biggerPwdSize < pwdSize {
			biggerPwdSize = pwdSize
		}

		if (lowCase.MatchString(pwd) && upCase.MatchString(pwd) && numbers.MatchString(pwd) && symbols.MatchString(pwd) && whitespace.MatchString(pwd) && pwdSize > 19) || (lowCase.MatchString(pwd) && upCase.MatchString(pwd) && numbers.MatchString(pwd) && symbols.MatchString(pwd) && pwdSize > 24) {
			strength = "Excelent"
			nExcelent += 1
		} else if (lowCase.MatchString(pwd) && upCase.MatchString(pwd) && numbers.MatchString(pwd) && symbols.MatchString(pwd) && whitespace.MatchString(pwd) && pwdSize > 11) || (lowCase.MatchString(pwd) && upCase.MatchString(pwd) && numbers.MatchString(pwd) && symbols.MatchString(pwd) && pwdSize > 14) {
			strength = "Very Strong"
			nVStrong += 1
		} else if (lowCase.MatchString(pwd) && upCase.MatchString(pwd) && numbers.MatchString(pwd) && symbols.MatchString(pwd) && pwdSize > 7) || (lowCase.MatchString(pwd) && upCase.MatchString(pwd) && numbers.MatchString(pwd) && symbols.MatchString(pwd) && whitespace.MatchString(pwd) && pwdSize > 7) {
			strength = "Strong"
			nStrong += 1
		} else {
			if checkVeryWeak(pwd, dicList) != "" {
				strength = "Very Weak"
				nVWeak += 1
			} else {
				strength = "Weak"
				nWeak += 1
			}
		}

		fmt.Fprintln(writer, pwd, "\t", pwdSize, "\t", strength)
		writer.Flush()
	}

	averagePwdSize := sum(averagePwd) / len(averagePwd)

	fmt.Printf("\n\nStatistics:\n- Total passwords analysed: %d.\n- Number of Excelent Passwords: %d.\n- Number of Very Strong Passwords: %d.\n- Number of Strong Passwords: %d.\n- Number of Weak Passwords: %d.\n- Number of Very Weak Passwords: %d.\n", len(averagePwd), nExcelent, nVStrong, nStrong, nWeak, nVWeak)
	fmt.Printf("- The smaller password has the size of %d characters.\n- The biggest password has the size of %d characters.\n- The average size is %d.", smallerPwdSize, biggerPwdSize, averagePwdSize)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func main() {

	if len(os.Args[1:]) != 2 {
		fmt.Println("[+] Usage: ./pwdAnalysis <passwords to be analysed>.txt <dictionary of weak passwords>.txt")
		os.Exit(0)
	}

	passwordAnalysis(os.Args[1], os.Args[2])

}

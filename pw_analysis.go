package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/Xe/x/entropy"
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

	var strength = ""
	smallerPwdSize := 1000
	biggerPwdSize := 0
	var averagePwd = []int{}
	nExcelent := 0
	nStrong := 0
	nWeak := 0
	nBadD := 0
	nBad := 0
	nPoor := 0

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
	fmt.Fprintln(writer, "\nPASSWORD\tLENGTH\tENTROPY\tSTRENGTH\n")

	for scanner.Scan() {
		pwd := scanner.Text()
		pwdSize := len(scanner.Text())

		averagePwd = append(averagePwd, pwdSize)

		entropyScore := entropy.Shannon(pwd)

		if smallerPwdSize > pwdSize {
			smallerPwdSize = pwdSize
		}

		if biggerPwdSize < pwdSize {
			biggerPwdSize = pwdSize
		}

		if entropyScore >= 100 {
			strength = "Excelent"
			nExcelent += 1
		} else if entropyScore > 65 {
			strength = "Strong"
			nStrong += 1
		} else {
			if checkVeryWeak(pwd, dicList) != "" {
				strength = "Bad - Dictionary"
				nBadD += 1
			} else if entropyScore > 40 {
				strength = "Weak"
				nWeak += 1
			} else if entropyScore <= 0 {
				strength = "Bad"
				nBad += 1
			} else {
				strength = "Poor"
				nPoor += 1
			}
		}

		fmt.Fprintln(writer, pwd, "\t", pwdSize, "\t", entropyScore, "\t", strength)
		writer.Flush()
	}

	averagePwdSize := sum(averagePwd) / len(averagePwd)

	fmt.Printf("\n\nStatistics:\n- Total passwords analysed: %d.\n- Number of excelent passwords: %d.\n- Number of strong passwords: %d.\n- Number of weak passwords: %d.\n- Number of poor passwords: %d.\n- Number of bad passwords: %d.\n- Number of bad passwords contained in the dictionary: %d.\n", len(averagePwd), nExcelent, nStrong, nWeak, nPoor, nBad, nBadD)
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

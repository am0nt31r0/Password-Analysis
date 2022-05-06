# Password-Analysis
Script in Go that analyzes a list of passwords based in its entropy and in weak passwords from a dictionary. Useful for penetration tests and phishing campaigns.

## Considerations
The script requires two txt files to execute successfully. The first one contains the passwords to be analysed - one password per line - and the second must be a dictionary of weak passwords, like rockyou.txt.

The entropy score used to classify the passwords is the same that is used by KeePassXC:
https://keepassxc.org/blog/2020-08-15-keepassxc-password-healthcheck/

Download the entropy source code to your src directory:
- https://github.com/Xe/x/tree/master/entropy (Thanks Xe for the Shannon entropy equation implementation)
- For example:
```
└─$ tree src 
src
├── github.com
│   └── Xe
│       └── x
│           └── entropy
│               ├── doc.go
│               ├── shannon.go
│               └── shannon_test.go
```

## Download and Compilation
```
git clone https://github.com/am0nt31r0/Password-Analysis.git
cd Password-Analysis
go build -ldflags "-w -s" -o pwdAnalysis
```
- Move the binary to bin directory.
```
└─$ tree go
go/
├── bin
│   ├── pwdAnalysis
├── src
|   ├── github.com
│       └── Xe
│           └── x
│               └── entropy
│                   ├── doc.go
│                   ├── shannon.go
│                   └── shannon_test.go
|   ├── Password Analysis
|       ├── pw_analysis.go
```

## Usage
```
pwdAnalysis pwds.txt rockyou.txt                                  
[+] Password Analysis Software

Starting password analysis of: pwds.txt
Using rockyou.txt as dictionary of comparison.


PASSWORD                                LENGTH                                  ENTROPY                                 STRENGTH

12345 a                                  7                                       21                                      Poor
user123                                  7                                       21                                      Bad - Dictionary
passwordstrong                           14                                      56                                      Weak
thisIsit?                                9                                       27                                      Poor
saQWE1                                   6                                       18                                      Poor
sa#X#2 34                                9                                       27                                      Poor
QWE123!"# $sds1                          15                                      60                                      Weak
123456789                                9                                       36                                      Bad - Dictionary
QWE123!"# $sds1zxczxc4"3r1 a             28                                      140                                     Excelent


Statistics:
- Total passwords analysed: 9.
- Number of excelent passwords: 1.
- Number of strong passwords: 0.
- Number of weak passwords: 2.
- Number of poor passwords: 4.
- Number of bad passwords: 0.
- Number of bad passwords contained in the dictionary: 2.
- The smaller password has the size of 6 characters.
- The biggest password has the size of 28 characters.
- The average size is 11. 
```

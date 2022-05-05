# Password-Analysis
Go script to analyze a list of passwords. Useful for penetration tests and phishing campaigns.

## Considerations
The script requires two txt files to execute successfully. The first one contains the passwords to be analysed - one password per line - and the second must be a dictionary of weak passwords, like rockyou.txt.

## Compilation
```
go build -ldflags "-w -s" -o pwdAnalysis
```

## Usage
```
pwdAnalysis pwds.txt rockyou.txt                                  
[+] Password Analysis Software

Starting password analysis of: pwds.txt
Using rockyou.txt as dictionary of comparison.


PASSWORD                                LENGTH                                  STRENGTH

12345 a                                  7                                       Weak
user123                                  7                                       Very Weak
passwordstrong                           14                                      Weak
thisIsit?                                9                                       Weak
saQWE1                                   6                                       Weak
sa#X#2 34                                9                                       Strong
QWE123!"# $sds1                          15                                      Very Strong
123456789                                9                                       Very Weak
QWE123!"# $sds1zxczxc4"3r1 a             28                                      Excelent


Statistics:
- Total passwords analysed: 9.
- Number of Excelent Passwords: 1.
- Number of Very Strong Passwords: 1.
- Number of Strong Passwords: 1.
- Number of Weak Passwords: 4.
- Number of Very Weak Passwords: 2.
- The smaller password has the size of 6 characters.
- The biggest password has the size of 28 characters.
- The average size is 11.   
```

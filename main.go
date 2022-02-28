package main

import "fmt"

func main() {
	var stages = [5]string{"KYL REV KORN HS ---\\\\ \\\\__/ _________________/---\n", "REV KORN ---\\\\ \\\\__/ _________________/--- KYL HS\n", "KORN ---\\\\ \\\\__/ _________________/--- KYL REV HS\n", "KYL ---\\\\ \\\\__/ _________________/--- HS KORN REV\n", " ---\\\\ \\\\__/ _________________/--- HS KORN REV KYL"}

	fmt.Println(stages, "\nDONE!")
}

func Crossing() string {
	return "KYL REV KORN HS ---\\\\ \\\\__/ _________________/---\n REV KORN ---\\\\ \\\\__/ _________________/--- KYL HS\n KORN ---\\\\ \\\\__/ _________________/--- KYL REV HS\n KYL ---\\\\ \\\\__/ _________________/--- HS KORN REV\n ---\\\\ \\\\__/ _________________/--- HS KORN REV KYL"
}
package main

import "testing"

func TestCrossing(t *testing.T) {
	wanted := "KYL REV KORN HS ---\\\\ \\\\__/ _________________/---\n REV KORN ---\\\\ \\\\__/ _________________/--- KYL HS\n KORN ---\\\\ \\\\__/ _________________/--- KYL REV HS\n KYL ---\\\\ \\\\__/ _________________/--- HS KORN REV\n ---\\\\ \\\\__/ _________________/--- HS KORN REV KYL"

	stages := Crossing()

	if stages != wanted {
		t.Error("Test failed!\n Got:\n", stages, "\nbut wanted:\n", wanted)
	}
}
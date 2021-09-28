package mascot_test

import (
	"testing"

	mascot "github.com/avinashrulz/mygocodes/Mascot"
)

func TestMascot(t *testing.T) {
	if mascot.BestMascot() != "Avinash" {
		t.Fatal("Wrong Mascot :( ")
	}
}

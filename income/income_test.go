package income

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_incomeCalc(t *testing.T) {
	stdOut := os.Stdout
	r, w, _ := os.Pipe()

	os.Stdout = w
	IncomeCalc()
	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)
	os.Stdout = stdOut

	if !strings.Contains(output, "$34320.00") {
		t.Errorf("\n Expected result: $34320.00, instead got: %s", output)
	}
}

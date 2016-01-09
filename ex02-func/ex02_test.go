package ex02

import "testing"

func TestSum(t *testing.T) {
	actual := sum(4, 7)
	if actual != 11 {
		t.Fatalf("Incorrect result for 4+7: wanted 11, got %d", actual)
	}
}

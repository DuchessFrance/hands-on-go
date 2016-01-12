package ex02

import "testing"

func TestSum(t *testing.T) {
	actual := sum(4, 7)
	if actual != 11 {
		t.Fatalf("Incorrect result for 4+7: wanted 11, got %d", actual)
	}
}

func TestPredSucc(t *testing.T) {
	prev, next := predSucc(4)
	if prev != 3 {
		t.Fatalf("Incorrect result for predSucc.prev: wanted 3, got %d", prev)
	}
	if next != 5 {
		t.Fatalf("Incorrect result for predSucc.next: wanted 5, got %d", next)
	}
}

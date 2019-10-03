package luckyday

import "testing"

func doTestCase(t *testing.T, name string, winnings int, allInCount int, expected int) {
	rounds := solution(winnings, allInCount)
	if rounds == expected {
		t.Logf("[%v] SUCCESS: N=%v/K=%v => expected: %v, got: %v\n", name, winnings, allInCount, expected, rounds)
	} else {
		t.Errorf("[%v] FAILED: N=%v/K=%v => expected: %v, got: %v\n", name, winnings, allInCount, expected, rounds)
	}
}

func Test8_0(t *testing.T) {
	doTestCase(t, "8/0", 8, 0, 7)
}

func Test18_2(t *testing.T) {
	doTestCase(t, "18/2", 18, 2, 6)
}

func Test10_10(t *testing.T) {
	doTestCase(t, "10/10", 10, 10, 4)
}

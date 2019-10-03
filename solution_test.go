package luckyday

import "testing"

func doTestCase(t *testing.T, name string, winnings int, allInCount int, expected int) {
	rounds := solution(winnings, allInCount)
	logFunc := t.Errorf
	status := "FAILED"
	if rounds == expected {
		logFunc = t.Logf
		status = "SUCCESS"
	}
	logFunc("[%v] %v: N=%v/K=%v => expected: %v, got: %v\n", name, status, winnings, allInCount, expected, rounds)
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

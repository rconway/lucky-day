package luckyday

func solution(N int, K int) int {
	rounds := 0

	for N >= 2 {
		if K > 0 && N%2 == 0 {
			N = N / 2
			K--
		} else {
			N = N - 1
		}
		rounds++
	}

	return rounds
}

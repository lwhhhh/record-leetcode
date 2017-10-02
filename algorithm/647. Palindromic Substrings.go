package algorithm

type Pair struct {
	left  int
	right int
}

func countSubstrings(s string) int {
	queue := make(chan Pair, len(s)*2)

	var sum int
	go func() {
		for {
			select {
			case p := <-queue:
				sum++
				if p.left-1 >= 0 && p.right+1 < len(s) && s[p.left-1] == s[p.right+1] {
					queue <- Pair{p.left - 1, p.right + 1}
				}
			}
			if len(queue) == 0 {
				close(queue)
				break
			}
		}
	}()

	//go func() {
	for i := 0; i < len(s); i++ {
		queue <- Pair{left: i, right: i}
	}
	//}()

	//go func() {
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			queue <- Pair{left: i, right: i + 1}
		}
	}
	//}()

	return sum
}

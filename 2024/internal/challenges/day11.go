package challenges

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

type DayEleven struct{}

func run(stones []int, max int) int {
	type key struct {
		stone int
		bc    int
	}
	memo := make(map[key]int)
	var mmx sync.Mutex

	var blink func(int, int) int
	blink = func(stone, bc int) int {
		mmx.Lock()
		k := key{stone, bc}
		if count, exists := memo[k]; exists {
			mmx.Unlock()
			return count
		}
		mmx.Unlock()

		if bc == max {
			return 1
		}

		var count int
		if stone == 0 {
			count = blink(1, bc+1)
		} else if sa := fmt.Sprintf("%d", stone); len(sa)%2 == 0 {
			left, _ := strconv.Atoi(sa[:len(sa)/2])
			right, _ := strconv.Atoi(sa[len(sa)/2:])
			count = blink(left, bc+1) + blink(right, bc+1)
		} else {
			count = blink(stone*2024, bc+1)
		}

		mmx.Lock()
		memo[k] = count
		mmx.Unlock()

		return count
	}

	ch := make(chan int, len(stones))
	var wg sync.WaitGroup

	for i := 0; i < len(stones); i++ {
		wg.Add(1)

		go func(stone int) {
			defer wg.Done()
			ch <- blink(stone, 0)
		}(stones[i])
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	count := 0
	for c := range ch {
		count += c
	}

	return count
}

func (DayEleven) SolvePartOne(input []string) (int, error) {
	stones := []int{}
	for _, x := range strings.Split(input[0], " ") {
		val, _ := strconv.Atoi(x)
		stones = append(stones, val)
	}

	return run(stones, 25), nil
}

func (DayEleven) SolvePartTwo(input []string) (int, error) {
	stones := []int{}
	for _, x := range strings.Split(input[0], " ") {
		val, _ := strconv.Atoi(x)
		stones = append(stones, val)
	}

	return run(stones, 75), nil
}

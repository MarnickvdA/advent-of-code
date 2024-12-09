package challenges

import (
	"fmt"
	"strconv"
)

type DayNine struct{}

func (DayNine) SolvePartOne(input []string) (int, error) {
	var memory []string

	fillMemory(&input, &memory, nil)

	s, f := 0, len(memory)-1
	for s < f {
		if memory[s] != "." {
			s++
			continue
		}

		if memory[f] == "." {
			f--
			continue
		}

		// swap
		memory[s] = memory[f]
		memory[f] = "."
		s++
		f--
	}

	return checksum(memory), nil
}

type Space struct {
	free    []int
	indices []int
}

func (DayNine) SolvePartTwo(input []string) (int, error) {
	var memory []string
	spaces := Space{
		free:    []int{},
		indices: []int{},
	}
	maxId := fillMemory(&input, &memory, &spaces)
	tail := len(memory) - 1

ids:
	for id := maxId; id > 0; id-- {
		for memory[tail] != fmt.Sprintf("%d", id) {
			tail--
		}
		end := tail

		for memory[end] == memory[tail-1] {
			tail--
		}

		size := end - tail + 1

		for x := range spaces.free {
			if spaces.free[x] >= size && spaces.indices[x] < tail {
				for z := 0; z < size; z++ {
					memory[spaces.indices[x]+z] = memory[tail+z]
					memory[tail+z] = "."
				}

				spaces.free[x] -= size
				spaces.indices[x] += size
				continue ids
			}
		}
	}

	return checksum(memory), nil
}

func fillMemory(input *[]string, memory *[]string, spaces *Space) int {
	line := (*input)[0]

	isBlock := true
	id := 0
	index := 0
	for i, c := range line {
		v, _ := strconv.Atoi(string(c))

		var size int
		for size = 0; size < v; size++ {
			index++
			if isBlock {
				*memory = append(*memory, fmt.Sprint(id))
			} else {
				*memory = append(*memory, ".")
			}
		}

		if spaces != nil && !isBlock && size > 0 {
			spaces.free = append(spaces.free, size)
			spaces.indices = append(spaces.indices, index-size)
		}

		if isBlock {
			id++
		}

		isBlock = !isBlock
		i++
	}

	return id - 1
}

func checksum(m []string) int {
	sum := 0

	for i, v := range m {
		if v != "." {
			x, _ := strconv.Atoi(v)
			sum += i * x
		}
	}

	return sum
}

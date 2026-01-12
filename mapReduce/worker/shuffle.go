package worker

func Shuffle(mappedOutputs []map[string]int, shuffled *map[string][]int) {
	for _, mapped := range mappedOutputs {
		for word, count := range mapped {
			(*shuffled)[word] = append((*shuffled)[word], count)
		}
	}
}
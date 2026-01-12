package worker

import "sort"

type KV struct {
	Key string
	Val int
}

func Reducer(shuffled map[string][]int, reduced *map[string]int) {
	sortedKV := []*KV{}

	sum := func(arr []int) int {
		res := 0

		for _, n := range arr {
			res += n
		}

		return res
	}

	for word, counts := range shuffled {
		sortedKV = append(sortedKV, &KV{Key: word, Val: sum(counts)})
	}

	sort.Slice(sortedKV, func(i, j int) bool {
		return sortedKV[i].Val > sortedKV[j].Val
	})
	
	for _, kv := range sortedKV {
		(*reduced)[kv.Key] = kv.Val
	}
}

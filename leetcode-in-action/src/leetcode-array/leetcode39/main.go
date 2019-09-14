package main

import (
	"container/list"
	"fmt"
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	result := make([][]int, 0)
	oneResult := list.New()
	result = dfs(candidates, 0, oneResult, target, result)
	return result
}

func dfs(candidates []int, begin int, oneResult *list.List, surplus int, result [][]int) [][]int {
	if surplus == 0 {
		appendResult := make([]int, 0)

		for i := oneResult.Front(); i != nil; i = i.Next() {
			appendResult = append(appendResult, i.Value.(int))
		}
		result = append(result, appendResult)
		return result
	}
	if begin > len(candidates) || candidates[begin] > surplus {
		return result
	}
	for i := begin; i < len(candidates); i++ {
		if candidates[i] > surplus {
			return result
		}
		element := oneResult.PushBack(candidates[i])
		result = dfs(candidates, i, oneResult, surplus-candidates[i], result)
		oneResult.Remove(element)
	}
	return result
}

func main() {
	candidates := []int{2, 3, 6, 7}
	target := 7

	result := combinationSum(candidates, target)
	fmt.Println("%d", len(result))

	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result[i]); j++ {
			fmt.Print("%d, ", result[i][j])
		}
		fmt.Println()
	}
}

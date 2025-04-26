// https://leetcode.com/problems/find-all-possible-recipes-from-given-supplies/
package medium

import "math"

func FindAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
	n := len(recipes)
	// Map each supply to true
	avaiSupplies := make(map[string]bool, len(supplies))
	for _, s := range supplies {
		avaiSupplies[s] = true
	}

	// Map each recipe to its index
	recipeToIdx := make(map[string]int, n)
	for i := 0; i < n; i++ {
		recipeToIdx[recipes[i]] = i
	}

	// Build dependency graph from "ingredient recipe" -> "recipe that depends on it"
	// Only if ingredient is itself a recipe
	// inDegree[i] = how many recipe-ingredients the i-th recipe depends on
	depenGraph := make(map[string][]string)
	inDegree := make([]int, n)

	for i := 0; i < n; i++ {
		for _, ing := range ingredients[i] {
			// If this ingredient is a supply, it doesn't increase inDegree
			if avaiSupplies[ing] {
				continue
			}

			// If the ingredient is also a recipe, set up dependency
			if _, ok := recipeToIdx[ing]; ok {
				depenGraph[ing] = append(depenGraph[ing], recipes[i])
				inDegree[i]++
			} else {
				// If ingredient is neither a supply nor a known recipe, we can never make recipes[i]
				// Mark inDegree as "infinite" so it won't ever be processed
				inDegree[i] = math.MaxInt32
			}
		}
	}

	// Initialize queue with all recipes that have inDegree == 0
	queue := make([]int, 0)
	for i := 0; i < n; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	result := make([]string, 0)
	// Standard Kahn's Algorithm BFS for topological order
	for len(queue) > 0 {
		idx := queue[0]
		queue = queue[1:]
		recipe := recipes[idx]
		result = append(result, recipe)

		// If no one depends on this recipe, skip
		if depenGraph[recipe] == nil {
			continue
		}

		// Decrease inDegree of dependent recipes
		for _, depRecipe := range depenGraph[recipe] {
			depIdx := recipeToIdx[depRecipe]
			inDegree[depIdx]--
			if inDegree[depIdx] == 0 {
				queue = append(queue, depIdx)
			}
		}
	}

	return result
}

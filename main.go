package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/deckarep/golang-set"
)

func main() {
	dataset := [][]string{
		[]string{"pane", "cereali"},
		[]string{"pane", "latte", "caffè"},
		[]string{"latte", "cereali"},
		[]string{"latte", "caffè"},
		[]string{"pane", "latte", "cereali"},
		[]string{"latte", "cereali"},
		[]string{"pane", "cereali"},
		[]string{"pane", "latte", "te"},
		[]string{"pane", "latte", "cereali", "te"},
		[]string{"pane", "te"},
		[]string{"pane", "caffè"},
	}

	minimumSupport := 0.4
	minConfidence := 0.2

	apriori(dataset, minimumSupport, minConfidence)

}

func apriori(dataset [][]string, support float64, confidence float64) {

	elements := elements(dataset)
	freqSet := make(map[string]float64)
	largeSet := make(map[int]mapset.Set)
	oneCSet := returnItemsWithMinSupport(elements, dataset, support, &freqSet)
	currentLSet := oneCSet

	k := 2

	for currentLSet.Cardinality() != 0 {
		largeSet[k-1] = currentLSet
		currentLSet = joinSet(currentLSet, k)
		currentCSet := returnItemsWithMinSupport(currentLSet, dataset, support, &freqSet)
		currentLSet = currentCSet
		k = k + 1
	}
	fmt.Println(largeSet)
}

func returnItemsWithMinSupport(itemSet mapset.Set, dataset [][]string, minSupport float64, freqSet *map[string]float64) mapset.Set {

	localItemSet := mapset.NewSet()
	localSet := make(map[string]float64)

	for _, item := range itemSet.ToSlice() {
		dkey := strings.Split(item.(string), "-")
		sort.Strings(dkey)
		for _, line := range dataset {
			if contains(line, dkey) {
				key := strings.Join(dkey, "-")
				(*freqSet)[key] += 1.0
				localSet[key] += 1.0
			}
		}
	}

	for item, count := range localSet {
		support := count / float64(len(dataset))

		if support >= minSupport {
			localItemSet.Add(item)
		}
	}

	return localItemSet

}

func joinSet(itemSet mapset.Set, length int) mapset.Set {

	ret := mapset.NewSet()

	for _, i := range itemSet.ToSlice() {
		for _, j := range itemSet.ToSlice() {
			i := i.(string)
			j := j.(string)

			i_a := strings.Split(i, "-")
			j_a := strings.Split(j, "-")

			dkey := (union(i_a, j_a))
			if len(dkey) == length {
				sort.Strings(dkey)
				key := strings.Join(dkey, "-")
				ret.Add(key)

			}
		}
	}
	return ret
}

func union(a []string, b []string) []string {

	ret := mapset.NewSet()

	for _, v := range a {
		ret.Add(v)
	}
	for _, v := range b {
		ret.Add(v)
	}
	rets := []string{}
	for _, v := range ret.ToSlice() {
		rets = append(rets, v.(string))
	}
	return rets
}

func elements(dataset [][]string) mapset.Set {

	ret := mapset.NewSet()

	for i := 0; i < len(dataset); i++ {
		for j := 0; j < len(dataset[i]); j++ {
			if ret.Contains(dataset[i][j]) == false {
				ret.Add(dataset[i][j])
			}
		}
	}
	return ret
}

func contains_dataset(s [][]string, e []string) bool {
	ret := false
	for _, v := range s {
		ret = contains(v, e)
		if ret == true {
			break
		}
	}
	return ret
}

func contains_element(s []string, e string) bool {
	ret := false
	for _, a := range s {
		if a == e {
			ret = true
			break
		}
	}
	return ret
}

func contains(s []string, e []string) bool {
	count := 0
	if len(s) < len(e) {
		return false
	}
	mm := make(map[string]bool)
	for _, a := range e {
		mm[a] = true
	}

	for _, a := range s {
		if _, ok := mm[a]; ok {
			count += 1
		}
	}
	return count == len(e)
}

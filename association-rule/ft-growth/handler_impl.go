package ftgrowth

import (
	"fmt"
	"log"
	"math"
	"sort"
	"time"

	"github.com/pkg/profile"
)

// Item represents an item.
type Item int

// Process -
// Base on https://github.com/cpearce/arm-go
func (h handler) Process() (result []string) {
	defer profile.Start().Stop()

	itemizer, frequency, numMatches := countItems(h.dataset)

	log.Println("Generating frequent itemsets via fpGrowth")
	start := time.Now()
	itemsWithCount := generateFrequentItemsets(h.dataset, h.minimumSupport, itemizer, frequency, numMatches)
	log.Printf("fpGrowth generated %d frequent patterns in %s",
		len(itemsWithCount), time.Since(start))

	log.Println("Itemsets:")
	printItemsets(itemsWithCount, itemizer, numMatches)

	log.Println("Generating association rules...")
	start = time.Now()
	rules := generateRules(itemsWithCount, numMatches, h.minConfidence, h.minLift)
	numRules := countRules(rules)
	log.Printf("Generated %d association rules in %s", numRules, time.Since(start))

	printRules(rules, itemizer)
	for _, chunk := range rules {
		for _, rule := range chunk {
			first := true
			for _, item := range rule.Antecedent {
				if !first {
					fmt.Printf(" ")
				}
				first = false
				fmt.Print(itemizer.toStr(item))
			}
			fmt.Print(" => ")
			first = true
			for _, item := range rule.Consequent {
				if !first {
					fmt.Printf(" ")
				}
				first = false
				fmt.Print(itemizer.toStr(item))
			}
			fmt.Printf(",%f,%f,%f\n", rule.Confidence, rule.Lift, rule.Support)
		}
	}
	return nil
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func printItemsets(itemsets []itemsetWithCount, itemizer *Itemizer, numMatches int) {

	fmt.Println("Itemset,Support")
	n := float64(numMatches)
	for _, iwc := range itemsets {
		first := true
		for _, item := range iwc.itemset {
			if !first {
				fmt.Printf(" ")
			}
			first = false
			fmt.Print(itemizer.toStr(item))
		}
		fmt.Printf(" %f\n", float64(iwc.count)/n)
	}
}

func printRules(rules [][]Rule, itemizer *Itemizer) {

	fmt.Println("Antecedent => Consequent,Confidence,Lift,Support")
	for _, chunk := range rules {
		for _, rule := range chunk {
			first := true
			for _, item := range rule.Antecedent {
				if !first {
					fmt.Printf(" ")
				}
				first = false
				fmt.Print(itemizer.toStr(item))
			}
			fmt.Print(" => ")
			first = true
			for _, item := range rule.Consequent {
				if !first {
					fmt.Printf(" ")
				}
				first = false
				fmt.Print(itemizer.toStr(item))
			}
			fmt.Printf(",%f,%f,%f\n", rule.Confidence, rule.Lift, rule.Support)
		}
	}
}

func countRules(rules [][]Rule) int {
	n := 0
	for _, chunk := range rules {
		n += len(chunk)
	}
	return n
}

func countItems(dataSet [][]string) (*Itemizer, *itemCount, int) {
	frequency := makeCounts()
	itemizer := newItemizer()

	numMatches := 0
	for _, set := range dataSet {
		numMatches++
		itemizer.forEachItem(
			set,
			func(item Item) {
				frequency.increment(item, 1)
			})
	}

	return &itemizer, &frequency, numMatches
}

func generateFrequentItemsets(dataSet [][]string, minimumSupport float64, itemizer *Itemizer, frequency *itemCount, numMatches int) []itemsetWithCount {
	minCount := max(1, int(math.Ceil(minimumSupport*float64(numMatches))))

	tree := newTree()
	for _, set := range dataSet {
		match := itemizer.filter(
			set,
			func(i Item) bool {
				return frequency.get(i) >= minCount
			})

		if len(match) == 0 {
			continue
		}
		// Sort by decreasing frequency, tie break lexicographically.
		sort.SliceStable(match, func(i, j int) bool {
			a := match[i]
			b := match[j]
			if frequency.get(a) == frequency.get(b) {
				return itemizer.cmp(a, b)
			}
			return frequency.get(a) > frequency.get(b)
		})
		tree.Insert(match, 1)
	}

	return fpGrowth(tree, make([]Item, 0), minCount)
}

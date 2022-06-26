package main

import (
	"fmt"

	"github.com/LeandroLS/valoop"
)

func breadthFirstSearch() {
	graph := make(map[string][]string)
	graph["Leandro"] = []string{"Alice", "Bob", "Claire"}
	graph["Bob"] = []string{"Anuj", "Peggy"}
	graph["Alice"] = []string{"Peggy"}
	graph["Claire"] = []string{"Thom", "Jonny"}
	graph["Anuj"] = []string{}
	graph["Peggy"] = []string{}
	graph["Thom"] = []string{}
	graph["Jonny"] = []string{}
	var queue []string
	var verifiedQueue []string
	queue = append(queue, graph["Leandro"]...)
	for len(queue) >= 1 {
		person := queue[0]
		queue = queue[1:]

		personIsVerified := valoop.StrSliceContains(verifiedQueue, person)
		if !personIsVerified {
			if person == "Jonny" {
				fmt.Println("Jonny is seller")
			} else {
				fmt.Printf("%s is not seller\n", person)
				queue = append(queue, graph[person]...)
				verifiedQueue = append(verifiedQueue, person)
			}
		}
	}
}

func main() {
	breadthFirstSearch()
}

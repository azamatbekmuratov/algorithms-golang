package main

import (
	"fmt"
)

func enqueue(queue []string, element string) []string {
	queue = append(queue, element)
	// fmt.Println("Enqueued:", element)
	return queue
}

func dequeue(queue []string) []string {
	// element := queue[0]
	// fmt.Println("Dequeued: ", element)
	return queue[1:]
}

func main() {
	graph := make(map[string][]string)
	graph["you"] = []string{"kami", "alice", "bob", "azamat", "claire"}
	graph["bob"] = []string{"anuj", "peggy"}
	graph["alice"] = []string{"peggy"}
	graph["claire"] = []string{"thom", "jonny"}
	graph["anuj"] = []string{}
	graph["peggy"] = []string{}
	graph["thom"] = []string{}
	graph["jonny"] = []string{}

	var queue []string
	for key, _ := range graph {
		queue = enqueue(queue, key)
	}
	fmt.Println((queue))
	searched := []string{}

	for len(queue) > 0 {
		if queue[0] == "azamat " {
			fmt.Printf("Found the person to sell apple: %v\n", queue[0])
			break
		}
		for _, sVal := range searched {
			if sVal != queue[0] {
				for _, value := range graph[queue[0]] {
					if sVal != value {
						fmt.Printf("Adding person to queue: %v\n", value)
						queue = enqueue(queue, value)
					}
				}
			}
		}
		searched = append(searched, queue[0])
		queue = dequeue(queue)
	}
}

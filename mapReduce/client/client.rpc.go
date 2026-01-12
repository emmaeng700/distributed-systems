package main

import (
	"fmt"
	"log"
	"net/rpc"
	"sync"
)

type ChunkArgs struct {
	Input    []string
	WorkerId int
}

type ChunkResult struct {
	L, R int
}

type PhaseOneArgs struct {
	ChunkedInput []string
}

type PhaseOneReply struct {
	MappedOutputs []map[string]int
}

type PhaseTwoArgs struct {
	MappedOutputs  []map[string]int
	OutputFilePath string
}

type PhaseTwoReply struct{}

func main() {
	var wg sync.WaitGroup
	client, err := rpc.Dial("tcp", "localhost:8080")

	if err != nil {
		fmt.Println("Could not connect rpc")
		return
	}

	defer client.Close()

	const NUM_OF_WORKERS = 3
	inputFilepath := "/Users/fnuworsu/Distributed Systems/mapReduce/client/input.txt"
	var input []string

	err = client.Call("Coordinator.GetInput", &inputFilepath, &input)

	if err != nil {
		log.Fatal(err)
		return
	}

	mappedCh := make(chan map[string]int, NUM_OF_WORKERS)
	outputFilePath := "/Users/fnuworsu/Distributed Systems/mapReduce/client/result.txt"

	for workerId := 0; workerId < NUM_OF_WORKERS; workerId++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			args := &ChunkArgs{Input: input, WorkerId: id}
			var res ChunkResult

			err := client.Call("Coordinator.Chunk", args, &res)
			if err != nil {
				log.Fatal(err)
				return
			}

			p1Args := &PhaseOneArgs{ChunkedInput: input[res.L:res.R]}
			p1Reply := &PhaseOneReply{}
			err = client.Call("Coordinator.PhaseOneWorker", p1Args, p1Reply)
			if err != nil {
				log.Fatal(err)
				return
			}

			if len(p1Reply.MappedOutputs) > 0 {
				mappedCh <- p1Reply.MappedOutputs[0]
			} else {
				mappedCh <- map[string]int{}
			}
		}(workerId)
	}

	wg.Wait()
	close(mappedCh)

	var allMapped []map[string]int
	for m := range mappedCh {
		allMapped = append(allMapped, m)
	}

	p2Args := &PhaseTwoArgs{MappedOutputs: allMapped, OutputFilePath: outputFilePath}
	p2Reply := &PhaseTwoReply{}
	err = client.Call("Coordinator.PhaseTwoWorker", p2Args, p2Reply)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("All workers completed!")
}

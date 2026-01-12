package coordinator

import (
	"bufio"
	"fmt"
	"mapReduce/worker"
	"os"
	"strings"
)

const NUM_OF_WORKERS = 3

type Coordinator struct{}

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

func (coord *Coordinator) GetInput(filepath *string, result *[]string) error {
	file, err := os.Open(*filepath)

	if err != nil {
		return err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		*result = append(*result, strings.TrimSpace(line))
	}

	err = scanner.Err()

	if err != nil {
		return err
	}

	return nil
}

func (coord *Coordinator) Chunk(args *ChunkArgs, result *ChunkResult) error {
	inputLen := len(args.Input)
	chunk_size := inputLen / NUM_OF_WORKERS
	start := args.WorkerId * chunk_size
	end := 0

	if args.WorkerId == NUM_OF_WORKERS-1 {
		end = inputLen
	} else {
		end = start + chunk_size
	}

	*result = ChunkResult{L: start, R: end}
	return nil
}

func (coord *Coordinator) PhaseOneWorker(args *PhaseOneArgs, reply *PhaseOneReply) error {
	mapped := make(map[string]int)

	for _, line := range args.ChunkedInput {
		worker.Mapper(line, &mapped)
	}

	reply.MappedOutputs = append(reply.MappedOutputs, mapped)
	return nil
}

func (coord *Coordinator) PhaseTwoWorker(args *PhaseTwoArgs, reply *PhaseTwoReply) error {
	shuffled := make(map[string][]int)
	reduced := make(map[string]int)

	worker.Shuffle(args.MappedOutputs, &shuffled)
	worker.Reducer(shuffled, &reduced)

	file, err := os.Create(args.OutputFilePath)

	if err != nil {
		return err
	}

	defer file.Close()

	for key, val := range reduced {
		line := fmt.Sprintf("(%s, %d)", key, val)
		_, err := file.WriteString(line + "\n")

		if err != nil {
			return err
		}
	}

	return nil
}

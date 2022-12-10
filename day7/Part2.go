package main

import "adventcodingchallenge_2022/utility"

type Part2 struct {
	answer int
}

func (alg *Part2) Process(data []string) (error, interface{}) {

	sizeToConsider := 14
	queue := utility.NewFixedSizeStringQueue(sizeToConsider)
	for i := 0; i < len(data); i++ {
		if i > (sizeToConsider - 1) {
			duplicates := utility.FindDuplicates(queue.PeakAll())
			if duplicates == nil {
				alg.answer = i
				break
			}
		}
		queue.Enqueue(data[i])
	}
	return nil, alg.answer

}

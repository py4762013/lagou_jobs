package main

import (
	"lagou_jobs/pipeline"
	"lagou_jobs/spider"
	"log"
	"sync"
)

var (
	kds = []string{
		//"magento",
		"golang",
	}
	citys = []string{
		//"北京",
		/*"上海",
		"广州",
		"深圳",
		"杭州",
		"成都",*/
		"长沙",
	}

	initResults = []spider.InitResult{}
	loopResults = []spider.LoopResult{}
	jobPipeline = pipeline.NewJobPipeline()

	wg sync.WaitGroup
)

func main() {
	for _, kd := range kds {
		for _, city := range citys {
			wg.Add(1)
			go func(city string, kd string) {
				defer wg.Done()
				initResult, err := spider.InitJobs(city, 1, kd)
				if err != nil {
					log.Fatalln(err)
				}

				initResults = append(initResults, initResult...)
				loopResults = append(loopResults, spider.LoopJobs())
			}(city, kd)
		}
	}

	wg.Wait()

	jobPipeline.Push()

	log.Printf("Init Result: %v", initResults)
	log.Printf("Loop Result: %v", loopResults)
}

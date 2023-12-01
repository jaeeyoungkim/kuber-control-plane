package main

import (
	"encoding/json"
	"fmt"
	"my.com/lib"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		for {
			lib.GetRawMetrics()
			time.Sleep(time.Second * 5)
		}
	}()

	for {
		total := lib.GetApiserverCurrentInflightRequests()
		marshal, err := json.Marshal(total)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(marshal))
		time.Sleep(time.Second * 5)
	}
}

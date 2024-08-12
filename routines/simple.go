package routines

import (
	"sync"

	"github.com/satioO/basics/v2/collections"
)

func Process(list collections.List[uint], wg *sync.WaitGroup) {
	go func() {
		list.Print()
		wg.Done()
	}()
}

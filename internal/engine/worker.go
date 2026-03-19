package engine

import "sync"

func Worker(cmdCh <-chan Cmd, wg *sync.WaitGroup) {
	defer wg.Done()

	for cmd := range cmdCh {
		cmd.Handle()
	}
}

package errgroup

import "sync"

// PanicGroup panic group
type PanicGroup struct {
	wg sync.WaitGroup
}

// Wait wait
func (g *PanicGroup) Wait() {
	g.wg.Wait()
}

// Go go
func (g *PanicGroup) Go(f func() error) {
	g.wg.Add(1)
	go func() {
		defer g.wg.Done()
		if err := f(); err != nil {
			panic(err)
		}
	}()
}

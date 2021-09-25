/*
 * This package was originally written by @austintraver on GitHub.
 * Thanks so much for the awesome work! I've made some modifications to
 * make it work with the rest of the project.
 */

package ui

import (
	"fmt"
	"sync/atomic"
	"time"
)

var states = []rune{'|', '/', '-', '\\', '|'}

// Spinner main type
type Spinner struct {
	frames []rune
	index  int
	active uint64
	text   string
}

func New(text string) *Spinner {
	s := &Spinner{
		text: ("\r\033[K") + text,
	}
	s.Set(states)
	return s
}

func (target *Spinner) Set(frames []rune) {
	target.frames = frames
}

func (target *Spinner) Start() *Spinner {
	if atomic.LoadUint64(&target.active) > 0 {
		return target
	}
	atomic.StoreUint64(&target.active, 1)
	go func() {
		for atomic.LoadUint64(&target.active) > 0 {
			fmt.Printf(target.text, target.next())
			time.Sleep(100 * time.Millisecond)
		}
	}()
	return target
}

func clear() (n int, err error) {
	return fmt.Print("\r\033[K")
}

func (target *Spinner) Stop() (ok bool, err error) {
	x := atomic.SwapUint64(&target.active, 0)
	if x > 0 {
		_, err = clear()
		ok = true
	}
	return
}

func (target *Spinner) next() (char string) {
	char = string(target.frames[target.index%len(target.frames)])
	target.index++
	return
}

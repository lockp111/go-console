package console

import (
	"fmt"
	"sync"
)

var cs = &consoleSpace{}

type consoleSpace struct {
	index int
	total int
	mux   sync.Mutex
}

func (p *consoleSpace) Println(output string) int {
	p.mux.Lock()
	defer p.mux.Unlock()

	if p.index < p.total {
		p.Down(p.total - p.index)
	}

	p.total++
	p.index = p.total
	fmt.Println(output)
	return p.index - 1
}
func (p *consoleSpace) Up(n int) {
	fmt.Printf("\x1b[%dA\r", n)
}

func (p *consoleSpace) Down(n int) {
	fmt.Printf("\x1b[%dB\r", n)
}

func (p *consoleSpace) Progress(index int, output string) {
	p.mux.Lock()
	defer p.mux.Unlock()

	abs := p.index - index
	switch {
	case abs > 0:
		fmt.Printf("\x1b[%dA\x1b[K", abs)
	case abs < 0:
		fmt.Printf("\x1b[%dB\x1b[K", -abs)
	}
	fmt.Printf("%s\x1b[K", output)
	if p.total-index >= 0 {
		fmt.Printf("\x1b[%dB\x1b[K", p.total-index)
	} else {
		p.total = index
	}
}

package console

import (
	"fmt"
	"sync"
)

var cs = &consoleSpace{}

type consoleSpace struct {
	mux   sync.Mutex
	index int
	last  int
}

func (p *consoleSpace) Println(output interface{}) int {
	p.mux.Lock()
	defer p.mux.Unlock()

	p.end()
	p.last++
	p.index = p.last
	fmt.Println(output)
	return p.index - 1
}

func (p *consoleSpace) up(n int) {
	if p.index-n < 0 {
		p.index = 0
		fmt.Printf("\x1b[%dF\r", p.index)
		return
	}
	p.index -= n
	fmt.Printf("\x1b[%dF\r", n)

}

func (p *consoleSpace) down(n int) {
	fmt.Printf("\x1b[%dE\r", n)
	p.index += n
	if p.index > p.last {
		p.last = p.index
	}
}

func (p *consoleSpace) end() {
	n := p.last - p.index
	switch {
	case n == 0:
		return
	case n < 0:
		fmt.Printf("\x1b[%dE\r", p.index-p.last)
		p.last = p.index
		return
	default:
		fmt.Printf("\x1b[%dE\r", p.last-p.index)
		p.index = p.last
	}
}

func (p *consoleSpace) Progress(n int, output string) {
	p.mux.Lock()
	defer p.mux.Unlock()

	abs := p.index - n
	switch {
	case abs > 0:
		p.up(abs)
	case abs < 0:
		p.down(-abs)
	}

	fmt.Printf("\x1b[2K%s", output)
	p.end()
}

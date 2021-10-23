package console

func Println(output string) int {
	return cs.Println(output)
}

func Up(n int) {
	cs.Up(n)
}

func Down(n int) {
	cs.Down(n)
}

func Progress(index int, output string) {
	cs.Progress(index, output)
}

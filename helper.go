package console

func Println(output string) int {
	return cs.Println(output)
}

func Progress(index int, output string) {
	cs.Progress(index, output)
}

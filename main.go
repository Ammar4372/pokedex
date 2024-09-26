package main

type command struct {
	name        string
	description string
	callback    func() error
}

var offset int

func main() {
	start_repl()
}

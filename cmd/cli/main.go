package main

func main() {
	cli := NewCli()
	go cli.ReadLoop()
	cli.WriteLoop()
}

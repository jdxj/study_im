package main

func main() {
	sign := make(chan bool)
	client := NewClient()
	client.Run(sign)
}

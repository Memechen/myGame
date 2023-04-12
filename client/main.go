package client

func main() {
	c := NewClient()
	c.Run()
	select {}
}

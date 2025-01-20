package main

func main() {
	server := NewAPIServer(":8888")
	server.Run()
}

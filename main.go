package main

import "code/cmd"

// @title Go-Web
// @version 0.0.1
// @description 学习Golang
func main() {
	defer cmd.Clean()
	cmd.Start()
}

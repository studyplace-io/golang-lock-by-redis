package main

import "redis-practice/jobtest"

func main() {
	jobtest.Run()
	select {}
}

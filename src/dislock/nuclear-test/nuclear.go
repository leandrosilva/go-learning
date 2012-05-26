package main

import (
	"dislock/lock"
	"flag"
	"fmt"
)

func main() {
	var uuid = flag.String("u", "nuclear_suff", "Lock UUID")
	var client = flag.String("c", "homer_simpson", "Client ID")

	flag.Parse()

	var l = lock.New(*uuid, *client)

	fmt.Println("Creating new lock: {UUID:", l.UUID, ", Client:", l.Client, ", Acquired:", l.Acquired, "}")
	fmt.Println("Trying to acquire this lock")
	l.Acquire(5)
	fmt.Println("Could it acquire this lock:", l.Acquired)

	fmt.Println("Type <RETURN> to QUIT...")
	var key int
	fmt.Scanf(">> %v", key)

	if l.Acquired {
		fmt.Println("Releasing this lock...")
		l.Release()
		fmt.Println("This lock was released")
	}
}

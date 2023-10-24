package main

type User struct {
	Name string `ormazing:"primary key"`
	Age  int
}

var dial, _ = di

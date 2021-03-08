package main

import (
	"fmt"

	"github.com/nafeem-evatix/repohandler"
)

func main() {
	fmt.Println("hello, world !")

	inmemory := repohandler.NewInMemoryDB()       // initializing in memory database
	data := struct{ Name string }{Name: "Evatix"} // test data

	// creating multiple datas
	inmemory.Create(data)
	inmemory.Create(data)
	inmemory.Create(data)

	fmt.Println(inmemory.Get())            // returns all data created earlier
	fmt.Println(inmemory.GetSingle("aaa")) // error No Data Found
}

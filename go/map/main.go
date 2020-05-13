package main

import (
	"fmt"
	"sync"
)

func main() {
	m := make(map[string]int, 100)
	fmt.Println(len(m))

	var syncM sync.Map
	syncM.Load('1')

}

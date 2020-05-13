package main

import "fmt"
/*
0 [1 2 3 4]
1 [1 2 3 4]
delete i: 2  value: [1 2 4]
2 [1 2 4]
3 [1 2 4]
注意小心越界
*/
func main(){
	var infoList =[]int{1,2,3,4}
	for i := range infoList {
		if i == 2 {
			infoList = append(infoList[:i], infoList[i+1:]...)
			fmt.Println("delete i:",i," value:",infoList)
		}
		fmt.Println(i,infoList)
	}
}

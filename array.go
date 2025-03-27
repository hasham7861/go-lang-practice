package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println(a)

	a[4] = 100
	fmt.Println("set:", a)
    fmt.Println("get:", a[4])
    fmt.Println("len:", len(a))

    // to declare array and intialize all values
    b := [5]int{1,2,3,4,5}
    fmt.Println("dcl:", b)

    // compiler can auto count the array values and set a size
    b = [...]int{1,2,3,4,5}

    // if you assign value by index, then index value in between the set index and previous element will be 0
    b = [...]int{100, 3: 400, 500}
    fmt.Println("idx:", b)

    var twoD [2][3]int
    for i:=0; i<2; i++ {
    	for j:=0; j<3; j++{
    		twoD[i][j] = i +j
    	}
    }

    // you can intialize 2d array like so, with column size first and then row size
    twoD = [2][3]int{
    	{1,2,3},
    	{1,2,3},
    }

    fmt.Println("2d: ", twoD)
}
package main

import "fmt"

func main() {

	nl1, nl2 := getLength()
	//fmt.Println(nl1,nl2)
	na1, na2 := createArrays(nl1, nl2)
	fmt.Println("Your Arrays are as follows:", na1, "and ", na2)

	fmt.Println("Your new merged array is: ", append(na1, na2...))
}

func getLength() (l1 int, l2 int) {

	//l1, l2 := 0,0

	for l1 == 0 {
		fmt.Println("Enter the length of your 1st array:")
		fmt.Scan(&l1)
		if l1 == 0 {
			fmt.Println("Can't enter 0 as length of array. Try Again!")
		}
	}

	for l2 == 0 {
		fmt.Println("Enter the length of your 2nd array:")
		fmt.Scan(&l2)
		if l2 == 0 {
			fmt.Println("Can't enter 0 as length of array. Try Again!")
		}

	}
	return
}

func createArrays(al1 int, al2 int) ([]int, []int) {

	/* can i pass the returned values of getLength() as arguments in createArrays()? if so how? i tried calling
	getLength within createArrays like so...
	func createArrays(getLength ())
	Of course ths is wrong but is there any way I can do this rather than calling the function in main and then passing
	those values down to createArrays like how its currently doing?
	can we use closure here?
	*/

	array1 := make([]int, al1)
	array2 := make([]int, al2)
	fmt.Println("Lets start with the 1st array elements!")
	for i := 0; i < al1; i++ {
		fmt.Println("Enter value", i, ":")
		fmt.Scan(&array1[i])
		if array1[i] == 0 {
			fmt.Println("Cannot enter 0. Try again")
			i--
		}
	}

	fmt.Println("Lets start with the 2nd array elements!")
	for j := 0; j < al2; j++ {
		fmt.Println("Enter value", j, ":")
		fmt.Scan(&array2[j])
		if array2[j] == 0 {
			fmt.Println("Cannot enter 0. Try again")
			j--
		}
	}

	return array1, array2
}

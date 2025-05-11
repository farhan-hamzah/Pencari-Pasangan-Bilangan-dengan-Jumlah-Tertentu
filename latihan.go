package main
import "fmt"
const NMAX int = 100
func main(){
	var A[NMAX]int
	var B[NMAX]int
	var n, i, j, target, c int
	var found bool
	found = false
	fmt.Scan(&n, &target)
	for i = 0; i < n; i++{
		fmt.Scan(&A[i])
	}
	for i = 0; i < n && found == false; i++{
		c++
		for j = i+1; j < n && found == false; j++{
			if A[i]+A[j] == target{
				B[i] = A[i]
				B[i+1] = A[j]
				c++
				found = true
			}
		}
	}
	if found {
		for i = 0; i < c; i++ {
			fmt.Print(B[i], " ")
		}
		fmt.Println()
	} else {
		fmt.Println("Tidak ditemukan pasangan.")
	}
}

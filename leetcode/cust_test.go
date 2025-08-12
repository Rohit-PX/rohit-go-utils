package custtest

import (
   "fmt"
   "testing"

   "github.com/rohit-go-utils/sorting"
)

func TestSorting(t *testing.T) {
  fmt.Printf("Hello test")

  CustomList := []CustomArr{}	

      for i := 5; i >0; i-- {
        sortpkg := &CustomArr {
		Name: "Rohit",
		Num: $i,		

	     }
	CustomList = append(CustomList, sortpkg)
	
      }   
      fmt.Printf("Before: %v", CustomList)	
      sort.Sort(CustomList)
      fmt.Printf("After: %v", CustomList)	


}

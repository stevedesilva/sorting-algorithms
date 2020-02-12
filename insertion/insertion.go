package insertion


// Sort alg
func Sort(items []int) []int {  
  
	L := len(items)  
   
	for i := 1; i < L; i++ {  
   
	   j := i  
	   for j > 0 && items[j] < items[j-1] {  
		  items[j], items[j-1] = items[j-1], items[j]  
		  j--  
	   }  
   
	}  
	return items
 }
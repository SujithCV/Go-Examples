package main
import "fmt"
func main(){
a:=[5]int {1,2,3,4,5}

fmt.Println(a)
fmt.Println("Length of array- ",len(a))
TwoD :=[2][3] int {{1,2,3},{4,5,6}}
for i:=0;i<2;i++{
  for j:=0;j<3;j++{
    fmt.Println(TwoD[i][j])
  }
}

}

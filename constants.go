package main
import(
"fmt"
)

const s string="Constant"

func main(){

const n=5000000
const d=3e22/n
fmt.Println(s)
fmt.Println(d)
fmt.Println(int64(d))
fmt.Println(n)

}

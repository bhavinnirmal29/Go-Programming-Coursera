package main
import (
	"fmt"
	"strings"
	"os"
	"bufio"
)
var s string
func main(){
	var s1 string
	fmt.Printf("Enter String : ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s1 = scanner.Text()
	fmt.Printf("Entered String :  %s\n", s1)
	s=strings.TrimSpace(strings.ToLower(s1))
	if (strings.HasPrefix(s,"i") && strings.Contains(s,"a") && strings.HasSuffix(s,"n")){
		fmt.Printf("Found!")
	}else{
		fmt.Printf("Not Found!")
	}
}
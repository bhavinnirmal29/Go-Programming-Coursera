package main
import (
	"fmt"
)
type Animal interface{
	Eat()
	Speak()
	Move()
}
type Cow struct{}

func (cow Cow) Eat(){
	fmt.Printf("food : grass\n")
}

func (cow Cow) Speak(){
	fmt.Printf("noise : moo\n")
}

func (cow Cow) Move(){
	fmt.Printf("locomotion : walk\n")
}

type Bird struct{

}

func (bird Bird) Eat(){
	fmt.Printf("food : worms\n")
}

func (bird Bird) Speak(){
	fmt.Printf("noise : peep\n")
}

func (bird Bird) Move(){
	fmt.Printf("locomotion : fly\n")
}

type Snake struct{

}

func (snake Snake) Eat(){
	fmt.Printf("food : mice\n")
}

func (snake Snake) Speak(){
	fmt.Printf("noise : hsss\n")
}

func (snake Snake) Move(){
	fmt.Printf("locomotion : slither\n")
}

var choice string
func main(){
	var str1,str2 string
	cow:=Cow{}
	bird:=Bird{}
	snake:=Snake{}
	an_list:=make(map[string]Animal)
	fmt.Printf("Type any letter/digit to Exit.\n")
	for {
		fmt.Printf("Enter Command as string : newanimal or query \n ")
		fmt.Scan(&choice)
		if(choice=="newanimal"){
			fmt.Printf("Enter Name & type : ")
			fmt.Scan(&str1)
			fmt.Scan(&str2)
			if str2 == "cow" {
				an_list[str1] = cow
			} else if str2 == "bird" {
				an_list[str1] = bird
			} else if str2 == "snake" {
				an_list[str1] = snake
			}else{
				continue
			}
			fmt.Printf("\nCreated It! \n")
		}else if(choice=="query"){
			fmt.Printf("Enter Name & Query : ")
			fmt.Scan(&str1)
			fmt.Scan(&str2)
			ani,ok:=an_list[str1]
			if !ok {
				fmt.Printf("\nIncorrect")
			}else{
				fmt.Printf("\nName : %s  ",str1)
				if str2 == "eat" {
					ani.Eat()
				} else if str2 == "move" {
					ani.Move()
				} else if str2 == "speak" {
					ani.Speak()
				} else {
					fmt.Println("Incorrect input. Try again")
				}
			}
		}else{
			break
		}
	}
}
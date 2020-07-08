package main
import (
	"fmt"
)
type Animal struct{
	food string
	locomotion string
	noise string
}
func (a *Animal) Eat(name string){
	if(name=="cow"){
		a.food="grass"
	}
	if(name=="bird"){
		a.food="worms"
	}
	if(name=="snake"){
		a.food="mice"
	}
}
func (a *Animal) Move(name string){
	if(name=="cow"){
		a.locomotion="walk"
	}
	if(name=="bird"){
		a.locomotion="fly"
	}
	if(name=="snake"){
		a.locomotion="slither"
	}
}
func (a *Animal) Speak(name string){
	if(name=="cow"){
		a.noise="moo"
	}
	if(name=="bird"){
		a.noise="peep"
	}
	if(name=="snake"){
		a.noise="hsss"
	}
}
var name,method string
func main(){
	var a Animal
	fmt.Printf("Type Any String other than choice to Exit. \n")
	for {
		fmt.Printf("Enter strings > ")
		fmt.Scan(&name)
		fmt.Scan(&method)
		if(method=="Eat"){
			a.Eat(name)
			fmt.Printf("\nName : %s   %s : %s\n",name,method,a.food)
		}else if(method=="Move"){
			a.Move(name)
			fmt.Printf("\nName : %s   %s : %s\n",name,method,a.locomotion)
		}else if(method=="Speak"){
			a.Speak(name)
			fmt.Printf("\nName : %s   %s : %s\n",name,method,a.noise)
		}else{
			break
		}
	}
}
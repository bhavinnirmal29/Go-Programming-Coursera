package main
import (
	"fmt"
	"math"
)
var a,v0,s0 float64
var t float64
func GenDisplaceFn(a1,v01,s01 float64) func(float64) float64{
	fn:=func(t float64) float64{
		return 0.5 * (a1*math.Pow(t,2))+(v01*t)+s01
	}
	return fn
}
func main(){
	fmt.Printf("Enter values for acceleration,initial velocity, & initial displacement (Space Separated) : ")
	fmt.Scan(&a)
	fmt.Scan(&v0)
	fmt.Scan(&s0)
	fmt.Printf("\nEnter value for time : ")
	fmt.Scan(&t)
	fn:=GenDisplaceFn(a,v0,s0)
	fmt.Printf("%f",fn(t))
}
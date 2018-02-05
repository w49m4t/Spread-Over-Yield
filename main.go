package main

import (
	"fmt"
	"math"
)

==begin
These are the net present value functions of a Tbond. Shuold be updated to
be able to project it dynamically
==end

func pv(c, r, d float64, t float64) float64{
    var cv float64 = -c
    var j float64 = t-1
    for i:=1.0; i<j; i++ {
        cv+=(c*r)/(math.Pow((1.0+d),i))
    }
    cv+=(c*(1+r))/(math.Pow((1.0+d),t))
    return cv
}

func pvd(c, r, d float64, t float64) float64{
    var cv float64 = -c
    var j float64 = t-1
    for i:=1.0; i<j; i++ {
        
        cv+=(i*c*r)/(math.Pow((1+d),i+1.0))
    }
    cv+=(t*(1+r)*c)/(math.Pow((1+d),t+1))
    return -cv
}

func soy(target, c, r, d, t float64) float64{
    spr:=0.01 //initial guess, change this to start from a different point
    i:=0
    for ; math.Abs(target-pv(c,r,d+spr,t))>=0.001; {
        spr-=pv(c,r,d+spr,t)/pvd(c,r,d+spr,t) //tolerance is now set to 0.001, change if you re looking for a more accurate value
        i+=1
        if i>3000{ 
            fmt.Println("error") //reaching the total error number
            break }
    }

    return spr
}

func main() {
	//soy calculates the IRR as a spread over yield of a Tbond. Cashflow is automatically populated
	
	z:=soy(0.0,100.0,0.03,0.02,25.0)
	k:=pv(100.0,0.03,0.02+z,25.0)
	fmt.Println(z)
	fmt.Println(k)
}

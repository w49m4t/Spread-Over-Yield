package main

import (
	"fmt"
	"math"
)

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
    spr:=0.01
    i:=0
    for ; math.Abs(target-pv(c,r,d+spr,t))>=0.001; {
        spr-=pv(c,r,d+spr,t)/pvd(c,r,d+spr,t)
        i+=1
        if i>3000{ 
            fmt.Println("error")
            break }
    }

    return spr
}

func main() {
	//fmt.Println("Hello, 世界")
	
	z:=soy(0.0,100.0,0.03,0.02,25.0)
	k:=pv(100.0,0.03,0.02+z,25.0)
	fmt.Println(z)
	fmt.Println(k)
}

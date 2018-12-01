package main

import (
  "fmt"
  "./gistool"
)

func main() {
 var x float64 = 12128777.9792
 var y float64 = 4040203.41477
 lon,lat,_ := gistool.MCT_2_BD09(x,y)
 lonb,latb := gistool.BD09_2_GCJ03(lon,lat)
 fmt.Printf("%f,%f\n",lon,lat)
 fmt.Printf("%f,%f\n",lonb,latb)
}

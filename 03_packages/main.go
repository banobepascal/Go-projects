package main

import (
	"fmt" 
	"math"
	"github.com/BanobePascal/practice/03_packages/strutil"
	"math/rand"
	"time"
)

func main() {
	
	epoch := time.Now().Unix()
	fmt.Println(epoch)

	rand.Seed(epoch)
	fmt.Println(rand.Intn(100))

	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Ceil(2.7))
	fmt.Println(math.Sqrt(2.7))
	fmt.Println(strutil.Reverse("hello"))
}
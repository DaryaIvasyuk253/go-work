package main

import ("fmt"
	"os"
	"time"
	"strconv"
)

var inpt string
var except int
var counter int
var success int
var fibonacci int
var nullparam int
func main(){

	for i:=0; i<1 ; i++  {
		go scan1()
		time.Sleep(time.Second * 10)

		if inpt == "" {
			nullparam++
			counter++
			fmt.Println(counter, " : ", fibonacci)
			fibonacci++
			except++
			fmt.Println("exception")
		}
		inpt, err:= strconv.Atoi(inpt)
		fmt.Println(err)

		if inpt == fibonacci {
			counter++
			fmt.Println(counter, " : ", fibonacci)
			fibonacci++
			 success++
		}
		if    inpt != fibonacci && counter==0 {
			counter++
			fmt.Println(counter, " : ", fibonacci)
			fibonacci++
			except++
			fmt.Println("exception")
		}

	}

	a:= 0
	b:= 1
	for i:=0; ;i++ {

		if except == 3 {
			os.Exit(0)
		}
		if except == 0 && success == 10 {
			os.Exit(0)

		}
		go scan2()
		time.Sleep(time.Second * 10)
		inpt, err:= strconv.Atoi(inpt)
		fmt.Println(err)
		if inpt == fibonacci {
			counter++
			success++
			fmt.Println(counter, " : ", fibonacci)
			fibonacci = a + b
			a = b
			b = fibonacci

		}else  {
			counter++
			except++
			fmt.Println("exception")
			fmt.Println(counter, " : ", fibonacci)
			fibonacci = a + b
			a = b
			b = fibonacci
		}
	}
}

func scan1()  {
	fmt.Scanln(&inpt)
}

func scan2()  {
	fmt.Scanln(&inpt)

}



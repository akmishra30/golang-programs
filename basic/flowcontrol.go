package main

import (
	p "fmt"
	"runtime"
	con "strconv"
	"time"
)

func main() {
	// for loop function
	forloop()
	//while loop function
	whileloop()
	//ifelse
	ifelse()
	//switch statement
	switch_statement()
}

func forloop() {
	p.Println("------Inside forloop------")
	for i := 0; i < 5; i++ {
		p.Println(con.Itoa(i)) //con.Itoa is used to convert interger to string
	}

	//infinite loop
	//for {}
	sum := 1
	for sum < 100 { //init and post statements are optional
		sum += sum
	}
	p.Println(sum)
}

func whileloop() {
	p.Println("------Inside whileloop------")

	sum := 1
	for sum < 5 {
		p.Println(sum)
		sum += sum
	}
	p.Println("Summation : " + con.Itoa(sum))
}

func ifelse() {
	p.Println("------Inside ifelse------")
	count := 1
	if count > 0 { //only if condition
		count = 5
	}

	p.Println(count)

	if v := count * 10; v > 10 { // if condition with short statement
		p.Println(v)
	}
	/*
		if/else condition.
		Note: variable defined in if block can be used in else block as well
			 The same varibale can't be use outside if-else block
	*/
	if tmp := count * 10; tmp > 10 {
		p.Println("if block : ", tmp)
	} else {
		p.Println("else block : ", tmp)
	}

}

/*
In GO switch case is similar to other language except break statement in each case.
GO itself adds break statement automatically. Also, GO cases need not be contant.
You can put condition as well in GO case.
*/
func switch_statement() {
	p.Println("------Inside switch_statement------")
	p.Println("------Simple switch example------")
	v := runtime.GOOS
	switch v {
	case "darwin":
		println("Mac OS")
	case "linux":
		println("Linux")
	default:
		println(v)
	}

	p.Println("------Switch case with evaluation order.------")
	month := time.Now().Month()

	switch time.November { //evaluation exp
	case month + 0: //conditional expression
		p.Println("It's a running month..!!")
	case month + 1:
		p.Println("Next month will be ..!!")
	default:
		p.Println("Still bit far ..!!")
	}

	p.Println("-----Switch case with no condition------")

	t := time.Now()
	switch {
	case t.Hour() < 12:
		p.Println("Good Morning..!!")
	case t.Hour() < 17:
		p.Println("Good Afternoon..!!")
	case t.Hour() < 20:
		p.Println("Good Evening..!!")
	default:
		p.Println("Good Night..!!")
	}

}

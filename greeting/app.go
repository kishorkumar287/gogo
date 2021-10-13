package main

import (
	"fmt"
)

var m map[string]int = make(map[string]int)

type login struct {
	password string
	desig    int
	//1-admin,2-customer
}

var l map[string]login = make(map[string]login)

func customer() {
	var p string
	var q int
	//fmt.Print(m)
	for k, v := range m {
		if v > 0 {
			fmt.Println(k, "Quatity is", v)
		}
	}
	fmt.Print("Enter the product you want to buy:")
	fmt.Scanln(&p)
	fmt.Print("Enter the product quantity you want to buy:")
	fmt.Scanln(&q)
	if m[p] != 0 && m[p] >= q {
		m[p] -= q
		fmt.Println("product purchased, Thanks")
	} else {
		fmt.Println("Please enter correct values")
	}

}

func owner() {
	var p string
	var q int

	//fmt.Print(m)
	for k, v := range m {
		fmt.Println(k, "Quatity is", v)
	}
	fmt.Print("Enter the product you want to update:")
	fmt.Scanln(&p)
	fmt.Print("Enter the product quantity you want to update:")
	fmt.Scanln(&q)
	m[p] = q
	fmt.Println("Product updated Sucessfully, available products are:")
	for k, v := range m {
		fmt.Println(k, "Quatity is", v)
	}
	fmt.Println("************")

}

func signin() {

	var user string
	var pass string
	fmt.Println("Enter username and password:")

	fmt.Scanln(&user)
	fmt.Scanln(&pass)
	if l[user].desig == 0 || l[user].password != pass {
		fmt.Println("Enter valid username and pasword")
	} else {
		for {
			//fmt.Print("Two")
			if l[user].desig == 1 {
				owner()
			} else {
				customer()
			}
		}
	}
}

func signup() {
	fmt.Println("Enter username and password:")
	var name string
	var password string
	fmt.Scanln(&name)
	fmt.Scanln(&password)
	if l[name].desig == 0 {
		l[name] = login{password, 2}
		signin()
	} else {
		fmt.Println("Sorry username already exist")
		signup()
	}
}
func main() {
	l["kishor"] = login{"kishor", 1}
	l["customer"] = login{"cust", 2}
	m["KitKat"] = 5
	m["Kit"] = 5
	m["Kat"] = 5
	var n int
	fmt.Println("Enter 1 for Signup and 2 for Signin:")
	fmt.Scanln(&n)
	if n == 1 {
		signup()
	} else {
		signin()
	}

}

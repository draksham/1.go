// package main
// import ("fmt")

// func main() {
//   fmt.Println("Hello World!")
// } 


// package main
// import ("fmt")

// func main() {
//   var i,j = 10,20

//   fmt.Print(i ," and ", j)
// }



// package main
// import ("fmt")

// func main() {
//   var i,j string = "Hello","World"

//   fmt.Println(i,j)
// }

// package main
// import ("fmt")

// func main() {
//   var arr1 = [3]int{1,2,3}
//   arr2 := [5]int{4,5,6,7,8}

//   fmt.Println(arr1)
//   fmt.Println(arr2)
// }


// package main
// import ("fmt")

// func main() {
//   var cars = [4]string{"Volvo", "BMW", "Ford", "Mazda"}
//   fmt.Print(cars)
// }


// package main
// import ("fmt")

// func main() {
//   prices := [3]int{10,20,30}

//   fmt.Println(prices[0])
//   fmt.Println(prices[2])
// }


// package main
// import ("fmt")

// func main() {
//   prices := [3]int{10,20,30}

//   prices[2] = 50
//   fmt.Println(prices)
// }


// package main
// import ("fmt")

// func main() {
//   arr1 := [5]int{} //not initialized
//   arr2 := [5]int{1,2} //partially initialized
//   arr3 := [5]int{1,2,3,4,5} //fully initialized

//   fmt.Println(arr1)
//   fmt.Println(arr2)
//   fmt.Println(arr3)
// }


// package main
// import ("fmt")

// func main() {
//   arr1 := [5]int{} //not initialized
//   arr2 := [5]int{1,2} //partially initialized
//   arr3 := [5]int{1,2,3,4,5} //fully initialized

//   fmt.Println(arr1)
//   fmt.Println(arr2)
//   fmt.Println(arr3)
// }



// package main
// import ("fmt")

// func main() {
//   arr1 := [5]int{1:10,2:40}

//   fmt.Println(arr1)
// }




// package main
// import ("fmt")

// func main() {
//   day := 4

//   switch day {
//   case 1:
//     fmt.Println("Monday")
//   case 2:
//     fmt.Println("Tuesday")
//   case 3:
//     fmt.Println("Wednesday")
//   case 4:
//     fmt.Println("Thursday")
//   case 5:
//     fmt.Println("Friday")
//   case 6:
//     fmt.Println("Saturday")
//   case 7:
//     fmt.Println("Sunday")
//   }
// }


// package main
// import ("fmt")

// func main() {
//   for i:=0; i < 5; i++ {
//     fmt.Println(i)
//   }
// }


// package main
// import ("fmt")

// func myMessage() {
//   fmt.Println("I just got executed!")
// }

// func main() {
//   myMessage() // call the function
// }



// package main
// import ("fmt")

// func familyName(fname string) {
//   fmt.Println("Hello", fname, "Refsnes")
// }

// func main() {
//   familyName("Liam")
//   familyName("Jenny")
//   familyName("Anja")
// }


// package main
// import ("fmt")

// func testcount(x int) int {
//   if x == 10 {
//     return 0
//   }
//   fmt.Println(x)
//   return testcount(x)
// }

// func main(){
//   testcount(1)
// }



// package main
// import ("fmt")

// type Person struct {
//   name string
//   age int
//   job string
//   salary int
// }

// func main() {
//   var pers1 Person
//   // var pers2 Person

//   // Pers1 specification
//   pers1.name = "Hege"
//   pers1.age = 45
//   pers1.job = "Teacher"
//   pers1.salary = 6000

//   // Pers2 specification
//   // pers2.name = "Cecilie"
//   // pers2.age = 24
//   // pers2.job = "Marketing"
//   // pers2.salary = 4500

//   // Access and print Pers1 info
//   fmt.Println("Name: ", pers1.name)
//   fmt.Println("Age: ", pers1.age)
//   fmt.Println("Job: ", pers1.job)
//   fmt.Println("Salary: ", pers1.salary)
// }

//   // Access and print Pers2 info
// //   fmt.Println("Name: ", pers2.name)
// //   fmt.Println("Age: ", pers2.age)
// //   fmt.Println("Job: ", pers2.job)
// //   fmt.Println("Salary: ", pers2.salary)
// // }


// package main
// import ("fmt")

// type Person struct {
//   name string
//   age int
//   job string
//   salary int
// }

// func main() {
//   var pers1 Person
//   var pers2 Person

//   // Pers1 specification
//   pers1.name = "Hege"
//   pers1.age = 45
//   pers1.job = "Teacher"
//   pers1.salary = 6000

//   // Pers2 specification
//   pers2.name = "Cecilie"
//   pers2.age = 24
//   pers2.job = "Marketing"
//   pers2.salary = 4500

//   // Print Pers1 info by calling a function
//   printPerson(pers1)

//   // Print Pers2 info by calling a function
//   printPerson(pers2)
// }

// func printPerson(pers Person) {
//   fmt.Println("Name: ", pers.name)
//   fmt.Println("Age: ", pers.age)
//   fmt.Println("Job: ", pers.job)
//   fmt.Println("Salary: ", pers.salary)
// }




// package main
// import ("fmt")

// func main() {
//   var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
//   b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}

//   fmt.Printf("a\t%v\n", a)
//   fmt.Printf("b\t%v\n", b)
// }



// package main
// import ("fmt")

// func main() {
//   var a = make(map[string]string) // The map is empty now
//   a["brand"] = "Ford"
//   a["model"] = "Mustang"
//   a["year"] = "1964"
//                                  // a is no longer empty
//   b := make(map[string]int)
//   b["Oslo"] = 1
//   b["Bergen"] = 2
//   b["Trondheim"] = 3
//   b["Stavanger"] = 4

//   fmt.Printf("a\t%v\n", a)
//   fmt.Printf("b\t%v\n", b)
// }



// package main
// import ("fmt")

// func main() {
//   var a = make(map[string]string)
//   var b map[string]string

//   fmt.Println(a == nil)
//   fmt.Println(b == nil)
// }



// package main
// import ("fmt")

// func main() {
//   var a = make(map[string]string)
//   a["brand"] = "Ford"
//   a["model"] = "Mustang"
//   a["year"] = "1964"

//   fmt.Printf(a["brand"])
// }

package main
import ("fmt")

func main() {
  var a = make(map[string]string)
  a["brand"] = "Ford"
  a["model"] = "Mustang"
  a["year"] = "1964"

  fmt.Println(a)

  a["year"] = "1970" // Updating an element
  a["color"] = "red" // Adding an element

  fmt.Println(a)
}
# domz2_1

    package main

    import (
    "fmt"
    "github.com/Hrukem/domz2_1"
    "time"
    )

    func main() {
    c := domz2_1.New()

	fmt.Println("\nset three value with lifetime 5 second")
	c.Set("1", 12345, 5)
	c.Set("2", 54321, 5)
	c.Set("3", "Golang", 5)
	fmt.Println(c.Get("1"))
	fmt.Println(c.Get("2"))
	fmt.Println(c.Get("3"))

	time.Sleep(time.Second * 2)
	fmt.Println("\nfrom 3 second")
	fmt.Println(c.Get("1"))
	fmt.Println(c.Get("2"))
	fmt.Println(c.Get("3"))

	fmt.Println("\ndelete key \"2\"")
	c.Delete("2")
	fmt.Println(c.Get("1"))
	fmt.Println(c.Get("2"))
	fmt.Println(c.Get("3"))

	fmt.Println("\nset new value 2 with new lifetime 6 second")
	c.Set("2", "Ninja", 6)
	fmt.Println(c.Get("1"), "lifetime 5 sec.")
	fmt.Println(c.Get("2"), "lifetime 5+6 sec.")
	fmt.Println(c.Get("3"), "lifetime 5 sec.")

	time.Sleep(time.Second * 4)
	fmt.Println("\nfrom 6 second")
	fmt.Println(c.Get("1"), "lifetime 5 sec.")
	fmt.Println(c.Get("2"), "lifetime 5+6 sec.")
	fmt.Println(c.Get("3"), "lifetime 5 sec.")

	time.Sleep(time.Second * 2)
	fmt.Println("\nfrom 8 second")
	fmt.Println(c.Get("1"), "lifetime 5 sec.")
	fmt.Println(c.Get("2"), "lifetime 5+6 sec.")
	fmt.Println(c.Get("3"), "lifetime 5 sec.")

	time.Sleep(time.Second * 4)
	fmt.Println("\nfrom 12 second")
	fmt.Println(c.Get("1"), "lifetime 5 sec.")
	fmt.Println(c.Get("2"), "lifetime 5+6 sec.")
	fmt.Println(c.Get("3"), "lifetime 5 sec.")

	fmt.Println("\nEnd program")
}
package main

import "fmt"

type sender struct {
	rateLimit int
	user
}

type user struct {
	name   string
	number int
}

func test(s sender) {
	fmt.Println("Sender name:", s.name)
	fmt.Println("Sender number:", s.number)
	fmt.Println("Sender rateLimit", s.rateLimit)
	fmt.Println("=====================")
}

func main() {
	test(sender{
		rateLimit: 10000,
		user: user{
			name:   "Deborah",
			number: 180555558790,
		},
	})
	test(sender{
		rateLimit: 5000,
		user: user{
			name:   "Sarah",
			number: 1905558790,
		},
	})

	test(sender{
		rateLimit: 1000,
		user: user{
			name:   "Sally",
			number: 1905558790,
		},
	})

}

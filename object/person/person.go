package person

import "fmt"

type ContactInfo struct {
	Email string
	Zip   int
}

type Person struct {
	FirstName string
	LastName  string
	// contact contactInfo
	ContactInfo // the same with "contactInfo contactInfo"
}

func (p Person) Print() {
	fmt.Printf("%+v\n", p)
}

func (p *Person) Update(name string) {
	(*p).FirstName = name
}

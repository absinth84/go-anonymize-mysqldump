package main

import (
	"syreclabs.com/go/faker"
)

func generateUsername() string {
	return (faker.Internet().UserName() + faker.Lorem().Characters(6))
}

func generatePassword() string {
	// TODO encrypt this value
	return faker.Internet().Password(8, 14)
}

func generateEmail() string {
	return faker.Internet().SafeEmail()
}

func generateURL() string {
	return faker.Internet().Url()
}

func generateName() string {
	return faker.Name().Name()
}

func generateFirstName() string {
	return faker.Name().FirstName()
}

func generateLastName() string {
	return faker.Name().LastName()
}

func generateParagraph() string {
	return faker.Lorem().Sentence(3)
}

func generateIPv4() string {
	return faker.Internet().IpV4Address()
}

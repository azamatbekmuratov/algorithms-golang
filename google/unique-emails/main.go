package main

import (
	"fmt"
)

type UEmailSet map[string]struct{}

//add Email to the set
func (uEmailSet UEmailSet) add(email string) {
	uEmailSet[email] = struct{}{}
}

//removes the email from the set
func (uEmailSet UEmailSet) remove(email string) {
	delete(uEmailSet, email)
}

func numUniqueEmails(emails []string) int {
	uEmails := UEmailSet{}

	for i := 0; i < len(emails); i++ {
		email := []rune(emails[i])
		tmpEmail := []rune{}
		domain := []rune{}

		for _, ch := range email {
			if string(ch) == "+" || string(ch) == "@" {
				break
			}
			if string(ch) != "." {
				tmpEmail = append(tmpEmail, ch)
			}
		}

		for j := len(email) - 1; j > 0; j-- {
			if email[j] != '@' {
				domain = append(domain, email[j])
			} else {
				break
			}
		}

		//reverse the domain
		for i := 0; i < len(domain); i++ {
			for k := len(domain) - 1; k > 0; k-- {
				domain[i], domain[k] = domain[k], domain[i]
			}
		}
		tmpEmail = append(tmpEmail, domain...)

		uEmails.add(string(tmpEmail))
	}
	fmt.Println(uEmails)
	return len(uEmails)
}

func main() {
	input := []string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"}
	fmt.Println(numUniqueEmails(input))
}

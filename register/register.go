package register

import "fmt"

//Register registers a new user to twitter.
func Register() error {
	createIdentity()
	fmt.Println("we register a new user")

	return nil
}

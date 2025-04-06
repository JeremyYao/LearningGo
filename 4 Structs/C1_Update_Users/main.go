package main

type Membership struct {
	Type             string
	MessageCharLimit int
}

type User struct {
	Name string
	Membership
}

func newUser(name string, membershipType string) User {
	new_user := User{
		Name: name,
		Membership: Membership{
			Type:             membershipType,
			MessageCharLimit: 100,
		},
	}

	if membershipType == "premium" {
		new_user.MessageCharLimit = 1000
	}

	return new_user
}

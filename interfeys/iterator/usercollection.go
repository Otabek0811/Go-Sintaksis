package main

type userCollection struct {
	users []*User
}

func (u *userCollection) createIterator() Iterator {
	return &UserIterator{
		users: u.users,
	}
}

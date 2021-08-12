package main

type userIterator struct {
	index          int
	userCollection *userCollection
}

func (u *userIterator) hasNext() bool {
	if u.index < len(u.userCollection.users) {
		return true
	}
	return false
}

func (u *userIterator) getNext() *user {
	if u.hasNext() {
		user := u.userCollection.users[u.index]
		u.index++
		return user
	}
	return nil
}

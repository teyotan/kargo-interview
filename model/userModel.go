package model

type User struct {
	ID   int
	Name string
}

func GetUserData() []User {
	var user User
	var userList []User

	nameList := []string{"Pikachu", "Doraemon", "Doggy A", "Doggy B"}

	userList = make([]User, 0)
	for i := 1; i < 5; i++ {
		user = User{
			ID:   i,
			Name: nameList[i-1],
		}
		userList = append(userList, user)
	}

	return userList
}

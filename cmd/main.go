package main

import (
	fnc "User/internal/functions"
	st "User/storage"
)

func main() {
	var users []st.User
	fnc.UserInfo(&users)
	fnc.Printer(users)
}
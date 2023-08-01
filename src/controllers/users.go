package usersController

import "net/http"

func Create(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("User created"))
}

func Index(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("User index"))
}

func Show(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("User showed"))
}

func Update(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("User updated"))
}

func Delete(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("User deleted"))
}

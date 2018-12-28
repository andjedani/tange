// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graph

type Token struct {
	Msg     string `json:"msg"`
	Success bool   `json:"success"`
}

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

type UserAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type FetchUser struct {
	ID uint32 `json:"id"`
}

type NewUser struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
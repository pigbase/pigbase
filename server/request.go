package main

type OpCodeRequest int8

const (
	// Authorization
	LOGIN = iota

	// Accounts
	CREATE_USER
	CHANGE_NAME_OF_USER
	CHANGE_PASSWORD_OF_USER

	// Collections
	CREATE_COLLECTION
	DELETE_COLLECTION
	EDIT_COLLECTION_NAME

	// Adding value
	ADD_VALUE_TO_COLLECTION
	GET_VALUE_OF_COLLECTION
	EDIT_VALUE_OF_COLLECTION
)

type Request struct {
	Op OpCodeRequest `json:"op"` // Opcode of Request

}

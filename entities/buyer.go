package entities

import "time"

type Buyer struct {
	ID			int
	UserID		int
	User		User
	Nik			string
	Name		string
	Address		string
	PhoneNumber	string
	Status		string
	CreatedAt	time.Time
	UpdatedAt	time.Time
}
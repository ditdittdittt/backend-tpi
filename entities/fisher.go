package entities

import "time"

type Fisher struct {
	ID			int
	UserID		int
	User		User
	Nik			string
	Name		string
	Address		string
	ShipType	string
	AbkTotal	int
	PhoneNumber	string
	Status		string
	CreatedAt	time.Time
	UpdatedAt	time.Time
}
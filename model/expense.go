package model

type ExpenseTracker struct {
	Id string	`csv:"Id"`
	Date string	`csv:"Date"`
	Description string `csv:"Description"`
	Amount string `csv:"Amount"`
}
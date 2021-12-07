package entity

import "errors"

type Transaction struct {
	ID           string
	AccountId    string
	Amount       float64
	CreditCard   CreditCard
	Status       string
	ErrorMessage string
}

func NewTransaction() *Transaction {
	return &Transaction{}
}

func (t *Transaction) IsValid() error {
	if t.Amount > 1000 {
		return errors.New("you dont have limit for this transaction")
	}

	if t.Amount < 1 {
		return errors.New("the amount most be greater than one")
	}

	return nil
}

func (t *Transaction) SetCreditCard(card CreditCard) {
	t.CreditCard = card
}

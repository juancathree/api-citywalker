package domain

type Expenses map[string][]Expense

type Expense struct {
	Amount      float64 `validate:"required" json:"amount" bson:"amount"`
	IsPersonal  bool    `validate:"required" json:"isPersonal" bson:"isPersonal"`
	Description string  `validate:"required,max=30" json:"description" bson:"description"`
}

func (e *Expenses) Initialize(userID *string) {
	*e = make(Expenses)
	(*e)[*userID] = make([]Expense, 0)
}

package entity

type Bankaccount struct {
	ID        string  `gorm:"primaryKey;not null"`
	Account   string  `gorm:"primary_key; type:varchar(50)" json:"account"`
	Bank      string  `json:"bank" gorm:"type:varchar(255)"`
	Name      string  `json:"name" gorm:"type:varchar(255)"`
	Telephone string  `json:"telephone" gorm:"type:varchar(255)"`
	Money     float64 `gorm:"type:float" json:"money"`
	// Transactions []Transaction `gorm:"foreignKey:UserId"`
}

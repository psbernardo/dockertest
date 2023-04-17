package entities

type Person struct {
	ID       int    `gorm:"type:int"`
	Name     string `gorm:"type:varchar(100) NOT NULL"`
	LastName string `gorm:"type:varchar(100) NOT NULL"`
	Age      int    `gorm:"type:varchar(100) NOT NULL"`
}

func (c *Person) TableName() string {
	return "Person"
}

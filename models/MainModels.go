package models

type User struct {
	Id       uint   `gorm:"primaryKey" json:"id"`
	Name     string `gorm:"type:varchar(255)" json:"name"`
	Username string `gorm:"type:varchar(16)" json:"username"`
	Password []byte `gorm:"type:bytea" json:"-"`
}

type Teacher struct {
	Id       uint   `gorm:"primaryKey" json:"id"`
	Name     string `gorm:"type:varchar(255)" json:"name"`
	Username string `gorm:"type:varchar(16)" json:"username"`
	Password []byte `gorm:"type:bytea" json:"-"`
	Course   []Course
}

type Course struct {
	Id      uint `gorm:"primaryKey" json:"id"`
	Content []Content
	Quizz   []Quizz
}

type Content struct {
	Img         string
	Tittle      string
	Description string
}

type Quizz struct {
	Img         string
	Description string
	Answer      []Answer
	Result      int
}

type Answer struct {
	Option1 string
	Option2 string
	Option3 string
	Option4 string
}

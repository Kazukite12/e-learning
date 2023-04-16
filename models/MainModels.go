package models

type User struct {
	ID       int      `gorm:"primaryKey" json:"id"`
	Name     string   `gorm:"type:varchar(255)" json:"name"`
	Username string   `gorm:"type:varchar(16)" json:"username"`
	Password []byte   `gorm:"type:bytea" json:"-"`
	Course   []Course `json:"courses" gorm:"many2many:user_courses"`
	CourseID []int    `json:"course_id" gorm:"-"`
}

type Course struct {
	ID          int    `gorm:"primaryKey" json:"id"`
	Tittle      string `gorm:"type:varchar(16)" json:"tittle"`
	Img         string `gorm:"type:varchar(255)" json:"img"`
	Description string `gorm:"type:varchar(255)" json:"description"`
	Key         string `json:"key"`
}

type UserCourse struct {
	UserID   int `json:"user_id"`
	CourseID int `json:"course_id"`
}

package user

import (
    "gorm.io/gorm"
    "github.com/Retual004/task-service/" // ??? какой путь здесь писать?
)

type User struct {
    gorm.Model
    Email    string `json:"email"`
    Password string `json:"password"`
    DeletedAt gorm.DeletedAt `json:"deleted_at"`
    Tasks []task.Task `gorm:"ForeignKey:UserID"`
}
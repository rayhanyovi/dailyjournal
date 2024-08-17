package models

import "time"


type User struct {
    ID        int64     `xorm:"pk autoincr 'id'"`
    Name      string    `xorm:"'name'"`
    Email     string    `xorm:"'email'"`
    Password  string    `xorm:"'password'"`
    CreatedAt time.Time `xorm:"'created_at'"`
    UpdatedAt time.Time `xorm:"'updated_at'"`
}


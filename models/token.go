package models

import (
    "time"
    "gorm.io/gorm"
    "encoding/json"
    "github.com/stewbawka/go-auth/jwt"
)

var (
    TokenTTL = time.Minute * 5
)

type Token struct {
    ID     uint   `json:"id" gorm:"primary_key"`
    UserID int `json:"user_id"`
    User User
    Token string `json:"token"`
    CreatedAt    time.Time `json:"created_at"`
    UpdatedAt    time.Time `json:"updated_at"`
}

func (t *Token) BeforeSave(tx *gorm.DB) (err error) {
    payload, err := json.Marshal(t.User)
    if err != nil {
        return err
    }
    t.Token, err = jwt.CreateToken(TokenTTL, string(payload))
    return err
}

package models

import (
    "fmt"
    "time"
    "gorm.io/gorm"
    "crypto/rand"
    "encoding/base64"
    "golang.org/x/crypto/argon2"
)

var (
    DB *gorm.DB
    saltLength uint32 = 16
    iterations uint32 = 3
    memory uint32 = 64 * 1024
    parallelism uint8 = 3
    keyLength uint32 = 32
)

type User struct {
    ID     uint   `json:"id" gorm:"primary_key"`
    Email string `json:"email"`
    FirstName string `json:"first_name"`
    LastName string `json:"last_name"`
    Password string `json:"-" gorm:"-"`
    HashedPassword string `json:"-"`
    CreatedAt    time.Time `json:"created_at"`
    UpdatedAt    time.Time `json:"updated_at"`

}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
    fmt.Println("before save")
    if u.Password != "" {
        fmt.Println("updating password")
        if u.HashedPassword, err = HashPassword(u.Password); err != nil {
            return err
        }
    }
    fmt.Println("done")
    return
}

func HashPassword(password string) (encodedHash string, err error) {

    salt, err := generateRandomBytes(saltLength)
    if err != nil {
        return "", err
    }
    hash := argon2.IDKey([]byte(password), salt, iterations, memory, parallelism, keyLength)
    b64Salt := base64.RawStdEncoding.EncodeToString(salt)
    b64Hash := base64.RawStdEncoding.EncodeToString(hash)
    encodedHash = fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, memory, iterations, parallelism, b64Salt, b64Hash)
    return encodedHash, nil
}

func generateRandomBytes(n uint32) ([]byte, error) {
    b := make([]byte, n)
    _, err := rand.Read(b)
    if err != nil {
        return nil, err
    }
    return b, nil
}

package models

import (
    "fmt"
    "time"
    "gorm.io/gorm"
    "crypto/rand"
    "encoding/base64"
    "golang.org/x/crypto/argon2"
    "errors"
    "strings"
    "crypto/subtle"
)

var (
    ErrInvalidHash = errors.New("the encoded hash is not in the correct format")
    ErrIncompatibleVersion = errors.New("incompatible version of argon2")
)

type ArgonParams struct {
    memory uint32
    iterations uint32
    parallelism uint8
    saltLength uint32
    keyLength uint32
}

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
    if u.Password != "" {
        if u.HashedPassword, err = HashPassword(u.Password); err != nil {
            return err
        }
    }
    return
}

func (u *User) ComparePasswordAndHash(password string) (match bool, err error) {
    p, salt, hash, err := decodeHash(u.HashedPassword)
    if err != nil {
        return false, err
    }

    otherHash := argon2.IDKey([]byte(password), salt, p.iterations, p.memory, p.parallelism, p.keyLength)
    if subtle.ConstantTimeCompare(hash, otherHash) == 1 {
        return true, nil
    }
    return false, nil
}

func HashPassword(password string) (encodedHash string, err error) {

    p := &ArgonParams{
        memory: 64 * 1024,
        iterations: 3,
        parallelism: 3,
        saltLength: 16,
        keyLength: 32,
    }
    salt, err := generateRandomBytes(p.saltLength)
    if err != nil {
        return "", err
    }
    hash := argon2.IDKey([]byte(password), salt, p.iterations, p.memory, p.parallelism, p.keyLength)
    b64Salt := base64.RawStdEncoding.EncodeToString(salt)
    b64Hash := base64.RawStdEncoding.EncodeToString(hash)
    encodedHash = fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, p.memory, p.iterations, p.parallelism, b64Salt, b64Hash)
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

func decodeHash(encodedHash string) (p *ArgonParams, salt, hash []byte, err error) {
    vals := strings.Split(encodedHash, "$")
    if len(vals) != 6 {
        return nil, nil, nil, ErrInvalidHash
    }

    var version int
    _, err = fmt.Sscanf(vals[2], "v=%d", &version)
    if err != nil {
        return nil, nil, nil, err
    }

    if version != argon2.Version {
        return nil, nil, nil, ErrIncompatibleVersion
    }

    p = &ArgonParams{}
    _, err = fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &p.memory, &p.iterations, &p.parallelism)
    if err != nil {
        return nil, nil, nil, err
    }

    salt, err = base64.RawStdEncoding.Strict().DecodeString(vals[4])
    if err != nil {
        return nil, nil, nil, err
    }

    p.saltLength = uint32(len(salt))

    hash, err = base64.RawStdEncoding.Strict().DecodeString(vals[5])
    if err != nil {
        return nil, nil, nil, err
    }

    p.keyLength = uint32(len(hash))

    return p, salt, hash, nil
}

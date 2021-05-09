package jwt

import (
    "fmt"
    "io/ioutil"
    "github.com/lestrrat-go/jwx/jwk"
)

type Keypair struct {
    privateKey []byte
    publicKey []byte
}

var (
    JWTKeypair *Keypair
    Jwks *jwk.Key
)

func LoadKeypair(path string) {
    privateKey, err := ioutil.ReadFile(fmt.Sprintf("%s/private-key.pem", path))
    if err != nil  {
        panic(err.Error())
    }

    publicKey, err := ioutil.ReadFile(fmt.Sprintf("%s/public-key.pem", path))
    if err != nil  {
        panic(err.Error())
    }

    keypair := Keypair{privateKey: privateKey, publicKey: publicKey}
    JWTKeypair = &keypair
    GetPublicJwks()
}

func GetPublicJwks() () {
    set, err := jwk.New(JWTKeypair.publicKey)
    if err != nil {
        panic(err.Error())
    }

    err = jwk.AssignKeyID(set)
    if err != nil {
        panic(err.Error())
    }
    Jwks = &set
}


package jwt

import (
    "fmt"
    "io/ioutil"
    "github.com/lestrrat-go/jwx/jwk"
    "encoding/pem"
    "crypto/x509"
    "errors"
)

type Keypair struct {
    privateKey []byte
    publicKey []byte
}

type JWKeys struct {
    Keys [1]jwk.Key `json:"keys"`
}

var (
    JWTKeypair *Keypair
    Jwks *JWKeys
    JwkKid string
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
    block, _ := pem.Decode([]byte(JWTKeypair.publicKey))
    if block == nil {
        panic(errors.New("failed to parse PEM block containing the key"))
    }

    pub, err := x509.ParsePKIXPublicKey(block.Bytes)
    if err != nil {
        panic(err.Error())
    }

    set, err := jwk.New(pub)
    if err != nil {
        panic(err.Error())
    }

    err = jwk.AssignKeyID(set)
    if err != nil {
        panic(err.Error())
    }
    keysArr := [1]jwk.Key{set}
    JwkKid = set.KeyID()
    Jwks = &JWKeys{Keys: keysArr}
}


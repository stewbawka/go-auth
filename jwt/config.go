package jwt

import (
    "fmt"
    "io/ioutil"
)

type Keypair struct {
    privateKey []byte
    publicKey []byte
}

var (
    JWTKeypair *Keypair
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
}

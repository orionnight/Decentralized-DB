package services

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func CreateKS(password string) {
	ks := keystore.NewKeyStore("./keystore", keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := ks.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex()) // 0x20F8D42FB0F667F2E53930fed426f225752453b3
	fmt.Println(ks.Accounts()[0])      // 0x20F8D42FB0F667F2E53930fed426f225752453b3
	fmt.Println(ks.Wallets()[0])       // 0x20F8D42FB0F667F2E53930fed426f225752453b3
}

func importKs(filename string, password string) {
	file := "./keystore/" + filename
	ks := keystore.NewKeyStore("./keystore", keystore.StandardScryptN, keystore.StandardScryptP)
	jsonBytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	account, err := ks.Import(jsonBytes, password, password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex()) // 0x20F8D42FB0F667F2E53930fed426f225752453b3

	if err := os.Remove(file); err != nil {
		log.Fatal(err)
	}
}

// func main() {
// 	createKs("secret")
// 	importKs("UTC--2022-02-20T04-37-25.914168271Z--16632e5403a2daed33c98247d02789b83e3cc313", "secret")
// }

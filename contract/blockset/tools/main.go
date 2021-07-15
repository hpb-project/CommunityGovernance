package main

import (
	"crypto/ecdsa"
	"flag"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/xueqianLu/testsol/govermentSet/contracts"
	"log"
	"math/big"
	"strings"
)

var (
	defaultgas, _      = new(big.Int).SetString("100000", 10)
	defaultgasprice, _ = new(big.Int).SetString("18000000000", 10)
)

func getAddrFromPrivkey(priv *ecdsa.PrivateKey) common.Address {
	publicKey := priv.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Println("from address", fromAddress.String())
	return fromAddress
}

func main() {
	url := flag.String("u", "http://127.0.0.1:8545", "rpc url")
	senderPrivKey := flag.String("priv", "", "Sender private key in hex")
	addr := flag.String("addr", "","contract address")
	set := flag.String("set","","set value with key,val pair")
	get := flag.String("get", "", "the key witch to query")
	flag.Parse()

	client := NewHttpClient(*url)
	chainId := client.ChainID()
	contractAddress := common.HexToAddress(*addr)

	if len(*senderPrivKey) == 0 {
		log.Fatal("sender private key required.")
	}

	privateKey, err := crypto.HexToECDSA(*senderPrivKey)
	if err != nil {
		log.Fatal(err)
	}
	fromAddress := getAddrFromPrivkey(privateKey)
	nonce := client.GetNonce(fromAddress.String())
	if err != nil {
		nonce = 0
	}

	fmt.Println("NonceAt", nonce)


	m := strings.Split(*set, ",")

	if len(m) == 2 {
		// set k with val
		k,val := m[0], m[1]
		v,_ := big.NewInt(0).SetString(val, 10)
		fmt.Printf("set k(%s) = (%d)\n ", k, v)
		parsed, err := abi.JSON(strings.NewReader(contracts.ContractsABI))
		if err != nil {
			log.Fatal(err)
		}
		data, err := parsed.Pack("setValue", k, v)
		if err != nil {
			log.Fatal(err)
		}

		tx := types.NewTransaction(nonce, contractAddress, big.NewInt(0), defaultgas.Uint64(), defaultgasprice, data)
		signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainId), privateKey)
		if err != nil {
			log.Fatal(err)
		}
		txhash, err := client.SendSignedTx(signedTx)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("tx sent: %s\n", txhash)

		return
	}
	if len(*get) > 0 {
		err = getValue(*addr, *get, client.eth)
		if err != nil {
			log.Fatal("get value error with ", err)
		}
	}
}
//
//func setValue(conAddr string, key string, val uint, client *ethclient.Client) error {
//	cAddr := common.HexToAddress(conAddr)
//	defaultOpt := &bind.TransactOpts{}
//
//	govermentSet, err := contracts.NewContracts(cAddr, client)
//	if err != nil {
//		//log.Println("newErc20 failed")
//		return err
//	}
//
//
//	value, err := govermentSet.GetValue(defaultOpt, key)
//	if err != nil {
//		//log.Println("coin name failed")
//		return err
//	}
//	log.Printf(" key %s ===> %d", value)
//	return nil
//}
//
//

func getValue(conAddr string, key string, client *ethclient.Client) error {
	cAddr := common.HexToAddress(conAddr)
	defaultOpt := &bind.CallOpts{}

	govermentSet, err := contracts.NewContracts(cAddr, client)
	if err != nil {
		//log.Println("newErc20 failed")
		return err
	}

	value, err := govermentSet.GetValue(defaultOpt, key)
	if err != nil {
		//log.Println("coin name failed")
		return err
	}
	log.Printf(" key %s ===> %d", key, value.Uint64())
	return nil
}



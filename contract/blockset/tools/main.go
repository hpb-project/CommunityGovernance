package main

import (
	"crypto/ecdsa"
	"flag"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"strings"
)

var (
	defaultgas, _      = new(big.Int).SetString("10000000", 10)
	defaultgasprice, _ = new(big.Int).SetString("18000000000", 10)

	user = privateAccount{}
)

type privateAccount struct {
	privk   *ecdsa.PrivateKey
	addr    common.Address
	chainid *big.Int
}

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

func SimpleInfo(conAddr string, client *ethclient.Client) {
	fmt.Println("current simple info")
	GetOwner(conAddr, client)
	GetAdmins(conAddr, client)
}

func main() {
	url := flag.String("u", "http://127.0.0.1:8545", "rpc url")
	senderPrivKey := flag.String("priv", "", "Sender private key in hex")
	addr := flag.String("addr", "", "contract address")
	// admin
	addAdmin := flag.String("addAdmin", "", "new admin address")
	delAdmin := flag.String("delAdmin", "", "delete admin address")
	//admins := flag.Bool("listAdmis", false, "list all admins")

	// owner
	chowner := flag.String("owner", "", "new owner address")

	// proposal
	resetProposal := flag.String("reset", "", "reset proposal key,val pair")
	addProposal := flag.String("proposal", "", "set value with key,val pair")
	voteProposal := flag.String("vote", "", "vote proposal key")
	// get value
	get := flag.String("get", "", "the key witch to query")
	flag.Parse()

	var err error
	client := NewHttpClient(*url)
	chainId := client.ChainID()
	if len(*addr) == 0 {
		log.Fatal("no contract address")
	}

	SimpleInfo(*addr, client.eth)

	if len(*senderPrivKey) > 0 {
		privateKey, err := crypto.HexToECDSA(*senderPrivKey)
		if err != nil {
			log.Fatal(err)
		}
		fromAddress := getAddrFromPrivkey(privateKey)

		user.privk = privateKey
		user.addr = fromAddress
		user.chainid = chainId
	}
	if len(*addAdmin) > 0 {
		err = AddAdmin(*addr, *addAdmin, client.eth)
		if err != nil {
			log.Fatal("err ", err)
		}
	}

	if len(*delAdmin) > 0 {
		err = DelAdmin(*addr, *delAdmin, client.eth)
		if err != nil {
			log.Fatal("err ", err)
		}
	}

	if len(*chowner) > 0 {
		err = ChangeOwner(*addr, *chowner, client.eth)
		if err != nil {
			log.Fatal("err ", err)
		}
	}

	if len(*resetProposal) > 0 {
		m := strings.Split(*resetProposal, ",")
		k, val := m[0], m[1]
		err = ResetProposal(*addr, k, val, client.eth)
		if err != nil {
			log.Fatal("err ", err)
		}
	}

	if len(*addProposal) > 0 {
		m := strings.Split(*addProposal, ",")
		k, val := m[0], m[1]
		err = AddProposal(*addr, k, val, client.eth)
		if err != nil {
			log.Fatal("err ", err)
		}
	}

	if len(*voteProposal) > 0 {
		err = VoteProposal(*addr, *voteProposal, client.eth)
		if err != nil {
			log.Fatal("err ", err)
		}
	}

	if len(*get) > 0 {
		err = GetValue(*addr, *get, client.eth)
		if err != nil {
			log.Fatal("err ", err)
		}
	}
}

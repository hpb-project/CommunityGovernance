package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"gotest.tools/assert"
	"log"
	"sync"
	"testing"
)

var (
	contractAddr = ""
	rpc          = "http://114.242.26.15:8006"
	privk        = ""
	client       = NewHttpClient(rpc)
	once         = sync.Once{}
)

func firstinit() {
	privateKey, _ := crypto.HexToECDSA(privk)
	fromAddress := getAddrFromPrivkey(privateKey)
	user.privk = privateKey
	user.addr = fromAddress
	user.chainid = client.ChainID()

	contractAddr, _ = deploy()
	log.Println("deploy new contract at ", contractAddr)
}

func deploy() (string, error) {
	addr, tx, err := Deploy(client.eth)
	if err != nil {
		log.Println("deploy failed ", err)
		return "", err
	}
	log.Println("deploy tx :", tx.Hash())
	return addr.String(), nil
	//	return "", errors.New("deploy failed")
}

func testGetAdmins(t *testing.T, n int) {
	admins, err := GetAdmins(contractAddr, client.eth)
	assert.Assert(t, err == nil, fmt.Sprintf("get admins failed:%s", err))
	assert.Assert(t, len(admins) == n, fmt.Sprintf("miss match admins number %d", len(admins)))
}

func TestAddAdmin(t *testing.T) {
	once.Do(firstinit)
	nAdmin := "0x69DC6E2990C73B658DcbAB841c630F082AAAeD2D"
	err := AddAdmin(contractAddr, nAdmin, client.eth)
	assert.Assert(t, err == nil, fmt.Sprintf("add admin failed:%s", err))
	testGetAdmins(t, 1)
}

func TestDelAdmin(t *testing.T) {
	once.Do(firstinit)
	nAdmin := "0x69DC6E2990C73B658DcbAB841c630F082AAAeD2D"
	err := DelAdmin(contractAddr, nAdmin, client.eth)
	assert.Assert(t, err == nil, fmt.Sprintf("del admin failed:%s", err))
	testGetAdmins(t, 0)
}

package main

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/xueqianLu/testsol/govermentSet/contracts"
	"log"
	"math/big"
)

func GetAdmins(conAddr string, client *ethclient.Client) error {
	cAddr := common.HexToAddress(conAddr)
	defaultOpt := &bind.CallOpts{}

	govermentSet, err := contracts.NewContracts(cAddr, client)
	if err != nil {
		//log.Println("newErc20 failed")
		return err
	}

	admins, err := govermentSet.GetAdmins(defaultOpt)
	if err != nil {
		//log.Println("coin name failed")
		return err
	}
	log.Printf(" total admins %d\n", len(admins))
	for _, a := range admins {
		log.Printf("admins-->%s\n", a.String())
	}
	return nil
}

func GetOwner(conAddr string, client *ethclient.Client) error {
	cAddr := common.HexToAddress(conAddr)
	defaultOpt := &bind.CallOpts{}

	govermentSet, err := contracts.NewContracts(cAddr, client)
	if err != nil {
		//log.Println("newErc20 failed")
		return err
	}

	owner, err := govermentSet.GetOwner(defaultOpt)
	if err != nil {
		//log.Println("coin name failed")
		return err
	}
	log.Printf(" current owner is %s\n", owner.String())
	return nil
}

func GetValue(conAddr string, key string, client *ethclient.Client) error {
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

func GetThreshold(conAddr string, key string, client *ethclient.Client) error {
	cAddr := common.HexToAddress(conAddr)
	defaultOpt := &bind.CallOpts{}

	govermentSet, err := contracts.NewContracts(cAddr, client)
	if err != nil {
		//log.Println("newErc20 failed")
		return err
	}

	if len(key) == 0 {
		value, err := govermentSet.GetThreshold(defaultOpt)
		if err != nil {
			//log.Println("coin name failed")
			return err
		}
		log.Printf(" threshold %d", key, value.Uint64())
	} else {
		value, err := govermentSet.GetProposolThreshold(defaultOpt, key)
		if err != nil {
			//log.Println("coin name failed")
			return err
		}
		log.Printf(" proposal %s ===> %d", key, value.Uint64())
	}
	return nil
}

func signFunc(addr common.Address, tx *types.Transaction) (*types.Transaction, error) {
	return types.SignTx(tx, types.NewEIP155Signer(user.chainid), user.privk)
}

func ChangeOwner(conAddr string, newOwner string, client *ethclient.Client) error {
	cAddr := common.HexToAddress(conAddr)
	addr := common.HexToAddress(newOwner)
	defaultOpt := &bind.TransactOpts{
		From:   user.addr,
		Signer: signFunc,
	}

	govermentSet, err := contracts.NewContracts(cAddr, client)
	if err != nil {
		//log.Println("newErc20 failed")
		return err
	}

	tx, err := govermentSet.ChangeOwner(defaultOpt, addr)
	if err != nil {
		//log.Println("coin name failed")
		return err
	}
	log.Printf(" change new owner tx %s\n", tx.Hash().String())
	return nil
}

func AddAdmin(conAddr string, admin string, client *ethclient.Client) error {
	cAddr := common.HexToAddress(conAddr)
	addr := common.HexToAddress(admin)
	defaultOpt := &bind.TransactOpts{
		From:   user.addr,
		Signer: signFunc,
	}

	govermentSet, err := contracts.NewContracts(cAddr, client)
	if err != nil {
		//log.Println("newErc20 failed")
		return err
	}

	tx, err := govermentSet.AddAdmin(defaultOpt, addr)
	if err != nil {
		//log.Println("coin name failed")
		return err
	}
	log.Printf(" add admin tx %s\n", tx.Hash().String())
	return nil
}

func DelAdmin(conAddr string, admin string, client *ethclient.Client) error {
	cAddr := common.HexToAddress(conAddr)
	addr := common.HexToAddress(admin)
	defaultOpt := &bind.TransactOpts{
		From:   user.addr,
		Signer: signFunc,
	}

	govermentSet, err := contracts.NewContracts(cAddr, client)
	if err != nil {
		//log.Println("newErc20 failed")
		return err
	}

	tx, err := govermentSet.DeleteAdmin(defaultOpt, addr)
	if err != nil {
		//log.Println("coin name failed")
		return err
	}
	log.Printf(" delete admin tx %s\n", tx.Hash().String())
	return nil
}

func AddProposal(conAddr string, key, val string, client *ethclient.Client) error {
	cAddr := common.HexToAddress(conAddr)
	pv, _ := new(big.Int).SetString(val, 10)
	defaultOpt := &bind.TransactOpts{
		From:   user.addr,
		Signer: signFunc,
	}

	govermentSet, err := contracts.NewContracts(cAddr, client)
	if err != nil {
		//log.Println("newErc20 failed")
		return err
	}

	tx, err := govermentSet.AddProposal(defaultOpt, key, pv)
	if err != nil {
		//log.Println("coin name failed")
		return err
	}
	log.Printf(" add proposal tx %s\n", tx.Hash().String())
	return nil
}

func ResetProposal(conAddr string, key, val string, client *ethclient.Client) error {
	cAddr := common.HexToAddress(conAddr)
	pv, _ := new(big.Int).SetString(val, 10)
	defaultOpt := &bind.TransactOpts{
		From:   user.addr,
		Signer: signFunc,
	}

	govermentSet, err := contracts.NewContracts(cAddr, client)
	if err != nil {
		//log.Println("newErc20 failed")
		return err
	}

	tx, err := govermentSet.ResetProposal(defaultOpt, key, pv)
	if err != nil {
		//log.Println("coin name failed")
		return err
	}
	log.Printf(" reset proposal tx %s\n", tx.Hash().String())
	return nil
}

func VoteProposal(conAddr string, key string, client *ethclient.Client) error {
	cAddr := common.HexToAddress(conAddr)
	defaultOpt := &bind.TransactOpts{
		From:   user.addr,
		Signer: signFunc,
	}

	govermentSet, err := contracts.NewContracts(cAddr, client)
	if err != nil {
		//log.Println("newErc20 failed")
		return err
	}

	tx, err := govermentSet.VoteProposal(defaultOpt, key)
	if err != nil {
		//log.Println("coin name failed")
		return err
	}
	log.Printf(" vote proposal tx %s\n", tx.Hash().String())
	return nil
}

func SetThreshold(conAddr string, val string, client *ethclient.Client) error {
	cAddr := common.HexToAddress(conAddr)
	pv, _ := new(big.Int).SetString(val, 10)
	defaultOpt := &bind.TransactOpts{
		From:   user.addr,
		Signer: signFunc,
	}

	govermentSet, err := contracts.NewContracts(cAddr, client)
	if err != nil {
		//log.Println("newErc20 failed")
		return err
	}

	tx, err := govermentSet.SetThreshold(defaultOpt, pv)
	if err != nil {
		//log.Println("coin name failed")
		return err
	}
	log.Printf(" vote proposal tx %s\n", tx.Hash().String())
	return nil
}

package main

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/xueqianLu/testsol/govermentSet/contracts/blockset"
	"github.com/xueqianLu/testsol/govermentSet/contracts/proxy"
	"log"
	"math/big"
	"time"
)

func waitTx(tx *types.Transaction, client *ethclient.Client) error {
	try := 10
	for try > 0 {
		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			try--
			time.Sleep(time.Millisecond * 400)
			continue
		} else {
			if receipt.Status == 1 {
				return nil
			} else {
				return errors.New("deploy failed")
			}
		}
	}
	return errors.New("deploy failed")
}

func GetAdmins(conAddr string, client *ethclient.Client) ([]common.Address, error) {
	cAddr := common.HexToAddress(conAddr)
	defaultOpt := &bind.CallOpts{}

	govermentSet, err := blockset.NewBlockset(cAddr, client)
	if err != nil {
		//log.Println("newErc20 failed")
		return nil, err
	}

	admins, err := govermentSet.GetAdmins(defaultOpt)
	if err != nil {
		//log.Println("coin name failed")
		return nil, err
	}
	log.Printf(" total admins %d\n", len(admins))
	for _, a := range admins {
		log.Printf("admins-->%s\n", a.String())
	}
	return admins, nil
}

func GetOwner(conAddr string, client *ethclient.Client) error {
	cAddr := common.HexToAddress(conAddr)
	defaultOpt := &bind.CallOpts{}

	govermentSet, err := blockset.NewBlockset(cAddr, client)
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
func ProxyGetValue(conAddr string, key string, client *ethclient.Client) error {
	cAddr := common.HexToAddress(conAddr)
	defaultOpt := &bind.CallOpts{}

	Proxy, err := proxy.NewProxy(cAddr, client)
	if err != nil {
		//log.Println("newErc20 failed")
		return err
	}

	value, err := Proxy.GetValue(defaultOpt, key)
	if err != nil {
		//log.Println("coin name failed")
		return err
	}
	log.Printf(" key %s ===> %d", key, value.Uint64())
	return nil
}

func ProxyGetContract(conAddr string, client *ethclient.Client) error {
	cAddr := common.HexToAddress(conAddr)
	defaultOpt := &bind.CallOpts{}

	Proxy, err := proxy.NewProxy(cAddr, client)
	if err != nil {
		//log.Println("newErc20 failed")
		return err
	}

	addr, err := Proxy.Getcontract(defaultOpt)
	if err != nil {
		//log.Println("coin name failed")
		return err
	}
	log.Printf(" current contract addr is %s", addr)
	return nil
}

func ProxySetContract(paddr string, blocksetAddr string, client *ethclient.Client) error {
	cAddr := common.HexToAddress(paddr)
	addr := common.HexToAddress(blocksetAddr)
	defaultOpt := getTransactOpts()

	Proxy, err := proxy.NewProxy(cAddr, client)
	if err != nil {
		//log.Println("newErc20 failed")
		return err
	}

	tx, err := Proxy.Setcontract(defaultOpt, addr)
	if err != nil {
		//log.Println("coin name failed")
		return err
	}
	waitTx(tx, client)
	log.Printf(" set contract tx %s\n", tx.Hash().String())
	return nil
}

func GetValue(conAddr string, key string, client *ethclient.Client) error {
	cAddr := common.HexToAddress(conAddr)
	defaultOpt := &bind.CallOpts{}

	govermentSet, err := blockset.NewBlockset(cAddr, client)
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

	govermentSet, err := blockset.NewBlockset(cAddr, client)
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
		log.Printf(" threshold %d", value.Uint64())
	} else {
		value, err := govermentSet.GetProposalThreshold(defaultOpt, key)
		if err != nil {
			//log.Println("coin name failed")
			return err
		}
		log.Printf(" proposal %s threshold %d", key, value.Uint64())
	}
	return nil
}

func signFunc(addr common.Address, tx *types.Transaction) (*types.Transaction, error) {
	return types.SignTx(tx, types.NewEIP155Signer(user.chainid), user.privk)
}

func ChangeOwner(conAddr string, newOwner string, client *ethclient.Client) error {
	cAddr := common.HexToAddress(conAddr)
	addr := common.HexToAddress(newOwner)
	defaultOpt := getTransactOpts()

	govermentSet, err := blockset.NewBlockset(cAddr, client)
	if err != nil {
		//log.Println("newErc20 failed")
		return err
	}

	tx, err := govermentSet.ChangeOwner(defaultOpt, addr)
	if err != nil {
		//log.Println("coin name failed")
		return err
	}
	waitTx(tx, client)
	log.Printf(" change new owner tx %s\n", tx.Hash().String())
	return nil
}

func AddAdmin(conAddr string, admin string, client *ethclient.Client) error {
	cAddr := common.HexToAddress(conAddr)
	addr := common.HexToAddress(admin)
	defaultOpt := getTransactOpts()

	govermentSet, err := blockset.NewBlockset(cAddr, client)
	if err != nil {
		//log.Println("newErc20 failed")
		return err
	}

	tx, err := govermentSet.AddAdmin(defaultOpt, addr)
	if err != nil {
		//log.Println("coin name failed")
		return err
	}
	waitTx(tx, client)
	log.Printf(" add admin tx %s\n", tx.Hash().String())
	return nil
}

func DelAdmin(conAddr string, admin string, client *ethclient.Client) error {
	cAddr := common.HexToAddress(conAddr)
	addr := common.HexToAddress(admin)
	defaultOpt := getTransactOpts()
	govermentSet, err := blockset.NewBlockset(cAddr, client)
	if err != nil {
		//log.Println("newErc20 failed")
		return err
	}

	tx, err := govermentSet.DeleteAdmin(defaultOpt, addr)
	if err != nil {
		//log.Println("coin name failed")
		return err
	}
	waitTx(tx, client)
	log.Printf(" delete admin tx %s\n", tx.Hash().String())
	return nil
}

func AddProposal(conAddr string, key, val string, client *ethclient.Client) error {
	cAddr := common.HexToAddress(conAddr)
	pv, _ := new(big.Int).SetString(val, 10)
	defaultOpt := getTransactOpts()

	govermentSet, err := blockset.NewBlockset(cAddr, client)
	if err != nil {
		//log.Println("newErc20 failed")
		return err
	}

	tx, err := govermentSet.AddProposal(defaultOpt, key, pv)
	if err != nil {
		//log.Println("coin name failed")
		return err
	}
	waitTx(tx, client)
	log.Printf(" add proposal tx %s\n", tx.Hash().String())
	return nil
}

func ResetProposal(conAddr string, key, val string, client *ethclient.Client) error {
	cAddr := common.HexToAddress(conAddr)
	pv, _ := new(big.Int).SetString(val, 10)
	defaultOpt := getTransactOpts()

	govermentSet, err := blockset.NewBlockset(cAddr, client)
	if err != nil {
		//log.Println("newErc20 failed")
		return err
	}

	tx, err := govermentSet.ResetProposal(defaultOpt, key, pv)
	if err != nil {
		//log.Println("coin name failed")
		return err
	}
	waitTx(tx, client)
	log.Printf(" reset proposal tx %s\n", tx.Hash().String())
	return nil
}

func VoteProposal(conAddr string, key string, client *ethclient.Client) error {
	cAddr := common.HexToAddress(conAddr)
	defaultOpt := getTransactOpts()

	govermentSet, err := blockset.NewBlockset(cAddr, client)
	if err != nil {
		//log.Println("newErc20 failed")
		return err
	}

	tx, err := govermentSet.VoteProposal(defaultOpt, key)
	if err != nil {
		//log.Println("coin name failed")
		return err
	}
	waitTx(tx, client)
	log.Printf(" vote proposal tx %s\n", tx.Hash().String())
	return nil
}

func SetThreshold(conAddr string, val string, client *ethclient.Client) error {
	cAddr := common.HexToAddress(conAddr)
	pv, _ := new(big.Int).SetString(val, 10)
	defaultOpt := getTransactOpts()

	govermentSet, err := blockset.NewBlockset(cAddr, client)
	if err != nil {
		//log.Println("newErc20 failed")
		return err
	}

	tx, err := govermentSet.SetThreshold(defaultOpt, pv)
	if err != nil {
		//log.Println("coin name failed")
		return err
	}
	waitTx(tx, client)
	log.Printf(" vote proposal tx %s\n", tx.Hash().String())
	return nil
}

func Deploy(client *ethclient.Client) (common.Address, *types.Transaction, error) {
	defaultOpt := getTransactOpts()
	addr, tx, _, err := blockset.DeployBlockset(defaultOpt, client)
	if err != nil {
		return common.Address{}, nil, err
	}
	waitTx(tx, client)
	return addr, tx, nil
}

func getTransactOpts() *bind.TransactOpts {
	return &bind.TransactOpts{
		From:     user.addr,
		Signer:   signFunc,
		Context:  context.Background(),
		GasPrice: defaultgasprice,
		GasLimit: defaultgas.Uint64(),
	}
}

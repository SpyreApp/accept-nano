package main

import (
	"github.com/SpyreApp/accept-nano/internal/nano"
	"github.com/cenkalti/log"
	"github.com/shopspring/decimal"
)

func sendAll(account, destination, privateKey string) error {
	log.Debugln("sending from", account)
	// Get the account info for the payment address
	info, err := node.AccountInfo(account)
	if err != nil {
		return err
	}
	if info.Balance.IsZero() {
		return nil
	}

	work, err := nano.GenerateWork(info.Frontier, true)
	if err != nil {
		return err
	}

	block, err := node.BlockCreate(info.Frontier, account, config.Representative, decimal.Zero, destination, privateKey, work)
	if err != nil {
		return err
	}
	hash, err := node.Process(block)
	if err != nil {
		return err
	}
	log.Debugln("published new block:", hash)
	return nil
}

func sendTax(account, destination, privateKey string, balance decimal.Decimal) error {
	log.Debugln("sending from", account)
	info, err := node.AccountInfo(account)
	if err != nil {
		return err
	}
	if info.Balance.IsZero() {
		return nil
	}
	work, err := nano.GenerateWork(info.Frontier, true)
	if err != nil {
		return err
	}

	block, err := node.BlockCreate(info.Frontier, account, config.Representative, balance, destination, privateKey, work)
	if err != nil {
		return err
	}
	hash, err := node.Process(block)
	if err != nil {
		return err
	}
	log.Debugln("published new block:", hash)
	return nil
}

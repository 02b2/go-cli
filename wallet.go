package main

import (
	"crypto/sha256"
	"log"
	"strconv"

	"fyne.io/fyne/v2/widget"
	"github.com/SixofClubsss/dReams/rpc"
	dero "github.com/deroproject/derohe/rpc"
)

const (
	WALLET_MAINNET_DEFAULT   = "127.0.0.1:10103"
	WALLET_TESTNET_DEFAULT   = "127.0.0.1:40403"
	WALLET_SIMULATOR_DEFAULT = "127.0.0.1:30000"
)

var (
	walletAddress      string
	walletConnect      bool
	passHash           [32]byte
	rpcLoginInput      *widget.Entry
	walletCheckBox     *widget.Check
	walletBalance      *widget.Label
	walletAddressLabel *widget.Label
	kill_process       bool
	quit               chan struct{}
	daemonStatusLabel  *widget.Label
)



func isWalletConnected() {
    if walletConnect {
        walletCheckBox.SetChecked(true)
    } else {
        walletCheckBox.SetChecked(false)
    }
}


func GetAddress() error {
	rpcClientW, ctx, cancel := rpc.SetWalletClient(walletAddress, rpcLoginInput.Text)
	defer cancel()

	var result *dero.GetAddress_Result
	err := rpcClientW.CallFor(ctx, &result, "GetAddress")

	if err != nil {
		walletConnect = false
		walletCheckBox.SetChecked(false)
		log.Println("[GetAddress]", err)
		return nil
	}

	address := len(result.Address)
	if address == 66 {
		walletConnect = true
		walletCheckBox.SetChecked(true)
		walletAddressLabel.SetText("Address: " + result.Address)
		log.Println("Wallet Connected")
		log.Println("Address: " + result.Address)
		data := []byte(rpcLoginInput.Text)
		passHash = sha256.Sum256(data)
	}

	return err
}



func GetBalance() error {
	rpcClientW, ctx, cancel := rpc.SetWalletClient(walletAddress, rpcLoginInput.Text)
	defer cancel()

	var result *dero.GetBalance_Result
	err := rpcClientW.CallFor(ctx, &result, "GetBalance")

	if err != nil {
		log.Println("[GetBalance]", err)
		return nil
	}

	atomic := float64(result.Unlocked_Balance)
	div := atomic / 100000
	str := strconv.FormatFloat(div, 'f', 5, 64)
	walletBalance.SetText("Balance: " + str)
	log.Println("Wallet Balance: " + str)

	return err
}

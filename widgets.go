package main

import (
	"log"

	"fyne.io/fyne/v2"

	"fyne.io/fyne/v2/widget"
)

// / declare some widgets
var (
	primes   = []string{"MAINNET", "TESTNET", "SIMULATOR", "CUSTOM"} /// set select menu
	dropDown = widget.NewSelect(primes, func(s string) {             /// do when select changes
		log.Println("Daemon Set To:", s)
	})

	rpcWalletInput = widget.NewEntry()
	contractInput  = widget.NewMultiLineEntry()

	daemonCheckBox = widget.NewCheck("Daemon Connected", func(value bool) {

	})
	
	
)

func rpcLoginEdit() fyne.Widget { /// user:pass password entry
	rpcLoginInput.SetPlaceHolder("RPC user:pass")
	rpcLoginInput.Resize(fyne.NewSize(360, 45))
	rpcLoginInput.Move(fyne.NewPos(10, 650))

	return rpcLoginInput
}

func rpcWalletEdit() fyne.Widget { /// wallet rpc address entry
	rpcWalletInput.SetPlaceHolder("Wallet RPC Address")
	rpcWalletInput.Resize(fyne.NewSize(250, 45))
	rpcWalletInput.Move(fyne.NewPos(10, 700))

	return rpcWalletInput
}

func rpcConnectButton() fyne.Widget { /// wallet connect button
	button := widget.NewButton("Connect", func() { /// do on pressed
		walletAddress = rpcWalletInput.Text
		GetAddress()
	})
	button.Resize(fyne.NewSize(100, 42))
	button.Move(fyne.NewPos(270, 702))

	return button
}


func daemonSelectOption() fyne.Widget { /// daemon select menu
	dropDown.SetSelectedIndex(0)
	dropDown.Resize(fyne.NewSize(180, 45))
	dropDown.Move(fyne.NewPos(10, 550))

	return dropDown
}

func daemonConnectBox() fyne.Widget { /// daemon check box
	daemonCheckBox.Resize(fyne.NewSize(30, 30))
	daemonCheckBox.Move(fyne.NewPos(3, 595))
	daemonCheckBox.Disable()

	return daemonCheckBox
}








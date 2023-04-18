package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"
	"time"

	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)


func main() {

	// Define and parse command-line arguments for network type and wallet address
	networkType := flag.String("network", "mainnet", "Network type: mainnet, testnet, or simulator")
	walletAddressArg := flag.String("wallet", "", "Wallet address")
	rpcUserArg := flag.String("rpcuser", "", "Wallet RPC username")
	rpcPasswordArg := flag.String("rpcpassword", "", "Wallet RPC password")
	flag.Parse()

	a := app.New()
	w := a.NewWindow("GUI")
	daemonCheckBox = widget.NewCheck("Daemon Connected", func(bool) {})
    daemonCheckBox.Disable()
	// Handle ctrl-c close
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		kill_process = true
		stopLoop()
		fmt.Println()
		os.Exit(1)
	}()

	
	
	

	// Fetch daemon and wallet status on interval
	fetchLoop()

	

	// Set up wallet connection, input and balance widgets
	rpcLoginInput = widget.NewEntry()
	rpcLoginInput.SetPlaceHolder("Enter Wallet RPC Username:Password")
	walletCheckBox = widget.NewCheck("Wallet Connected", func(bool) {})
	walletCheckBox.Disable()
	walletBalance = widget.NewLabel("Balance: ")
	walletAddressLabel = widget.NewLabel("Address: ")
	daemonStatusLabel = widget.NewLabel("")
	walletAddressInput := widget.NewEntry()
	walletAddressInput.SetPlaceHolder("Enter Wallet Address")
	
	

		// Use the wallet address from the command-line argument if provided
		if *walletAddressArg != "" {
			walletAddressInput.SetText(*walletAddressArg)
		}

	// Create a dropdown for selecting the network
	networkSelect := widget.NewSelect([]string{"Mainnet", "Testnet", "Simulator"}, func(selected string) {
		switch selected {
		case "Mainnet":
			
			daemonAddress = DAEMON_MAINNET_DEFAULT
		case "Testnet":
			
			daemonAddress = DAEMON_TESTNET_DEFAULT
		case "Simulator":
			
			daemonAddress = DAEMON_SIMULATOR_DEFAULT
		}
	})
	networkSelect.SetSelected("Mainnet")

		// Set the default network based on the command-line argument
		switch *networkType {
		case "mainnet":
			networkSelect.SetSelected("Mainnet")
		case "testnet":
			networkSelect.SetSelected("Testnet")
		case "simulator":
			networkSelect.SetSelected("Simulator")
		default:
			fmt.Println("Invalid network type. Please use 'mainnet', 'testnet', or 'simulator'.")
			return
		}

		// New button for opening text files
		btnOpenTxtFile := widget.NewButton("Open .txt files", func() {
			fileDialog := dialog.NewFileOpen(func(r fyne.URIReadCloser, _ error) {
				data, _ := ioutil.ReadAll(r)
				result := fyne.NewStaticResource("name", data)
				entry := widget.NewMultiLineEntry()
				entry.SetText(string(result.StaticContent))
				w := fyne.CurrentApp().NewWindow(string(result.StaticName))
				w.SetContent(container.NewScroll(entry))
				w.Resize(fyne.NewSize(400, 400))
				w.Show()
			}, w)
			fileDialog.SetFilter(storage.NewExtensionFileFilter([]string{".txt"}))
			fileDialog.Show()
		})

		btnConnect := widget.NewButton("Connect and Get Balance", func() {
		// Use the input wallet address if provided, otherwise use the default address
		if walletAddressInput.Text != "" {
			walletAddress = *walletAddressArg
		}

		err := GetAddress()
		if err == nil {
			err = GetBalance()
		}
	})

	// Check if the wallet address and RPC login details are provided, and connect to the daemon and wallet automatically
	if *walletAddressArg != "" && *rpcUserArg != "" && (*rpcPasswordArg != "" || *networkType == "simulator") {
		walletAddress = walletAddressInput.Text
		rpcLoginInput.Text = fmt.Sprintf("%s:%s", *rpcUserArg, *rpcPasswordArg)

		err := GetAddress()
		if err == nil {
			err = GetBalance()
		}
	}

	// Add the wallet connection and balance widgets to a container
	walletBox := container.NewVBox(rpcLoginInput, networkSelect, walletAddressInput, walletCheckBox, walletBalance, walletAddressLabel, btnConnect)
	// User input form
	input1 := widget.NewEntry()
	input2 := widget.NewEntry()
	input3 := widget.NewEntry()

	label1 := widget.NewLabel("Name:")
	label2 := widget.NewLabel("Email:")
	label3 := widget.NewLabel("Phone:")

	nameLabel := widget.NewLabel("")
	emailLabel := widget.NewLabel("")
	phoneLabel := widget.NewLabel("")
	

	btn1 := widget.NewButton("Submit", func() {
		name := input1.Text
		email := input2.Text
		phone := input3.Text

		nameLabel.SetText(name)
		emailLabel.SetText(email)
		phoneLabel.SetText(phone)
	})

	btn2 := widget.NewButton("Quit", func() {
		a.Quit()
	})

	
	quitBtnBox := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), btn2)

	grid := container.New(layout.NewGridLayoutWithColumns(2),
		label1, input1, nameLabel,
		label2, input2, emailLabel,
		label3, input3, phoneLabel,
		layout.NewSpacer(),
		btn1,
		layout.NewSpacer(),
		daemonCheckBox,
		daemonStatusLabel,
		
		
	)
	for i := 0; i < 6; i++ {
		fmt.Println(walletAddressLabel.Text, walletBalance.Text, daemonCheckBox.Text)
		time.Sleep(5 * time.Second)
		}
		fmt.Println("Now CLosing Gracefully ;)")
		a.Quit() //uncomment to shut app after 6 prints

	// Set the walletBox and the existing form elements as the content of the window
	w.SetContent(container.NewVBox(walletBox, btnOpenTxtFile, grid, quitBtnBox))

	w.Resize(fyne.NewSize(600, 600))
	w.ShowAndRun()
}



package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/SixofClubsss/dReams/rpc"
)


type GetInfoResponse struct {
	Status string `json:"status"`
	Height int    `json:"height"`
}

func daemonStatus() {
	if daemonConnect {
		daemonStatusLabel.SetText("Daemon Connected")
	} else {
		daemonStatusLabel.SetText("Daemon Disconnected")
	}
}


func isDaemonConnected() {
    if daemonConnect {
        daemonCheckBox.SetChecked(true)
    } else {
        daemonCheckBox.SetChecked(false)
    }
}

func daemonOption() fyne.Widget {
	// Create a dropdown for selecting the network
	dropDown := widget.NewSelect([]string{"Mainnet", "Testnet", "Simulator"}, func(selected string) {
		switch selected {
		case "Mainnet":
			daemonAddress = DAEMON_MAINNET_DEFAULT
		case "Testnet":
			daemonAddress = DAEMON_TESTNET_DEFAULT
		case "Simulator":
			daemonAddress = DAEMON_SIMULATOR_DEFAULT
		}
	})
	dropDown.SetSelected("Mainnet")

	return dropDown
}



func PingDaemon() error {
	rpcClientD, ctx, cancel := rpc.SetDaemonClient(daemonAddress)
	defer cancel()

	var result GetInfoResponse
	err := rpcClientD.CallFor(ctx, &result, "get_info")

	if err != nil {
		log.Println("[Ping]", err)
		return err
	}

	// Check if the daemon is connected based on the 'status' field or other fields in the response
	// Adjust the condition according to the actual response structure
	if result.Status == "OK" {
		daemonConnect = true
	} else {
		daemonConnect = false
	}

	return nil
}

func init() { /// Handle ctrl-c close
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		kill_process = true
		stopLoop()
		fmt.Println()
		os.Exit(1)
	}()
}


func fetchLoop() { /// ping daemon and get height loop
    var ticker = time.NewTicker(6 * time.Second)
    quit = make(chan struct{})
    go func() {
        for {
            select {
			case <-ticker.C:
                err := PingDaemon()
                if err == nil {
                    daemonConnect = true
                } else {
                    daemonConnect = false
                }
                isDaemonConnected()
                isWalletConnected()
				daemonStatus()
                GetHeight()
            
            case <-quit: /// exit loop
                log.Println("[Wallet]] Exiting...")
                ticker.Stop()
                return
            }
        }
    }()
}


func stopLoop() { /// to exit loop
	quit <- struct{}{}
}

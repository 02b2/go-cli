package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"

	"github.com/SixofClubsss/dReams/rpc"
	dero "github.com/deroproject/derohe/rpc"
)

const (
	DAEMON_MAINNET_DEFAULT   = "127.0.0.1:10102"
	DAEMON_TESTNET_DEFAULT   = "127.0.0.1:40403"
	DAEMON_SIMULATOR_DEFAULT = "127.0.0.1:20000"
)

var (
	daemonAddress string
	daemonConnect bool
	currentHeight string
)

func Ping() error { /// ping blockchain for connection
	rpcClientD, ctx, cancel := rpc.SetDaemonClient(daemonAddress)
	defer cancel()

	var result string
	err := rpcClientD.CallFor(ctx, &result, "DERO.Ping")
	if err != nil {
		daemonConnect = false
		return nil
	}

	if result == "Pong " {
		daemonConnect = true
	} else {
		daemonConnect = false
	}

	return err
}

func GetHeight() error { /// get current height and displays
	rpcClientD, ctx, cancel := rpc.SetDaemonClient(daemonAddress)
	defer cancel()

	var result *dero.Daemon_GetHeight_Result
	err := rpcClientD.CallFor(ctx, &result, "DERO.GetHeight")

	if err != nil {
		return nil
	}
	

	return err
}


func getSCcode(scid string) error { /// get sc code and print in terminal
	rpcClientD, ctx, cancel := rpc.SetDaemonClient(daemonAddress)
	defer cancel()

	var result *dero.GetSC_Result
	params := dero.GetSC_Params{
		SCID:      scid,
		Code:      true,
		Variables: false,
	}
	err := rpcClientD.CallFor(ctx, &result, "DERO.GetSC", params)

	if err != nil {
		log.Println("[getSCcode]", err)
		return nil
	}

	fmt.Println(result.Code)

	return err
}

func findKey(i interface{}) (text string) {
	switch v := i.(type) {
	case uint64:
		text = strconv.Itoa(int(v))
	case string:
		text = v
	case float64:
		text = strconv.Itoa(int(v))
	default:

	}

	return
}

func SortStringMap(m map[string]interface{}) (str string) {
	keys := make([]string, 0, len(m))

	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		if m[k] == m["C"] {
			/// skipping C
		} else {
			str = str + k + " " + findKey(m[k]) + " \n"
		}
	}

	return
}

func SortUintMap(m map[uint64]interface{}) (str string) {
	keys := make([]uint64, 0, len(m))

	for k := range m {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

	for _, k := range keys {
		str = str + strconv.Itoa(int(k)) + " " + findKey(m[k]) + " \n"

	}

	return
}

package main

import (
	"fmt"
	"log"

	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/ethereum/go-ethereum/common"
)

func main() {
	//parse config
	configs, err := conf.ParseConfigFile("config.toml")
	if err != nil {
		log.Fatal(err)
	}
	config := &configs[0]

	//connect peer
	cli, err := client.Dial(config)
	if err != nil {
		log.Fatal(err)
	}

	//create a contract instance
	contractAddr := "0x63f6a200106d5c1f04d1f0889ca60cf8af40f27c"
	instance, err := NewHelloWorld(common.HexToAddress(contractAddr), cli)
	if err != nil {
		log.Fatal(err)
	}

	//call function of smart contract
	name, err := instance.GetName(cli.GetCallOpts())
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(name)
}


package main

import "core"

func main() {
	bc := core.NewBlockchain()
	bc.SendData("Send 1 BTC to Jim")
	bc.SendData("Send 2 ETH to Jack")

	bc.Print()

}

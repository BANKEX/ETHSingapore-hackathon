package eventHandlers

import "encoding/hex"

type Event struct {
	Signature string
	Handler   func(data []byte)
}

var Data map[string][]string

var Handlers = []Event{
	Event{"CoinDeposited(address,uint256)", HandleDeposit},
}

func HandleDeposit(data []byte) {
	Data[generateGuid()] = append(Data[generateGuid()], hex.EncodeToString(data))
}

func generateGuid() string {
	return ""
}
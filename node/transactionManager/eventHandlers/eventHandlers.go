package eventHandlers

type Event struct {
	Signature string
	Handler   func(data []byte)
}

var Handlers = []Event{
	Event{"CoinDeposited(address,uint256)", HandleDeposit},
}

func HandleDeposit(data []byte) {

}

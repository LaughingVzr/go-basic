package interfacep

import "fmt"

type USB interface {
	Name() string
	Connector
}

type Hub interface{
	 func()
}

type Connector interface {
	Connect()
}

type PhoneConnector struct {
	name string
}

type WasherConnector struct {
	name string
}

func (pc PhoneConnector) Name() string {
	return pc.name
}

func (pc PhoneConnector) Connect() {
	fmt.Println("Connect:", pc.name)
}

func (wash WasherConnector) Connect() {
	fmt.Println("Connect:", wash.name)
}

func Disconnect(usb interface{}) {
	switch v := usb.(type) {
	case PhoneConnector:
		fmt.Println("Disconnect.", v.name)
	default:
		fmt.Println("Unknown Device.")
	}
}

func Invoke() {
	pc := PhoneConnector{"PhoneConnector"}
	var conn Connector
	conn = Connector(pc)
	conn.Connect()

	pc.name = "pc"
	conn.Connect()

	fmt.Println(pc)
}

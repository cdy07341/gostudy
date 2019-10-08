// interface
package main

import (
	"fmt"
)

func main() {
	a := PhoneConnector{"PhoneConnector"}
	a.Connect()
	disConnect(a)
	//--------------------------

}

type USB interface {
	Name() string
	// Connect()
	Connecter
}

type Connecter interface {
	Connect()
}

type PhoneConnector struct {
	name string
}

func (pc PhoneConnector) Name() string {
	return pc.name
}
func (pc PhoneConnector) Connect() {
	fmt.Println("Connected:", pc.name)
}
func disConnect(usb USB) {
	// if pc, ok := usb.(PhoneConnector); ok {
	// 	fmt.Println("disConnected:", pc.name)
	// 	return
	// }
	// fmt.Println("unkown device")

	//传入的参数为(usb intreface{})
	switch v := usb.(type) {
	case PhoneConnector:
		fmt.Println("Disconnected:", v.name)
	default:
		fmt.Println("unkown device")
	}
}

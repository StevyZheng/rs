package net

type Interface struct {
	Name    string
	Mac     string
	Ip      string
	Netmask string
	Gateway string
}

type Net struct {
	Interfaces []Interface
}

func (n Net) GetNetInfo() {

}

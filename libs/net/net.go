package net

import (
	"github.com/emirpasic/gods/lists/arraylist"
	"rs/libs/common"
)

type Interface struct {
	Name    string
	Mac     string
	Ip      string
	Netmask string
	Gateway string
}

type Net struct {
	Interfaces arraylist.List
}

func (n *Net) InitNet() {
	baseNicPath := "/sys/class/net/"
	files, err := common.ListFiles(baseNicPath)
	if err != nil {
		return
	}
	it := files.Iterator()
	for it.Next() {
		nameTmp := it.Value().(string)
		if nameTmp != "lo" {
			interfaceT := Interface{
				Name: nameTmp,
			}
			n.Interfaces.Add(interfaceT)
		}
	}
}

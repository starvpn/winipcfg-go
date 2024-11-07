package winipcfg

import (
	"fmt"
	"testing"
)

func TestEnablingInterface(t *testing.T) {
	//fmt.Println(CheckInterfaces("StarVPN"))
	fmt.Println(FindInterfaceStatus("StarVPN"))
}

// 查看网卡状态
func TestCheckInterfaces(t *testing.T) {
	// netsh interface show interface
}

// RenameInterface
func TestRenameInterface(t *testing.T) {
	// netsh interface set interface name="以太网" newname="Ethernet"
	err := RenameInterface("本地", "111")
	if err != nil {
		t.Error(err)
	}
}

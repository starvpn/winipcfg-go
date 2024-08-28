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

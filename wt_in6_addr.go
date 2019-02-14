/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"fmt"
	"net"
)

// https://docs.microsoft.com/en-us/windows/desktop/api/in6addr/ns-in6addr-in6_addr
// IN6_ADDR defined in in6addr.h
type wtIn6Addr struct {
	Byte [16]uint8 // Windows type [16]UCHAR
}

func (addr *wtIn6Addr) toNetIp() net.IP {

	if addr == nil {
		return nil
	}

	return net.IP{
		byte(addr.Byte[0]),
		byte(addr.Byte[1]),
		byte(addr.Byte[2]),
		byte(addr.Byte[3]),
		byte(addr.Byte[4]),
		byte(addr.Byte[5]),
		byte(addr.Byte[6]),
		byte(addr.Byte[7]),
		byte(addr.Byte[8]),
		byte(addr.Byte[9]),
		byte(addr.Byte[10]),
		byte(addr.Byte[11]),
		byte(addr.Byte[12]),
		byte(addr.Byte[13]),
		byte(addr.Byte[14]),
		byte(addr.Byte[15]),
	}
}

func netIpToWtIn6Addr(ip net.IP) (*wtIn6Addr, error) {

	ip6 := ip.To16()

	if ip6 == nil {
		return nil, fmt.Errorf("Input IP isn't a valid IPv6 address.")
	}

	in6_addr := wtIn6Addr{}

	for i := 0; i < 16; i++ {
		in6_addr.Byte[i] = ip6[i]
	}

	return &in6_addr, nil
}

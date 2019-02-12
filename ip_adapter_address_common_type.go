/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import "fmt"

type IpAdapterAddressCommonType struct {

	// The interface the address belongs to.
	Interface Interface

	// TODO: Documentation missing. What is it?
	Length uint32

	// The address.
	Address SockaddrInet
}

func ipAdapterAddressFromWtDnsServerAddress(ifc Interface, wta *wtIpAdapterDnsServerAddressXp) (*IpAdapterAddressCommonType,
	error) {
	if wta == nil {
		return nil, nil
	} else {
		return ipAdapterAddressFromLengthAddress(ifc, wta.Length, &wta.Address)
	}
}

func ipAdapterAddressFromWtWinsServerAddress(ifc Interface, wta *wtIpAdapterWinsServerAddressLh) (*IpAdapterAddressCommonType,
	error) {
	if wta == nil {
		return nil, nil
	} else {
		return ipAdapterAddressFromLengthAddress(ifc, wta.Length, &wta.Address)
	}
}

func ipAdapterAddressFromWtGatewayAddress(ifc Interface, wta *wtIpAdapterGatewayAddressLh) (*IpAdapterAddressCommonType,
	error) {
	if wta == nil {
		return nil, nil
	} else {
		return ipAdapterAddressFromLengthAddress(ifc, wta.Length, &wta.Address)
	}
}

func ipAdapterAddressFromLengthAddress(ifc Interface, length uint32, wtsa *wtSocketAddress) (*IpAdapterAddressCommonType,
	error) {

	sainet, err := sockaddrInetFromWtSocketAddress(wtsa)

	if err != nil {
		return nil, err
	}

	act := IpAdapterAddressCommonType{Interface: ifc, Length: length, Address: *sainet}

	return &act, nil
}

func (a *IpAdapterAddressCommonType) commonTypeAddressString() string {
	if a == nil {
		return ""
	} else {
		return fmt.Sprintf("Length: %d; Address: [%s]", a.Length, a.Address.String())
	}
}

func (a *IpAdapterAddressCommonType) String() string {
	return a.commonTypeAddressString()
}
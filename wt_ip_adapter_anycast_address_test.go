/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"testing"
	"unsafe"
)

func Test_IP_ADAPTER_ANYCAST_ADDRESS_XP_Size(t *testing.T) {

	const Actual_IP_ADAPTER_ANYCAST_ADDRESS_XP_Size = unsafe.Sizeof(IP_ADAPTER_ANYCAST_ADDRESS_XP{})

	if Actual_IP_ADAPTER_ANYCAST_ADDRESS_XP_Size != wtIpAdapterAnycastAddressXp_Size {
		t.Errorf("Size of IP_ADAPTER_ANYCAST_ADDRESS_XP is %d, although %d is expected.", Actual_IP_ADAPTER_ANYCAST_ADDRESS_XP_Size, wtIpAdapterAnycastAddressXp_Size)
	}
}

func Test_IP_ADAPTER_ANYCAST_ADDRESS_XP_Offsets(t *testing.T) {

	s := IP_ADAPTER_ANYCAST_ADDRESS_XP{}
	sp := uintptr(unsafe.Pointer(&s))

	offset := uintptr(unsafe.Pointer(&s.Flags)) - sp

	if offset != wtIpAdapterAnycastAddressXp_Flags_Offset {
		t.Errorf("IP_ADAPTER_ANYCAST_ADDRESS_XP.Flags offset is %d although %d is expected", offset, wtIpAdapterAnycastAddressXp_Flags_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.Next)) - sp

	if offset != wtIpAdapterAnycastAddressXp_Next_Offset {
		t.Errorf("IP_ADAPTER_ANYCAST_ADDRESS_XP.Next offset is %d although %d is expected", offset, wtIpAdapterAnycastAddressXp_Next_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.Address)) - sp

	if offset != wtIpAdapterAnycastAddressXp_Address_Offset {
		t.Errorf("IP_ADAPTER_ANYCAST_ADDRESS_XP.Address offset is %d although %d is expected", offset, wtIpAdapterAnycastAddressXp_Address_Offset)
		return
	}
}

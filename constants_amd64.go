/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

const (
	wtIpAdapterAddressesLh_Size = 448

	wtIpAdapterAddressesLh_IfIndex_Offset                = 4
	wtIpAdapterAddressesLh_Next_Offset                   = 8
	wtIpAdapterAddressesLh_AdapterName_Offset           = 16
	wtIpAdapterAddressesLh_FirstUnicastAddress_Offset     = 24
	wtIpAdapterAddressesLh_FirstAnycastAddress_Offset     = 32
	wtIpAdapterAddressesLh_FirstMulticastAddress_Offset   = 40
	wtIpAdapterAddressesLh_FirstDnsServerAddress_Offset   = 48
	wtIpAdapterAddressesLh_DnsSuffix_Offset               = 56
	wtIpAdapterAddressesLh_Description_Offset             = 64
	wtIpAdapterAddressesLh_FriendlyName_Offset            = 72
	wtIpAdapterAddressesLh_PhysicalAddress_Offset         = 80
	wtIpAdapterAddressesLh_PhysicalAddressLength_Offset  = 88
	wtIpAdapterAddressesLh_Flags_Offset                  = 92
	wtIpAdapterAddressesLh_Mtu_Offset                    = 96
	wtIpAdapterAddressesLh_IfType_Offset                 = 100
	wtIpAdapterAddressesLh_OperStatus_Offset              = 104
	wtIpAdapterAddressesLh_Ipv6IfIndex_Offset             = 108
	wtIpAdapterAddressesLh_ZoneIndices_Offset             = 112
	wtIpAdapterAddressesLh_FirstPrefix_Offset             = 176
	wtIpAdapterAddressesLh_TransmitLinkSpeed_Offset       = 184
	wtIpAdapterAddressesLh_ReceiveLinkSpeed_Offset        = 192
	wtIpAdapterAddressesLh_FirstWinsServerAddress_Offset  = 200
	wtIpAdapterAddressesLh_FirstGatewayAddress_Offset     = 208
	wtIpAdapterAddressesLh_Ipv4Metric_Offset              = 216
	wtIpAdapterAddressesLh_Ipv6Metric_Offset             = 220
	wtIpAdapterAddressesLh_Luid_Offset                   = 224
	wtIpAdapterAddressesLh_Dhcpv4Server_Offset           = 232
	wtIpAdapterAddressesLh_CompartmentId_Offset          = 248
	wtIpAdapterAddressesLh_NetworkGuid_Offset            = 252
	wtIpAdapterAddressesLh_ConnectionType_Offset         = 268
	wtIpAdapterAddressesLh_TunnelType_Offset             = 272
	wtIpAdapterAddressesLh_Dhcpv6Server_Offset           = 280
	wtIpAdapterAddressesLh_Dhcpv6ClientDuid_Offset       = 296
	wtIpAdapterAddressesLh_Dhcpv6ClientDuidLength_Offset = 428
	wtIpAdapterAddressesLh_Dhcpv6Iaid_Offset             = 432
	wtIpAdapterAddressesLh_FirstDnsSuffix_Offset         = 440

	wtIpAdapterAnycastAddressXp_Size = 32

	wtIpAdapterAnycastAddressXp_Flags_Offset   = 4
	wtIpAdapterAnycastAddressXp_Next_Offset    = 8
	wtIpAdapterAnycastAddressXp_Address_Offset = 16

	wtSocketAddress_Size = 16

	wtSocketAddress_iSockaddrLength_Offset = 8

	wtIpAdapterDnsServerAddressXp_Size = 32

	wtIpAdapterDnsServerAddressXp_Reserved_Offset = 4
	wtIpAdapterDnsServerAddressXp_Next_Offset = 8
	wtIpAdapterDnsServerAddressXp_Address_Offset = 16

	wtIpAdapterDnsSuffix_Size = 520

	wtIpAdapterDnsSuffix_String_Offset = 8

	IP_ADAPTER_GATEWAY_ADDRESS_LH_Size = 32

	IP_ADAPTER_GATEWAY_ADDRESS_LH_Reserved_Offset = 4
	IP_ADAPTER_GATEWAY_ADDRESS_LH_Next_Offset = 8
	IP_ADAPTER_GATEWAY_ADDRESS_LH_Address_Offset = 16

	IP_ADAPTER_MULTICAST_ADDRESS_XP_Size = 32

	IP_ADAPTER_MULTICAST_ADDRESS_XP_Flags_Offset = 4
	IP_ADAPTER_MULTICAST_ADDRESS_XP_Next_Offset = 8
	IP_ADAPTER_MULTICAST_ADDRESS_XP_Address_Offset = 16

	IP_ADAPTER_PREFIX_XP_Size = 40

	IP_ADAPTER_PREFIX_XP_Flags_Offset = 4
	IP_ADAPTER_PREFIX_XP_Next_Offset = 8
	IP_ADAPTER_PREFIX_XP_Address_Offset = 16
	IP_ADAPTER_PREFIX_XP_PrefixLength_Offset = 32

	IP_ADAPTER_UNICAST_ADDRESS_LH_Size = 64

	IP_ADAPTER_UNICAST_ADDRESS_LH_Flags_Offset = 4
	IP_ADAPTER_UNICAST_ADDRESS_LH_Next_Offset = 8
	IP_ADAPTER_UNICAST_ADDRESS_LH_Address_Offset = 16
	IP_ADAPTER_UNICAST_ADDRESS_LH_PrefixOrigin_Offset = 32
	IP_ADAPTER_UNICAST_ADDRESS_LH_SuffixOrigin_Offset = 36
	IP_ADAPTER_UNICAST_ADDRESS_LH_DadState_Offset = 40
	IP_ADAPTER_UNICAST_ADDRESS_LH_ValidLifetime_Offset = 44
	IP_ADAPTER_UNICAST_ADDRESS_LH_PreferredLifetime_Offset = 48
	IP_ADAPTER_UNICAST_ADDRESS_LH_LeaseLifetime_Offset = 52
	IP_ADAPTER_UNICAST_ADDRESS_LH_OnLinkPrefixLength_Offset = 56

	IP_ADAPTER_WINS_SERVER_ADDRESS_LH_Size = 32

	IP_ADAPTER_WINS_SERVER_ADDRESS_LH_Reserved_Offset = 4
	IP_ADAPTER_WINS_SERVER_ADDRESS_LH_Next_Offset = 8
	IP_ADAPTER_WINS_SERVER_ADDRESS_LH_Address_Offset = 16
)

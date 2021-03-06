// Code generated by "stringer -type=API_TYPE"; DO NOT EDIT.

package zebra

import "fmt"

const _API_TYPE_name = "FRR_INTERFACE_ADDINTERFACE_ADDINTERFACE_DELETEINTERFACE_ADDRESS_ADDINTERFACE_ADDRESS_DELETEINTERFACE_UPINTERFACE_DOWNIPV4_ROUTE_ADDIPV4_ROUTE_DELETEIPV6_ROUTE_ADDIPV6_ROUTE_DELETEREDISTRIBUTE_ADDREDISTRIBUTE_DELETEREDISTRIBUTE_DEFAULT_ADDREDISTRIBUTE_DEFAULT_DELETEIPV4_NEXTHOP_LOOKUPIPV6_NEXTHOP_LOOKUPIPV4_IMPORT_LOOKUPIPV6_IMPORT_LOOKUPINTERFACE_RENAMEROUTER_ID_ADDROUTER_ID_DELETEROUTER_ID_UPDATEHELLOIPV4_NEXTHOP_LOOKUP_MRIBVRF_UNREGISTERINTERFACE_LINK_PARAMSNEXTHOP_REGISTERNEXTHOP_UNREGISTERNEXTHOP_UPDATEMESSAGE_MAXFRR_BFD_DEST_REPLAYFRR_REDISTRIBUTE_IPV4_ADDFRR_REDISTRIBUTE_IPV4_DELFRR_REDISTRIBUTE_IPV6_ADDFRR_REDISTRIBUTE_IPV6_DELFRR_VRF_UNREGISTERFRR_VRF_ADDFRR_VRF_DELETEFRR_INTERFACE_VRF_UPDATEFRR_BFD_CLIENT_REGISTERFRR_INTERFACE_ENABLE_RADVFRR_INTERFACE_DISABLE_RADVFRR_IPV4_NEXTHOP_LOOKUP_MRIBFRR_INTERFACE_LINK_PARAMSFRR_MPLS_LABELS_ADDFRR_MPLS_LABELS_DELETEFRR_IPV4_NEXTHOP_ADDFRR_IPV4_NEXTHOP_DELETEFRR_IPV6_NEXTHOP_ADDFRR_IPV6_NEXTHOP_DELETEFRR_IPMR_ROUTE_STATSFRR_LABEL_MANAGER_CONNECTFRR_GET_LABEL_CHUNKFRR_RELEASE_LABEL_CHUNKFRR_PW_ADDFRR_PW_DELETEFRR_PW_SETFRR_PW_UNSETFRR_PW_STATUS_UPDATE"

var _API_TYPE_index = [...]uint16{0, 17, 30, 46, 67, 91, 103, 117, 131, 148, 162, 179, 195, 214, 238, 265, 284, 303, 321, 339, 355, 368, 384, 400, 405, 429, 443, 464, 480, 498, 512, 523, 542, 567, 592, 617, 642, 660, 671, 685, 709, 732, 757, 783, 811, 836, 855, 877, 897, 920, 940, 963, 983, 1008, 1027, 1050, 1060, 1073, 1083, 1095, 1115}

func (i API_TYPE) String() string {
	if i >= API_TYPE(len(_API_TYPE_index)-1) {
		return fmt.Sprintf("API_TYPE(%d)", i)
	}
	return _API_TYPE_name[_API_TYPE_index[i]:_API_TYPE_index[i+1]]
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v2/listener/listener.proto

/*
Package listener is a generated protocol buffer package.

It is generated from these files:
	envoy/api/v2/listener/listener.proto

It has these top-level messages:
	Filter
	FilterChainMatch
	FilterChain
	ListenerFilter
*/
package listener

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import envoy_api_v2_core "github.com/cilium/cilium/pkg/envoy/envoy/api/v2/core"
import envoy_api_v2_auth "github.com/cilium/cilium/pkg/envoy/envoy/api/v2/auth"
import envoy_api_v2_core1 "github.com/cilium/cilium/pkg/envoy/envoy/api/v2/core"
import google_protobuf4 "github.com/golang/protobuf/ptypes/struct"
import google_protobuf "github.com/golang/protobuf/ptypes/wrappers"
import _ "github.com/lyft/protoc-gen-validate/validate"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Filter struct {
	// The name of the filter to instantiate. The name must match a supported
	// filter. The built-in filters are:
	//
	// [#comment:TODO(mattklein123): Auto generate the following list]
	// * :ref:`envoy.client_ssl_auth<config_network_filters_client_ssl_auth>`
	// * :ref:`envoy.echo <config_network_filters_echo>`
	// * :ref:`envoy.http_connection_manager <config_http_conn_man>`
	// * :ref:`envoy.mongo_proxy <config_network_filters_mongo_proxy>`
	// * :ref:`envoy.ratelimit <config_network_filters_rate_limit>`
	// * :ref:`envoy.redis_proxy <config_network_filters_redis_proxy>`
	// * :ref:`envoy.tcp_proxy <config_network_filters_tcp_proxy>`
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Filter specific configuration which depends on the filter being
	// instantiated. See the supported filters for further documentation.
	Config *google_protobuf4.Struct `protobuf:"bytes,2,opt,name=config" json:"config,omitempty"`
	// [#not-implemented-hide:]
	DeprecatedV1 *Filter_DeprecatedV1 `protobuf:"bytes,3,opt,name=deprecated_v1,json=deprecatedV1" json:"deprecated_v1,omitempty"`
}

func (m *Filter) Reset()                    { *m = Filter{} }
func (m *Filter) String() string            { return proto.CompactTextString(m) }
func (*Filter) ProtoMessage()               {}
func (*Filter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Filter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Filter) GetConfig() *google_protobuf4.Struct {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *Filter) GetDeprecatedV1() *Filter_DeprecatedV1 {
	if m != nil {
		return m.DeprecatedV1
	}
	return nil
}

// [#not-implemented-hide:]
type Filter_DeprecatedV1 struct {
	Type string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
}

func (m *Filter_DeprecatedV1) Reset()                    { *m = Filter_DeprecatedV1{} }
func (m *Filter_DeprecatedV1) String() string            { return proto.CompactTextString(m) }
func (*Filter_DeprecatedV1) ProtoMessage()               {}
func (*Filter_DeprecatedV1) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *Filter_DeprecatedV1) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

// Specifies the match criteria for selecting a specific filter chain for a
// listener.
type FilterChainMatch struct {
	// If non-empty, the SNI domains to consider. May contain a wildcard prefix,
	// e.g. ``*.example.com``.
	//
	// .. attention::
	//
	//   See the :ref:`FAQ entry <faq_how_to_setup_sni>` on how to configure SNI for more
	//   information.
	SniDomains []string `protobuf:"bytes,1,rep,name=sni_domains,json=sniDomains" json:"sni_domains,omitempty"`
	// If non-empty, an IP address and prefix length to match addresses when the
	// listener is bound to 0.0.0.0/:: or when use_original_dst is specified.
	// [#not-implemented-hide:]
	PrefixRanges []*envoy_api_v2_core.CidrRange `protobuf:"bytes,3,rep,name=prefix_ranges,json=prefixRanges" json:"prefix_ranges,omitempty"`
	// If non-empty, an IP address and suffix length to match addresses when the
	// listener is bound to 0.0.0.0/:: or when use_original_dst is specified.
	// [#not-implemented-hide:]
	AddressSuffix string `protobuf:"bytes,4,opt,name=address_suffix,json=addressSuffix" json:"address_suffix,omitempty"`
	// [#not-implemented-hide:]
	SuffixLen *google_protobuf.UInt32Value `protobuf:"bytes,5,opt,name=suffix_len,json=suffixLen" json:"suffix_len,omitempty"`
	// The criteria is satisfied if the source IP address of the downstream
	// connection is contained in at least one of the specified subnets. If the
	// parameter is not specified or the list is empty, the source IP address is
	// ignored.
	// [#not-implemented-hide:]
	SourcePrefixRanges []*envoy_api_v2_core.CidrRange `protobuf:"bytes,6,rep,name=source_prefix_ranges,json=sourcePrefixRanges" json:"source_prefix_ranges,omitempty"`
	// The criteria is satisfied if the source port of the downstream connection
	// is contained in at least one of the specified ports. If the parameter is
	// not specified, the source port is ignored.
	// [#not-implemented-hide:]
	SourcePorts []*google_protobuf.UInt32Value `protobuf:"bytes,7,rep,name=source_ports,json=sourcePorts" json:"source_ports,omitempty"`
	// Optional destination port to consider when use_original_dst is set on the
	// listener in determining a filter chain match.
	// [#not-implemented-hide:]
	DestinationPort *google_protobuf.UInt32Value `protobuf:"bytes,8,opt,name=destination_port,json=destinationPort" json:"destination_port,omitempty"`
}

func (m *FilterChainMatch) Reset()                    { *m = FilterChainMatch{} }
func (m *FilterChainMatch) String() string            { return proto.CompactTextString(m) }
func (*FilterChainMatch) ProtoMessage()               {}
func (*FilterChainMatch) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *FilterChainMatch) GetSniDomains() []string {
	if m != nil {
		return m.SniDomains
	}
	return nil
}

func (m *FilterChainMatch) GetPrefixRanges() []*envoy_api_v2_core.CidrRange {
	if m != nil {
		return m.PrefixRanges
	}
	return nil
}

func (m *FilterChainMatch) GetAddressSuffix() string {
	if m != nil {
		return m.AddressSuffix
	}
	return ""
}

func (m *FilterChainMatch) GetSuffixLen() *google_protobuf.UInt32Value {
	if m != nil {
		return m.SuffixLen
	}
	return nil
}

func (m *FilterChainMatch) GetSourcePrefixRanges() []*envoy_api_v2_core.CidrRange {
	if m != nil {
		return m.SourcePrefixRanges
	}
	return nil
}

func (m *FilterChainMatch) GetSourcePorts() []*google_protobuf.UInt32Value {
	if m != nil {
		return m.SourcePorts
	}
	return nil
}

func (m *FilterChainMatch) GetDestinationPort() *google_protobuf.UInt32Value {
	if m != nil {
		return m.DestinationPort
	}
	return nil
}

// A filter chain wraps a set of match criteria, an option TLS context, a set of filters, and
// various other parameters.
type FilterChain struct {
	// The criteria to use when matching a connection to this filter chain.
	FilterChainMatch *FilterChainMatch `protobuf:"bytes,1,opt,name=filter_chain_match,json=filterChainMatch" json:"filter_chain_match,omitempty"`
	// The TLS context for this filter chain.
	TlsContext *envoy_api_v2_auth.DownstreamTlsContext `protobuf:"bytes,2,opt,name=tls_context,json=tlsContext" json:"tls_context,omitempty"`
	// A list of individual network filters that make up the filter chain for
	// connections established with the listener. Order matters as the filters are
	// processed sequentially as connection events happen. Note: If the filter
	// list is empty, the connection will close by default.
	Filters []*Filter `protobuf:"bytes,3,rep,name=filters" json:"filters,omitempty"`
	// Whether the listener should expect a PROXY protocol V1 header on new
	// connections. If this option is enabled, the listener will assume that that
	// remote address of the connection is the one specified in the header. Some
	// load balancers including the AWS ELB support this option. If the option is
	// absent or set to false, Envoy will use the physical peer address of the
	// connection as the remote address.
	UseProxyProto *google_protobuf.BoolValue `protobuf:"bytes,4,opt,name=use_proxy_proto,json=useProxyProto" json:"use_proxy_proto,omitempty"`
	// [#not-implemented-hide:] filter chain metadata.
	Metadata *envoy_api_v2_core1.Metadata `protobuf:"bytes,5,opt,name=metadata" json:"metadata,omitempty"`
	// See :ref:`base.TransportSocket<envoy_api_msg_core.TransportSocket>` description.
	TransportSocket *envoy_api_v2_core1.TransportSocket `protobuf:"bytes,6,opt,name=transport_socket,json=transportSocket" json:"transport_socket,omitempty"`
}

func (m *FilterChain) Reset()                    { *m = FilterChain{} }
func (m *FilterChain) String() string            { return proto.CompactTextString(m) }
func (*FilterChain) ProtoMessage()               {}
func (*FilterChain) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *FilterChain) GetFilterChainMatch() *FilterChainMatch {
	if m != nil {
		return m.FilterChainMatch
	}
	return nil
}

func (m *FilterChain) GetTlsContext() *envoy_api_v2_auth.DownstreamTlsContext {
	if m != nil {
		return m.TlsContext
	}
	return nil
}

func (m *FilterChain) GetFilters() []*Filter {
	if m != nil {
		return m.Filters
	}
	return nil
}

func (m *FilterChain) GetUseProxyProto() *google_protobuf.BoolValue {
	if m != nil {
		return m.UseProxyProto
	}
	return nil
}

func (m *FilterChain) GetMetadata() *envoy_api_v2_core1.Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *FilterChain) GetTransportSocket() *envoy_api_v2_core1.TransportSocket {
	if m != nil {
		return m.TransportSocket
	}
	return nil
}

type ListenerFilter struct {
	// The name of the filter to instantiate. The name must match a supported
	// filter. The built-in filters are:
	//
	// [#comment:TODO(mattklein123): Auto generate the following list]
	// * :ref:`envoy.listener.original_dst <config_listener_filters_original_dst>`
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Filter specific configuration which depends on the filter being
	// instantiated. See the supported filters for further documentation.
	Config *google_protobuf4.Struct `protobuf:"bytes,2,opt,name=config" json:"config,omitempty"`
}

func (m *ListenerFilter) Reset()                    { *m = ListenerFilter{} }
func (m *ListenerFilter) String() string            { return proto.CompactTextString(m) }
func (*ListenerFilter) ProtoMessage()               {}
func (*ListenerFilter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ListenerFilter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ListenerFilter) GetConfig() *google_protobuf4.Struct {
	if m != nil {
		return m.Config
	}
	return nil
}

func init() {
	proto.RegisterType((*Filter)(nil), "envoy.api.v2.listener.Filter")
	proto.RegisterType((*Filter_DeprecatedV1)(nil), "envoy.api.v2.listener.Filter.DeprecatedV1")
	proto.RegisterType((*FilterChainMatch)(nil), "envoy.api.v2.listener.FilterChainMatch")
	proto.RegisterType((*FilterChain)(nil), "envoy.api.v2.listener.FilterChain")
	proto.RegisterType((*ListenerFilter)(nil), "envoy.api.v2.listener.ListenerFilter")
}

func init() { proto.RegisterFile("envoy/api/v2/listener/listener.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 683 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0xc5, 0x4d, 0x48, 0x9b, 0x71, 0xd2, 0x46, 0xa3, 0xa2, 0x5a, 0xa1, 0x8f, 0x28, 0x02, 0x11,
	0xb1, 0xb0, 0x55, 0x77, 0xc1, 0x02, 0x21, 0x44, 0x5a, 0xf1, 0x90, 0x5a, 0x54, 0x39, 0x6d, 0x17,
	0x6c, 0xcc, 0xd4, 0xbe, 0x4e, 0x47, 0x38, 0x33, 0xd6, 0xcc, 0x38, 0x4d, 0xff, 0x88, 0x0f, 0x60,
	0x03, 0x2b, 0x96, 0x6c, 0xf9, 0x01, 0x16, 0xec, 0xf8, 0x0b, 0xe4, 0xb1, 0x1d, 0x92, 0x36, 0xaa,
	0xba, 0x61, 0x77, 0x1f, 0xe7, 0x9c, 0x1c, 0xdf, 0x7b, 0x33, 0xe8, 0x11, 0xb0, 0x31, 0xbf, 0x72,
	0x48, 0x42, 0x9d, 0xb1, 0xeb, 0xc4, 0x54, 0x2a, 0x60, 0x20, 0xa6, 0x81, 0x9d, 0x08, 0xae, 0x38,
	0x7e, 0xa0, 0x51, 0x36, 0x49, 0xa8, 0x3d, 0x76, 0xed, 0xb2, 0xd9, 0xde, 0x99, 0x23, 0x07, 0x5c,
	0x80, 0x43, 0xc2, 0x50, 0x80, 0x94, 0x39, 0xaf, 0xbd, 0x39, 0x07, 0x20, 0xa9, 0xba, 0x70, 0x02,
	0x10, 0x6a, 0x61, 0x57, 0xd3, 0xcf, 0x89, 0x84, 0xb2, 0x3b, 0xe4, 0x7c, 0x18, 0x83, 0xa3, 0xb3,
	0xf3, 0x34, 0x72, 0xa4, 0x12, 0x69, 0x50, 0x72, 0xb7, 0xaf, 0x77, 0x2f, 0x05, 0x49, 0x12, 0x10,
	0xe5, 0x2f, 0x6f, 0x8c, 0x49, 0x4c, 0x43, 0xa2, 0xc0, 0x29, 0x83, 0xa2, 0xb1, 0x3e, 0xe4, 0x43,
	0xae, 0x43, 0x27, 0x8b, 0xf2, 0x6a, 0xf7, 0xa7, 0x81, 0x6a, 0xaf, 0x69, 0xac, 0x40, 0xe0, 0x2d,
	0x54, 0x65, 0x64, 0x04, 0x96, 0xd1, 0x31, 0x7a, 0xf5, 0x7e, 0xfd, 0xdb, 0x9f, 0xef, 0x95, 0xaa,
	0x58, 0xea, 0x18, 0x9e, 0x2e, 0x63, 0x07, 0xd5, 0x02, 0xce, 0x22, 0x3a, 0xb4, 0x96, 0x3a, 0x46,
	0xcf, 0x74, 0x37, 0xec, 0xdc, 0x89, 0x5d, 0x3a, 0xb1, 0x07, 0xda, 0xa7, 0x57, 0xc0, 0xf0, 0x00,
	0x35, 0x43, 0x48, 0x04, 0x04, 0x44, 0x41, 0xe8, 0x8f, 0x77, 0xad, 0x8a, 0xe6, 0x3d, 0xb5, 0x17,
	0xce, 0xd4, 0xce, 0x5d, 0xd8, 0x07, 0x53, 0xca, 0xd9, 0x6e, 0x7f, 0xc9, 0x32, 0xbc, 0x46, 0x38,
	0x53, 0x69, 0x77, 0x51, 0x63, 0x16, 0x81, 0x31, 0xaa, 0xaa, 0xab, 0xa4, 0x30, 0xed, 0xe9, 0xb8,
	0xfb, 0xa5, 0x82, 0x5a, 0xb9, 0xda, 0xfe, 0x05, 0xa1, 0xec, 0x88, 0xa8, 0xe0, 0x02, 0xef, 0x20,
	0x53, 0x32, 0xea, 0x87, 0x7c, 0x44, 0x28, 0x93, 0x96, 0xd1, 0xa9, 0xf4, 0xea, 0x1e, 0x92, 0x8c,
	0x1e, 0xe4, 0x15, 0xfc, 0x0a, 0x35, 0x13, 0x01, 0x11, 0x9d, 0xf8, 0x82, 0xb0, 0x21, 0x48, 0xab,
	0xd2, 0xa9, 0xf4, 0x4c, 0x77, 0x73, 0xde, 0x6e, 0xb6, 0x2c, 0x7b, 0x9f, 0x86, 0xc2, 0xcb, 0x40,
	0x5e, 0x23, 0xa7, 0xe8, 0x44, 0xe2, 0xc7, 0x68, 0xb5, 0x38, 0x03, 0x5f, 0xa6, 0x51, 0x44, 0x27,
	0x56, 0x55, 0xdb, 0x6a, 0x16, 0xd5, 0x81, 0x2e, 0xe2, 0xe7, 0x08, 0xe5, 0x6d, 0x3f, 0x06, 0x66,
	0xdd, 0xd7, 0x53, 0xd9, 0xbc, 0x31, 0xcd, 0xd3, 0x77, 0x4c, 0xed, 0xb9, 0x67, 0x24, 0x4e, 0xc1,
	0xab, 0xe7, 0xf8, 0x43, 0x60, 0xf8, 0x3d, 0x5a, 0x97, 0x3c, 0x15, 0x01, 0xf8, 0xf3, 0x6e, 0x6b,
	0x77, 0x70, 0x8b, 0x73, 0xe6, 0xf1, 0xac, 0xe7, 0x97, 0xa8, 0x51, 0xea, 0x71, 0xa1, 0xa4, 0xb5,
	0x5c, 0xe8, 0xdc, 0x66, 0xc7, 0x2c, 0x74, 0x32, 0x02, 0x7e, 0x83, 0x5a, 0x21, 0x48, 0x45, 0x19,
	0x51, 0x94, 0x33, 0xad, 0x62, 0xad, 0xdc, 0xe1, 0x9b, 0xd6, 0x66, 0x58, 0x99, 0x52, 0xf7, 0x6b,
	0x05, 0x99, 0x33, 0x6b, 0xc3, 0xa7, 0x08, 0x47, 0x3a, 0xf5, 0x83, 0x2c, 0xf7, 0x47, 0xd9, 0x1e,
	0xf5, 0xa2, 0x4d, 0xf7, 0xc9, 0xad, 0x47, 0xf4, 0x6f, 0xed, 0x5e, 0x2b, 0xba, 0x7e, 0x08, 0x6f,
	0x91, 0xa9, 0x62, 0xe9, 0x07, 0x9c, 0x29, 0x98, 0xa8, 0xe2, 0x98, 0xaf, 0xe9, 0x65, 0x7f, 0x58,
	0xfb, 0x80, 0x5f, 0x32, 0xa9, 0x04, 0x90, 0xd1, 0x49, 0x2c, 0xf7, 0x73, 0xb8, 0x87, 0xd4, 0x34,
	0xc6, 0x2f, 0xd0, 0x72, 0xae, 0x5e, 0xde, 0xca, 0xd6, 0xad, 0xae, 0xfa, 0xd5, 0x1f, 0xbf, 0x76,
	0xee, 0x79, 0x25, 0x07, 0xf7, 0xd1, 0x5a, 0x2a, 0xb3, 0x35, 0xf2, 0xc9, 0x95, 0xaf, 0x47, 0xa4,
	0xcf, 0xc5, 0x74, 0xdb, 0x37, 0xe6, 0xd6, 0xe7, 0x3c, 0xce, 0xa7, 0xd6, 0x4c, 0x25, 0x1c, 0x67,
	0x8c, 0x63, 0xfd, 0x3e, 0x3d, 0x43, 0x2b, 0x23, 0x50, 0x24, 0x24, 0x8a, 0x14, 0x87, 0xf4, 0x70,
	0xc1, 0x05, 0x1c, 0x15, 0x10, 0x6f, 0x0a, 0xc6, 0x47, 0xa8, 0xa5, 0x04, 0x61, 0x32, 0x5b, 0x97,
	0x2f, 0x79, 0xf0, 0x09, 0x94, 0x55, 0xd3, 0x02, 0xdd, 0x05, 0x02, 0x27, 0x25, 0x74, 0xa0, 0x91,
	0xde, 0x9a, 0x9a, 0x2f, 0x74, 0x3f, 0xa2, 0xd5, 0xc3, 0xe2, 0x6b, 0xff, 0xcf, 0x6b, 0xd2, 0x5f,
	0xfd, 0xfc, 0x7b, 0xdb, 0xf8, 0xb0, 0x52, 0xce, 0xf4, 0xbc, 0xa6, 0x81, 0x7b, 0x7f, 0x03, 0x00,
	0x00, 0xff, 0xff, 0xee, 0x15, 0xf4, 0xc9, 0xc8, 0x05, 0x00, 0x00,
}
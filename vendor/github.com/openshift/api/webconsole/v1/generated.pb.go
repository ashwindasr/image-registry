// Code generated by protoc-gen-gogo.
// source: github.com/openshift/api/webconsole/v1/generated.proto
// DO NOT EDIT!

/*
	Package v1 is a generated protocol buffer package.

	It is generated from these files:
		github.com/openshift/api/webconsole/v1/generated.proto

	It has these top-level messages:
		ClusterInfo
		ExtensionsConfiguration
		FeaturesConfiguration
		WebConsoleConfiguration
*/
package v1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

func (m *ClusterInfo) Reset()                    { *m = ClusterInfo{} }
func (*ClusterInfo) ProtoMessage()               {}
func (*ClusterInfo) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{0} }

func (m *ExtensionsConfiguration) Reset()                    { *m = ExtensionsConfiguration{} }
func (*ExtensionsConfiguration) ProtoMessage()               {}
func (*ExtensionsConfiguration) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{1} }

func (m *FeaturesConfiguration) Reset()                    { *m = FeaturesConfiguration{} }
func (*FeaturesConfiguration) ProtoMessage()               {}
func (*FeaturesConfiguration) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{2} }

func (m *WebConsoleConfiguration) Reset()                    { *m = WebConsoleConfiguration{} }
func (*WebConsoleConfiguration) ProtoMessage()               {}
func (*WebConsoleConfiguration) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{3} }

func init() {
	proto.RegisterType((*ClusterInfo)(nil), "github.com.openshift.api.webconsole.v1.ClusterInfo")
	proto.RegisterType((*ExtensionsConfiguration)(nil), "github.com.openshift.api.webconsole.v1.ExtensionsConfiguration")
	proto.RegisterType((*FeaturesConfiguration)(nil), "github.com.openshift.api.webconsole.v1.FeaturesConfiguration")
	proto.RegisterType((*WebConsoleConfiguration)(nil), "github.com.openshift.api.webconsole.v1.WebConsoleConfiguration")
}
func (m *ClusterInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClusterInfo) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.ConsolePublicURL)))
	i += copy(dAtA[i:], m.ConsolePublicURL)
	dAtA[i] = 0x12
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.MasterPublicURL)))
	i += copy(dAtA[i:], m.MasterPublicURL)
	dAtA[i] = 0x1a
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.LoggingPublicURL)))
	i += copy(dAtA[i:], m.LoggingPublicURL)
	dAtA[i] = 0x22
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.MetricsPublicURL)))
	i += copy(dAtA[i:], m.MetricsPublicURL)
	dAtA[i] = 0x2a
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.LogoutPublicURL)))
	i += copy(dAtA[i:], m.LogoutPublicURL)
	return i, nil
}

func (m *ExtensionsConfiguration) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExtensionsConfiguration) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ScriptURLs) > 0 {
		for _, s := range m.ScriptURLs {
			dAtA[i] = 0xa
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.StylesheetURLs) > 0 {
		for _, s := range m.StylesheetURLs {
			dAtA[i] = 0x12
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.Properties) > 0 {
		keysForProperties := make([]string, 0, len(m.Properties))
		for k := range m.Properties {
			keysForProperties = append(keysForProperties, string(k))
		}
		github_com_gogo_protobuf_sortkeys.Strings(keysForProperties)
		for _, k := range keysForProperties {
			dAtA[i] = 0x1a
			i++
			v := m.Properties[string(k)]
			mapSize := 1 + len(k) + sovGenerated(uint64(len(k))) + 1 + len(v) + sovGenerated(uint64(len(v)))
			i = encodeVarintGenerated(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintGenerated(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintGenerated(dAtA, i, uint64(len(v)))
			i += copy(dAtA[i:], v)
		}
	}
	return i, nil
}

func (m *FeaturesConfiguration) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FeaturesConfiguration) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.InactivityTimeoutMinutes))
	dAtA[i] = 0x10
	i++
	if m.ClusterResourceOverridesEnabled {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i++
	return i, nil
}

func (m *WebConsoleConfiguration) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WebConsoleConfiguration) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.ServingInfo.Size()))
	n1, err := m.ServingInfo.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	dAtA[i] = 0x12
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.ClusterInfo.Size()))
	n2, err := m.ClusterInfo.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	dAtA[i] = 0x1a
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.Features.Size()))
	n3, err := m.Features.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	dAtA[i] = 0x22
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.Extensions.Size()))
	n4, err := m.Extensions.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	return i, nil
}

func encodeFixed64Generated(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Generated(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintGenerated(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ClusterInfo) Size() (n int) {
	var l int
	_ = l
	l = len(m.ConsolePublicURL)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.MasterPublicURL)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.LoggingPublicURL)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.MetricsPublicURL)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.LogoutPublicURL)
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *ExtensionsConfiguration) Size() (n int) {
	var l int
	_ = l
	if len(m.ScriptURLs) > 0 {
		for _, s := range m.ScriptURLs {
			l = len(s)
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	if len(m.StylesheetURLs) > 0 {
		for _, s := range m.StylesheetURLs {
			l = len(s)
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	if len(m.Properties) > 0 {
		for k, v := range m.Properties {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovGenerated(uint64(len(k))) + 1 + len(v) + sovGenerated(uint64(len(v)))
			n += mapEntrySize + 1 + sovGenerated(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *FeaturesConfiguration) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovGenerated(uint64(m.InactivityTimeoutMinutes))
	n += 2
	return n
}

func (m *WebConsoleConfiguration) Size() (n int) {
	var l int
	_ = l
	l = m.ServingInfo.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.ClusterInfo.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Features.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Extensions.Size()
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func sovGenerated(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *ClusterInfo) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ClusterInfo{`,
		`ConsolePublicURL:` + fmt.Sprintf("%v", this.ConsolePublicURL) + `,`,
		`MasterPublicURL:` + fmt.Sprintf("%v", this.MasterPublicURL) + `,`,
		`LoggingPublicURL:` + fmt.Sprintf("%v", this.LoggingPublicURL) + `,`,
		`MetricsPublicURL:` + fmt.Sprintf("%v", this.MetricsPublicURL) + `,`,
		`LogoutPublicURL:` + fmt.Sprintf("%v", this.LogoutPublicURL) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ExtensionsConfiguration) String() string {
	if this == nil {
		return "nil"
	}
	keysForProperties := make([]string, 0, len(this.Properties))
	for k := range this.Properties {
		keysForProperties = append(keysForProperties, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForProperties)
	mapStringForProperties := "map[string]string{"
	for _, k := range keysForProperties {
		mapStringForProperties += fmt.Sprintf("%v: %v,", k, this.Properties[k])
	}
	mapStringForProperties += "}"
	s := strings.Join([]string{`&ExtensionsConfiguration{`,
		`ScriptURLs:` + fmt.Sprintf("%v", this.ScriptURLs) + `,`,
		`StylesheetURLs:` + fmt.Sprintf("%v", this.StylesheetURLs) + `,`,
		`Properties:` + mapStringForProperties + `,`,
		`}`,
	}, "")
	return s
}
func (this *FeaturesConfiguration) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&FeaturesConfiguration{`,
		`InactivityTimeoutMinutes:` + fmt.Sprintf("%v", this.InactivityTimeoutMinutes) + `,`,
		`ClusterResourceOverridesEnabled:` + fmt.Sprintf("%v", this.ClusterResourceOverridesEnabled) + `,`,
		`}`,
	}, "")
	return s
}
func (this *WebConsoleConfiguration) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&WebConsoleConfiguration{`,
		`ServingInfo:` + strings.Replace(strings.Replace(this.ServingInfo.String(), "HTTPServingInfo", "github_com_openshift_api_config_v1.HTTPServingInfo", 1), `&`, ``, 1) + `,`,
		`ClusterInfo:` + strings.Replace(strings.Replace(this.ClusterInfo.String(), "ClusterInfo", "ClusterInfo", 1), `&`, ``, 1) + `,`,
		`Features:` + strings.Replace(strings.Replace(this.Features.String(), "FeaturesConfiguration", "FeaturesConfiguration", 1), `&`, ``, 1) + `,`,
		`Extensions:` + strings.Replace(strings.Replace(this.Extensions.String(), "ExtensionsConfiguration", "ExtensionsConfiguration", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringGenerated(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *ClusterInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ClusterInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClusterInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConsolePublicURL", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConsolePublicURL = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MasterPublicURL", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MasterPublicURL = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LoggingPublicURL", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LoggingPublicURL = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MetricsPublicURL", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MetricsPublicURL = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LogoutPublicURL", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LogoutPublicURL = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ExtensionsConfiguration) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ExtensionsConfiguration: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExtensionsConfiguration: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ScriptURLs", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ScriptURLs = append(m.ScriptURLs, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StylesheetURLs", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StylesheetURLs = append(m.StylesheetURLs, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Properties", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var keykey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				keykey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			var stringLenmapkey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLenmapkey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLenmapkey := int(stringLenmapkey)
			if intStringLenmapkey < 0 {
				return ErrInvalidLengthGenerated
			}
			postStringIndexmapkey := iNdEx + intStringLenmapkey
			if postStringIndexmapkey > l {
				return io.ErrUnexpectedEOF
			}
			mapkey := string(dAtA[iNdEx:postStringIndexmapkey])
			iNdEx = postStringIndexmapkey
			if m.Properties == nil {
				m.Properties = make(map[string]string)
			}
			if iNdEx < postIndex {
				var valuekey uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenerated
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					valuekey |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				var stringLenmapvalue uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenerated
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLenmapvalue |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLenmapvalue := int(stringLenmapvalue)
				if intStringLenmapvalue < 0 {
					return ErrInvalidLengthGenerated
				}
				postStringIndexmapvalue := iNdEx + intStringLenmapvalue
				if postStringIndexmapvalue > l {
					return io.ErrUnexpectedEOF
				}
				mapvalue := string(dAtA[iNdEx:postStringIndexmapvalue])
				iNdEx = postStringIndexmapvalue
				m.Properties[mapkey] = mapvalue
			} else {
				var mapvalue string
				m.Properties[mapkey] = mapvalue
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *FeaturesConfiguration) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: FeaturesConfiguration: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FeaturesConfiguration: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field InactivityTimeoutMinutes", wireType)
			}
			m.InactivityTimeoutMinutes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.InactivityTimeoutMinutes |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClusterResourceOverridesEnabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.ClusterResourceOverridesEnabled = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *WebConsoleConfiguration) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: WebConsoleConfiguration: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WebConsoleConfiguration: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServingInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ServingInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClusterInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ClusterInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Features", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Features.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Extensions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Extensions.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGenerated(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthGenerated
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGenerated
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipGenerated(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthGenerated = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerated   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/openshift/api/webconsole/v1/generated.proto", fileDescriptorGenerated)
}

var fileDescriptorGenerated = []byte{
	// 732 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0xcf, 0x6b, 0xdb, 0x48,
	0x14, 0xc7, 0x2d, 0x7b, 0x03, 0xc9, 0x18, 0x12, 0xa3, 0xdd, 0x25, 0xc6, 0x07, 0xd9, 0xf8, 0xb0,
	0xeb, 0xcb, 0x4a, 0xc4, 0x59, 0x42, 0x08, 0x84, 0x65, 0xe5, 0xcd, 0xd2, 0x80, 0x43, 0xc2, 0x24,
	0xa1, 0x50, 0x7a, 0x91, 0x95, 0x67, 0x79, 0x6a, 0x69, 0x46, 0x9d, 0x19, 0xb9, 0xf5, 0xad, 0xd0,
	0x7f, 0xa0, 0x7f, 0x56, 0xe8, 0x29, 0xd0, 0x4b, 0x4e, 0xa6, 0x51, 0x29, 0xf4, 0xcf, 0x28, 0x92,
	0x55, 0xeb, 0x87, 0xe3, 0x24, 0xd0, 0xdb, 0xcc, 0xbc, 0xf7, 0xf9, 0xbe, 0xf7, 0x66, 0xde, 0x1b,
	0xb4, 0xe7, 0x10, 0x39, 0x0a, 0x06, 0xba, 0xcd, 0x3c, 0x83, 0xf9, 0x40, 0xc5, 0x88, 0x0c, 0xa5,
	0x61, 0xf9, 0xc4, 0x78, 0x03, 0x03, 0x9b, 0x51, 0xc1, 0x5c, 0x30, 0x26, 0x3b, 0x86, 0x03, 0x14,
	0xb8, 0x25, 0xe1, 0x4a, 0xf7, 0x39, 0x93, 0x4c, 0xfd, 0x23, 0xe5, 0xf4, 0x05, 0xa7, 0x5b, 0x3e,
	0xd1, 0x53, 0x4e, 0x9f, 0xec, 0x34, 0xfe, 0xca, 0xe8, 0x3b, 0xcc, 0x61, 0x46, 0x8c, 0x0f, 0x82,
	0x61, 0xbc, 0x8b, 0x37, 0xf1, 0x6a, 0x2e, 0xdb, 0xe8, 0xae, 0x4c, 0xc7, 0x66, 0x74, 0x48, 0x9c,
	0x7b, 0x52, 0x69, 0xfc, 0x3d, 0xde, 0x17, 0x3a, 0x61, 0x91, 0x97, 0x67, 0xd9, 0x23, 0x42, 0x81,
	0x4f, 0x0d, 0x7f, 0xec, 0x44, 0x07, 0xc2, 0xf0, 0x40, 0x5a, 0xf7, 0x51, 0xc6, 0x2a, 0x8a, 0x07,
	0x54, 0x12, 0x0f, 0x96, 0x80, 0xbd, 0xc7, 0x00, 0x61, 0x8f, 0xc0, 0xb3, 0x96, 0xb8, 0xdd, 0x55,
	0x5c, 0x20, 0x89, 0x6b, 0x10, 0x2a, 0x85, 0xe4, 0x45, 0xa8, 0xfd, 0xb5, 0x8c, 0xaa, 0x3d, 0x37,
	0x10, 0x12, 0xf8, 0x31, 0x1d, 0x32, 0xf5, 0x3f, 0x54, 0x4b, 0x2e, 0xf5, 0x2c, 0x18, 0xb8, 0xc4,
	0xbe, 0xc4, 0xfd, 0xba, 0xd2, 0x52, 0x3a, 0x1b, 0x66, 0xfd, 0x7a, 0xd6, 0x2c, 0x85, 0xb3, 0x66,
	0xad, 0x57, 0xb0, 0xe3, 0x25, 0x42, 0xfd, 0x17, 0x6d, 0x79, 0x56, 0xa4, 0x99, 0x8a, 0x94, 0x63,
	0x91, 0xed, 0x44, 0x64, 0xeb, 0x24, 0x6f, 0xc6, 0x45, 0xff, 0x28, 0x11, 0x97, 0x39, 0x0e, 0xa1,
	0x4e, 0xaa, 0x51, 0xc9, 0x27, 0xd2, 0x2f, 0xd8, 0xf1, 0x12, 0x11, 0xa9, 0x78, 0x20, 0x39, 0xb1,
	0x45, 0xaa, 0xf2, 0x4b, 0x5e, 0xe5, 0xa4, 0x60, 0xc7, 0x4b, 0x44, 0x54, 0x8e, 0xcb, 0x1c, 0x16,
	0xc8, 0x54, 0x64, 0x2d, 0x5f, 0x4e, 0x3f, 0x6f, 0xc6, 0x45, 0xff, 0xf6, 0xc7, 0x32, 0xda, 0x3e,
	0x7a, 0x2b, 0x81, 0x0a, 0xc2, 0xa8, 0xe8, 0xc5, 0x3d, 0x16, 0x70, 0x4b, 0x12, 0x46, 0x55, 0x1d,
	0x21, 0x61, 0x73, 0xe2, 0xcb, 0x4b, 0xdc, 0x17, 0x75, 0xa5, 0x55, 0xe9, 0x6c, 0x98, 0x9b, 0xe1,
	0xac, 0x89, 0xce, 0x17, 0xa7, 0x38, 0xe3, 0xa1, 0x1e, 0xa0, 0x4d, 0x21, 0xa7, 0x2e, 0x88, 0x11,
	0xc0, 0x9c, 0x29, 0xc7, 0x8c, 0x1a, 0xce, 0x9a, 0x9b, 0xe7, 0x39, 0x0b, 0x2e, 0x78, 0xaa, 0xef,
	0x15, 0x84, 0x7c, 0xce, 0x7c, 0xe0, 0x92, 0x80, 0xa8, 0x57, 0x5a, 0x95, 0x4e, 0xb5, 0x7b, 0xaa,
	0x3f, 0x6d, 0xc8, 0xf4, 0x15, 0x15, 0xe8, 0x67, 0x0b, 0xc5, 0x23, 0x2a, 0xf9, 0xd4, 0x54, 0x93,
	0x7b, 0x41, 0xa9, 0x01, 0x67, 0xc2, 0x36, 0x0e, 0xd1, 0x56, 0x01, 0x51, 0x6b, 0xa8, 0x32, 0x86,
	0xe9, 0xbc, 0xd7, 0x70, 0xb4, 0x54, 0x7f, 0x43, 0x6b, 0x13, 0xcb, 0x0d, 0x60, 0xde, 0x3a, 0x78,
	0xbe, 0x39, 0x28, 0xef, 0x2b, 0xed, 0x6f, 0x0a, 0xfa, 0xfd, 0x7f, 0xb0, 0x64, 0xc0, 0xa1, 0x70,
	0x95, 0x2f, 0x51, 0x9d, 0x50, 0xcb, 0x96, 0x64, 0x42, 0xe4, 0xf4, 0x82, 0x78, 0xc0, 0x02, 0x79,
	0x42, 0x68, 0x20, 0x41, 0xc4, 0xd2, 0x15, 0xb3, 0x95, 0xa4, 0x56, 0x3f, 0x5e, 0xe1, 0x87, 0x57,
	0x2a, 0xa8, 0xaf, 0x51, 0xd3, 0x9e, 0xcf, 0x0a, 0x06, 0xc1, 0x02, 0x6e, 0xc3, 0xe9, 0x04, 0x38,
	0x27, 0x57, 0x51, 0x11, 0xd6, 0xc0, 0x85, 0xab, 0x38, 0xd7, 0x75, 0xf3, 0xcf, 0x24, 0x48, 0xb3,
	0xf7, 0xb0, 0x3b, 0x7e, 0x4c, 0xaf, 0xfd, 0xa9, 0x82, 0xb6, 0x9f, 0xc3, 0x20, 0x99, 0xb9, 0x7c,
	0xb1, 0xaf, 0x50, 0x55, 0x00, 0x9f, 0x10, 0xea, 0x44, 0xa3, 0x1b, 0xd7, 0x57, 0xed, 0xee, 0xae,
	0x7e, 0xcb, 0xf9, 0xcf, 0x16, 0xbd, 0xe3, 0xb3, 0x8b, 0x8b, 0xb3, 0xf3, 0x14, 0x35, 0x7f, 0x4d,
	0xf2, 0xad, 0x66, 0x0e, 0x71, 0x56, 0x3c, 0x8a, 0x65, 0xa7, 0xdf, 0x44, 0x5c, 0xe6, 0x83, 0xb1,
	0xf2, 0x7d, 0x93, 0xf9, 0x61, 0xd2, 0x58, 0x99, 0x43, 0x9c, 0x15, 0x57, 0xc7, 0x68, 0x7d, 0x98,
	0xbc, 0x6e, 0x3c, 0xf2, 0xd5, 0xee, 0xe1, 0x53, 0x03, 0xdd, 0xdb, 0x15, 0x66, 0x2d, 0x09, 0xb9,
	0xfe, 0xc3, 0x8c, 0x17, 0x01, 0x54, 0x81, 0x10, 0x2c, 0xba, 0x3a, 0xfe, 0x1b, 0xaa, 0xdd, 0x7f,
	0x7e, 0x72, 0x1e, 0xd2, 0xfe, 0x4f, 0x1d, 0x70, 0x26, 0x8c, 0xd9, 0xb9, 0xbe, 0xd3, 0x4a, 0x37,
	0x77, 0x5a, 0xe9, 0xf6, 0x4e, 0x2b, 0xbd, 0x0b, 0x35, 0xe5, 0x3a, 0xd4, 0x94, 0x9b, 0x50, 0x53,
	0x6e, 0x43, 0x4d, 0xf9, 0x1c, 0x6a, 0xca, 0x87, 0x2f, 0x5a, 0xe9, 0x45, 0x79, 0xb2, 0xf3, 0x3d,
	0x00, 0x00, 0xff, 0xff, 0x62, 0xf7, 0x90, 0x27, 0x37, 0x07, 0x00, 0x00,
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/resources/customer.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v2/enums"
	_ "google.golang.org/genproto/googleapis/api/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// A customer.
type Customer struct {
	// The resource name of the customer.
	// Customer resource names have the form:
	//
	// `customers/{customer_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The ID of the customer.
	Id *wrappers.Int64Value `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	// Optional, non-unique descriptive name of the customer.
	DescriptiveName *wrappers.StringValue `protobuf:"bytes,4,opt,name=descriptive_name,json=descriptiveName,proto3" json:"descriptive_name,omitempty"`
	// The currency in which the account operates.
	// A subset of the currency codes from the ISO 4217 standard is
	// supported.
	CurrencyCode *wrappers.StringValue `protobuf:"bytes,5,opt,name=currency_code,json=currencyCode,proto3" json:"currency_code,omitempty"`
	// The local timezone ID of the customer.
	TimeZone *wrappers.StringValue `protobuf:"bytes,6,opt,name=time_zone,json=timeZone,proto3" json:"time_zone,omitempty"`
	// The URL template for constructing a tracking URL out of parameters.
	TrackingUrlTemplate *wrappers.StringValue `protobuf:"bytes,7,opt,name=tracking_url_template,json=trackingUrlTemplate,proto3" json:"tracking_url_template,omitempty"`
	// The URL template for appending params to the final URL
	FinalUrlSuffix *wrappers.StringValue `protobuf:"bytes,11,opt,name=final_url_suffix,json=finalUrlSuffix,proto3" json:"final_url_suffix,omitempty"`
	// Whether auto-tagging is enabled for the customer.
	AutoTaggingEnabled *wrappers.BoolValue `protobuf:"bytes,8,opt,name=auto_tagging_enabled,json=autoTaggingEnabled,proto3" json:"auto_tagging_enabled,omitempty"`
	// Whether the Customer has a Partners program badge. If the Customer is not
	// associated with the Partners program, this will be false. For more
	// information, see https://support.google.com/partners/answer/3125774.
	HasPartnersBadge *wrappers.BoolValue `protobuf:"bytes,9,opt,name=has_partners_badge,json=hasPartnersBadge,proto3" json:"has_partners_badge,omitempty"`
	// Whether the customer is a manager.
	Manager *wrappers.BoolValue `protobuf:"bytes,12,opt,name=manager,proto3" json:"manager,omitempty"`
	// Whether the customer is a test account.
	TestAccount *wrappers.BoolValue `protobuf:"bytes,13,opt,name=test_account,json=testAccount,proto3" json:"test_account,omitempty"`
	// Call reporting setting for a customer.
	CallReportingSetting *CallReportingSetting `protobuf:"bytes,10,opt,name=call_reporting_setting,json=callReportingSetting,proto3" json:"call_reporting_setting,omitempty"`
	// Conversion tracking setting for a customer.
	ConversionTrackingSetting *ConversionTrackingSetting `protobuf:"bytes,14,opt,name=conversion_tracking_setting,json=conversionTrackingSetting,proto3" json:"conversion_tracking_setting,omitempty"`
	// Remarketing setting for a customer.
	RemarketingSetting *RemarketingSetting `protobuf:"bytes,15,opt,name=remarketing_setting,json=remarketingSetting,proto3" json:"remarketing_setting,omitempty"`
	// Reasons why the customer is not eligible to use PaymentMode.CONVERSIONS. If
	// the list is empty, the customer is eligible. This field is read-only.
	PayPerConversionEligibilityFailureReasons []enums.CustomerPayPerConversionEligibilityFailureReasonEnum_CustomerPayPerConversionEligibilityFailureReason `protobuf:"varint,16,rep,packed,name=pay_per_conversion_eligibility_failure_reasons,json=payPerConversionEligibilityFailureReasons,proto3,enum=google.ads.googleads.v2.enums.CustomerPayPerConversionEligibilityFailureReasonEnum_CustomerPayPerConversionEligibilityFailureReason" json:"pay_per_conversion_eligibility_failure_reasons,omitempty"`
	XXX_NoUnkeyedLiteral                      struct{}                                                                                                      `json:"-"`
	XXX_unrecognized                          []byte                                                                                                        `json:"-"`
	XXX_sizecache                             int32                                                                                                         `json:"-"`
}

func (m *Customer) Reset()         { *m = Customer{} }
func (m *Customer) String() string { return proto.CompactTextString(m) }
func (*Customer) ProtoMessage()    {}
func (*Customer) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b507c5c606ef490, []int{0}
}

func (m *Customer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Customer.Unmarshal(m, b)
}
func (m *Customer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Customer.Marshal(b, m, deterministic)
}
func (m *Customer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Customer.Merge(m, src)
}
func (m *Customer) XXX_Size() int {
	return xxx_messageInfo_Customer.Size(m)
}
func (m *Customer) XXX_DiscardUnknown() {
	xxx_messageInfo_Customer.DiscardUnknown(m)
}

var xxx_messageInfo_Customer proto.InternalMessageInfo

func (m *Customer) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *Customer) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Customer) GetDescriptiveName() *wrappers.StringValue {
	if m != nil {
		return m.DescriptiveName
	}
	return nil
}

func (m *Customer) GetCurrencyCode() *wrappers.StringValue {
	if m != nil {
		return m.CurrencyCode
	}
	return nil
}

func (m *Customer) GetTimeZone() *wrappers.StringValue {
	if m != nil {
		return m.TimeZone
	}
	return nil
}

func (m *Customer) GetTrackingUrlTemplate() *wrappers.StringValue {
	if m != nil {
		return m.TrackingUrlTemplate
	}
	return nil
}

func (m *Customer) GetFinalUrlSuffix() *wrappers.StringValue {
	if m != nil {
		return m.FinalUrlSuffix
	}
	return nil
}

func (m *Customer) GetAutoTaggingEnabled() *wrappers.BoolValue {
	if m != nil {
		return m.AutoTaggingEnabled
	}
	return nil
}

func (m *Customer) GetHasPartnersBadge() *wrappers.BoolValue {
	if m != nil {
		return m.HasPartnersBadge
	}
	return nil
}

func (m *Customer) GetManager() *wrappers.BoolValue {
	if m != nil {
		return m.Manager
	}
	return nil
}

func (m *Customer) GetTestAccount() *wrappers.BoolValue {
	if m != nil {
		return m.TestAccount
	}
	return nil
}

func (m *Customer) GetCallReportingSetting() *CallReportingSetting {
	if m != nil {
		return m.CallReportingSetting
	}
	return nil
}

func (m *Customer) GetConversionTrackingSetting() *ConversionTrackingSetting {
	if m != nil {
		return m.ConversionTrackingSetting
	}
	return nil
}

func (m *Customer) GetRemarketingSetting() *RemarketingSetting {
	if m != nil {
		return m.RemarketingSetting
	}
	return nil
}

func (m *Customer) GetPayPerConversionEligibilityFailureReasons() []enums.CustomerPayPerConversionEligibilityFailureReasonEnum_CustomerPayPerConversionEligibilityFailureReason {
	if m != nil {
		return m.PayPerConversionEligibilityFailureReasons
	}
	return nil
}

// Call reporting setting for a customer.
type CallReportingSetting struct {
	// Enable reporting of phone call events by redirecting them via Google
	// System.
	CallReportingEnabled *wrappers.BoolValue `protobuf:"bytes,1,opt,name=call_reporting_enabled,json=callReportingEnabled,proto3" json:"call_reporting_enabled,omitempty"`
	// Whether to enable call conversion reporting.
	CallConversionReportingEnabled *wrappers.BoolValue `protobuf:"bytes,2,opt,name=call_conversion_reporting_enabled,json=callConversionReportingEnabled,proto3" json:"call_conversion_reporting_enabled,omitempty"`
	// Customer-level call conversion action to attribute a call conversion to.
	// If not set a default conversion action is used. Only in effect when
	// call_conversion_reporting_enabled is set to true.
	CallConversionAction *wrappers.StringValue `protobuf:"bytes,9,opt,name=call_conversion_action,json=callConversionAction,proto3" json:"call_conversion_action,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CallReportingSetting) Reset()         { *m = CallReportingSetting{} }
func (m *CallReportingSetting) String() string { return proto.CompactTextString(m) }
func (*CallReportingSetting) ProtoMessage()    {}
func (*CallReportingSetting) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b507c5c606ef490, []int{1}
}

func (m *CallReportingSetting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CallReportingSetting.Unmarshal(m, b)
}
func (m *CallReportingSetting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CallReportingSetting.Marshal(b, m, deterministic)
}
func (m *CallReportingSetting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallReportingSetting.Merge(m, src)
}
func (m *CallReportingSetting) XXX_Size() int {
	return xxx_messageInfo_CallReportingSetting.Size(m)
}
func (m *CallReportingSetting) XXX_DiscardUnknown() {
	xxx_messageInfo_CallReportingSetting.DiscardUnknown(m)
}

var xxx_messageInfo_CallReportingSetting proto.InternalMessageInfo

func (m *CallReportingSetting) GetCallReportingEnabled() *wrappers.BoolValue {
	if m != nil {
		return m.CallReportingEnabled
	}
	return nil
}

func (m *CallReportingSetting) GetCallConversionReportingEnabled() *wrappers.BoolValue {
	if m != nil {
		return m.CallConversionReportingEnabled
	}
	return nil
}

func (m *CallReportingSetting) GetCallConversionAction() *wrappers.StringValue {
	if m != nil {
		return m.CallConversionAction
	}
	return nil
}

// A collection of customer-wide settings related to Google Ads Conversion
// Tracking.
type ConversionTrackingSetting struct {
	// The conversion tracking id used for this account. This id is automatically
	// assigned after any conversion tracking feature is used. If the customer
	// doesn't use conversion tracking, this is 0. This field is read-only.
	ConversionTrackingId *wrappers.Int64Value `protobuf:"bytes,1,opt,name=conversion_tracking_id,json=conversionTrackingId,proto3" json:"conversion_tracking_id,omitempty"`
	// The conversion tracking id of the customer's manager. This is set when the
	// customer is opted into cross account conversion tracking, and it overrides
	// conversion_tracking_id. This field can only be managed through the Google
	// Ads UI. This field is read-only.
	CrossAccountConversionTrackingId *wrappers.Int64Value `protobuf:"bytes,2,opt,name=cross_account_conversion_tracking_id,json=crossAccountConversionTrackingId,proto3" json:"cross_account_conversion_tracking_id,omitempty"`
	XXX_NoUnkeyedLiteral             struct{}             `json:"-"`
	XXX_unrecognized                 []byte               `json:"-"`
	XXX_sizecache                    int32                `json:"-"`
}

func (m *ConversionTrackingSetting) Reset()         { *m = ConversionTrackingSetting{} }
func (m *ConversionTrackingSetting) String() string { return proto.CompactTextString(m) }
func (*ConversionTrackingSetting) ProtoMessage()    {}
func (*ConversionTrackingSetting) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b507c5c606ef490, []int{2}
}

func (m *ConversionTrackingSetting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConversionTrackingSetting.Unmarshal(m, b)
}
func (m *ConversionTrackingSetting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConversionTrackingSetting.Marshal(b, m, deterministic)
}
func (m *ConversionTrackingSetting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConversionTrackingSetting.Merge(m, src)
}
func (m *ConversionTrackingSetting) XXX_Size() int {
	return xxx_messageInfo_ConversionTrackingSetting.Size(m)
}
func (m *ConversionTrackingSetting) XXX_DiscardUnknown() {
	xxx_messageInfo_ConversionTrackingSetting.DiscardUnknown(m)
}

var xxx_messageInfo_ConversionTrackingSetting proto.InternalMessageInfo

func (m *ConversionTrackingSetting) GetConversionTrackingId() *wrappers.Int64Value {
	if m != nil {
		return m.ConversionTrackingId
	}
	return nil
}

func (m *ConversionTrackingSetting) GetCrossAccountConversionTrackingId() *wrappers.Int64Value {
	if m != nil {
		return m.CrossAccountConversionTrackingId
	}
	return nil
}

// Remarketing setting for a customer.
type RemarketingSetting struct {
	// The Google global site tag.
	GoogleGlobalSiteTag  *wrappers.StringValue `protobuf:"bytes,1,opt,name=google_global_site_tag,json=googleGlobalSiteTag,proto3" json:"google_global_site_tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *RemarketingSetting) Reset()         { *m = RemarketingSetting{} }
func (m *RemarketingSetting) String() string { return proto.CompactTextString(m) }
func (*RemarketingSetting) ProtoMessage()    {}
func (*RemarketingSetting) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b507c5c606ef490, []int{3}
}

func (m *RemarketingSetting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemarketingSetting.Unmarshal(m, b)
}
func (m *RemarketingSetting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemarketingSetting.Marshal(b, m, deterministic)
}
func (m *RemarketingSetting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemarketingSetting.Merge(m, src)
}
func (m *RemarketingSetting) XXX_Size() int {
	return xxx_messageInfo_RemarketingSetting.Size(m)
}
func (m *RemarketingSetting) XXX_DiscardUnknown() {
	xxx_messageInfo_RemarketingSetting.DiscardUnknown(m)
}

var xxx_messageInfo_RemarketingSetting proto.InternalMessageInfo

func (m *RemarketingSetting) GetGoogleGlobalSiteTag() *wrappers.StringValue {
	if m != nil {
		return m.GoogleGlobalSiteTag
	}
	return nil
}

func init() {
	proto.RegisterType((*Customer)(nil), "google.ads.googleads.v2.resources.Customer")
	proto.RegisterType((*CallReportingSetting)(nil), "google.ads.googleads.v2.resources.CallReportingSetting")
	proto.RegisterType((*ConversionTrackingSetting)(nil), "google.ads.googleads.v2.resources.ConversionTrackingSetting")
	proto.RegisterType((*RemarketingSetting)(nil), "google.ads.googleads.v2.resources.RemarketingSetting")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/resources/customer.proto", fileDescriptor_9b507c5c606ef490)
}

var fileDescriptor_9b507c5c606ef490 = []byte{
	// 878 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x96, 0xdf, 0x6e, 0xe4, 0x34,
	0x14, 0xc6, 0x95, 0x29, 0xec, 0xb6, 0xee, 0x9f, 0xad, 0xbc, 0xc3, 0x2a, 0xdb, 0x5d, 0xad, 0xda,
	0xc2, 0x4a, 0x45, 0x48, 0x19, 0x34, 0x2c, 0x20, 0x06, 0xb8, 0x98, 0x56, 0xdd, 0xb2, 0x08, 0xa1,
	0x21, 0xed, 0xf6, 0xa2, 0xaa, 0x64, 0x79, 0x92, 0x33, 0x5e, 0xab, 0x8e, 0x1d, 0xd9, 0x4e, 0xa1,
	0x88, 0x4b, 0xae, 0x78, 0x0c, 0xb8, 0xe3, 0x51, 0x78, 0x05, 0x5e, 0x00, 0xf1, 0x08, 0x5c, 0xa1,
	0x38, 0x71, 0xa6, 0xd5, 0x74, 0x9a, 0xd9, 0xab, 0xf1, 0xd8, 0xe7, 0xfb, 0x7d, 0xb6, 0x73, 0xce,
	0x49, 0xd0, 0xc7, 0x4c, 0x29, 0x26, 0xa0, 0x47, 0x53, 0xd3, 0xab, 0x86, 0xe5, 0xe8, 0xb2, 0xdf,
	0xd3, 0x60, 0x54, 0xa1, 0x13, 0x30, 0xbd, 0xa4, 0x30, 0x56, 0x65, 0xa0, 0xa3, 0x5c, 0x2b, 0xab,
	0xf0, 0x4e, 0x15, 0x16, 0xd1, 0xd4, 0x44, 0x8d, 0x22, 0xba, 0xec, 0x47, 0x8d, 0x62, 0xeb, 0x6c,
	0x1e, 0x14, 0x64, 0x91, 0x4d, 0x81, 0x24, 0xa7, 0x57, 0x24, 0x07, 0x4d, 0x12, 0x25, 0x2f, 0x41,
	0x1b, 0xae, 0x24, 0x01, 0xc1, 0x19, 0x1f, 0x73, 0xc1, 0xed, 0x15, 0x99, 0x50, 0x2e, 0x0a, 0x0d,
	0x44, 0x03, 0x35, 0x4a, 0x56, 0xf6, 0x5b, 0xcf, 0x6a, 0xb6, 0xfb, 0x37, 0x2e, 0x26, 0xbd, 0x1f,
	0x35, 0xcd, 0x73, 0xd0, 0xa6, 0x5e, 0x7f, 0xea, 0xbd, 0x73, 0xde, 0xa3, 0x52, 0x2a, 0x4b, 0x2d,
	0x57, 0xb2, 0x5e, 0xdd, 0xfd, 0x0d, 0xa1, 0xe5, 0x83, 0xda, 0x1e, 0xbf, 0x8f, 0xd6, 0xfd, 0x9e,
	0x89, 0xa4, 0x19, 0x84, 0xc1, 0x76, 0xb0, 0xb7, 0x12, 0xaf, 0xf9, 0xc9, 0xef, 0x69, 0x06, 0xf8,
	0x23, 0xd4, 0xe1, 0x69, 0xb8, 0xb4, 0x1d, 0xec, 0xad, 0xf6, 0x9f, 0xd4, 0x07, 0x8e, 0xbc, 0x79,
	0xf4, 0x4a, 0xda, 0xcf, 0x5e, 0x9c, 0x52, 0x51, 0x40, 0xdc, 0xe1, 0x29, 0x3e, 0x42, 0x9b, 0x29,
	0x98, 0x44, 0xf3, 0xdc, 0xf2, 0xcb, 0x1a, 0xfa, 0x8e, 0x93, 0x3e, 0x9d, 0x91, 0x1e, 0x5b, 0xcd,
	0x25, 0xab, 0xb4, 0x0f, 0xae, 0xa9, 0x9c, 0xeb, 0x10, 0xad, 0x27, 0x85, 0xd6, 0x20, 0x93, 0x2b,
	0x92, 0xa8, 0x14, 0xc2, 0x77, 0x17, 0xa0, 0xac, 0x79, 0xc9, 0x81, 0x4a, 0x01, 0x7f, 0x81, 0x56,
	0x2c, 0xcf, 0x80, 0xfc, 0xac, 0x24, 0x84, 0xf7, 0x16, 0x90, 0x2f, 0x97, 0xe1, 0x67, 0x4a, 0x02,
	0x1e, 0xa1, 0xf7, 0xac, 0xa6, 0xc9, 0x05, 0x97, 0x8c, 0x14, 0x5a, 0x10, 0x0b, 0x59, 0x2e, 0xa8,
	0x85, 0xf0, 0xfe, 0x02, 0x98, 0x87, 0x5e, 0xfa, 0x5a, 0x8b, 0x93, 0x5a, 0x88, 0x5f, 0xa2, 0xcd,
	0x09, 0x97, 0x54, 0x38, 0x9c, 0x29, 0x26, 0x13, 0xfe, 0x53, 0xb8, 0xba, 0x00, 0x6c, 0xc3, 0xa9,
	0x5e, 0x6b, 0x71, 0xec, 0x34, 0xf8, 0x3b, 0xd4, 0xa5, 0x85, 0x55, 0xc4, 0x52, 0xc6, 0xca, 0xdd,
	0x81, 0xa4, 0x63, 0x01, 0x69, 0xb8, 0xec, 0x58, 0x5b, 0x33, 0xac, 0x7d, 0xa5, 0x44, 0x45, 0xc2,
	0xa5, 0xee, 0xa4, 0x92, 0x1d, 0x56, 0x2a, 0xfc, 0x0d, 0xc2, 0x6f, 0xa8, 0x21, 0x39, 0xd5, 0x56,
	0x82, 0x36, 0x64, 0x4c, 0x53, 0x06, 0xe1, 0x4a, 0x2b, 0x6b, 0xf3, 0x0d, 0x35, 0xa3, 0x5a, 0xb4,
	0x5f, 0x6a, 0xf0, 0x0b, 0x74, 0x3f, 0xa3, 0x92, 0x32, 0xd0, 0xe1, 0x5a, 0xab, 0xdc, 0x87, 0xe2,
	0xaf, 0xd1, 0x9a, 0x05, 0x63, 0x09, 0x4d, 0x12, 0x55, 0x48, 0x1b, 0xae, 0xb7, 0x4a, 0x57, 0xcb,
	0xf8, 0x61, 0x15, 0x8e, 0x33, 0xf4, 0x28, 0xa1, 0x42, 0x10, 0x0d, 0xb9, 0xd2, 0xb6, 0xbc, 0x0e,
	0x03, 0xb6, 0xfc, 0x0d, 0x91, 0x03, 0x7d, 0x1e, 0xb5, 0x96, 0x6a, 0x74, 0x40, 0x85, 0x88, 0xbd,
	0xfe, 0xb8, 0x92, 0xc7, 0xdd, 0xe4, 0x96, 0x59, 0xfc, 0x0b, 0x7a, 0x72, 0xad, 0x50, 0x9b, 0x04,
	0xf1, 0x9e, 0x1b, 0xce, 0xf3, 0xab, 0x45, 0x3c, 0x1b, 0xca, 0x49, 0x0d, 0xf1, 0xc6, 0x8f, 0x93,
	0x79, 0x4b, 0x78, 0x82, 0x1e, 0x6a, 0xc8, 0xa8, 0xbe, 0x80, 0x1b, 0x27, 0x7d, 0xe0, 0x5c, 0x3f,
	0x5d, 0xc0, 0x35, 0x9e, 0xaa, 0xbd, 0x1d, 0xd6, 0x33, 0x73, 0xf8, 0x9f, 0x00, 0x45, 0x6f, 0xd5,
	0x97, 0x4c, 0xb8, 0xb9, 0xbd, 0xb4, 0xb7, 0xd1, 0xff, 0x35, 0x98, 0xbb, 0x09, 0xd7, 0xf6, 0x22,
	0xdf, 0x77, 0x46, 0xf4, 0x6a, 0x04, 0x7a, 0x7a, 0x09, 0x87, 0x53, 0xf4, 0xcb, 0x8a, 0x1c, 0x3b,
	0xf0, 0xa1, 0x2c, 0xb2, 0xb7, 0x16, 0xc5, 0x1f, 0xe6, 0x0b, 0x46, 0x9a, 0xdd, 0x3f, 0x3a, 0xa8,
	0x7b, 0xdb, 0xf3, 0xc7, 0xa3, 0x99, 0xc4, 0xf2, 0x75, 0x16, 0xb4, 0x66, 0xe8, 0xcd, 0xdc, 0xf1,
	0x95, 0x06, 0x68, 0xc7, 0x11, 0xaf, 0xdd, 0xe8, 0x2c, 0xbc, 0xd3, 0x0a, 0x7f, 0x56, 0x42, 0xa6,
	0x47, 0x9b, 0xb1, 0x89, 0xeb, 0x8d, 0x5f, 0xb3, 0xa1, 0x49, 0xd9, 0xff, 0xeb, 0xa2, 0xbe, 0xbb,
	0xd9, 0x74, 0x6f, 0xd2, 0x87, 0x4e, 0xb9, 0xfb, 0x77, 0x80, 0x1e, 0xcf, 0xcd, 0x58, 0xfc, 0x03,
	0x7a, 0x74, 0x5b, 0x51, 0x70, 0x7f, 0x55, 0x77, 0xbe, 0x32, 0xba, 0xb3, 0xe9, 0xfe, 0x2a, 0xc5,
	0x17, 0xe8, 0x83, 0x44, 0x2b, 0x63, 0x7c, 0x5b, 0x20, 0x73, 0x0c, 0x3a, 0xed, 0x06, 0xdb, 0x0e,
	0x54, 0xf7, 0x8b, 0x83, 0x5b, 0xcc, 0x76, 0x19, 0xc2, 0xb3, 0x85, 0x51, 0x9e, 0xaa, 0xa2, 0x12,
	0x26, 0xd4, 0x98, 0x0a, 0x62, 0xb8, 0x85, 0xb2, 0xe9, 0xd6, 0xa7, 0x6a, 0x79, 0x03, 0x54, 0x8b,
	0x47, 0x4e, 0x7a, 0xcc, 0x2d, 0x9c, 0x50, 0xb6, 0xff, 0x5f, 0x80, 0x9e, 0x27, 0x2a, 0x6b, 0x2f,
	0xd4, 0xfd, 0xf5, 0x26, 0xe7, 0x4b, 0xf8, 0x28, 0x38, 0xfb, 0xb6, 0xd6, 0x30, 0x25, 0xa8, 0x64,
	0x91, 0xd2, 0xac, 0xc7, 0x40, 0x3a, 0x6b, 0xff, 0x79, 0x91, 0x73, 0x73, 0xc7, 0x27, 0xcc, 0x97,
	0xcd, 0xe8, 0xf7, 0xce, 0xd2, 0xd1, 0x70, 0xf8, 0x67, 0x67, 0xe7, 0xa8, 0x42, 0x0e, 0x53, 0x13,
	0x55, 0xc3, 0x72, 0x74, 0xda, 0x8f, 0x62, 0x1f, 0xf9, 0x97, 0x8f, 0x39, 0x1f, 0xa6, 0xe6, 0xbc,
	0x89, 0x39, 0x3f, 0xed, 0x9f, 0x37, 0x31, 0xff, 0x76, 0x9e, 0x57, 0x0b, 0x83, 0xc1, 0x30, 0x35,
	0x83, 0x41, 0x13, 0x35, 0x18, 0x9c, 0xf6, 0x07, 0x83, 0x26, 0x6e, 0x7c, 0xcf, 0x6d, 0xf6, 0x93,
	0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0xa4, 0x52, 0x7c, 0xe3, 0x6e, 0x09, 0x00, 0x00,
}
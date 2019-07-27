// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/change_status.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/ie310mu/ie310go/forks/github.com/golang/protobuf/proto"
	wrappers "github.com/ie310mu/ie310go/forks/github.com/golang/protobuf/ptypes/wrappers"
	enums "github.com/ie310mu/ie310go/forks/google.golang.org/genproto/googleapis/ads/googleads/v1/enums"
	_ "github.com/ie310mu/ie310go/forks/google.golang.org/genproto/googleapis/api/annotations"
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

// Describes the status of returned resource.
type ChangeStatus struct {
	// The resource name of the change status.
	// Change status resource names have the form:
	//
	// `customers/{customer_id}/changeStatus/{change_status_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Time at which the most recent change has occurred on this resource.
	LastChangeDateTime *wrappers.StringValue `protobuf:"bytes,3,opt,name=last_change_date_time,json=lastChangeDateTime,proto3" json:"last_change_date_time,omitempty"`
	// Represents the type of the changed resource. This dictates what fields
	// will be set. For example, for AD_GROUP, campaign and ad_group fields will
	// be set.
	ResourceType enums.ChangeStatusResourceTypeEnum_ChangeStatusResourceType `protobuf:"varint,4,opt,name=resource_type,json=resourceType,proto3,enum=google.ads.googleads.v1.enums.ChangeStatusResourceTypeEnum_ChangeStatusResourceType" json:"resource_type,omitempty"`
	// The Campaign affected by this change.
	Campaign *wrappers.StringValue `protobuf:"bytes,5,opt,name=campaign,proto3" json:"campaign,omitempty"`
	// The AdGroup affected by this change.
	AdGroup *wrappers.StringValue `protobuf:"bytes,6,opt,name=ad_group,json=adGroup,proto3" json:"ad_group,omitempty"`
	// Represents the status of the changed resource.
	ResourceStatus enums.ChangeStatusOperationEnum_ChangeStatusOperation `protobuf:"varint,8,opt,name=resource_status,json=resourceStatus,proto3,enum=google.ads.googleads.v1.enums.ChangeStatusOperationEnum_ChangeStatusOperation" json:"resource_status,omitempty"`
	// The AdGroupAd affected by this change.
	AdGroupAd *wrappers.StringValue `protobuf:"bytes,9,opt,name=ad_group_ad,json=adGroupAd,proto3" json:"ad_group_ad,omitempty"`
	// The AdGroupCriterion affected by this change.
	AdGroupCriterion *wrappers.StringValue `protobuf:"bytes,10,opt,name=ad_group_criterion,json=adGroupCriterion,proto3" json:"ad_group_criterion,omitempty"`
	// The CampaignCriterion affected by this change.
	CampaignCriterion *wrappers.StringValue `protobuf:"bytes,11,opt,name=campaign_criterion,json=campaignCriterion,proto3" json:"campaign_criterion,omitempty"`
	// The Feed affected by this change.
	Feed *wrappers.StringValue `protobuf:"bytes,12,opt,name=feed,proto3" json:"feed,omitempty"`
	// The FeedItem affected by this change.
	FeedItem *wrappers.StringValue `protobuf:"bytes,13,opt,name=feed_item,json=feedItem,proto3" json:"feed_item,omitempty"`
	// The AdGroupFeed affected by this change.
	AdGroupFeed *wrappers.StringValue `protobuf:"bytes,14,opt,name=ad_group_feed,json=adGroupFeed,proto3" json:"ad_group_feed,omitempty"`
	// The CampaignFeed affected by this change.
	CampaignFeed *wrappers.StringValue `protobuf:"bytes,15,opt,name=campaign_feed,json=campaignFeed,proto3" json:"campaign_feed,omitempty"`
	// The AdGroupBidModifier affected by this change.
	AdGroupBidModifier   *wrappers.StringValue `protobuf:"bytes,16,opt,name=ad_group_bid_modifier,json=adGroupBidModifier,proto3" json:"ad_group_bid_modifier,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ChangeStatus) Reset()         { *m = ChangeStatus{} }
func (m *ChangeStatus) String() string { return proto.CompactTextString(m) }
func (*ChangeStatus) ProtoMessage()    {}
func (*ChangeStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a34c24bda75f035, []int{0}
}

func (m *ChangeStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangeStatus.Unmarshal(m, b)
}
func (m *ChangeStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangeStatus.Marshal(b, m, deterministic)
}
func (m *ChangeStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeStatus.Merge(m, src)
}
func (m *ChangeStatus) XXX_Size() int {
	return xxx_messageInfo_ChangeStatus.Size(m)
}
func (m *ChangeStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeStatus proto.InternalMessageInfo

func (m *ChangeStatus) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *ChangeStatus) GetLastChangeDateTime() *wrappers.StringValue {
	if m != nil {
		return m.LastChangeDateTime
	}
	return nil
}

func (m *ChangeStatus) GetResourceType() enums.ChangeStatusResourceTypeEnum_ChangeStatusResourceType {
	if m != nil {
		return m.ResourceType
	}
	return enums.ChangeStatusResourceTypeEnum_UNSPECIFIED
}

func (m *ChangeStatus) GetCampaign() *wrappers.StringValue {
	if m != nil {
		return m.Campaign
	}
	return nil
}

func (m *ChangeStatus) GetAdGroup() *wrappers.StringValue {
	if m != nil {
		return m.AdGroup
	}
	return nil
}

func (m *ChangeStatus) GetResourceStatus() enums.ChangeStatusOperationEnum_ChangeStatusOperation {
	if m != nil {
		return m.ResourceStatus
	}
	return enums.ChangeStatusOperationEnum_UNSPECIFIED
}

func (m *ChangeStatus) GetAdGroupAd() *wrappers.StringValue {
	if m != nil {
		return m.AdGroupAd
	}
	return nil
}

func (m *ChangeStatus) GetAdGroupCriterion() *wrappers.StringValue {
	if m != nil {
		return m.AdGroupCriterion
	}
	return nil
}

func (m *ChangeStatus) GetCampaignCriterion() *wrappers.StringValue {
	if m != nil {
		return m.CampaignCriterion
	}
	return nil
}

func (m *ChangeStatus) GetFeed() *wrappers.StringValue {
	if m != nil {
		return m.Feed
	}
	return nil
}

func (m *ChangeStatus) GetFeedItem() *wrappers.StringValue {
	if m != nil {
		return m.FeedItem
	}
	return nil
}

func (m *ChangeStatus) GetAdGroupFeed() *wrappers.StringValue {
	if m != nil {
		return m.AdGroupFeed
	}
	return nil
}

func (m *ChangeStatus) GetCampaignFeed() *wrappers.StringValue {
	if m != nil {
		return m.CampaignFeed
	}
	return nil
}

func (m *ChangeStatus) GetAdGroupBidModifier() *wrappers.StringValue {
	if m != nil {
		return m.AdGroupBidModifier
	}
	return nil
}

func init() {
	proto.RegisterType((*ChangeStatus)(nil), "google.ads.googleads.v1.resources.ChangeStatus")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/change_status.proto", fileDescriptor_9a34c24bda75f035)
}

var fileDescriptor_9a34c24bda75f035 = []byte{
	// 602 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xdd, 0x6a, 0xdb, 0x3e,
	0x14, 0xc7, 0x6d, 0xff, 0xfd, 0x27, 0x4a, 0xd2, 0x0f, 0x41, 0xc1, 0x94, 0x32, 0xda, 0x8d, 0x42,
	0xaf, 0xe4, 0xa5, 0x63, 0x6c, 0x73, 0x07, 0x9b, 0xdb, 0x6d, 0x65, 0x1d, 0x6b, 0x4b, 0x5a, 0x72,
	0x31, 0x02, 0x46, 0x8d, 0x54, 0x4f, 0x10, 0x4b, 0x46, 0x92, 0x5b, 0xf2, 0x00, 0x7b, 0x91, 0x5d,
	0xee, 0x51, 0xf6, 0x28, 0x7b, 0x88, 0x31, 0x2c, 0x4b, 0x6a, 0xd8, 0xc8, 0xe6, 0x5e, 0xe5, 0x38,
	0xe7, 0xf7, 0x75, 0x8e, 0x6c, 0x81, 0xa7, 0x99, 0x10, 0xd9, 0x84, 0x46, 0x98, 0xa8, 0xa8, 0x2e,
	0xab, 0xea, 0xa6, 0x1f, 0x49, 0xaa, 0x44, 0x29, 0xc7, 0x54, 0x45, 0xe3, 0xcf, 0x98, 0x67, 0x34,
	0x55, 0x1a, 0xeb, 0x52, 0xa1, 0x42, 0x0a, 0x2d, 0xe0, 0x4e, 0x8d, 0x45, 0x98, 0x28, 0xe4, 0x69,
	0xe8, 0xa6, 0x8f, 0x3c, 0x6d, 0xf3, 0x60, 0x9e, 0x32, 0xe5, 0x65, 0xfe, 0x9b, 0x6a, 0x2a, 0x0a,
	0x2a, 0xb1, 0x66, 0x82, 0xd7, 0xfa, 0x9b, 0xaf, 0xee, 0x43, 0x76, 0x9e, 0xa9, 0x9e, 0x16, 0xd4,
	0x0a, 0x6c, 0x39, 0x81, 0x82, 0x45, 0x98, 0x73, 0xa1, 0x8d, 0xba, 0x8d, 0xbf, 0xf9, 0xc0, 0x76,
	0xcd, 0xd3, 0x55, 0x79, 0x1d, 0xdd, 0x4a, 0x5c, 0x14, 0x54, 0xda, 0xfe, 0xc3, 0x2f, 0x2d, 0xd0,
	0x3d, 0x32, 0x1e, 0x17, 0xc6, 0x02, 0x3e, 0x02, 0x3d, 0xef, 0xc2, 0x71, 0x4e, 0xc3, 0x60, 0x3b,
	0xd8, 0x6b, 0x0f, 0xba, 0xee, 0xcf, 0x53, 0x9c, 0x53, 0x78, 0x06, 0x36, 0x26, 0x58, 0xe9, 0xd4,
	0xa6, 0x23, 0x58, 0xd3, 0x54, 0xb3, 0x9c, 0x86, 0x8b, 0xdb, 0xc1, 0x5e, 0x67, 0x7f, 0xcb, 0x6e,
	0x0a, 0x39, 0x57, 0x74, 0xa1, 0x25, 0xe3, 0xd9, 0x10, 0x4f, 0x4a, 0x3a, 0x80, 0x15, 0xb5, 0xf6,
	0x7c, 0x83, 0x35, 0xbd, 0x64, 0x39, 0x85, 0xd3, 0x19, 0xd7, 0x6a, 0xb6, 0x70, 0x69, 0x3b, 0xd8,
	0x5b, 0xd9, 0xbf, 0x44, 0xf3, 0xb6, 0x6f, 0xb6, 0x83, 0x66, 0x93, 0x0f, 0x2c, 0xff, 0x72, 0x5a,
	0xd0, 0xb7, 0xbc, 0xcc, 0xe7, 0x36, 0xef, 0x66, 0xa9, 0x9e, 0xe0, 0x73, 0xd0, 0x1a, 0xe3, 0xbc,
	0xc0, 0x2c, 0xe3, 0xe1, 0x7f, 0x0d, 0xe2, 0x7b, 0x34, 0x7c, 0x06, 0x5a, 0x98, 0xa4, 0x99, 0x14,
	0x65, 0x11, 0x2e, 0x37, 0x60, 0xfe, 0x8f, 0xc9, 0x71, 0x05, 0x86, 0xb7, 0x60, 0xd5, 0x4f, 0x5b,
	0x9f, 0x6c, 0xd8, 0x32, 0xf3, 0x9e, 0xde, 0x63, 0xde, 0x33, 0xf7, 0x22, 0xfd, 0x31, 0xac, 0xef,
	0x0c, 0x56, 0x9c, 0x8d, 0x3d, 0xdc, 0x97, 0xa0, 0xe3, 0x12, 0xa7, 0x98, 0x84, 0xed, 0x06, 0xa1,
	0xdb, 0x36, 0x74, 0x42, 0xe0, 0x09, 0x80, 0x9e, 0x3d, 0x96, 0x4c, 0x53, 0xc9, 0x04, 0x0f, 0x41,
	0x03, 0x91, 0x35, 0x2b, 0x72, 0xe4, 0x58, 0xf0, 0x03, 0x80, 0x6e, 0x8f, 0x33, 0x5a, 0x9d, 0x06,
	0x5a, 0xeb, 0x8e, 0x77, 0x27, 0xf6, 0x18, 0x2c, 0x5d, 0x53, 0x4a, 0xc2, 0x6e, 0x03, 0xba, 0x41,
	0xc2, 0x17, 0xa0, 0x5d, 0xfd, 0xa6, 0x4c, 0xd3, 0x3c, 0xec, 0x35, 0x39, 0xf5, 0x0a, 0xfe, 0x5e,
	0xd3, 0x1c, 0xbe, 0x06, 0x3d, 0xbf, 0x05, 0xe3, 0xba, 0xd2, 0x80, 0xde, 0xb1, 0x0b, 0x78, 0x57,
	0x99, 0x27, 0xa0, 0xe7, 0x67, 0x37, 0x0a, 0xab, 0x0d, 0x14, 0xba, 0x8e, 0x62, 0x24, 0xce, 0xc0,
	0x86, 0x0f, 0x71, 0xc5, 0x48, 0x9a, 0x0b, 0xc2, 0xae, 0x19, 0x95, 0xe1, 0x5a, 0x93, 0x0f, 0xd0,
	0x86, 0x39, 0x64, 0xe4, 0xa3, 0xe5, 0x1d, 0xfe, 0x0c, 0xc0, 0xee, 0x58, 0xe4, 0xe8, 0x9f, 0xb7,
	0xdd, 0xe1, 0xfa, 0xec, 0xab, 0x76, 0x5e, 0xe9, 0x9f, 0x07, 0x9f, 0x4e, 0x2c, 0x2f, 0x13, 0x13,
	0xcc, 0x33, 0x24, 0x64, 0x16, 0x65, 0x94, 0x1b, 0x77, 0x77, 0xab, 0x15, 0x4c, 0xfd, 0xe5, 0xee,
	0x3d, 0xf0, 0xd5, 0xd7, 0x85, 0xc5, 0xe3, 0x24, 0xf9, 0xb6, 0xb0, 0x73, 0x5c, 0x4b, 0x26, 0x44,
	0xa1, 0xba, 0xac, 0xaa, 0x61, 0x1f, 0xb9, 0x6f, 0x59, 0x7d, 0x77, 0x98, 0x51, 0x42, 0xd4, 0xc8,
	0x63, 0x46, 0xc3, 0xfe, 0xc8, 0x63, 0x7e, 0x2c, 0xec, 0xd6, 0x8d, 0x38, 0x4e, 0x88, 0x8a, 0x63,
	0x8f, 0x8a, 0xe3, 0x61, 0x3f, 0x8e, 0x3d, 0xee, 0x6a, 0xd9, 0x84, 0x7d, 0xf2, 0x2b, 0x00, 0x00,
	0xff, 0xff, 0x28, 0xd9, 0xc6, 0x5e, 0x27, 0x06, 0x00, 0x00,
}

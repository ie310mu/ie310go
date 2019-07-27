// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/type/color.proto

package color

import (
	fmt "fmt"
	math "math"

	proto "github.com/ie310mu/ie310go/forks/github.com/golang/protobuf/proto"
	wrappers "github.com/ie310mu/ie310go/forks/github.com/golang/protobuf/ptypes/wrappers"
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

// Represents a color in the RGBA color space. This representation is designed
// for simplicity of conversion to/from color representations in various
// languages over compactness; for example, the fields of this representation
// can be trivially provided to the constructor of "java.awt.Color" in Java; it
// can also be trivially provided to UIColor's "+colorWithRed:green:blue:alpha"
// method in iOS; and, with just a little work, it can be easily formatted into
// a CSS "rgba()" string in JavaScript, as well.
//
// Note: this proto does not carry information about the absolute color space
// that should be used to interpret the RGB value (e.g. sRGB, Adobe RGB,
// DCI-P3, BT.2020, etc.). By default, applications SHOULD assume the sRGB color
// space.
//
// Example (Java):
//
//      import com.google.type.Color;
//
//      // ...
//      public static java.awt.Color fromProto(Color protocolor) {
//        float alpha = protocolor.hasAlpha()
//            ? protocolor.getAlpha().getValue()
//            : 1.0;
//
//        return new java.awt.Color(
//            protocolor.getRed(),
//            protocolor.getGreen(),
//            protocolor.getBlue(),
//            alpha);
//      }
//
//      public static Color toProto(java.awt.Color color) {
//        float red = (float) color.getRed();
//        float green = (float) color.getGreen();
//        float blue = (float) color.getBlue();
//        float denominator = 255.0;
//        Color.Builder resultBuilder =
//            Color
//                .newBuilder()
//                .setRed(red / denominator)
//                .setGreen(green / denominator)
//                .setBlue(blue / denominator);
//        int alpha = color.getAlpha();
//        if (alpha != 255) {
//          result.setAlpha(
//              FloatValue
//                  .newBuilder()
//                  .setValue(((float) alpha) / denominator)
//                  .build());
//        }
//        return resultBuilder.build();
//      }
//      // ...
//
// Example (iOS / Obj-C):
//
//      // ...
//      static UIColor* fromProto(Color* protocolor) {
//         float red = [protocolor red];
//         float green = [protocolor green];
//         float blue = [protocolor blue];
//         FloatValue* alpha_wrapper = [protocolor alpha];
//         float alpha = 1.0;
//         if (alpha_wrapper != nil) {
//           alpha = [alpha_wrapper value];
//         }
//         return [UIColor colorWithRed:red green:green blue:blue alpha:alpha];
//      }
//
//      static Color* toProto(UIColor* color) {
//          CGFloat red, green, blue, alpha;
//          if (![color getRed:&red green:&green blue:&blue alpha:&alpha]) {
//            return nil;
//          }
//          Color* result = [[Color alloc] init];
//          [result setRed:red];
//          [result setGreen:green];
//          [result setBlue:blue];
//          if (alpha <= 0.9999) {
//            [result setAlpha:floatWrapperWithValue(alpha)];
//          }
//          [result autorelease];
//          return result;
//     }
//     // ...
//
//  Example (JavaScript):
//
//     // ...
//
//     var protoToCssColor = function(rgb_color) {
//        var redFrac = rgb_color.red || 0.0;
//        var greenFrac = rgb_color.green || 0.0;
//        var blueFrac = rgb_color.blue || 0.0;
//        var red = Math.floor(redFrac * 255);
//        var green = Math.floor(greenFrac * 255);
//        var blue = Math.floor(blueFrac * 255);
//
//        if (!('alpha' in rgb_color)) {
//           return rgbToCssColor_(red, green, blue);
//        }
//
//        var alphaFrac = rgb_color.alpha.value || 0.0;
//        var rgbParams = [red, green, blue].join(',');
//        return ['rgba(', rgbParams, ',', alphaFrac, ')'].join('');
//     };
//
//     var rgbToCssColor_ = function(red, green, blue) {
//       var rgbNumber = new Number((red << 16) | (green << 8) | blue);
//       var hexString = rgbNumber.toString(16);
//       var missingZeros = 6 - hexString.length;
//       var resultBuilder = ['#'];
//       for (var i = 0; i < missingZeros; i++) {
//          resultBuilder.push('0');
//       }
//       resultBuilder.push(hexString);
//       return resultBuilder.join('');
//     };
//
//     // ...
type Color struct {
	// The amount of red in the color as a value in the interval [0, 1].
	Red float32 `protobuf:"fixed32,1,opt,name=red,proto3" json:"red,omitempty"`
	// The amount of green in the color as a value in the interval [0, 1].
	Green float32 `protobuf:"fixed32,2,opt,name=green,proto3" json:"green,omitempty"`
	// The amount of blue in the color as a value in the interval [0, 1].
	Blue float32 `protobuf:"fixed32,3,opt,name=blue,proto3" json:"blue,omitempty"`
	// The fraction of this color that should be applied to the pixel. That is,
	// the final pixel color is defined by the equation:
	//
	//   pixel color = alpha * (this color) + (1.0 - alpha) * (background color)
	//
	// This means that a value of 1.0 corresponds to a solid color, whereas
	// a value of 0.0 corresponds to a completely transparent color. This
	// uses a wrapper message rather than a simple float scalar so that it is
	// possible to distinguish between a default value and the value being unset.
	// If omitted, this color object is to be rendered as a solid color
	// (as if the alpha value had been explicitly given with a value of 1.0).
	Alpha                *wrappers.FloatValue `protobuf:"bytes,4,opt,name=alpha,proto3" json:"alpha,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Color) Reset()         { *m = Color{} }
func (m *Color) String() string { return proto.CompactTextString(m) }
func (*Color) ProtoMessage()    {}
func (*Color) Descriptor() ([]byte, []int) {
	return fileDescriptor_8454902cb8a2d34f, []int{0}
}

func (m *Color) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Color.Unmarshal(m, b)
}
func (m *Color) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Color.Marshal(b, m, deterministic)
}
func (m *Color) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Color.Merge(m, src)
}
func (m *Color) XXX_Size() int {
	return xxx_messageInfo_Color.Size(m)
}
func (m *Color) XXX_DiscardUnknown() {
	xxx_messageInfo_Color.DiscardUnknown(m)
}

var xxx_messageInfo_Color proto.InternalMessageInfo

func (m *Color) GetRed() float32 {
	if m != nil {
		return m.Red
	}
	return 0
}

func (m *Color) GetGreen() float32 {
	if m != nil {
		return m.Green
	}
	return 0
}

func (m *Color) GetBlue() float32 {
	if m != nil {
		return m.Blue
	}
	return 0
}

func (m *Color) GetAlpha() *wrappers.FloatValue {
	if m != nil {
		return m.Alpha
	}
	return nil
}

func init() {
	proto.RegisterType((*Color)(nil), "google.type.Color")
}

func init() { proto.RegisterFile("google/type/color.proto", fileDescriptor_8454902cb8a2d34f) }

var fileDescriptor_8454902cb8a2d34f = []byte{
	// 228 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0x31, 0x4b, 0xc5, 0x30,
	0x10, 0xc7, 0x49, 0xfb, 0xea, 0x70, 0x6f, 0x50, 0x82, 0x60, 0x50, 0x90, 0x87, 0xd3, 0x9b, 0x12,
	0x54, 0x70, 0x71, 0x7b, 0x82, 0xae, 0xa5, 0x88, 0x83, 0x93, 0x69, 0x3d, 0xa3, 0x10, 0x7b, 0x21,
	0x6d, 0x15, 0xbf, 0x8e, 0x9f, 0xd2, 0x51, 0x72, 0xe9, 0x83, 0x2e, 0xe1, 0x72, 0xbf, 0xdf, 0x25,
	0xff, 0x83, 0x13, 0x47, 0xe4, 0x3c, 0x9a, 0xf1, 0x27, 0xa0, 0xe9, 0xc8, 0x53, 0xd4, 0x21, 0xd2,
	0x48, 0x72, 0x9d, 0x81, 0x4e, 0xe0, 0xf4, 0x7c, 0xb6, 0x18, 0xb5, 0xd3, 0x9b, 0xf9, 0x8e, 0x36,
	0x04, 0x8c, 0x43, 0x96, 0x2f, 0xbe, 0xa0, 0xba, 0x4b, 0xb3, 0xf2, 0x08, 0xca, 0x88, 0xaf, 0x4a,
	0x6c, 0xc4, 0xb6, 0x68, 0x52, 0x29, 0x8f, 0xa1, 0x72, 0x11, 0xb1, 0x57, 0x05, 0xf7, 0xf2, 0x45,
	0x4a, 0x58, 0xb5, 0x7e, 0x42, 0x55, 0x72, 0x93, 0x6b, 0x79, 0x09, 0x95, 0xf5, 0xe1, 0xdd, 0xaa,
	0xd5, 0x46, 0x6c, 0xd7, 0x57, 0x67, 0x7a, 0x4e, 0xb0, 0xff, 0x54, 0xdf, 0x7b, 0xb2, 0xe3, 0x93,
	0xf5, 0x13, 0x36, 0xd9, 0xdc, 0xbd, 0xc0, 0x61, 0x47, 0x9f, 0x7a, 0x11, 0x75, 0x07, 0x1c, 0xa4,
	0x4e, 0x33, 0xb5, 0x78, 0xbe, 0x99, 0x91, 0x23, 0x6f, 0x7b, 0xa7, 0x29, 0x3a, 0xe3, 0xb0, 0xe7,
	0x17, 0x4d, 0x46, 0x36, 0x7c, 0x0c, 0x8b, 0xed, 0x6f, 0xf9, 0xfc, 0x13, 0xe2, 0xb7, 0x28, 0x1f,
	0x1e, 0xeb, 0xf6, 0x80, 0xdd, 0xeb, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0xdc, 0x38, 0x5a, 0x5f,
	0x28, 0x01, 0x00, 0x00,
}

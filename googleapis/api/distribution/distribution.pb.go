// Code generated by protoc-gen-go.
// source: google.golang.org/genproto/googleapis/api/distribution/distribution.proto
// DO NOT EDIT!

/*
Package google_api is a generated protocol buffer package.

It is generated from these files:
	google.golang.org/genproto/googleapis/api/distribution/distribution.proto

It has these top-level messages:
	Distribution
*/
package google_api // import "google.golang.org/genproto/googleapis/api/distribution"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/serviceconfig"
import _ "github.com/golang/protobuf/ptypes/any"
import _ "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Distribution contains summary statistics for a population of values and,
// optionally, a histogram representing the distribution of those values across
// a specified set of histogram buckets.
//
// The summary statistics are the count, mean, sum of the squared deviation from
// the mean, the minimum, and the maximum of the set of population of values.
//
// The histogram is based on a sequence of buckets and gives a count of values
// that fall into each bucket.  The boundaries of the buckets are given either
// explicitly or by specifying parameters for a method of computing them
// (buckets of fixed width or buckets of exponentially increasing width).
//
// Although it is not forbidden, it is generally a bad idea to include
// non-finite values (infinities or NaNs) in the population of values, as this
// will render the `mean` and `sum_of_squared_deviation` fields meaningless.
type Distribution struct {
	// The number of values in the population. Must be non-negative.
	Count int64 `protobuf:"varint,1,opt,name=count" json:"count,omitempty"`
	// The arithmetic mean of the values in the population. If `count` is zero
	// then this field must be zero.
	Mean float64 `protobuf:"fixed64,2,opt,name=mean" json:"mean,omitempty"`
	// The sum of squared deviations from the mean of the values in the
	// population.  For values x_i this is:
	//
	//     Sum[i=1..n]((x_i - mean)^2)
	//
	// Knuth, "The Art of Computer Programming", Vol. 2, page 323, 3rd edition
	// describes Welford's method for accumulating this sum in one pass.
	//
	// If `count` is zero then this field must be zero.
	SumOfSquaredDeviation float64 `protobuf:"fixed64,3,opt,name=sum_of_squared_deviation,json=sumOfSquaredDeviation" json:"sum_of_squared_deviation,omitempty"`
	// If specified, contains the range of the population values. The field
	// must not be present if the `count` is zero.
	Range *Distribution_Range `protobuf:"bytes,4,opt,name=range" json:"range,omitempty"`
	// Defines the histogram bucket boundaries.
	BucketOptions *Distribution_BucketOptions `protobuf:"bytes,6,opt,name=bucket_options,json=bucketOptions" json:"bucket_options,omitempty"`
	// If `bucket_options` is given, then the sum of the values in `bucket_counts`
	// must equal the value in `count`.  If `bucket_options` is not given, no
	// `bucket_counts` fields may be given.
	//
	// Bucket counts are given in order under the numbering scheme described
	// above (the underflow bucket has number 0; the finite buckets, if any,
	// have numbers 1 through N-2; the overflow bucket has number N-1).
	//
	// The size of `bucket_counts` must be no greater than N as defined in
	// `bucket_options`.
	//
	// Any suffix of trailing zero bucket_count fields may be omitted.
	BucketCounts []int64 `protobuf:"varint,7,rep,packed,name=bucket_counts,json=bucketCounts" json:"bucket_counts,omitempty"`
}

func (m *Distribution) Reset()                    { *m = Distribution{} }
func (m *Distribution) String() string            { return proto.CompactTextString(m) }
func (*Distribution) ProtoMessage()               {}
func (*Distribution) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Distribution) GetRange() *Distribution_Range {
	if m != nil {
		return m.Range
	}
	return nil
}

func (m *Distribution) GetBucketOptions() *Distribution_BucketOptions {
	if m != nil {
		return m.BucketOptions
	}
	return nil
}

// The range of the population values.
type Distribution_Range struct {
	// The minimum of the population values.
	Min float64 `protobuf:"fixed64,1,opt,name=min" json:"min,omitempty"`
	// The maximum of the population values.
	Max float64 `protobuf:"fixed64,2,opt,name=max" json:"max,omitempty"`
}

func (m *Distribution_Range) Reset()                    { *m = Distribution_Range{} }
func (m *Distribution_Range) String() string            { return proto.CompactTextString(m) }
func (*Distribution_Range) ProtoMessage()               {}
func (*Distribution_Range) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

// A Distribution may optionally contain a histogram of the values in the
// population.  The histogram is given in `bucket_counts` as counts of values
// that fall into one of a sequence of non-overlapping buckets.  The sequence
// of buckets is described by `bucket_options`.
//
// A bucket specifies an inclusive lower bound and exclusive upper bound for
// the values that are counted for that bucket.  The upper bound of a bucket
// is strictly greater than the lower bound.
//
// The sequence of N buckets for a Distribution consists of an underflow
// bucket (number 0), zero or more finite buckets (number 1 through N - 2) and
// an overflow bucket (number N - 1).  The buckets are contiguous:  the lower
// bound of bucket i (i > 0) is the same as the upper bound of bucket i - 1.
// The buckets span the whole range of finite values: lower bound of the
// underflow bucket is -infinity and the upper bound of the overflow bucket is
// +infinity.  The finite buckets are so-called because both bounds are
// finite.
//
// `BucketOptions` describes bucket boundaries in one of three ways.  Two
// describe the boundaries by giving parameters for a formula to generate
// boundaries and one gives the bucket boundaries explicitly.
//
// If `bucket_boundaries` is not given, then no `bucket_counts` may be given.
type Distribution_BucketOptions struct {
	// Exactly one of these three fields must be set.
	//
	// Types that are valid to be assigned to Options:
	//	*Distribution_BucketOptions_LinearBuckets
	//	*Distribution_BucketOptions_ExponentialBuckets
	//	*Distribution_BucketOptions_ExplicitBuckets
	Options isDistribution_BucketOptions_Options `protobuf_oneof:"options"`
}

func (m *Distribution_BucketOptions) Reset()                    { *m = Distribution_BucketOptions{} }
func (m *Distribution_BucketOptions) String() string            { return proto.CompactTextString(m) }
func (*Distribution_BucketOptions) ProtoMessage()               {}
func (*Distribution_BucketOptions) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 1} }

type isDistribution_BucketOptions_Options interface {
	isDistribution_BucketOptions_Options()
}

type Distribution_BucketOptions_LinearBuckets struct {
	LinearBuckets *Distribution_BucketOptions_Linear `protobuf:"bytes,1,opt,name=linear_buckets,json=linearBuckets,oneof"`
}
type Distribution_BucketOptions_ExponentialBuckets struct {
	ExponentialBuckets *Distribution_BucketOptions_Exponential `protobuf:"bytes,2,opt,name=exponential_buckets,json=exponentialBuckets,oneof"`
}
type Distribution_BucketOptions_ExplicitBuckets struct {
	ExplicitBuckets *Distribution_BucketOptions_Explicit `protobuf:"bytes,3,opt,name=explicit_buckets,json=explicitBuckets,oneof"`
}

func (*Distribution_BucketOptions_LinearBuckets) isDistribution_BucketOptions_Options()      {}
func (*Distribution_BucketOptions_ExponentialBuckets) isDistribution_BucketOptions_Options() {}
func (*Distribution_BucketOptions_ExplicitBuckets) isDistribution_BucketOptions_Options()    {}

func (m *Distribution_BucketOptions) GetOptions() isDistribution_BucketOptions_Options {
	if m != nil {
		return m.Options
	}
	return nil
}

func (m *Distribution_BucketOptions) GetLinearBuckets() *Distribution_BucketOptions_Linear {
	if x, ok := m.GetOptions().(*Distribution_BucketOptions_LinearBuckets); ok {
		return x.LinearBuckets
	}
	return nil
}

func (m *Distribution_BucketOptions) GetExponentialBuckets() *Distribution_BucketOptions_Exponential {
	if x, ok := m.GetOptions().(*Distribution_BucketOptions_ExponentialBuckets); ok {
		return x.ExponentialBuckets
	}
	return nil
}

func (m *Distribution_BucketOptions) GetExplicitBuckets() *Distribution_BucketOptions_Explicit {
	if x, ok := m.GetOptions().(*Distribution_BucketOptions_ExplicitBuckets); ok {
		return x.ExplicitBuckets
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Distribution_BucketOptions) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Distribution_BucketOptions_OneofMarshaler, _Distribution_BucketOptions_OneofUnmarshaler, _Distribution_BucketOptions_OneofSizer, []interface{}{
		(*Distribution_BucketOptions_LinearBuckets)(nil),
		(*Distribution_BucketOptions_ExponentialBuckets)(nil),
		(*Distribution_BucketOptions_ExplicitBuckets)(nil),
	}
}

func _Distribution_BucketOptions_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Distribution_BucketOptions)
	// options
	switch x := m.Options.(type) {
	case *Distribution_BucketOptions_LinearBuckets:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.LinearBuckets); err != nil {
			return err
		}
	case *Distribution_BucketOptions_ExponentialBuckets:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ExponentialBuckets); err != nil {
			return err
		}
	case *Distribution_BucketOptions_ExplicitBuckets:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ExplicitBuckets); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Distribution_BucketOptions.Options has unexpected type %T", x)
	}
	return nil
}

func _Distribution_BucketOptions_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Distribution_BucketOptions)
	switch tag {
	case 1: // options.linear_buckets
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Distribution_BucketOptions_Linear)
		err := b.DecodeMessage(msg)
		m.Options = &Distribution_BucketOptions_LinearBuckets{msg}
		return true, err
	case 2: // options.exponential_buckets
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Distribution_BucketOptions_Exponential)
		err := b.DecodeMessage(msg)
		m.Options = &Distribution_BucketOptions_ExponentialBuckets{msg}
		return true, err
	case 3: // options.explicit_buckets
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Distribution_BucketOptions_Explicit)
		err := b.DecodeMessage(msg)
		m.Options = &Distribution_BucketOptions_ExplicitBuckets{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Distribution_BucketOptions_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Distribution_BucketOptions)
	// options
	switch x := m.Options.(type) {
	case *Distribution_BucketOptions_LinearBuckets:
		s := proto.Size(x.LinearBuckets)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Distribution_BucketOptions_ExponentialBuckets:
		s := proto.Size(x.ExponentialBuckets)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Distribution_BucketOptions_ExplicitBuckets:
		s := proto.Size(x.ExplicitBuckets)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Specify a sequence of buckets that all have the same width (except
// overflow and underflow).  Each bucket represents a constant absolute
// uncertainty on the specific value in the bucket.
//
// Defines `num_finite_buckets + 2` (= N) buckets with these boundaries for
// bucket `i`:
//
//    Upper bound (0 <= i < N-1):     offset + (width * i).
//    Lower bound (1 <= i < N):       offset + (width * (i - 1)).
type Distribution_BucketOptions_Linear struct {
	// Must be greater than 0.
	NumFiniteBuckets int32 `protobuf:"varint,1,opt,name=num_finite_buckets,json=numFiniteBuckets" json:"num_finite_buckets,omitempty"`
	// Must be greater than 0.
	Width float64 `protobuf:"fixed64,2,opt,name=width" json:"width,omitempty"`
	// Lower bound of the first bucket.
	Offset float64 `protobuf:"fixed64,3,opt,name=offset" json:"offset,omitempty"`
}

func (m *Distribution_BucketOptions_Linear) Reset()         { *m = Distribution_BucketOptions_Linear{} }
func (m *Distribution_BucketOptions_Linear) String() string { return proto.CompactTextString(m) }
func (*Distribution_BucketOptions_Linear) ProtoMessage()    {}
func (*Distribution_BucketOptions_Linear) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 1, 0}
}

// Specify a sequence of buckets that have a width that is proportional to
// the value of the lower bound.  Each bucket represents a constant relative
// uncertainty on a specific value in the bucket.
//
// Defines `num_finite_buckets + 2` (= N) buckets with these boundaries for
// bucket i:
//
//    Upper bound (0 <= i < N-1):     scale * (growth_factor ^ i).
//    Lower bound (1 <= i < N):       scale * (growth_factor ^ (i - 1)).
type Distribution_BucketOptions_Exponential struct {
	// Must be greater than 0.
	NumFiniteBuckets int32 `protobuf:"varint,1,opt,name=num_finite_buckets,json=numFiniteBuckets" json:"num_finite_buckets,omitempty"`
	// Must be greater than 1.
	GrowthFactor float64 `protobuf:"fixed64,2,opt,name=growth_factor,json=growthFactor" json:"growth_factor,omitempty"`
	// Must be greater than 0.
	Scale float64 `protobuf:"fixed64,3,opt,name=scale" json:"scale,omitempty"`
}

func (m *Distribution_BucketOptions_Exponential) Reset() {
	*m = Distribution_BucketOptions_Exponential{}
}
func (m *Distribution_BucketOptions_Exponential) String() string { return proto.CompactTextString(m) }
func (*Distribution_BucketOptions_Exponential) ProtoMessage()    {}
func (*Distribution_BucketOptions_Exponential) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 1, 1}
}

// A set of buckets with arbitrary widths.
//
// Defines `size(bounds) + 1` (= N) buckets with these boundaries for
// bucket i:
//
//    Upper bound (0 <= i < N-1):     bounds[i]
//    Lower bound (1 <= i < N);       bounds[i - 1]
//
// There must be at least one element in `bounds`.  If `bounds` has only one
// element, there are no finite buckets, and that single element is the
// common boundary of the overflow and underflow buckets.
type Distribution_BucketOptions_Explicit struct {
	// The values must be monotonically increasing.
	Bounds []float64 `protobuf:"fixed64,1,rep,packed,name=bounds" json:"bounds,omitempty"`
}

func (m *Distribution_BucketOptions_Explicit) Reset()         { *m = Distribution_BucketOptions_Explicit{} }
func (m *Distribution_BucketOptions_Explicit) String() string { return proto.CompactTextString(m) }
func (*Distribution_BucketOptions_Explicit) ProtoMessage()    {}
func (*Distribution_BucketOptions_Explicit) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 1, 2}
}

func init() {
	proto.RegisterType((*Distribution)(nil), "google.api.Distribution")
	proto.RegisterType((*Distribution_Range)(nil), "google.api.Distribution.Range")
	proto.RegisterType((*Distribution_BucketOptions)(nil), "google.api.Distribution.BucketOptions")
	proto.RegisterType((*Distribution_BucketOptions_Linear)(nil), "google.api.Distribution.BucketOptions.Linear")
	proto.RegisterType((*Distribution_BucketOptions_Exponential)(nil), "google.api.Distribution.BucketOptions.Exponential")
	proto.RegisterType((*Distribution_BucketOptions_Explicit)(nil), "google.api.Distribution.BucketOptions.Explicit")
}

func init() {
	proto.RegisterFile("google.golang.org/genproto/googleapis/api/distribution/distribution.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 563 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x93, 0x5b, 0x8b, 0xd3, 0x40,
	0x14, 0xc7, 0x9b, 0x4d, 0x2f, 0x7a, 0x7a, 0xb1, 0x8e, 0xab, 0x84, 0x3c, 0x48, 0xd9, 0x45, 0x29,
	0xa8, 0x09, 0x54, 0xc1, 0x07, 0xdf, 0xea, 0xba, 0x54, 0x50, 0x76, 0x89, 0xe0, 0x93, 0x10, 0x26,
	0xc9, 0x24, 0x1d, 0x4c, 0x66, 0x62, 0x66, 0xd2, 0xed, 0xbe, 0xfb, 0xa5, 0xfc, 0x76, 0x92, 0x99,
	0xe9, 0x36, 0x7d, 0x10, 0xbb, 0x0f, 0x81, 0x73, 0xfd, 0xfd, 0xcf, 0x39, 0x64, 0xe0, 0x73, 0xc6,
	0x79, 0x96, 0x13, 0x2f, 0xe3, 0x39, 0x66, 0x99, 0xc7, 0xab, 0xcc, 0xcf, 0x08, 0x2b, 0x2b, 0x2e,
	0xb9, 0xaf, 0x53, 0xb8, 0xa4, 0xc2, 0xc7, 0x25, 0xf5, 0x13, 0x2a, 0x64, 0x45, 0xa3, 0x5a, 0x52,
	0xce, 0x0e, 0x1c, 0x4f, 0x95, 0x23, 0x30, 0x28, 0x5c, 0x52, 0xf7, 0x1e, 0x58, 0x41, 0xaa, 0x0d,
	0x8d, 0x49, 0xcc, 0x59, 0x4a, 0x33, 0x1f, 0x33, 0xc6, 0x25, 0x6e, 0xa8, 0x42, 0x63, 0x5d, 0x3f,
	0xa3, 0x72, 0x5d, 0x47, 0x5e, 0xcc, 0x0b, 0x5f, 0xe3, 0x7c, 0x95, 0x88, 0xea, 0xd4, 0x2f, 0xe5,
	0x6d, 0x49, 0x84, 0x8f, 0xd9, 0x6d, 0xf3, 0x99, 0x86, 0x0f, 0xff, 0x6f, 0x90, 0xb4, 0x20, 0x42,
	0xe2, 0xa2, 0xdc, 0x5b, 0xba, 0xf9, 0xec, 0xf7, 0x00, 0x46, 0x17, 0xad, 0xdd, 0xd0, 0x29, 0xf4,
	0x62, 0x5e, 0x33, 0xe9, 0x58, 0x33, 0x6b, 0x6e, 0x07, 0xda, 0x41, 0x08, 0xba, 0x05, 0xc1, 0xcc,
	0x39, 0x99, 0x59, 0x73, 0x2b, 0x50, 0x36, 0x7a, 0x0f, 0x8e, 0xa8, 0x8b, 0x90, 0xa7, 0xa1, 0xf8,
	0x55, 0xe3, 0x8a, 0x24, 0x61, 0x42, 0x36, 0x54, 0xed, 0xe2, 0xd8, 0xaa, 0xee, 0xa9, 0xa8, 0x8b,
	0xab, 0xf4, 0x9b, 0xce, 0x5e, 0xec, 0x92, 0xe8, 0x1d, 0xf4, 0x2a, 0xcc, 0x32, 0xe2, 0x74, 0x67,
	0xd6, 0x7c, 0xb8, 0x78, 0xee, 0xed, 0x0f, 0xe9, 0xb5, 0x67, 0xf1, 0x82, 0xa6, 0x2a, 0xd0, 0xc5,
	0xe8, 0x2b, 0x4c, 0xa2, 0x3a, 0xfe, 0x49, 0x64, 0xc8, 0x4b, 0x75, 0x2f, 0xa7, 0xaf, 0xda, 0x5f,
	0xfe, 0xb3, 0x7d, 0xa9, 0xca, 0xaf, 0x74, 0x75, 0x30, 0x8e, 0xda, 0x2e, 0x3a, 0x07, 0x13, 0x08,
	0xd5, 0x86, 0xc2, 0x19, 0xcc, 0xec, 0xb9, 0x1d, 0x8c, 0x74, 0xf0, 0xa3, 0x8a, 0xb9, 0xaf, 0xa0,
	0xa7, 0x66, 0x40, 0x53, 0xb0, 0x0b, 0xca, 0xd4, 0x4d, 0xac, 0xa0, 0x31, 0x55, 0x04, 0x6f, 0xcd,
	0x41, 0x1a, 0xd3, 0xfd, 0xd3, 0x85, 0xf1, 0x81, 0x24, 0xfa, 0x0e, 0x93, 0x9c, 0x32, 0x82, 0xab,
	0x50, 0x53, 0x85, 0x02, 0x0c, 0x17, 0x6f, 0x8e, 0x1b, 0xd9, 0xfb, 0xa2, 0x9a, 0x57, 0x9d, 0x60,
	0xac, 0x31, 0x3a, 0x2b, 0x10, 0x81, 0x27, 0x64, 0x5b, 0x72, 0x46, 0x98, 0xa4, 0x38, 0xbf, 0x83,
	0x9f, 0x28, 0xf8, 0xe2, 0x48, 0xf8, 0xa7, 0x3d, 0x61, 0xd5, 0x09, 0x50, 0x0b, 0xb8, 0x93, 0xf9,
	0x01, 0x53, 0xb2, 0x2d, 0x73, 0x1a, 0x53, 0x79, 0xa7, 0x61, 0x2b, 0x0d, 0xff, 0x78, 0x0d, 0xd5,
	0xbe, 0xea, 0x04, 0x8f, 0x76, 0x28, 0x43, 0x77, 0x13, 0xe8, 0xeb, 0xfd, 0xd0, 0x6b, 0x40, 0xac,
	0x2e, 0xc2, 0x94, 0x32, 0x2a, 0xc9, 0xc1, 0xa9, 0x7a, 0xc1, 0x94, 0xd5, 0xc5, 0xa5, 0x4a, 0xec,
	0xa6, 0x3a, 0x85, 0xde, 0x0d, 0x4d, 0xe4, 0xda, 0x9c, 0x5e, 0x3b, 0xe8, 0x19, 0xf4, 0x79, 0x9a,
	0x0a, 0x22, 0xcd, 0xaf, 0x67, 0x3c, 0x77, 0x03, 0xc3, 0xd6, 0xa2, 0xf7, 0x94, 0x3a, 0x87, 0x71,
	0x56, 0xf1, 0x1b, 0xb9, 0x0e, 0x53, 0x1c, 0x4b, 0x5e, 0x19, 0xc9, 0x91, 0x0e, 0x5e, 0xaa, 0x58,
	0x33, 0x8f, 0x88, 0x71, 0x4e, 0x8c, 0xb0, 0x76, 0xdc, 0x33, 0x78, 0xb0, 0x5b, 0xbe, 0x99, 0x2d,
	0xe2, 0x35, 0x4b, 0x1a, 0x21, 0xbb, 0x99, 0x4d, 0x7b, 0xcb, 0x87, 0x30, 0x30, 0xbf, 0xf2, 0xf2,
	0x05, 0x4c, 0x62, 0x5e, 0xb4, 0xae, 0xba, 0x7c, 0xdc, 0x3e, 0xeb, 0x75, 0xf3, 0x56, 0xaf, 0xad,
	0xa8, 0xaf, 0x1e, 0xed, 0xdb, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x87, 0xa9, 0x04, 0x41, 0xc6,
	0x04, 0x00, 0x00,
}
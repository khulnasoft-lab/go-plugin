// Protocol Buffers - Google's data interchange format
// Copyright 2008 Google Inc.  All rights reserved.
// Copyright 2022 Md Sulaiman.  All rights reserved.
// https://developers.google.com/protocol-buffers/

// Code generated by protoc-gen-go-plugin. DO NOT EDIT.
// versions:
// 	protoc-gen-go-plugin v0.1.0
// 	protoc               v4.25.2
// source: types/known/durationpb/duration.proto

package durationpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A Duration represents a signed, fixed-length span of time represented
// as a count of seconds and fractions of seconds at nanosecond
// resolution. It is independent of any calendar and concepts like "day"
// or "month". It is related to Timestamp in that the difference between
// two Timestamp values is a Duration and it can be added or subtracted
// from a Timestamp. Range is approximately +-10,000 years.
//
// # Examples
//
// Example 1: Compute Duration from two Timestamps in pseudo code.
//
//	Timestamp start = ...;
//	Timestamp end = ...;
//	Duration duration = ...;
//
//	duration.seconds = end.seconds - start.seconds;
//	duration.nanos = end.nanos - start.nanos;
//
//	if (duration.seconds < 0 && duration.nanos > 0) {
//	  duration.seconds += 1;
//	  duration.nanos -= 1000000000;
//	} else if (duration.seconds > 0 && duration.nanos < 0) {
//	  duration.seconds -= 1;
//	  duration.nanos += 1000000000;
//	}
//
// Example 2: Compute Timestamp from Timestamp + Duration in pseudo code.
//
//	Timestamp start = ...;
//	Duration duration = ...;
//	Timestamp end = ...;
//
//	end.seconds = start.seconds + duration.seconds;
//	end.nanos = start.nanos + duration.nanos;
//
//	if (end.nanos < 0) {
//	  end.seconds -= 1;
//	  end.nanos += 1000000000;
//	} else if (end.nanos >= 1000000000) {
//	  end.seconds += 1;
//	  end.nanos -= 1000000000;
//	}
//
// Example 3: Compute Duration from datetime.timedelta in Python.
//
//	td = datetime.timedelta(days=3, minutes=10)
//	duration = Duration()
//	duration.FromTimedelta(td)
//
// # JSON Mapping
//
// In JSON format, the Duration type is encoded as a string rather than an
// object, where the string ends in the suffix "s" (indicating seconds) and
// is preceded by the number of seconds, with nanoseconds expressed as
// fractional seconds. For example, 3 seconds with 0 nanoseconds should be
// encoded in JSON format as "3s", while 3 seconds and 1 nanosecond should
// be expressed in JSON format as "3.000000001s", and 3 seconds and 1
// microsecond should be expressed in JSON format as "3.000001s".
type Duration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Signed seconds of the span of time. Must be from -315,576,000,000
	// to +315,576,000,000 inclusive. Note: these bounds are computed from:
	// 60 sec/min * 60 min/hr * 24 hr/day * 365.25 days/year * 10000 years
	Seconds int64 `protobuf:"varint,1,opt,name=seconds,proto3" json:"seconds,omitempty"`
	// Signed fractions of a second at nanosecond resolution of the span
	// of time. Durations less than one second are represented with a 0
	// `seconds` field and a positive or negative `nanos` field. For durations
	// of one second or more, a non-zero value for the `nanos` field must be
	// of the same sign as the `seconds` field. Must be from -999,999,999
	// to +999,999,999 inclusive.
	Nanos int32 `protobuf:"varint,2,opt,name=nanos,proto3" json:"nanos,omitempty"`
}

func (x *Duration) ProtoReflect() protoreflect.Message {
	panic(`not implemented`)
}

func (x *Duration) GetSeconds() int64 {
	if x != nil {
		return x.Seconds
	}
	return 0
}

func (x *Duration) GetNanos() int32 {
	if x != nil {
		return x.Nanos
	}
	return 0
}

// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package go_testpkg is an autogenerated binder stub for package testpkg.
//   gobind -lang=go golang.org/x/mobile/bind/objc/testpkg
//
// File is generated by gobind. Do not edit.
package go_testpkg

import (
	"golang.org/x/mobile/bind/objc/testpkg"
	"golang.org/x/mobile/bind/seq"
)

func proxy_BytesAppend(out, in *seq.Buffer) {
	param_a := in.ReadByteArray()
	param_b := in.ReadByteArray()
	res := testpkg.BytesAppend(param_a, param_b)
	out.WriteByteArray(res)
}

func proxy_Hello(out, in *seq.Buffer) {
	param_s := in.ReadString()
	res := testpkg.Hello(param_s)
	out.WriteString(res)
}

func proxy_Hi(out, in *seq.Buffer) {
	testpkg.Hi()
}

func proxy_Int(out, in *seq.Buffer) {
	param_x := in.ReadInt32()
	testpkg.Int(param_x)
}

func proxy_NewS(out, in *seq.Buffer) {
	param_x := in.ReadFloat64()
	param_y := in.ReadFloat64()
	res := testpkg.NewS(param_x, param_y)
	out.WriteGoRef(res)
}

const (
	proxyS_Descriptor = "go.testpkg.S"
	proxyS_X_Get_Code = 0x00f
	proxyS_X_Set_Code = 0x01f
	proxyS_Y_Get_Code = 0x10f
	proxyS_Y_Set_Code = 0x11f
	proxyS_Sum_Code   = 0x00c
)

type proxyS seq.Ref

func proxyS_X_Set(out, in *seq.Buffer) {
	ref := in.ReadRef()
	v := in.ReadFloat64()
	ref.Get().(*testpkg.S).X = v
}

func proxyS_X_Get(out, in *seq.Buffer) {
	ref := in.ReadRef()
	v := ref.Get().(*testpkg.S).X
	out.WriteFloat64(v)
}

func proxyS_Y_Set(out, in *seq.Buffer) {
	ref := in.ReadRef()
	v := in.ReadFloat64()
	ref.Get().(*testpkg.S).Y = v
}

func proxyS_Y_Get(out, in *seq.Buffer) {
	ref := in.ReadRef()
	v := ref.Get().(*testpkg.S).Y
	out.WriteFloat64(v)
}

func proxyS_Sum(out, in *seq.Buffer) {
	ref := in.ReadRef()
	v := ref.Get().(*testpkg.S)
	res := v.Sum()
	out.WriteFloat64(res)
}

func init() {
	seq.Register(proxyS_Descriptor, proxyS_X_Set_Code, proxyS_X_Set)
	seq.Register(proxyS_Descriptor, proxyS_X_Get_Code, proxyS_X_Get)
	seq.Register(proxyS_Descriptor, proxyS_Y_Set_Code, proxyS_Y_Set)
	seq.Register(proxyS_Descriptor, proxyS_Y_Get_Code, proxyS_Y_Get)
	seq.Register(proxyS_Descriptor, proxyS_Sum_Code, proxyS_Sum)
}

func proxy_Sum(out, in *seq.Buffer) {
	param_x := in.ReadInt64()
	param_y := in.ReadInt64()
	res := testpkg.Sum(param_x, param_y)
	out.WriteInt64(res)
}

func init() {
	seq.Register("testpkg", 1, proxy_BytesAppend)
	seq.Register("testpkg", 2, proxy_Hello)
	seq.Register("testpkg", 3, proxy_Hi)
	seq.Register("testpkg", 4, proxy_Int)
	seq.Register("testpkg", 5, proxy_NewS)
	seq.Register("testpkg", 6, proxy_Sum)
}

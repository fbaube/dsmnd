package dsmnd

import (
	"go/types"
)

// See https://pkg.go.dev/go/types#BasicKind:
//
//   - type BasicKind int
//   - const (
//   - Invalid BasicKind = iota // type is invalid
//   - predeclared types
//   - Bool Int Int8 Int16 Int32 Int64 Uint Uint8 Uint16 Uint32 Uint64
//   - Uintptr Float32 Float64 Complex64 Complex128 String UnsafePointer
//   - types for untyped values
//   - UntypedBool UntypedInt UntypedRune UntypedFloat
//   - UntypedComplex UntypedString UntypedNil
//   - aliases
//   - Byte = Uint8
//   - Rune = Int32
//
// .
type GoBasicKind types.BasicKind

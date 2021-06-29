package e2e_test

import (
	"github.com/youta1119/golang-gen-getter/e2e_test/other"
)

//go:generate go run github.com/youta1119/golang-gen-getter/cmd/gen-getter -type=Structure
type Structure struct {
	id             string
	externalStruct *other.Other
	childStruct    ChildStructure
	childStructSlice       []*ChildStructure
}

func NewStructure(id string, externalStruct *other.Other, childStruct ChildStructure, childStructSlice []*ChildStructure) *Structure {
	return &Structure{id: id, externalStruct: externalStruct, childStruct: childStruct, childStructSlice: childStructSlice}
}

type ChildStructure struct {
	Foo string
}

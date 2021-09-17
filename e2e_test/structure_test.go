package e2e_test_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/youta1119/golang-gen-getter/e2e_test"
	"github.com/youta1119/golang-gen-getter/e2e_test/other"
	"testing"
)

func TestStructureGetter(t *testing.T) {
	s := e2e_test.NewStructure(
		"id",
		1,
		&other.Other{
			Foo: "other",
			Bar: "bar",
		},
		e2e_test.ChildStructure{
			Foo: "foo_child",

		},
		[]*e2e_test.ChildStructure{
			{
				Foo: "1",
			},
			{
				Foo: "2",
			},
		},
	)

	assert.EqualValues(t, "id", s.ID())
	assert.Equal(t, &other.Other{
		Foo: "other",
		Bar: "bar",
	}, s.ExternalStruct())
	assert.Equal(t, e2e_test.ChildStructure{
		Foo: "foo_child",
	}, s.ChildStruct())
	assert.Equal(t,[]*e2e_test.ChildStructure{
		{
			Foo: "1",
		},
		{
			Foo: "2",
		},
	}, s.ChildStructSlice())
}

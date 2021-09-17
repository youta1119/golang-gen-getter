package generator

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestGetterGenerator_Generate(t *testing.T) {
	type fields struct {
		PkgName      string
		TypeName     string
		Fields       []*Field
		GetterPrefix bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "should successful generate getter(GetterPrefix=true)",
			fields: fields{
				PkgName:  "example",
				TypeName: "Example",
				Fields: []*Field{
					{
						TypeName: "string",
						Name:     "foo",
					},
				},
				GetterPrefix: true,
			},
			want: strings.TrimSpace(`
package example

func (e *Example) GetFoo() string {
	return e.foo
}
`) + "\n",
		},

		{
			name: "should successful generate getter(GetterPrefix=false)",
			fields: fields{
				PkgName:  "example",
				TypeName: "Example",
				Fields: []*Field{
					{
						TypeName: "string",
						Name:     "foo",
					},
				},
				GetterPrefix: false,
			},
			want: strings.TrimSpace(`
package example

func (e *Example) Foo() string {
	return e.foo
}
`) + "\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GetterGenerator{
				PkgName:      tt.fields.PkgName,
				TypeName:     tt.fields.TypeName,
				Fields:       tt.fields.Fields,
				GetterPrefix: tt.fields.GetterPrefix,
			}
			got, err := g.Generate()
			require.NoError(t, err)
			assert.Equal(t, tt.want, string(got))
		})
	}
}

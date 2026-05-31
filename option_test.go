package opt_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/ulikunitz/opt"
)

func TestOption(t *testing.T) {
	type s struct {
		Name string
		A    int
	}
	type ts struct {
		Name  string           `json:",omitempty"`
		Size1 opt.Value[int]   `json:",omitzero"`
		Size2 opt.Value[int64] `json:",omitzero"`
		Size3 opt.Value[int8]
		S     opt.Value[s] `json:",omitzero"`
	}

	tests := []ts{
		{
			Name:  "1",
			Size1: opt.Val(10),
		},
		{
			Name: "2",
			S:    opt.Val(s{Name: "s", A: 1}),
		},
		{},
	}

	for _, tc := range tests {
		v := tc
		data, err := json.MarshalIndent(v, "", "  ")
		if err != nil {
			t.Fatalf("json.Marshal failed: %v", err)
		}

		t.Logf("\n%s", data)

		var v2 ts
		err = json.Unmarshal(data, &v2)
		if err != nil {
			t.Fatalf("json.Unmarshal failed: %v", err)
		}
		if v != v2 {
			t.Fatalf(
				"unmarshaled value does not match original: got %+v, want %+v",
				v2, v)
		}
	}
}

func Example() {
	type S struct {
		Name string
		A    opt.Value[int]     `json:",omitzero"`
		B    opt.Value[float64] `json:",omitzero"`
	}

	s := S{
		Name: "example",
		A:    opt.Val(42),
	}

	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", data)

	// Output:
	// {
	//   "Name": "example",
	//   "A": 42
	// }
}

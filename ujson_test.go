package ujson

import (
	"testing"
	"encoding/json"
	"log"
	"reflect"
)

type unmarshalTest struct {
	in  string
	ptr interface{}
	out interface{}
	err error
}


// testcases taken from encoding/json
var unmarshalTests = []unmarshalTest{
	// basic types
	{`true`, new(bool), true, nil},
	{`1`, new(int), 1, nil},
	{`1.2`, new(float64), 1.2, nil},
	{`-5`, new(int32), int32(-5), nil},
	{`"a\u1234"`, new(string), "a\u1234", nil},
	{`"http:\/\/"`, new(string), "http://", nil},
	{`"g-clef: \uD834\uDD1E"`, new(string), "g-clef: \U0001D11E", nil},
	{`"invalid: \uD834x\uDD1E"`, new(string), "invalid: \uFFFDx\uFFFD", nil},
	{"null", new(interface{}), nil, nil},
	// {`{"X": [1,2,3], "Y": 4}`, new(T), T{Y: 4}, &UnmarshalTypeError{"array", reflect.TypeOf("")}},
	// {`{"x": 1}`, new(tx), tx{}, &UnmarshalFieldError{"x", txType, txType.Field(0)}},

	// Z has a "-" tag.
	// {`{"Y": 1, "Z": 2}`, new(T), T{Y: 1}, nil},

	// syntax errors
	{`{"X": "foo", "Y"}`, nil, nil, &SyntaxError{"invalid character '}' after object key", 17}},
	{`[1, 2, 3+]`, nil, nil, &SyntaxError{"invalid character '+' after array element", 9}},

	// array tests
	{`[1, 2, 3]`, new([3]int), [3]int{1, 2, 3}, nil},
	{`[1, 2, 3]`, new([1]int), [1]int{1}, nil},
	{`[1, 2, 3]`, new([5]int), [5]int{1, 2, 3, 0, 0}, nil},

	// composite tests
	// {allValueIndent, new(All), allValue, nil},
	// {allValueCompact, new(All), allValue, nil},
	// {allValueIndent, new(*All), &allValue, nil},
	// {allValueCompact, new(*All), &allValue, nil},
	// {pallValueIndent, new(All), pallValue, nil},
	// {pallValueCompact, new(All), pallValue, nil},
	// {pallValueIndent, new(*All), &pallValue, nil},
	// {pallValueCompact, new(*All), &pallValue, nil},

	// unmarshal interface test
	// {`{"T":false}`, &um0, umtrue, nil}, // use "false" so test will fail if custom unmarshaler is not called
	// {`{"T":false}`, &ump, &umtrue, nil},
	// {`[{"T":false}]`, &umslice, umslice, nil},
	// {`[{"T":false}]`, &umslicep, &umslice, nil},
	// {`{"M":{"T":false}}`, &umstruct, umstruct, nil},
}

func TestUnmarshal(t *testing.T) {
	for i, tt := range unmarshalTests {
		in := []byte(tt.in)
		log.Printf("#%d: input %s", i, tt.in)
		// if err := checkValid(in, &scan); err != nil {
		// 	if !reflect.DeepEqual(err, tt.err) {
		// 		t.Errorf("#%d: checkValid: %#v", i, err)
		// 		continue
		// 	}
		// }
		// if tt.ptr == nil {
		// 	continue
		// }

		// v = new(right-type)
		// v := reflect.New(reflect.TypeOf(tt.ptr).Elem())
		var vv interface{}
		var err error
		if vv, err = Unmarshal([]byte(in)); !reflect.DeepEqual(err, tt.err) {
			t.Errorf("#%d: have error %v want %v", i, err, tt.err)
			continue
		}
		if !reflect.DeepEqual(vv, tt.out) {
			t.Errorf("#%d: mismatch %T and %T\nhave: %#+v\nwant: %#+v", i, vv, tt.out, vv, tt.out)
			data, _ := json.Marshal(vv)
			println(string(data))
			data, _ = json.Marshal(tt.out)
			println(string(data))
			continue
		}
	}
}

// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Large data benchmark.
// The JSON data is a summary of agl's changes in the
// go, webkit, and chromium open source projects.
// We benchmark converting between the JSON form
// and in-memory data structures.

package ujson

import (
	"compress/gzip"
	"io/ioutil"
	"os"
	"testing"
)

var codeJSON []byte

func codeInit() {
	f, err := os.Open("testdata/code.json.gz")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	gz, err := gzip.NewReader(f)
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(gz)
	if err != nil {
		panic(err)
	}

	codeJSON = data

	if _, err = Unmarshal(codeJSON); err != nil {
		panic("unmarshal code.json: " + err.Error())
	}
	// 
	// if data, err = Marshal(&codeStruct); err != nil {
	// 	panic("marshal code.json: " + err.Error())
	// }
	// 
	// if !bytes.Equal(data, codeJSON) {
	// 	println("different lengths", len(data), len(codeJSON))
	// 	for i := 0; i < len(data) && i < len(codeJSON); i++ {
	// 		if data[i] != codeJSON[i] {
	// 			println("re-marshal: changed at byte", i)
	// 			println("orig: ", string(codeJSON[i-10:i+10]))
	// 			println("new: ", string(data[i-10:i+10]))
	// 			break
	// 		}
	// 	}
	// 	panic("re-marshal code.json: different result")
	// }
}

func BenchmarkUjson(b *testing.B) {
	b.StopTimer()
	f, err := os.Open("testdata/small.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	data, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Unmarshal(data)
	}
	b.SetBytes(int64(len(data)))
}


func BenchmarkUjsonCodeUnmarshal(b *testing.B) {
	if codeJSON == nil {
		b.StopTimer()
		codeInit()
		b.StartTimer()
	}
	for i := 0; i < b.N; i++ {
		var err error
		if _, err = Unmarshal(codeJSON); err != nil {
			b.Fatal("Unmmarshal:", err)
		}
	}
	b.SetBytes(int64(len(codeJSON)))
}

// func BenchmarkCodeUnmarshalReuse(b *testing.B) {
// 	if codeJSON == nil {
// 		b.StopTimer()
// 		codeInit()
// 		b.StartTimer()
// 	}
// 	for i := 0; i < b.N; i++ {
// 		if _, err := Unmarshal(codeJSON); err != nil {
// 			b.Fatal("Unmmarshal:", err)
// 		}
// 	}
// 	b.SetBytes(int64(len(codeJSON)))
// }

// Copyright 2016 Eryx <evorui аt gmаil dοt cοm>, All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hflag

import (
	"os"
	"strconv"
	"strings"
)

const (
	Version = "0.9.0"
)

var (
	args = map[string]Bytex{}
)

func init() {

	if len(os.Args) < 2 {
		return
	}

	for i, k := range os.Args {

		if k[0] != '-' || len(k) < 2 {
			continue
		}

		k = strings.Trim(k, "-")

		if n := strings.Index(k, "="); n > 0 {
			if n+1 < len(k) {
				args[k[:n]] = Bytex(k[n+1:])
			} else {
				args[k[:n]] = Bytex("")
			}
			continue
		}

		if len(os.Args) <= (i+1) || os.Args[i+1][0] == '-' {
			args[k] = Bytex([]byte(""))
			continue
		}

		v := os.Args[i+1]

		args[k] = Bytex([]byte(v))
	}
}

func ValueOK(key string) (Bytex, bool) {

	if v, ok := args[key]; ok {
		return v, ok
	}

	return nil, false
}

func Value(key string) Bytex {

	if v, ok := ValueOK(key); ok {
		return v
	}

	return Bytex{}
}

func Each(fn func(key, val string)) {
	for k, v := range args {
		fn(k, v.String())
	}
}

// Universal Bytes
type Bytex []byte

// Bytes converts the value-bytes to bytes
func (bx Bytex) Bytes() []byte {
	return bx
}

// String converts the value-bytes to string
func (bx Bytex) String() string {
	return string(bx)
}

// Bool converts the value-bytes to bool
func (bx Bytex) Bool() bool {
	if len(bx) > 0 {
		if b, err := strconv.ParseBool(string(bx)); err == nil {
			return b
		}
	}
	return false
}

// Int64 converts the value-bytes to int64
func (bx Bytex) Int64() int64 {
	if len(bx) > 0 {
		if i64, err := strconv.ParseInt(string(bx), 10, 64); err == nil {
			return i64
		}
	}
	return 0
}

// Uint64 converts the value-bytes to uint64
func (bx Bytex) Uint64() uint64 {
	if len(bx) > 0 {
		if i64, err := strconv.ParseUint(string(bx), 10, 64); err == nil {
			return i64
		}
	}
	return 0
}

// Float64 converts the value-bytes to float64
func (bx Bytex) Float64() float64 {
	if len(bx) > 0 {
		if f64, err := strconv.ParseFloat(string(bx), 64); err == nil {
			return f64
		}
	}
	return 0
}

// Copyright (c) 2012 The 'objc' Package Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package objc

import "testing"

type SomeGoObject struct {
	Object `objc:"SomeGoObject : NSObject"`
}

func (sgo *SomeGoObject) Dealloc() {
}

func (sgo *SomeGoObject) FiftyFive() int64 {
	return 55
}

func (sgo *SomeGoObject) ThirtyTwo() int64 {
	return 32
}

func (sgo *SomeGoObject) Sum() int64 {
	return sgo.Send("thirtyTwo").Int() + sgo.Send("fiftyFive").Int()
}

func (sgo *SomeGoObject) GoSum() int64 {
	return sgo.ThirtyTwo() + sgo.FiftyFive()
}

func TestGoObjectCallObjC(t *testing.T) {
	c := NewClassFromStruct(SomeGoObject{})
	c.AddMethod("fiftyFive", (*SomeGoObject).FiftyFive)
	c.AddMethod("thirtyTwo", (*SomeGoObject).ThirtyTwo)
	c.AddMethod("goSum", (*SomeGoObject).GoSum)
	c.AddMethod("sum", (*SomeGoObject).Sum)
	RegisterClass(c)

	sgo := GetClass("SomeGoObject").Send("alloc").Send("init")
	if sgo.Send("goSum").Int() != sgo.Send("sum").Int() {
		t.Errorf("calculated sums do not match")
	}
}

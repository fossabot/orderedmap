// Copyright 2018, Goomba project Authors. All rights reserved.
//
// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with this
// work for additional information regarding copyright ownership.  The ASF
// licenses this file to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.  See the
// License for the specific language governing permissions and limitations
// under the License.

package orderedmap_test

import (
	"testing"

	"github.com/goombaio/orderedmap"
)

func TestOrderedMap_Put(t *testing.T) {
	m := orderedmap.NewOrderedMap()
	m.Put(5, "e")
	m.Put(6, "f")
	m.Put(7, "g")
	m.Put(3, "c")
	m.Put(4, "d")
	m.Put(1, "x")
	m.Put(2, "b")
	m.Put(1, "a") //overwrite
	m.Put(2, "b")
	structKey := complexType{"skey"}
	structValue := complexType{"svalue"}
	m.Put(structKey, structValue)
	m.Put(&structKey, &structValue)
	m.Put(true, false)

	table := []struct {
		key           interface{}
		expectedValue interface{}
		expectedFound bool
	}{
		{1, "a", true},
		{2, "b", true},
		{3, "c", true},
		{4, "d", true},
		{5, "e", true},
		{6, "f", true},
		{7, "g", true},
		{8, nil, false},
		{structKey, structValue, true},
		{&structKey, &structValue, true},
		{true, false, true},
	}

	for _, test := range table {
		actualValue, actualFound := m.Get(test.key)
		if actualValue != test.expectedValue || actualFound != test.expectedFound {
			t.Errorf("Got %v expected %v", actualValue, test.expectedValue)
		}
	}
}

func TestOrderedMap_Get(t *testing.T) {
	TestOrderedMap_Put(t)
}

func TestOrderedMap_Remove(t *testing.T) {
	m := orderedmap.NewOrderedMap()
	m.Put("bar", "foo")
	m.Put("foo", "bar")

	actualValue, actualFound := m.Get("foo")
	if actualValue != "bar" || !actualFound {
		t.Errorf("Got %v:%v expected %v:%v", actualValue, actualFound, "bar", true)
	}

	m.Remove("foo")
	actualValue, actualFound = m.Get("foo")
	if actualValue != nil || actualFound {
		t.Errorf("Got %v:%v expected %v:%v", actualValue, actualFound, nil, false)
	}

	m.Remove("foo") // already removed
	actualValue, actualFound = m.Get("foo")
	if actualValue != nil || actualFound {
		t.Errorf("Got %v:%v expected %v:%v", actualValue, actualFound, nil, false)
	}
}

func TestOrderedMap_Empty(t *testing.T) {
	m := orderedmap.NewOrderedMap()
	actualValue := m.Empty()
	if actualValue == false {
		t.Errorf("Got %v expected %v", actualValue, true)
	}

	m.Put("foo", "bar")
	actualValue = m.Empty()
	if actualValue == true {
		t.Errorf("Got %v expected %v", actualValue, false)
	}
}

func TestOrderedMap_Size(t *testing.T) {
	m := orderedmap.NewOrderedMap()
	m.Put(5, "e")
	m.Put(6, "f")
	m.Put(7, "g")
	m.Put(3, "c")
	m.Put(4, "d")
	m.Put(1, "x")
	m.Put(2, "b")
	m.Put(1, "a") //overwrite
	m.Put(2, "b")
	structKey := complexType{"skey"}
	structValue := complexType{"svalue"}
	m.Put(structKey, structValue)
	m.Put(&structKey, &structValue)
	m.Put(true, false)

	if actualSize := m.Size(); actualSize != 10 {
		t.Errorf("Got %v expected %v", actualSize, 10)
	}
}

func TestOrderedMap_Keys(t *testing.T) {
	m := orderedmap.NewOrderedMap()
	m.Put(5, "e")
	m.Put(6, "f")
	m.Put(7, "g")
	m.Put(3, "c")
	m.Put(4, "d")
	m.Put(1, "x")
	m.Put(2, "b")
	m.Put(1, "a") //overwrite
	m.Put(2, "b")
	structKey := complexType{"skey"}
	structValue := complexType{"svalue"}
	m.Put(structKey, structValue)
	m.Put(&structKey, &structValue)
	m.Put(true, false)

	actualKeys := m.Keys()
	expectedKeys := []interface{}{5, 6, 7, 3, 4, 1, 2, structKey, &structKey, true}
	if !sameElements(actualKeys, expectedKeys) {
		t.Errorf("Got %v expected %v", actualKeys, expectedKeys)
	}
}

func TestOrderedMap_Values(t *testing.T) {
	m := orderedmap.NewOrderedMap()
	m.Put(5, "e")
	m.Put(6, "f")
	m.Put(7, "g")
	m.Put(3, "c")
	m.Put(4, "d")
	m.Put(1, "x")
	m.Put(2, "b")
	m.Put(1, "a") //overwrite
	m.Put(2, "b")
	structKey := complexType{"skey"}
	structValue := complexType{"svalue"}
	m.Put(structKey, structValue)
	m.Put(&structKey, &structValue)
	m.Put(true, false)

	actualValues := m.Values()
	expectedValues := []interface{}{"e", "f", "g", "c", "d", "a", "b", structValue, &structValue, false}
	if !sameElements(actualValues, expectedValues) {
		t.Errorf("Got %v expected %v", actualValues, expectedValues)
	}
}

func BenchmarkOrderedMap_Put(b *testing.B) {
	m := orderedmap.NewOrderedMap()

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Put(i, i+1)
	}
}

var (
	resultBenchmarkOrderedMapGet1 interface{}
	resultBenchmarkOrderedMapGet2 bool
)

func BenchmarkOrderedMap_Get(b *testing.B) {
	m := orderedmap.NewOrderedMap()
	for i := 0; i < b.N; i++ {
		m.Put(i, i+1)
	}

	var (
		value interface{}
		found bool
	)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Get(i)
	}
	b.StopTimer()

	resultBenchmarkOrderedMapGet1 = value
	resultBenchmarkOrderedMapGet2 = found
}

func BenchmarkOrderedMap_Remove(b *testing.B) {
	m := orderedmap.NewOrderedMap()
	for i := 0; i < b.N; i++ {
		m.Put(i, i+1)
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Remove(i)
	}
}

var resultBenchmarkOrderedMapKeys []interface{}

func BenchmarkOrderedMap_Keys(b *testing.B) {
	m := orderedmap.NewOrderedMap()
	for i := 0; i < b.N; i++ {
		m.Put(i, i+1)
	}

	var keys []interface{}

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		keys = m.Keys()
	}
	b.StopTimer()

	resultBenchmarkOrderedMapKeys = keys
}

var resultBenchmarkOrderedMapValues []interface{}

func BenchmarkOrderedMap_Values(b *testing.B) {
	m := orderedmap.NewOrderedMap()
	for i := 0; i < b.N; i++ {
		m.Put(i, i+1)
	}

	var values []interface{}

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		values = m.Values()
	}
	b.StopTimer()

	resultBenchmarkOrderedMapValues = values
}

var resultBenchmarkOrderedMapSize int

func BenchmarkOrderedMap_Size(b *testing.B) {
	m := orderedmap.NewOrderedMap()
	for i := 0; i < b.N; i++ {
		m.Put(i, i+1)
	}

	var size int
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		size = m.Size()
	}
	b.StopTimer()
	resultBenchmarkOrderedMapSize = size
}

type complexType struct {
	foo string
}

func sameElements(a []interface{}, b []interface{}) bool {
	if len(a) != len(b) {
		return false
	}
	for _, av := range a {
		found := false
		for _, bv := range b {
			if av == bv {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

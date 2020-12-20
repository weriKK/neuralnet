package main

import (
	"reflect"
	"testing"
)

func TestMMul(t *testing.T) {

	m1 := m64{
		{0, -4, 6},
		{-3, 4, 1},
		{5, 2, 7},
	}

	m2 := m64{
		{8, -1, -2},
		{0, 4, 1},
		{6, 5, -3},
	}

	exp := m64{
		{36, 14, -22},
		{-18, 24, 7},
		{82, 38, -29},
	}

	out := m1.MMul(m2)

	if !reflect.DeepEqual(exp, out) {
		t.Fatalf("\nExpected:\t %v\nGot:\t\t %v", exp, out)
	}
}

func TestSMul(t *testing.T) {

	m1 := m64{
		{2, 4, 3},
		{1, 2, 1},
		{3, 4, 7},
	}

	m2 := m64{
		{1, 2, 3},
		{4, -3, 0},
		{2, 2, 0},
	}

	exp := m64{
		{2, 8, 9},
		{4, -6, 0},
		{6, 8, 0},
	}

	out := m1.SMul(m2)

	if !reflect.DeepEqual(exp, out) {
		t.Fatalf("\nExpected:\t %v\nGot:\t\t %v", exp, out)
	}
}

func TestTranspose(t *testing.T) {

	m1 := m64{
		{6, 4, 24},
		{1, 9, 8},
		{5, 5, 5},
	}

	exp := m64{
		{6, 1, 5},
		{4, 9, 5},
		{24, 8, 5},
	}

	out := m1.Transpose()

	if !reflect.DeepEqual(exp, out) {
		t.Fatalf("\nExpected:\t %v\nGot:\t\t %v", exp, out)
	}
}

func TestAdd(t *testing.T) {

	m1 := m64{
		{1, 2, 3},
		{3, 2, 1},
		{3, 3, 3},
	}

	m2 := m64{
		{9, 8, 7},
		{7, 8, 9},
		{7, 7, 7},
	}

	exp := m64{
		{10, 10, 10},
		{10, 10, 10},
		{10, 10, 10},
	}

	out := m1.Add(m2)

	if !reflect.DeepEqual(exp, out) {
		t.Fatalf("\nExpected:\t %v\nGot:\t\t %v", exp, out)
	}
}

func TestApplyFunc(t *testing.T) {

	mulfunc := func(n float64) float64 {
		return n * 2
	}

	m1 := m64{
		{6, 4, 24},
		{1, -9, 0},
		{5, 5, 5},
	}

	exp := m64{
		{12, 8, 48},
		{2, -18, 0},
		{10, 10, 10},
	}

	out := m1.ApplyFunc(mulfunc)

	if !reflect.DeepEqual(exp, out) {
		t.Fatalf("\nExpected:\t %v\nGot:\t\t %v", exp, out)
	}
}

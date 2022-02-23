package data

import (
	"math/rand"

	"github.com/google/uuid"
)

type A struct {
	SomeString  string
	SomeFloat   float64
	SomeIntData []int
}

type B struct {
	A               A
	SomeString      string
	OtherString     string
	SomeFloat32Data []float32
}

type BMap map[string]B

type C struct {
	SomeFloat64Data []float64
	SomeMap         BMap
}

func createString() string {
	return uuid.New().URN()
}

func createA() A {
	a := A{
		SomeString:  createString(),
		SomeFloat:   rand.Float64(),
		SomeIntData: make([]int, rand.Intn(50)+50),
	}
	for i := range a.SomeIntData {
		a.SomeIntData[i] = rand.Int()
	}
	return a
}

func createB() B {
	b := B{
		A:               createA(),
		SomeString:      createString(),
		OtherString:     createString(),
		SomeFloat32Data: make([]float32, rand.Intn(50)+50),
	}
	for i := range b.SomeFloat32Data {
		b.SomeFloat32Data[i] = rand.Float32()
	}
	return b
}

func createC() C {
	c := C{
		SomeFloat64Data: make([]float64, rand.Intn(50)+50),
		SomeMap:         make(map[string]B),
	}
	for i := range c.SomeFloat64Data {
		c.SomeFloat64Data[i] = rand.Float64()
	}
	for i := 0; i < 100; i++ {
		c.SomeMap[createString()] = createB()
	}
	return c
}

var SampleC C

func init() {
	SampleC = createC()
}

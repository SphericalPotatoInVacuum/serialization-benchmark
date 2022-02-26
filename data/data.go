package data

import (
	"math/rand"

	"github.com/SphericalPotatoInVacuum/serialization-benchmark/serializers/sproto"
	"github.com/google/uuid"
)

const size = 50

type A struct {
	SomeString    string
	SomeDouble    float64
	SomeInt32Data []int32
}

type B struct {
	SomeAData     []A
	SomeString    string
	OtherString   string
	SomeFloatData []float32
}

type BMap map[string]B

type C struct {
	SomeDoubleData []float64
	SomeMap        BMap
}

func createString() string {
	return uuid.New().URN()
}

func createA() A {
	a := A{
		SomeString:    createString(),
		SomeDouble:    rand.Float64(),
		SomeInt32Data: make([]int32, rand.Int31n(size)+size),
	}
	for i := range a.SomeInt32Data {
		a.SomeInt32Data[i] = rand.Int31()
	}
	return a
}

func createB() B {
	b := B{
		SomeAData:     make([]A, rand.Intn(size)+size),
		SomeString:    createString(),
		OtherString:   createString(),
		SomeFloatData: make([]float32, rand.Intn(size)+size),
	}
	for i := range b.SomeAData {
		b.SomeAData[i] = createA()
	}
	for i := range b.SomeFloatData {
		b.SomeFloatData[i] = rand.Float32()
	}
	return b
}

func createC() C {
	c := C{
		SomeDoubleData: make([]float64, rand.Intn(size)+size),
		SomeMap:        make(map[string]B),
	}
	for i := range c.SomeDoubleData {
		c.SomeDoubleData[i] = rand.Float64()
	}
	for i := 0; i < size; i++ {
		c.SomeMap[createString()] = createB()
	}
	return c
}

var SampleC C
var SampleProtoC sproto.C

func init() {
	rand.Seed(42)
	SampleC = createC()
	SampleProtoC = *toProtoC(SampleC)
}

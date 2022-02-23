package data

import "github.com/SphericalPotatoInVacuum/serialization-benchmark/serializers/sproto"

func toProtoC(input C) *sproto.C {
	out := sproto.C{
		SomeDoubleData: input.SomeDoubleData,
		SomeMap:        make(map[string]*sproto.C_B),
	}
	for k, v := range input.SomeMap {
		out.SomeMap[k] = toProtoB(v)
	}
	return &out
}

func toProtoB(input B) *sproto.C_B {
	out := sproto.C_B{
		SomeAData:     make([]*sproto.C_B_A, len(input.SomeAData)),
		SomeString:    input.SomeString,
		OtherString:   input.OtherString,
		SomeFloatData: input.SomeFloatData,
	}
	for i, v := range input.SomeAData {
		out.SomeAData[i] = toProtoA(v)
	}
	return &out
}

func toProtoA(input A) *sproto.C_B_A {
	out := sproto.C_B_A{
		SomeString:    input.SomeString,
		SomeDouble:    input.SomeDouble,
		SomeInt32Data: input.SomeInt32Data,
	}
	return &out
}

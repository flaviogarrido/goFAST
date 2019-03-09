// Copyright 2018 Alexander Poltoratskiy. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package fast_test

type decimalType struct {
	TemplateID           uint `fast:"*"`
	CopyDecimal          float64
	MandatoryDecimal     float64
	IndividualDecimal    float64
	IndividualDecimalOpt float64
}

type sequenceType struct {
	TemplateID    uint `fast:"*"`
	TestData      uint32
	OuterSequence []struct {
		OuterTestData uint32
		InnerSequence []struct {
			InnerTestData uint32
		}
	}
}

type byteVectorType struct {
	TemplateID      uint `fast:"*"`
	MandatoryVector []byte
	OptionalVector  []byte
}

type stringType struct {
	TemplateID       uint `fast:"*"`
	MandatoryAscii   string
	OptionalAscii    string
	MandatoryUnicode string
	OptionalUnicode  string
}

type integerType struct {
	TemplateID      uint `fast:"*"`
	MandatoryUint32 uint32
	OptionalUint32  uint32
	MandatoryUint64 uint64
	OptionalUint64  uint64
	MandatoryInt32  int32
	OptionalInt32   int32
	MandatoryInt64  int64
	OptionalInt64   int64
}

type groupType struct {
	TemplateID uint `fast:"*"`
	TestData   uint32
	OuterGroup struct {
		OuterTestData uint32
		InnerGroup    struct {
			InnerTestData uint32
		}
	}
}

var (
	decimalData1    = []byte{0xf8, 0x81, 0xfe, 0x4, 0x83, 0xff, 0xc, 0x8a, 0xfc, 0xa0, 0xff, 0x0, 0xef}
	decimalMessage1 = decimalType{
		TemplateID:           1,
		CopyDecimal:          5.15,
		MandatoryDecimal:     154.6,
		IndividualDecimal:    0.0032,
		IndividualDecimalOpt: 11.1,
	}

	sequenceData1    = []byte{0xc0, 0x82, 0x81, 0x81, 0x82, 0x83, 0x83, 0x84}
	sequenceMessage1 = sequenceType{
		TemplateID: 2,
		TestData:   1,
		OuterSequence: []struct {
			OuterTestData uint32
			InnerSequence []struct {
				InnerTestData uint32
			}
		}{
			{
				OuterTestData: 2,
				InnerSequence: []struct {
					InnerTestData uint32
				}{
					{InnerTestData: 3},
					{InnerTestData: 4},
				},
			},
		},
	}

	byteVectorData1    = []byte{0xc0, 0x83, 0x81, 0xc1, 0x82, 0xb3}
	byteVectorMessage1 = byteVectorType{
		TemplateID:      3,
		MandatoryVector: []byte{0xc1},
		OptionalVector:  []byte{0xb3},
	}

	stringData1    = []byte{0xc0, 0x84, 0x61, 0x62, 0xe3, 0x64, 0x65, 0xe6, 0x83, 0x67, 0x68, 0x69, 0x84, 0x6b, 0x6c, 0x6d}
	stringMessage1 = stringType{
		TemplateID:       4,
		MandatoryAscii:   "abc",
		OptionalAscii:    "def",
		MandatoryUnicode: "ghi",
		OptionalUnicode:  "klm",
	}

	integerData1    = []byte{0xc0, 0x85, 0x83, 0x85, 0x25, 0x20, 0x2f, 0x47, 0xfe, 0x25, 0x20, 0x2f, 0x48, 0x80, 0x85, 0x87, 0x8, 0x23, 0x51, 0x57, 0x8d, 0x8, 0x23, 0x51, 0x57, 0x8f}
	integerMessage1 = integerType{
		TemplateID:      5,
		MandatoryUint32: 3,
		OptionalUint32:  4,
		MandatoryUint64: 9999999998,
		OptionalUint64:  9999999999,
		MandatoryInt32:  5,
		OptionalInt32:   6,
		MandatoryInt64:  2222222221,
		OptionalInt64:   2222222222,
	}

	groupData1    = []byte{0xe0, 0x86, 0x81, 0x82, 0x83}
	groupMessage1 = groupType{
		TemplateID: 6,
		TestData:   1,
		OuterGroup: struct {
			OuterTestData uint32
			InnerGroup    struct {
				InnerTestData uint32
			}
		}{
			OuterTestData: 2,
			InnerGroup: struct {
				InnerTestData uint32
			}{
				InnerTestData: 3,
			},
		},
	}
)

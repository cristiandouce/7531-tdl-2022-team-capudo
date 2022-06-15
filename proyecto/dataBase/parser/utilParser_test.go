package parser

import (
	"testing"
	"fmt"
	"strings"
)

type tDatosParseToFloat32ErrorTest struct {
	input  string
	eParam error
	output  float32
	eTag string
}

var DatosParseToFloat32ErrorTest = []tDatosParseToFloat32ErrorTest{
	{input: "20.5", eParam: nil, output: 20.5, eTag: ""},
	{input: "-20.5", eParam: nil, output: -20.5, eTag: ""},
	{input: "2E0.5", eParam: nil, output: 0, eTag: "invalid syntax"},
	{input: "2E0.5", eParam: fmt.Errorf("Error"), output: 20.5, eTag: "Error"},
}

func TestParseToFloat32Error(t *testing.T) {
	for _, test := range DatosParseToFloat32ErrorTest {
		output,  e:= ParseToFloat32Error(test.input, test.eParam)
		if(test.output != output && !strings.Contains(e.Error(), test.eTag) ){
			t.Errorf("ParseToFloat32Error = (%v), se esperaba (%v)",output, test.input)
		}
	}
}

type tDatosParseToUint32ErrorTest struct {
	input  string
	eParam error
	output  uint32
	eTag string
}

var DatosParseToUint32ErrorTest = []tDatosParseToUint32ErrorTest{
	{input: "1789", eParam: nil, output: 1789, eTag: ""},
	{input: "58796231", eParam: nil, output: 58796231, eTag: ""},
	{input: "-58796231", eParam: nil, output: 0, eTag: "invalid syntax"},
	{input: "-58796231", eParam: fmt.Errorf("Error"), output: 58796231, eTag: "Error"},
}

func TestParseToUint32Error(t *testing.T) {
	for _, test := range DatosParseToUint32ErrorTest {
		output,  e:= ParseToUint32Error(test.input, test.eParam)
		if(test.output != output && !strings.Contains(e.Error(), test.eTag) ){
			t.Errorf("ParseToUint32Error = (%v), se esperaba (%v)",output, test.input)
		}
	}
}

type tDatosParseToIntErrorTest struct {
	input  string
	eParam error
	output  int
	eTag string
}

var DatosParseToIntErrorTest = []tDatosParseToIntErrorTest{
	{input: "17890", eParam: nil, output: 17890, eTag: ""},
	{input: "-58796231", eParam: nil, output: -58796231, eTag: ""},
	{input: "E58796231", eParam: nil, output: 0, eTag: "invalid syntax"},
	{input: "E58796231", eParam: fmt.Errorf("Error"), output: 0, eTag: "Error"},
}

func TestParseToIntError(t *testing.T) {
	for _, test := range DatosParseToIntErrorTest {
		output,  e:= ParseToIntError(test.input, test.eParam)
		if(test.output != output && !strings.Contains(e.Error(), test.eTag) ){
			t.Errorf("ParseToIntError = (%v), se esperaba (%v)",output, test.input)
		}
	}
}

type tDatosParseTimeUTCTest struct {
	input  string
	output string
	eTag string
}

var DatosParseTimeUTCTest = []tDatosParseTimeUTCTest{
	{input: "2021-04-10 20:48:15 UTC", output: "2021-04-10 20:48:15 +0000 UTC", eTag: ""},
	{input: "2021/04/10 20:48:15 UTC", output: "2021-04-10 20:48:15 +0000 UTC", eTag: "formato incorrecto de fecha"},
	{input: "2021-04-10 20/48/15 UTC", output: "2021-04-10 20:48:15 +0000 UTC", eTag: "formato incorrecto de hora"},
	{input: "2021/04/10 20/48/15 UTC", output: "2021-04-10 20:48:15 +0000 UTC", eTag: "formato incorrecto de fecha"},
}

func TestParseTimeUTCError(t *testing.T) {
	for _, test := range DatosParseTimeUTCTest {
		output,  e:= ParseTimeUTC(test.input)
		if(test.output != output.String() && !strings.Contains(e.Error(), test.eTag) ){
			t.Errorf("ParseTimeUTC = (%v), se esperaba (%v)",output, test.input)
		}
	}
}
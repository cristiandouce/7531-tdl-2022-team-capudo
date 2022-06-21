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

var datosParseToFloat32ErrorTest = []tDatosParseToFloat32ErrorTest{
	{input: "20.5", eParam: nil, output: 20.5, eTag: ""},
	{input: "-20.5", eParam: nil, output: -20.5, eTag: ""},
	{input: "2E0.5", eParam: nil, output: 0, eTag: "invalid syntax"},
	{input: "2E0.5", eParam: fmt.Errorf("Error"), output: 0, eTag: "Error"},
}

func TestParseToFloat32Error(t *testing.T) {
	for _, test := range datosParseToFloat32ErrorTest {
		output,  e:= ParseToFloat32Error(test.input, test.eParam)
		if(test.output != output){
			t.Errorf("ParseToFloat32Error = (%v), se esperaba (%v)",output, test.output)
		}
		if(e != nil){
			if(!strings.Contains(e.Error(), test.eTag) ){
				t.Errorf("ParseToFloat32Error = (%v), se esperaba (%v)",e.Error(), test.eTag)
			}
		}
		
	}
}

type tDatosParseToUint32ErrorTest struct {
	input  string
	eParam error
	output  uint32
	eTag string
}

var datosParseToUint32ErrorTest = []tDatosParseToUint32ErrorTest{
	{input: "1789", eParam: nil, output: 1789, eTag: ""},
	{input: "58796231", eParam: nil, output: 58796231, eTag: ""},
	{input: "-58796231", eParam: nil, output: 0, eTag: "invalid syntax"},
	{input: "-58796231", eParam: fmt.Errorf("Error"), output: 0, eTag: "Error"},
}

func TestParseToUint32Error(t *testing.T) {
	for _, test := range datosParseToUint32ErrorTest {
		output,  e:= ParseToUint32Error(test.input, test.eParam)
		if(test.output != output){
			t.Errorf("ParseToUint32Error = (%v), se esperaba (%v)",output, test.output)
		}
		if(e != nil){
			if(!strings.Contains(e.Error(), test.eTag) ){
				t.Errorf("ParseToUint32Error = (%v), se esperaba (%v)", e.Error(), test.eTag)
			}
		}
	}
}

type tDatosParseToUint16ErrorTest struct {
	input  string
	eParam error
	output  uint16
	eTag string
}

var datosParseToUint16ErrorTest = []tDatosParseToUint16ErrorTest{
	{input: "1789", eParam: nil, output: 1789, eTag: ""},
	{input: "1879", eParam: nil, output: 1879, eTag: ""},
	{input: "-1879", eParam: nil, output: 0, eTag: "invalid syntax"},
	{input: "-1879", eParam: fmt.Errorf("Error"), output: 0, eTag: "Error"},
}

func TestParseToUint16Error(t *testing.T) {
	for _, test := range datosParseToUint16ErrorTest {
		output,  e:= ParseToUint16Error(test.input, test.eParam)
		if(test.output != output){
			t.Errorf("ParseToUint16Error = (%v), se esperaba (%v)",output, test.output)
		} 
		if(e != nil){
			if(!strings.Contains(e.Error(), test.eTag) ){
				t.Errorf("ParseToUint16Error = (%v), se esperaba (%v)",e.Error(), test.eTag)
			}
		}
	}
}

type tDatosParseToIntErrorTest struct {
	input  string
	eParam error
	output  int
	eTag string
}

var datosParseToIntErrorTest = []tDatosParseToIntErrorTest{
	{input: "17890", eParam: nil, output: 17890, eTag: ""},
	{input: "-58796231", eParam: nil, output: -58796231, eTag: ""},
	{input: "E58796231", eParam: nil, output: 0, eTag: "invalid syntax"},
	{input: "E58796231", eParam: fmt.Errorf("Error"), output: 0, eTag: "Error"},
}

func TestParseToIntError(t *testing.T) {
	for _, test := range datosParseToIntErrorTest {
		output,  e:= ParseToIntError(test.input, test.eParam)
		if(test.output != output) {
			t.Errorf("ParseToIntError = (%v), se esperaba (%v)",output, test.output)
		}
		if(e != nil){
			if(!strings.Contains(e.Error(), test.eTag) ){
				t.Errorf("ParseToIntError = (%v), se esperaba (%v)", e.Error(), test.output)
			}
		}
	}
}

type tDatosParseTimeUTCTest struct {
	input  string
	output string
	eTag string
}

var datosParseTimeUTCTest = []tDatosParseTimeUTCTest{
	{input: "2021-04-10 20:48:15 UTC", output: "2021-04-10 20:48:15 +0000 UTC", eTag: ""},
	{input: "2021/04/10 20:48:15 UTC", output: "-0001-11-30 20:48:15 +0000 UTC", eTag: "formato incorrecto de fecha"},
	{input: "2021-04-10 20/48/15 UTC", output: "2021-04-10 00:00:00 +0000 UTC", eTag: "formato incorrecto de hora"},
	{input: "2021/04/10 20/48/15 UTC", output: "-0001-11-30 00:00:00 +0000 UTC", eTag: "formato incorrecto de fecha"},
}

func TestParseTimeUTCError(t *testing.T) {
	for _, test := range datosParseTimeUTCTest {
		output,  e:= ParseTimeUTC(test.input)
		if(test.output != output.String()){
			t.Errorf("ParseTimeUTC = (%v), se esperaba (%v)",output, test.output)
		}
		if(e != nil){
			if(!strings.Contains(e.Error(), test.eTag)){
				t.Errorf("ParseTimeUTC = (%v), se esperaba (%v)",e.Error(), test.eTag)
			}
		}
	}
}

type tDatosParseToInt16WithEmptyErrorTest struct {
	input  string
	eParam error
	output  int16
	eTag string
}

var datosParseToInt16WithEmptyErrorTest = []tDatosParseToInt16WithEmptyErrorTest{
	{input: "1", eParam: nil, output: 1, eTag: ""},
	{input: "40", eParam: nil, output: 40, eTag: ""},
	{input: "", eParam: nil, output: -1, eTag: ""},
	{input: "None", eParam: nil, output: -1, eTag: ""},
	{input: "-41", eParam: nil, output: 0, eTag: "invalid syntax"},
	{input: "41.0", eParam: nil, output: 0, eTag: "invalid syntax"},
	{input: "-42", eParam: fmt.Errorf("Error"), output: 0, eTag: "Error"},
}

func TestParseToInt16WithEmptyError(t *testing.T) {
	for _, test := range datosParseToInt16WithEmptyErrorTest {
		output,  e:= ParseToInt16WithEmptyError(test.input, test.eParam)
		if(test.output != output){
			t.Errorf("ParseToInt16WithEmptyError = (%v), se esperaba (%v)",output, test.output)	
		}
		if(e != nil ){
			if(!strings.Contains(e.Error(), test.eTag)){
				t.Errorf("ParseToInt16WithEmptyError = (%v), se esperaba (%v)",e.Error(), test.eTag)
			}
		} 
	}
}
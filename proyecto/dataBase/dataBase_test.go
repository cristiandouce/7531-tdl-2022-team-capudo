package dataBase

import (
	"testing"
)

const CANT_RECORRIDOS_PASS_TEST = 100

const CANT_BICICLETEROS_PASS_TEST = 100

func TestGetRecorridos(t *testing.T) {
	ouput, e := GetRecorridos()
	if (ouput == nil) {
		t.Errorf("testGetRecorridos no se pudieron recuperar los recorridos = (%v)", e)
	} else{
		if (len(*ouput) < CANT_RECORRIDOS_PASS_TEST){
			t.Errorf("testGetRecorridos recorridos insuficientes = (%v)", len(*ouput))
		}
	}
}

func TestGetBicicleteros(t *testing.T) {
	ouput, e := GetBicicleteros()
	if (ouput == nil) {
		t.Errorf("testGetBicicleteros no se pudieron recuperar los bicicleteros = (%v)", e)
	} else{
		if (len(*ouput) < CANT_BICICLETEROS_PASS_TEST){
			t.Errorf("testGetBicicleteros bicicleteros insuficientes = (%v)", len(*ouput))
		}
	}
}
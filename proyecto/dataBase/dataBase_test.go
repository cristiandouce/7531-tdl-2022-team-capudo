package dataBase

import (
	"testing"
)

const CANT_RECORRIDOS_PASS_TEST = 100

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

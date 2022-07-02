package database

import (
	"fmt"
	"testing"
)

const CANT_MIN_RECORRIDOS_PASS_TEST = 100

const CANT_MIN_BICICLETEROS_PASS_TEST = 100

const CANT_MIN_USUARIOS_PASS_TEST = 100

func TestGetRecorridos(t *testing.T) {
	ouput, e := GetRecorridos()
	if ouput == nil {
		t.Errorf("testGetRecorridos no se pudieron recuperar los recorridos = (%v)", e)
	} else {
		if len(*ouput) < CANT_MIN_RECORRIDOS_PASS_TEST {
			t.Errorf("testGetRecorridos recorridos insuficientes = (%v)", len(*ouput))
		}
	}
}

func TestGetBicicleteros(t *testing.T) {
	ouput, e := GetBicicleteros()
	if ouput == nil {
		t.Errorf("testGetBicicleteros no se pudieron recuperar los bicicleteros = (%v)", e)
	} else {
		if len(*ouput) < CANT_MIN_BICICLETEROS_PASS_TEST {
			t.Errorf("testGetBicicleteros bicicleteros insuficientes = (%v)", len(*ouput))
		}
	}
}
func TestGetUsuarios(t *testing.T) {
	ouput, e := GetUsuarios()
	if ouput == nil {
		t.Errorf("testGetUsuarios no se pudieron recuperar los usuarios = (%v)", e)
	} else {
		fmt.Println(len(*ouput))
		if len(*ouput) < CANT_MIN_USUARIOS_PASS_TEST {
			t.Errorf("testGetUsuarios usuarios insuficientes = (%v)", len(*ouput))
		}
	}
}

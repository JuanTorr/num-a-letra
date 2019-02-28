package numaletra

import (
	"testing"
)

func TestConvertir(t *testing.T) {
	s, err := IntLetra(0)
	if err != nil {
		t.Error(err)
	}
	if s != "cero" {
		t.Error(0, "!= cero; Res:", s)
	}

	/**************************/
	s, err = IntLetra(14)
	if err != nil {
		t.Error(err)
	}
	if s != "catorce" {
		t.Error(14, "!= catorce; Res:", s)
	}

	/**************************/
	s, err = IntLetra(22)
	if err != nil {
		t.Error(err)
	}
	if s != "veintidos" {
		t.Error(22, "!= veintidos; Res:", s)
	}

	/**************************/
	s, err = IntLetra(137)
	if err != nil {
		t.Error(err)
	}
	if s != "ciento treinta y siete" {
		t.Error(137, "!= ciento treinta y siete; Res:", s)
	}

	/**************************/
	s, err = IntLetra(300)
	if err != nil {
		t.Error(err)
	}
	if s != "trescientos" {
		t.Error(300, "!= trescientos; Res:", s)
	}

	/**************************/
	s, err = IntLetra(599)
	if err != nil {
		t.Error(err)
	}
	if s != "quinientos noventa y nueve" {
		t.Error(599, "!= quinientos noventa y nueve; Res:", s)
	}

	/**************************/
	s, err = IntLetra(1444) //1,444
	if err != nil {
		t.Error(err)
	}
	if s != "mil cuatrocientos cuarenta y cuatro" {
		t.Error(1444, "!= mil cuatrocientos cuarenta y cuatro; Res:", s)
	}

	/**************************/
	s, err = IntLetra(28319) //28,319
	if err != nil {
		t.Error(err)
	}
	if s != "veintiocho mil trescientos diecinueve" {
		t.Error(28319, "!= veintiocho mil trescientos diecinueve; Res:", s)
	}

	/**************************/
	s, err = IntLetra(256947) //256,947
	if err != nil {
		t.Error(err)
	}
	if s != "doscientos cincuenta y seis mil novecientos cuarenta y siete" {
		t.Error(256947, "!= doscientos cincuenta y seis mil novecientos cuarenta y siete; Res:", s)
	}
	/**************************/
	s, err = IntLetra(1569478) //1,569,478
	if err != nil {
		t.Error(err)
	}
	if s != "un millón quinientos sesenta y nueve mil cuatrocientos setenta y ocho" {
		t.Error(1569478, "!= un millón quinientos sesenta y nueve mil cuatrocientos setenta y ocho; Res:", s)
	}

	/**************************/
	s, err = IntLetra(67783156) //67,783,156
	if err != nil {
		t.Error(err)
	}
	if s != "sesenta y siete millones setecientos ochenta y tres mil ciento cincuenta y seis" {
		t.Error(67783156, "!= sesenta y siete millones setecientos ochenta y tres mil ciento cincuenta y seis; Res:", s)
	}

	/**************************/
	s, err = IntLetra(365782546) //365,782,546
	if err != nil {
		t.Error(err)
	}
	if s != "trescientos sesenta y cinco millones setecientos ochenta y dos mil quinientos cuarenta y seis" {
		t.Error(365782546, "!= trescientos sesenta y cinco millones setecientos ochenta y dos mil quinientos cuarenta y seis; Res:", s)
	}

	/**************************/
	s, err = IntLetra(1)
	if err != nil {
		t.Error(err)
	}
	if s != "uno" {
		t.Error(1, "!= uno; Res:", s)
	}

	/**************************/
	s, err = IntLetra(10000) //10,000
	if err != nil {
		t.Error(err)
	}
	if s != "diez mil" {
		t.Error(10000, "!= diez mil; Res:", s)
	}

	/**************************/
	s, err = IntLetra(1000000) //1,000,000
	if err != nil {
		t.Error(err)
	}
	if s != "un millón" {
		t.Error(1000000, "!= un millón; Res:", s)
	}

	/**************************/
	s, err = IntLetra(-1550)
	if err != nil {
		t.Error(err)
	}
	if s != "menos mil quinientos cincuenta" {
		t.Error(1550, "!= menos mil quinientos cincuenta; Res:", s)
	}

	/**************************/
	_, err = IntLetra(1000000000000) //1,000,000,000,000
	if err != ErrValorNoAdmitido {
		t.Error(err)
	}
}

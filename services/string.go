package services

import (
	"errors"
	"strconv"
)

type Chain struct {
	Type   string
	Value  string
	Length int
}

func ValidChain(c string) (bool, error) {
	if len(c) > 4 {
		t := c[0:2]
		l, _ := strconv.Atoi(c[2:4])
		v := c[4:]
		if t != "NN" && t != "TX" {
			return false, errors.New("el formato indicado no es correcto")
		}
		if l == len(v) {

			if t == "NN" {
				_, err := strconv.Atoi(v)
				if err != nil {
					return false, errors.New("el valor ingresado no coincide con el tipo indicado")
				}
			}
			return true, nil
		}
		return false, errors.New("la longitud indicada es incorrecta")

	}
	return false, errors.New("formato de cadena incorrecto")
}

func (r *Chain) ModChain(c string, v bool) (*Chain, error) {
	if v {
		r.Type = c[0:2]
		r.Length, _ = strconv.Atoi(c[2:4])
		r.Value = c[4:]

		return r, nil
	}
	return &Chain{"", "", 0}, nil
}

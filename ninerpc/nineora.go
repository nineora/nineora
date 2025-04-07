package ninerpc

import "github.com/nineora/nineora"

const (
	Nineora    = "nineora"
	NineoraGet = Nineora + "/get"
)

type NineoraService interface {
	Get() *nineora.Nineora
}

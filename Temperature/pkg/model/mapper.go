package model

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/src/gen"
)

func TemperatureToProto(m *Temperature) *gen.Temperature {
	return &gen.Temperature{
		Temperature: m.Temperature,
		Date:        timestamppb.New(m.Date),
	}
}

func TemperatureFromProto(m *gen.Temperature) *Temperature {
	return &Temperature{
		Temperature: m.Temperature,
		Date:        m.Date.AsTime(),
	}
}

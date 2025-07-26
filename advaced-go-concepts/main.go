package main

import (
	"errors"
	"fmt"
	"log"
)

var (
	ErrorNotImplemented = errors.New("not implemented")
	ErrorTruckNotFound  = errors.New("not found")
)

type Truck interface {
	LoadCargo() error
	UnloadCargo() error
}
type NormalTruck struct {
	id    string
	cargo int
}
type ElectricTruck struct {
	id      string
	cargo   int
	battery int
}

func (t *NormalTruck) LoadCargo() error {
	t.cargo += 2
	return nil
}
func (t *NormalTruck) UnloadCargo() error {
	t.cargo = 0
	return nil
}

func (e *ElectricTruck) LoadCargo() error {
	e.cargo += 2
	e.battery += 2
	return nil
}

func (e *ElectricTruck) UnloadCargo() error {
	e.cargo = 0
	e.battery = 0
	return nil
}

func processTruck(truck Truck) error {
	if err := truck.LoadCargo(); err != nil {
		return fmt.Errorf("Error loading cargo: %w \n", err)
	}
	return nil
}

func main() {
	if err := processTruck(&NormalTruck{id: "1"}); err != nil {
		log.Fatalf("Errror processing truck: %s", err)
	}
	if err := processTruck(&ElectricTruck{id: "1"}); err != nil {
		log.Fatalf("Errror processing truck: %s", err)
	}
}

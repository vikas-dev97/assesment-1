package main

import (
	"errors"
	"fmt"
)

type customerror struct {
	temprature  float32
	time         int
	errorType    string
	message       string
}
func (ce customerror) Error() string {
	return "temprature:" + ce.temprature + "Message:" + ce.message
}

	

type MicrowaveInterface interface {
	GrillMicrowave(temprature int) error
	BakeMicrowave(time int) error
}

// LG
type MicrowaveImpl1 struct {
	temprture    float32
	time       int
}

func (M *MicrowaveImpl1) GrillMicrowave(temprature int) error {
	if temprature<=100 degree
	return grillNotFound
}

func (M) *MicrowaveImpl1) BakeMicrowave(Temprature int) error {
	if temprature<=100
	return Bakedamaged
}

// samsung
type MicrowaveImpl2 struct {
	Microwave Temprture  float32
	Microwave Time       int
	MicrowaveName     string
	

func (M *MicrowaveImpl2) GrillMicrowave(Microwave time int) error {
     if time<=5
	return NotFound
}

func (M*MicrowaveImpl2)BakeMicrowave (Microwave time int) error {
	if time<=5 minute
	return Damaged
}

func main()  {
      	
  ``
	
}




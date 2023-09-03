package main

import (
	"fmt"
)

type Driven interface {
    Drive()
}

type Car struct{}

func (c *Car) Drive() {
    fmt.Println("Car is being driven")
}

type Driver struct {
    Age int
}

type CarProxy struct {
    car    Car
    driver *Driver
}

func NewCarProxy(driver *Driver) *CarProxy {
    return &CarProxy{Car{}, driver}
}

func (c *CarProxy) Drive() {
    if c.driver.Age >= 16 {
        c.car.Drive()
    } else {
        fmt.Println("Driver too young!")
    }
}

func main() {
	car := NewCarProxy(&Driver{12})
    car.Drive()
}
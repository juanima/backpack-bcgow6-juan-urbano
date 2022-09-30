package main

import "fmt"

func main() {
	var temperature int;
	var humidity int;
	var pressure float64;

	temperature, humidity, pressure = 9, 60, 1013.2;
	
	fmt.Printf("Temperature: %d degrees centigrade, humidity: %d and pressure: %f mb  in Bogotá\n", temperature, humidity, pressure);
}


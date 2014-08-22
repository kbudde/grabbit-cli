package main

import (
	"fmt"
	"os"
	"strings"
)

type OutputType int

const (
	Hostmonitor OutputType = iota
	Nagios      OutputType = iota
)

func getOutputType(outputString string) OutputType {
	switch strings.ToLower(outputString) {
	case "h", "hostmonitor":
		return Hostmonitor
	case "n", "nagios":
		return Nagios
	}
	panic("undefined output")
}

func printResults(value string, err error, result resultLevel, field string, outputType OutputType) {
	switch outputType {
	case Hostmonitor:
		outputHostmonitor(result, value, err)
	case Nagios:
		outputNagios(result, value, field, err)
	}
}

func outputHostmonitor(result resultLevel, value string, err error) {
	switch result {
	case OK:
		fmt.Printf("scriptres:Ok:%v\n", value)
	case Warn_Level:
		fmt.Printf("scriptres:Bad:%v\n", value)
	case Error_Level:
		fmt.Printf("scriptres:Bad:%v\n", value)
	case ErrorState:
		fmt.Printf("scriptres:Bad Contents:%v\n", err)
	case Unknown:
		fmt.Printf("scriptres:Unknown:%v\n", value)
	}
	os.Exit(0)

}

func outputNagios(result resultLevel, value string, field string, err error) {
	switch result {
	case OK:
		fmt.Printf("OK - Field: %v, value: %v\n", field, value)
		os.Exit(0)
	case Warn_Level:
		fmt.Printf("WARNING - Field: %v, value: %v\n", field, value)
		os.Exit(1)
	case Error_Level:
		fmt.Printf("CRITICAL - Field: %v, value: %v\n", field, value)
		os.Exit(2)
	case ErrorState:
		fmt.Printf("UNKNOWN - Field: %v, value: %v, Error: %v\n", field, value, err)
		os.Exit(3)
	case Unknown:
		fmt.Printf("UNKNOWN - Field: %v, value: %v, Error: %v\n", field, value, err)
		os.Exit(3)
	}

}

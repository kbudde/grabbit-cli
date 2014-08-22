package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/michaelklishin/rabbit-hole"
	"reflect"
	"strings"
)

func runCheck(check string, host string, username string, password string, vhost string, field string, queue string, nodeIndex int, exchangeName string) (string, error) {
	rmqc, err := rabbithole.NewClient(host, username, password)
	if err != nil {
		return "", errors.New("Connection failed")
	}
	switch strings.ToLower(check) {
	case "q", "queue":
		return handleQueueCheck(rmqc, vhost, queue, field)
	case "h", "host":
		return handleHostCheck(rmqc, nodeIndex, field)
	case "e", "exchange":
		return handleExchangeCheck(rmqc, vhost, exchangeName, field)
	case "":
		flag.PrintDefaults()
		return "", errors.New("Wrong Parameter")
	}
	flag.PrintDefaults()
	return "", errors.New("Wrong Parameter")
}

func handleExchangeCheck(rmqc *rabbithole.Client, vhost string, exchangeName string, field string) (string, error) {
	// information about individual exchange
	exchange, err := rmqc.GetExchange(vhost, exchangeName)
	if err != nil {
		return "", fmt.Errorf("GetExchange > %v", err)
	}
	return getFieldValue(exchange, field)
}

func handleQueueCheck(rmqc *rabbithole.Client, vhost string, queue string, field string) (string, error) {
	q, err := rmqc.GetQueue(vhost, queue)

	if err != nil {
		return "", err
	}
	return getFieldValue(q, field)
}

func handleHostCheck(rmqc *rabbithole.Client, nodeIndex int, field string) (string, error) {
	nodes, err := rmqc.ListNodes()
	if err != nil {
		return "", err
	}
	if nodeIndex >= len(nodes) {
		return "", errors.New("node index out of range")
	}
	return getFieldValue(&nodes[nodeIndex], field)
}

// Reflection. Get Value (as string) of field selected by name
func getFieldValue(f interface{}, fieldName string) (string, error) {
	val := reflect.ValueOf(f).Elem()

	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		switch strings.ToLower(fieldName) {
		case "print":
			fmt.Printf("Name: %s%s Value: %v\n", typeField.Name, strings.Repeat(" ", 40-len(typeField.Name)), valueField.Interface())
		case strings.ToLower(typeField.Name):
			return fmt.Sprintf("%v", valueField.Interface()), nil
		}
	}
	if fieldName == "print" {
		return "", fmt.Errorf("Just printing values")
	}
	return "", fmt.Errorf("Field not found: %v", fieldName)
}

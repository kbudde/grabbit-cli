package main

import (
	"fmt"
	"strconv"
)

type resultLevel int

const (
	OK          resultLevel = iota
	Warn_Level  resultLevel = iota
	Error_Level resultLevel = iota
	Unknown     resultLevel = iota
	ErrorState  resultLevel = iota
)

/**
* value: value to compare
* errObj: if not nil there was an error during getting the data.
* operator: how should the value be compared ("="" => equal value, ">" => value above ok, warn, error level, "<" => value below)
* okLevel, warnLevel, errorLevel: value to be compared with.
 */
func Compare(value string, errObj error, operator string, okLevel string, warnLevel string, errorLevel string) (resultLevel, error) {

	result := Unknown
	var compareError error

	if errObj != nil {
		result = ErrorState
		compareError = errObj
	} else {
		switch operator {
		case "=":
			if value == okLevel {
				result = OK
			} else if value == warnLevel {
				result = Warn_Level
			} else if value == errorLevel {
				result = Error_Level
			}
		case "<", ">":
			var err error
			result, err = compareFloat(value, operator, okLevel, warnLevel, errorLevel)
			if err != nil {
				compareError = err
			}
		}
	}
	return result, compareError
}

func compareFloat(value string, operator string, okLevel, warnLevel string, errorLevel string) (resultLevel, error) {
	fValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return ErrorState, fmt.Errorf("Error converting value '%v' to float", value)
	}

	fOKLevel, err := strconv.ParseFloat(okLevel, 64)
	if err != nil {
		return ErrorState, fmt.Errorf("Error converting ok value '%v' to float", okLevel)
	}

	fWarnLevel, err := strconv.ParseFloat(warnLevel, 64)
	if err != nil {
		if warnLevel == "" {
			fWarnLevel = fOKLevel //wanrlevel will no be used if not set
		} else {
			return ErrorState, fmt.Errorf("Error converting warn value '%v' to float", warnLevel)
		}
	}

	fErrorLevel, err := strconv.ParseFloat(errorLevel, 64)
	if err != nil {
		return ErrorState, fmt.Errorf("Error converting error value '%v' to float", errorLevel)
	}

	switch operator {
	case ">":
		if fValue > fOKLevel {
			return OK, nil
		} else if fValue > fWarnLevel {
			return Warn_Level, nil
		} else if fValue > fErrorLevel {
			return Error_Level, nil
		}
	case "<":
		if fValue < fOKLevel {
			return OK, nil
		} else if fValue < fWarnLevel {
			return Warn_Level, nil
		} else if fValue < fErrorLevel {
			return Error_Level, nil
		}
	}
	return Unknown, nil
}

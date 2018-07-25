package sdk

import (
	"errors"
)

type NetworkType uint8

// NetworkType enums
const (
	MAIN_NET NetworkType = 104
	TEST_NET NetworkType = 152
	MIJIN NetworkType = 96
	MIJIN_TEST NetworkType = 144
)

// Network error
var networkTypeError = errors.New("wrong raw NetworkType int")

// Get NetworkType by raw value
func NetworkTypeFromRaw(value int) (NetworkType, error){
	switch value {
	case 104:
		return MAIN_NET, nil
	case 152:
		return TEST_NET, nil
	case 96:
		return MIJIN, nil
	case 144:
		return MIJIN_TEST, nil
	default:
		return 0, networkTypeError
	}
}
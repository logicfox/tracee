package filters

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const (
	minIntVal int64 = math.MinInt64
	maxIntVal int64 = math.MaxInt64
)

type IntFilter struct {
	Equal    []int64
	NotEqual []int64
	Greater  int64
	Less     int64
	Is32Bit  bool
	enabled  bool
}

func NewIntFilter() *IntFilter {
	return newIntFilter(false)
}

func NewInt32Filter() *IntFilter {
	return newIntFilter(true)
}

func newIntFilter(is32Bit bool) *IntFilter {
	return &IntFilter{
		Equal:    []int64{},
		NotEqual: []int64{},
		Greater:  maxIntVal,
		Less:     minIntVal,
		Is32Bit:  is32Bit,
		enabled:  false,
	}
}

func (f *IntFilter) Enable() {
	f.enabled = true
}

func (f *IntFilter) Disable() {
	f.enabled = false
}

func (f *IntFilter) Enabled() bool {
	return f.enabled
}

func (filter *IntFilter) Parse(operatorAndValues string) error {
	if len(operatorAndValues) < 2 {
		return fmt.Errorf("invalid operator and/or values given to filter: %s", operatorAndValues)
	}
	valuesString := string(operatorAndValues[1:])
	operatorString := string(operatorAndValues[0])

	if operatorString == "!" {
		if len(operatorAndValues) < 3 {
			return fmt.Errorf("invalid operator and/or values given to filter: %s", operatorAndValues)
		}
		operatorString = operatorAndValues[0:2]
		valuesString = operatorAndValues[2:]
	}

	values := strings.Split(valuesString, ",")

	for i := range values {
		val, err := strconv.ParseInt(values[i], 10, 64)
		if err != nil {
			return fmt.Errorf("invalid filter value: %s", values[i])
		}
		if filter.Is32Bit && (val > math.MaxInt32) {
			return fmt.Errorf("filter value is too big: %s", values[i])
		}
		switch operatorString {
		case "=":
			filter.Equal = append(filter.Equal, val)
		case "!=":
			filter.NotEqual = append(filter.NotEqual, val)
		case ">":
			if (filter.Greater == maxIntVal) || (val > filter.Greater) {
				filter.Greater = val
			}
		case "<":
			if (filter.Less == minIntVal) || (val < filter.Less) {
				filter.Less = val
			}
		default:
			return fmt.Errorf("invalid filter operator: %s", operatorString)
		}
	}

	filter.Enable()

	return nil
}
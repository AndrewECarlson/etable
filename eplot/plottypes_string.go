// Code generated by "stringer -type=PlotTypes"; DO NOT EDIT.

package eplot

import (
	"errors"
	"strconv"
)

var _ = errors.New("dummy error")

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[XY-0]
	_ = x[Bar-1]
	_ = x[PlotTypesN-2]
}

const _PlotTypes_name = "XYBarPlotTypesN"

var _PlotTypes_index = [...]uint8{0, 2, 5, 15}

func (i PlotTypes) String() string {
	if i < 0 || i >= PlotTypes(len(_PlotTypes_index)-1) {
		return "PlotTypes(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _PlotTypes_name[_PlotTypes_index[i]:_PlotTypes_index[i+1]]
}

func (i *PlotTypes) FromString(s string) error {
	for j := 0; j < len(_PlotTypes_index)-1; j++ {
		if s == _PlotTypes_name[_PlotTypes_index[j]:_PlotTypes_index[j+1]] {
			*i = PlotTypes(j)
			return nil
		}
	}
	return errors.New("String: " + s + " is not a valid option for type: PlotTypes")
}

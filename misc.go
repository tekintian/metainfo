package metainfo

import (
	"reflect"
	"sort"
)

// Sorts the slice in place. Returns sl for convenience.
func Sort(sl interface{}, less interface{}) interface{} {
	sorter := sorter{
		sl:   reflect.ValueOf(sl),
		less: reflect.ValueOf(less),
	}
	sort.Sort(&sorter)
	return sorter.sl.Interface()
}

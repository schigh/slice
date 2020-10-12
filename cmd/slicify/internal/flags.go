package internal

import (
	"regexp"

	"github.com/schigh/slice"
)

const (
	mapOperation     = "map"
	eachOperation    = "each"
	tryEachOperation = "tryeach"
	ifEachOperation  = "ifeach"
	filterOperation  = "filter"
	chunkOperation   = "chunk"
	allOperation     = "all"
	byvalue          = "byvalue"
)

var (
	re *regexp.Regexp
)

func init() {
	// https://regex101.com/r/ydgkha/2
	re = regexp.MustCompile(`((?:map|each|filter|ifeach|tryeach|chunk)(?:\s(?:map|each|filter|ifeach|tryeach|chunk))*|all)(?:\sbyvalue)?`)
}

// OperationsFromFlags returns all operations in the flag
func OperationsFromFlags(flags string, forceByValue bool) ([]Operation, error) {
	matches := slice.String(re.FindAllString(flags, -1))

	var byValue bool
	if matches.Contains(byvalue) {
		byValue = true
	}

	byValue = byValue || forceByValue

	if matches.Contains(allOperation) {
		return allOps(!byValue), nil
	}

	var ops []Operation

	if matches.Contains(mapOperation) {
		ops = append(ops, Operation{
			Template: MapTmpl,
			ByRef:    !byValue,
		})
	}

	if matches.Contains(eachOperation) {
		ops = append(ops, Operation{
			Template: EachTmpl,
			ByRef:    !byValue,
		})
	}

	if matches.Contains("filter") {
		ops = append(ops, Operation{
			Template: FilterTmpl,
			ByRef:    !byValue,
		})
	}

	if matches.Contains("ifeach") {
		ops = append(ops, Operation{
			Template: IfEachTmpl,
			ByRef:    !byValue,
		})
	}

	if matches.Contains("tryeach") {
		ops = append(ops, Operation{
			Template: TryEachTmpl,
			ByRef:    !byValue,
		})
	}

	if matches.Contains("chunk") {
		ops = append(ops, Operation{
			Template: ChunkTmpl,
			ByRef:    !byValue,
		})
	}

	return ops, nil
}

func allOps(br bool) []Operation {
	return []Operation{
		{
			Template: MapTmpl,
			ByRef:    br,
			Name:     mapOperation,
		},
		{
			Template: FilterTmpl,
			ByRef:    br,
			Name:     filterOperation,
		},
		{
			Template: EachTmpl,
			ByRef:    br,
			Name:     eachOperation,
		},
		{
			Template: TryEachTmpl,
			ByRef:    br,
			Name:     tryEachOperation,
		},
		{
			Template: IfEachTmpl,
			ByRef:    br,
			Name:     ifEachOperation,
		},
		{
			Template: ChunkTmpl,
			ByRef:    br,
			Name:     chunkOperation,
		},
	}
}

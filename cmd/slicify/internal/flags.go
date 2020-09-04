package internal

import (
	"fmt"
	"regexp"
)

const (
	mapOperation     = "map"
	eachOperation    = "each"
	tryEachOperation = "tryeach"
	ifEachOperation  = "ifeach"
	filterOperation  = "filter"
	chunkOperation   = "chunk"
	allOperation     = "all"
)

var (
	re *regexp.Regexp
)

func init() {
	// https://regex101.com/r/ydgkha/1
	re = regexp.MustCompile(`(?m)(map|ifeach|tryeach|each|filter|chunk|all)(\(((?P<bv>byvalue)|(?P<cp>copy)|,)+\))?`)
}

// OperationsFromFlags returns all operations in the flag
func OperationsFromFlags(flags string) ([]Operation, error) {
	var ops []Operation
	matches := re.FindAllSubmatch([]byte(flags), -1)
	for _, m := range matches {
		operation := Operation{}
		operation.ByRef = len(m[4]) <= 0
		operation.Copy = len(m[5]) > 0
		op := string(m[1])
		switch op {
		case allOperation:
			return allOps(operation.ByRef, operation.Copy), nil
		case mapOperation:
			operation.Template = MapTmpl
		case filterOperation:
			operation.Template = FilterTmpl
		case eachOperation:
			operation.Template = EachTmpl
		case tryEachOperation:
			operation.Template = TryEachTmpl
		case ifEachOperation:
			operation.Template = IfEachTmpl
		case chunkOperation:
			operation.Template = ChunkTmpl
		default:
			return ops, fmt.Errorf("unknown operation: '%s'", op)
		}

		operation.Name = op

		ops = append(ops, operation)
	}

	return ops, nil
}

func allOps(br, cp bool) []Operation {
	return []Operation{
		{
			Template: MapTmpl,
			ByRef:    br,
			Copy:     cp,
			Name:     mapOperation,
		},
		{
			Template: FilterTmpl,
			ByRef:    br,
			Copy:     cp,
			Name:     filterOperation,
		},
		{
			Template: EachTmpl,
			ByRef:    br,
			Copy:     cp,
			Name:     eachOperation,
		},
		{
			Template: TryEachTmpl,
			ByRef:    br,
			Copy:     cp,
			Name:     tryEachOperation,
		},
		{
			Template: IfEachTmpl,
			ByRef:    br,
			Copy:     cp,
			Name:     ifEachOperation,
		},
		{
			Template: ChunkTmpl,
			ByRef:    br,
			Copy:     cp,
			Name:     chunkOperation,
		},
	}
}

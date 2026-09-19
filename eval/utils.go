package eval

import "lipi/object"

func isTruthy(condition object.Object) bool {
	switch condition := condition.(type) {
	case *object.Bool:
		return condition.Value
	case *object.Int:
		return condition.Value != 0
	case *object.Nil:
		return false
	default:
		return false
	}
}

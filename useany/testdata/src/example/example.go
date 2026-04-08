package example

// Empty interface in various positions.

var _ interface{} // want `use 'any' instead of 'interface\{\}'`

type T interface{} // want `use 'any' instead of 'interface\{\}'`

func emptyParam(_ interface{}) {} // want `use 'any' instead of 'interface\{\}'`

func emptyReturn() interface{} { return nil } // want `use 'any' instead of 'interface\{\}'`

func emptyLocal() {
	var _ interface{}       // want `use 'any' instead of 'interface\{\}'`
	_ = []interface{}{1, 2} // want `use 'any' instead of 'interface\{\}'`
}

func emptyMap() {
	_ = map[string]interface{}{} // want `use 'any' instead of 'interface\{\}'`
}

// Non-empty interfaces should NOT be flagged.

type Stringer interface {
	String() string
}

type ReadWriter interface {
	Read(p []byte) (int, error)
	Write(p []byte) (int, error)
}

// Using 'any' is fine.

var _ any

type U any

func anyParam(_ any) {}

func anyReturn() any { return nil }

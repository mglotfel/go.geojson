package geojson

import (
	"fmt"
)

// SetProperty provides the inverse of all the property functions
// and is here for consistency.
func (g *Geometry) SetProperty(key string, value interface{}) {
	if g.Properties == nil {
		g.Properties = make(map[string]interface{})
	}
	g.Properties[key] = value
}

// PropertyBool type asserts a property to `bool`.
func (g *Geometry) PropertyBool(key string) (bool, error) {
	if b, ok := (g.Properties[key]).(bool); ok {
		return b, nil
	}
	return false, fmt.Errorf("type assertion of `%s` to bool failed", key)
}

// PropertyInt type asserts a property to `int`.
func (g *Geometry) PropertyInt(key string) (int, error) {
	if i, ok := (g.Properties[key]).(int); ok {
		return i, nil
	}

	if i, ok := (g.Properties[key]).(float64); ok {
		return int(i), nil
	}

	return 0, fmt.Errorf("type assertion of `%s` to int failed", key)
}

// PropertyFloat64 type asserts a property to `float64`.
func (g *Geometry) PropertyFloat64(key string) (float64, error) {
	if i, ok := (g.Properties[key]).(float64); ok {
		return i, nil
	}
	return 0, fmt.Errorf("type assertion of `%s` to float64 failed", key)
}

// PropertyString type asserts a property to `string`.
func (g *Geometry) PropertyString(key string) (string, error) {
	if s, ok := (g.Properties[key]).(string); ok {
		return s, nil
	}
	return "", fmt.Errorf("type assertion of `%s` to string failed", key)
}

// PropertyMustBool guarantees the return of a `bool` (with optional default)
//
// useful when you explicitly want a `bool` in a single value return context:
//     myFunc(f.PropertyMustBool("param1"), f.PropertyMustBool("optional_param", true))
func (g *Geometry) PropertyMustBool(key string, def ...bool) bool {
	var defaul bool

	b, err := g.PropertyBool(key)
	if err == nil {
		return b
	}

	if len(def) > 0 {
		defaul = def[0]
	}

	return defaul
}

// PropertyMustInt guarantees the return of a `bool` (with optional default)
//
// useful when you explicitly want a `bool` in a single value return context:
//     myFunc(f.PropertyMustInt("param1"), f.PropertyMustInt("optional_param", 123))
func (g *Geometry) PropertyMustInt(key string, def ...int) int {
	var defaul int

	b, err := g.PropertyInt(key)
	if err == nil {
		return b
	}

	if len(def) > 0 {
		defaul = def[0]
	}

	return defaul
}

// PropertyMustFloat64 guarantees the return of a `bool` (with optional default)
//
// useful when you explicitly want a `bool` in a single value return context:
//     myFunc(f.PropertyMustFloat64("param1"), f.PropertyMustFloat64("optional_param", 10.1))
func (g *Geometry) PropertyMustFloat64(key string, def ...float64) float64 {
	var defaul float64

	b, err := g.PropertyFloat64(key)
	if err == nil {
		return b
	}

	if len(def) > 0 {
		defaul = def[0]
	}

	return defaul
}

// PropertyMustString guarantees the return of a `bool` (with optional default)
//
// useful when you explicitly want a `bool` in a single value return context:
//     myFunc(f.PropertyMustString("param1"), f.PropertyMustString("optional_param", "default"))
func (g *Geometry) PropertyMustString(key string, def ...string) string {
	var defaul string

	b, err := g.PropertyString(key)
	if err == nil {
		return b
	}

	if len(def) > 0 {
		defaul = def[0]
	}

	return defaul
}

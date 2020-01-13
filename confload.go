package confload

import (
	"flag"
	"fmt"
	"reflect"
	"strings"
)

// Loader type
type Loader int

// const for available loaders
const (
	VarEnvLoader Loader = iota
	FlagLoader
)

type Field struct {
	Name  []string
	Key   string
	Value reflect.Value
}

// Load configurations
func Load(cfg interface{}, loaders ...Loader) {
	fields := parseFields([]string{}, reflect.ValueOf(cfg).Elem())

	fmt.Println("fields", fields)

	for _, l := range loaders {
		switch l {
		case FlagLoader:
			flagLoader(fields)
		}
	}
}

func parseFields(prefix []string, ref reflect.Value) (fields []Field) {
	for i := 0; i < ref.Type().NumField(); i++ {
		value := ref.Field(i)
		typeField := ref.Type().Field(i)

		name := make([]string, len(prefix)+1)
		copy(name, prefix)
		name[len(name)-1] = typeField.Tag.Get("conf")

		switch value.Kind() {
		case reflect.Struct:
			fields = append(fields, parseFields(name, value)...)
		default:
			fields = append(fields, Field{
				Name:  name,
				Value: value,
			})
		}
	}

	return
}

func flagLoader(fields []Field) {
	for _, field := range fields {
		v := field.Value

		switch v.Kind() {
		case reflect.String:
			var val string
			flag.StringVar(&val, flagName(field.Name), v.String(), "")
			defer func() { v.SetString(val) }()
		case reflect.Int:
			var val int64
			flag.Int64Var(&val, flagName(field.Name), v.Int(), "")
			defer func() { v.SetInt(val) }()
		}
	}

	flag.Parse()
}

func flagName(name []string) string {
	return strings.Join(name, "-")
}

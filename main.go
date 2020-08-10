package structenv

import (
	"fmt"
	"os"
	"reflect"
)

const (
	TagName      = "env"
	ErrNotString = "поле с тегом env имеет не строковый тип"
	ErrEnvNotSet = "не установлена переменная окружения %v"
)

func setFieldFromEnv(field reflect.Value) error {
	value := os.Getenv(field.String())
	if value == "" {
		return EnvError{fmt.Errorf(ErrEnvNotSet, field)}
	}

	field.SetString(value)

	return nil
}

func traverseStruct(s reflect.Value) error {
	for i := 0; i < s.NumField(); i++ {
		if s.Type().Field(i).Type.Kind() == reflect.Struct {
			if err := traverseStruct(s.Field(i)); err != nil {
				return err
			}
		}

		if value, ok := s.Type().Field(i).Tag.Lookup(TagName); ok {
			if s.Field(i).Type().Kind() != reflect.String {
				return TypeError{fmt.Errorf(ErrNotString)}
			}

			if value != "" {
				s.Field(i).SetString(value)
			}

			if err := setFieldFromEnv(s.Field(i)); err != nil {
				return err
			}
		}
	}

	return nil
}

func SetFromEnvs(v interface{}) error {
	if err := traverseStruct(reflect.ValueOf(v).Elem()); err != nil {
		return err
	}

	return nil
}

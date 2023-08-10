package nveapi

import (
	"fmt"
	"reflect"
	"strconv"
)

func generateQueryParameterMap(r RequestQuery) (qm map[string]string, err error) {

	qm = make(map[string]string, 0)

	t := reflect.TypeOf(r)

	for i := 0; i < t.NumField(); i++ {

		var tqm map[string]string
		var accept bool

		field := t.Field(i)                                            //Get structField
		fieldValue := reflect.ValueOf(r).FieldByIndex(field.Index)     // Get the field value
		fieldTagQueryParameterValue := field.Tag.Get("queryParameter") // Get the query tag value
		fieldTagRequiredValue := field.Tag.Get("required")             // Get the required tag value

		tqm, accept, err = returnValueIfValid(&field, &fieldValue, fieldTagQueryParameterValue, fieldTagRequiredValue)
		if err != nil {
			return
		}

		if accept {
			for k, v := range tqm {
				qm[k] = v
			}
		}

	}

	return
}

func returnValueIfValid(f *reflect.StructField, v *reflect.Value, queryParamterName, required string) (value map[string]string, accept bool, err error) {

	value = make(map[string]string, 0)

	q := reflect.ValueOf(queryParamterName)
	if q.IsZero() {
		return
	}

	r := reflect.ValueOf(required)
	if r.IsZero() {
		required = "false"
	}

	rb, err := strconv.ParseBool(required)
	if err != nil {
		return
	}

	if v.IsZero() && !rb {
		accept = false
		return
	}

	if v.IsZero() && rb {
		err = fmt.Errorf("missing required value for field [%s]", f.Name)
		return
	}

	value[queryParamterName] = v.String()
	accept = true

	return
}

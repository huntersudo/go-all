package

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func GetMapKeys(m map[string][]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func Encode(s string) string {
	result := url.QueryEscape(s)
	result = strings.Replace(result, "+", "%20", -1)
	result = strings.Replace(result, "*", "%2A", -1)
	result = strings.Replace(result, "%7E", "~", -1)
	return result
}

func GetTimeStamp() string {
	now := time.Now()
	local, _ := time.LoadLocation("Local")
	timestamp := now.In(local).Format("2006-01-02T15:04:05Z")
	return timestamp
}

var t time.Time

func isZero(v reflect.Value) bool {
	//fmt.Printf("\n\nchecking isZero for value: %+v\n", v)
	switch v.Kind() {
	case reflect.Ptr:
		if v.IsNil() {
			return true
		}
		return false
	case reflect.Func, reflect.Map, reflect.Slice:
		return v.IsNil()
	case reflect.Array:
		z := true
		for i := 0; i < v.Len(); i++ {
			z = z && isZero(v.Index(i))
		}
		return z
	case reflect.Struct:
		if v.Type() == reflect.TypeOf(t) {
			if v.Interface().(time.Time).IsZero() {
				return true
			}
			return false
		}
		z := true
		for i := 0; i < v.NumField(); i++ {
			z = z && isZero(v.Field(i))
		}
		return z
	}
	// Compare other types directly:
	z := reflect.Zero(v.Type())
	//fmt.Printf("zero type for value: %+v\n\n\n", z)
	return v.Interface() == z.Interface()
}

//GetSignatureNonce returns random number
func GetSignatureNonce() int {
	return rand.Intn(999) + 1
}

func DeepCopyMap(source url.Values) url.Values {
	newmap := make(url.Values)
	for k, v := range source {
		if len(source[k]) > 1 {
			for i := range v {
				newmap.Add(k, v[i])
			}
		}
		newmap.Set(k, source.Get(k))
	}
	return newmap
}

// TODO-SML 仔细看 算是tag的一个小应用

func BuildQuery(opts interface{}) (url.Values, error) {
	optsValue := reflect.ValueOf(opts)
	if optsValue.Kind() == reflect.Ptr {
		optsValue = optsValue.Elem()
	}

	optsType := reflect.TypeOf(opts)
	if optsType.Kind() == reflect.Ptr {
		optsType = optsType.Elem()
	}

	params := url.Values{}

	if optsValue.Kind() == reflect.Struct {
		for i := 0; i < optsValue.NumField(); i++ {
			v := optsValue.Field(i)
			f := optsType.Field(i)
			qTag := f.Tag.Get("q")

			// if the field has a 'q' tag, it goes in the query string  TODO-SML 反射加 tag 实现需求
			if qTag != "" {
				tags := strings.Split(qTag, ",")

				// if the field is set, add it to the slice of query pieces
				if !isZero(v) {
				loop:
					switch v.Kind() {
					case reflect.Ptr:
						v = v.Elem()
						goto loop
					case reflect.String:
						params.Add(tags[0], v.String())
					case reflect.Int:
						params.Add(tags[0], strconv.FormatInt(v.Int(), 10))
					case reflect.Bool:
						params.Add(tags[0], strconv.FormatBool(v.Bool()))
					case reflect.Slice:
						switch v.Type().Elem() {
						case reflect.TypeOf(0):
							for i := 0; i < v.Len(); i++ {
								params.Add(tags[0], strconv.FormatInt(v.Index(i).Int(), 10))
							}
						default:
							for i := 0; i < v.Len(); i++ {
								params.Add(tags[0], v.Index(i).String())
							}
						}
					case reflect.Map:
						if v.Type().Key().Kind() == reflect.String && v.Type().Elem().Kind() == reflect.String {
							var s []string
							for _, k := range v.MapKeys() {
								value := v.MapIndex(k).String()
								s = append(s, fmt.Sprintf("'%s':'%s'", k.String(), value))
							}
							params.Add(tags[0], fmt.Sprintf("{%s}", strings.Join(s, ", ")))
						}
					}
				} else {
					// if the field has a 'required' tag, it can't have a zero-value
					if requiredTag := f.Tag.Get("required"); requiredTag == "true" {
						return nil, fmt.Errorf("Required query parameter [%s] not set.", f.Name)
					}
				}
			}
		}

		return params, nil
	}
	// Return an error if the underlying type of 'opts' isn't a struct.
	return nil, fmt.Errorf("Options type is not a struct.")
}

func DefaultOkCodes(method string) []int {
	switch {
	case method == "GET":
		return []int{200}
	case method == "POST":
		return []int{200, 201, 202}
	case method == "PUT":
		return []int{200, 201, 202}
	case method == "PATCH":
		return []int{200, 202, 204}
	case method == "DELETE":
		return []int{200, 202, 204}
	}

	return []int{}
}

func ResponseError(statusCode int, url string, method string, body []byte) error {
	var err error
	respErr := errors.ErrUnexpectedResponseCode{
		URL:    url,
		Method: method,
		Actual: statusCode,
		Body:   body,
	}
	switch statusCode {
	case http.StatusBadRequest:
		err = errors.ErrDefault400{respErr}
	case http.StatusUnauthorized:
		err = errors.ErrDefault401{respErr}
	case http.StatusForbidden:
		err = errors.ErrDefault403{respErr}
	case http.StatusNotFound:
		err = errors.ErrDefault404{respErr}
	case http.StatusMethodNotAllowed:
		err = errors.ErrDefault405{respErr}
	case http.StatusRequestTimeout:
		err = errors.ErrDefault408{respErr}
	case http.StatusConflict:
		err = errors.ErrDefault409{respErr}
	case 429:
		err = errors.ErrDefault429{respErr}
	case http.StatusInternalServerError:
		err = errors.ErrDefault500{respErr}
	case http.StatusServiceUnavailable:
		err = errors.ErrDefault503{respErr}
	}
	if err == nil {
		err = respErr
	}
	return err
}

func UpperCaseConvertor(originalString string) string {
	upper := strings.ToUpper(originalString)
	upper = strings.Replace(upper, "-", "_", -1)
	return upper
}

func BuildRequestBody(opts interface{}, parent string) (map[string]interface{}, error) {
	optsValue := reflect.ValueOf(opts)
	if optsValue.Kind() == reflect.Ptr {
		optsValue = optsValue.Elem()
	}

	optsType := reflect.TypeOf(opts)
	if optsType.Kind() == reflect.Ptr {
		optsType = optsType.Elem()
	}

	optsMap := make(map[string]interface{})
	if optsValue.Kind() == reflect.Struct {
		//fmt.Printf("optsValue.Kind() is a reflect.Struct: %+v\n", optsValue.Kind())
		for i := 0; i < optsValue.NumField(); i++ {
			v := optsValue.Field(i)
			f := optsType.Field(i)

			if f.Name != strings.Title(f.Name) {
				//fmt.Printf("Skipping field: %s...\n", f.Name)
				continue
			}

			//fmt.Printf("Starting on field: %s...\n", f.Name)

			zero := isZero(v)
			//fmt.Printf("v is zero?: %v\n", zero)

			// if the field has a required tag that's set to "true"
			if requiredTag := f.Tag.Get("required"); requiredTag == "true" {
				//fmt.Printf("Checking required field [%s]:\n\tv: %+v\n\tisZero:%v\n", f.Name, v.Interface(), zero)
				// if the field's value is zero, return a missing-argument error
				if zero {
					// if the field has a 'required' tag, it can't have a zero-value
					err := errors.ErrMissingInput{}
					err.Argument = f.Name
					return nil, err
				}
			}

			if xorTag := f.Tag.Get("xor"); xorTag != "" {
				//fmt.Printf("Checking `xor` tag for field [%s] with value %+v:\n\txorTag: %s\n", f.Name, v, xorTag)
				xorField := optsValue.FieldByName(xorTag)
				var xorFieldIsZero bool
				if reflect.ValueOf(xorField.Interface()) == reflect.Zero(xorField.Type()) {
					xorFieldIsZero = true
				} else {
					if xorField.Kind() == reflect.Ptr {
						xorField = xorField.Elem()
					}
					xorFieldIsZero = isZero(xorField)
				}
				if !(zero != xorFieldIsZero) {
					err := errors.ErrMissingInput{}
					err.Argument = fmt.Sprintf("%s/%s", f.Name, xorTag)
					err.Info = fmt.Sprintf("Exactly one of %s and %s must be provided", f.Name, xorTag)
					return nil, err
				}
			}

			if orTag := f.Tag.Get("or"); orTag != "" {
				//fmt.Printf("Checking `or` tag for field with:\n\tname: %+v\n\torTag:%s\n", f.Name, orTag)
				//fmt.Printf("field is zero?: %v\n", zero)
				if zero {
					orField := optsValue.FieldByName(orTag)
					var orFieldIsZero bool
					if reflect.ValueOf(orField.Interface()) == reflect.Zero(orField.Type()) {
						orFieldIsZero = true
					} else {
						if orField.Kind() == reflect.Ptr {
							orField = orField.Elem()
						}
						orFieldIsZero = isZero(orField)
					}
					if orFieldIsZero {
						err := errors.ErrMissingInput{}
						err.Argument = fmt.Sprintf("%s/%s", f.Name, orTag)
						err.Info = fmt.Sprintf("At least one of %s and %s must be provided", f.Name, orTag)
						return nil, err
					}
				}
			}

			jsonTag := f.Tag.Get("json")
			if jsonTag == "-" {
				continue
			}

			if v.Kind() == reflect.Slice || (v.Kind() == reflect.Ptr && v.Elem().Kind() == reflect.Slice) {
				sliceValue := v
				if sliceValue.Kind() == reflect.Ptr {
					sliceValue = sliceValue.Elem()
				}

				for i := 0; i < sliceValue.Len(); i++ {
					element := sliceValue.Index(i)
					if element.Kind() == reflect.Struct || (element.Kind() == reflect.Ptr && element.Elem().Kind() == reflect.Struct) {
						_, err := BuildRequestBody(element.Interface(), "")
						if err != nil {
							return nil, err
						}
					}
				}
			}
			if v.Kind() == reflect.Struct || (v.Kind() == reflect.Ptr && v.Elem().Kind() == reflect.Struct) {
				if zero {
					//fmt.Printf("value before change: %+v\n", optsValue.Field(i))
					if jsonTag != "" {
						jsonTagPieces := strings.Split(jsonTag, ",")
						if len(jsonTagPieces) > 1 && jsonTagPieces[1] == "omitempty" {
							if v.CanSet() {
								if !v.IsNil() {
									if v.Kind() == reflect.Ptr {
										v.Set(reflect.Zero(v.Type()))
									}
								}
								//fmt.Printf("value after change: %+v\n", optsValue.Field(i))
							}
						}
					}
					continue
				}

				//fmt.Printf("Calling BuildRequestBody with:\n\tv: %+v\n\tf.Name:%s\n", v.Interface(), f.Name)
				_, err := BuildRequestBody(v.Interface(), f.Name)
				if err != nil {
					return nil, err
				}
			}
		}

		//fmt.Printf("opts: %+v \n", opts)

		b, err := json.Marshal(opts)
		if err != nil {
			return nil, err
		}

		//fmt.Printf("string(b): %s\n", string(b))

		err = json.Unmarshal(b, &optsMap)
		if err != nil {
			return nil, err
		}

		//fmt.Printf("optsMap: %+v\n", optsMap)

		if parent != "" {
			optsMap = map[string]interface{}{parent: optsMap}
		}
		//fmt.Printf("optsMap after parent added: %+v\n", optsMap)
		return optsMap, nil
	}
	// Return an error if the underlying type of 'opts' isn't a struct.
	return nil, fmt.Errorf("Options type is not a struct.")
}

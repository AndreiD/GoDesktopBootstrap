package utils

import (
	"context"
	"godesktop/utils/log"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"time"
)

// Close error checking for defer close
func Close(c io.Closer) {
	err := c.Close()
	if err != nil {
		log.Fatal(err)
	}
}

// GetFloat .
func GetFloat(unk interface{}) float64 {
	var floatType = reflect.TypeOf(float64(0))
	var stringType = reflect.TypeOf("")

	switch i := unk.(type) {
	case float64:
		return i
	case float32:
		return float64(i)
	case int64:
		return float64(i)
	case int32:
		return float64(i)
	case int:
		return float64(i)
	case uint64:
		return float64(i)
	case uint32:
		return float64(i)
	case uint:
		return float64(i)
	case string:
		output, err := strconv.ParseFloat(i, 64)
		if err != nil {
			log.Error(err)
			return 0
		}
		return output
	default:
		v := reflect.ValueOf(unk)
		v = reflect.Indirect(v)
		if v.Type().ConvertibleTo(floatType) {
			fv := v.Convert(floatType)
			return fv.Float()
		} else if v.Type().ConvertibleTo(stringType) {
			sv := v.Convert(stringType)
			s := sv.String()
			output, err := strconv.ParseFloat(s, 64)
			if err != nil {
				log.Error(err)
				return 0
			}
			return output
		} else {
			log.Errorf("can't convert %v to float64", v.Type())
			return 0
		}
	}
}

// GetRequest executes a GET request
func GetRequest(url string) ([]byte, error) {

	client := http.Client{Timeout: 20 * time.Second}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	reqWithDeadline := req.WithContext(ctx)
	response, err := client.Do(reqWithDeadline)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(response.Body)
	return data, err

}

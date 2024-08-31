package main

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gorilla/mux"
)

// Domain Layer
type Counter struct {
	Count int `json:"count"`
}

func NewService() *Counter {
	return &Counter{}
}

type service interface {
	Get(ctx context.Context, id string) (int, error)
	Add(ctx context.Context, c Counter) error
}

// Transport Layer

// queryDecoder extracts URL query parameters into the given struct
func queryDecoder(r *http.Request, inp interface{}) error {
	q := r.URL.Query()
	inpVal := reflect.ValueOf(inp)
	if inpVal.Kind() != reflect.Ptr || inpVal.IsNil() {
		return errors.New("inp must be a non-nil pointer to a struct")
	}

	inpVal = inpVal.Elem()
	if inpVal.Kind() != reflect.Struct {
		return errors.New("inp must be a pointer to a struct")
	}

	inpType := inpVal.Type()
	for i := 0; i < inpType.NumField(); i++ {
		field := inpType.Field(i)
		tag := field.Tag.Get("url")
		if tag == "" {
			continue
		}

		values, ok := q[tag]
		if ok && len(values) > 0 {
			value := values[0]
			switch field.Type.Kind() {
			case reflect.Int:
				d, err := strconv.Atoi(value)
				if err != nil {
					return err
				}
				inpVal.Field(i).SetInt(int64(d))
			case reflect.String:
				inpVal.Field(i).SetString(value)
			case reflect.Float32:
				d, err := strconv.ParseFloat(value, 32)
				if err != nil {
					return err
				}
				inpVal.Field(i).SetFloat(d)
			case reflect.Float64:
				d, err := strconv.ParseFloat(value, 64)
				if err != nil {
					return err
				}
				inpVal.Field(i).SetFloat(d)
			case reflect.Bool:
				d, err := strconv.ParseBool(value)
				if err != nil {
					return err
				}
				inpVal.Field(i).SetBool(d)
			default:
				return errors.New("unsupported field type: " + field.Type.Kind().String())
			}
		}
	}
	return nil
}

type transport struct {
	svc service
}

func NewTransport(svc service) *transport {
	return &transport{
		svc: svc,
	}
}

type getRequest struct {
	Id string `url:"id"`
}

type getResponse struct {
	Count   int  `json:"count"`
	Success bool `json:"success"`
}
type addRequest struct {
	Counter Counter `json:"counter"`
}

type addResponse struct {
	Success bool `json:"success"`
}

// Get handler retrieves data based on the ID
func (trans *transport) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	req := &getRequest{}
	if err := queryDecoder(r, req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := req.Id
	data, err := trans.svc.Get(context.Background(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(getResponse{Success: true, Count: data})
}
func (trans *transport) Add(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	req := addRequest{}
	json.NewDecoder(r.Body).Decode(&req)
	err := trans.svc.Add(context.Background(), req.Counter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(addResponse{Success: true})
}

// Service Layer

// Get retrieves the count value based on the ID (for simplicity, ID is not used)
func (svc *Counter) Get(ctx context.Context, id string) (int, error) {
	return svc.Count, nil
}

// Add increments the counter
func (svc *Counter) Add(ctx context.Context, c Counter) error {
	svc.Count += c.Count
	return nil
}

func main() {
	r := mux.NewRouter()
	svc := NewService()
	ts := NewTransport(svc)
	r.HandleFunc("/get", ts.Get).Methods("GET")
	r.HandleFunc("/add", ts.Add).Methods("POST")
	http.ListenAndServe(":5000", r)
}

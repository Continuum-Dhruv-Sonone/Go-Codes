package processor

import (
	"errors"
	"io"
	"reflect"
	"strings"
	"testing"
)

// testReader implement the io.Reader interface
type testReader struct{}

//Read is for implementing the io.Reader getting error scanning
func (testReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("Error Reading")
}

func TestProcessInput(t *testing.T) {

	type args struct {
		f io.Reader
	}
	tests := []struct {
		name       string
		args       args
		wantInputs [][]string
		wantErr    bool
	}{
		{
			"Success",
			args{strings.NewReader("HELLO WINTER IS COMMING")},
			[][]string{{"HELLO", "WINTER", "IS", "COMMING"}},
			false,
		},
		{
			"Error",
			args{testReader{}},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotInputs, err := ProcessInput(tt.args.f)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProcessInput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInputs, tt.wantInputs) {
				t.Errorf("ProcessInput() = %v, want %v", gotInputs, tt.wantInputs)
			}
		})
	}
}

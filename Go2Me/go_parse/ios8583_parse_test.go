package go_parse

import (
	"reflect"
	"testing"
)

func TestDataMessage8583BCD(t *testing.T) {
	type args struct {
		mitStr string
		data   interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DataMessage8583BCD(tt.args.mitStr, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("DataMessage8583BCD() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DataMessage8583BCD() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDataMessage8583ASCII(t *testing.T) {
	type args struct {
		mitStr string
		data   interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DataMessage8583ASCII(tt.args.mitStr, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("DataMessage8583ASCII() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DataMessage8583ASCII() = %v, want %v", got, tt.want)
			}
		})
	}
}

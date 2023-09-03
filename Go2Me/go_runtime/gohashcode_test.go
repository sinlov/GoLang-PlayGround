package go_runtime

import "testing"

func Test_getObjHashCode(t *testing.T) {
	type args struct {
		obj *Obj
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := getObjHashCode(tt.args.obj)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. getObjHashCode() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. getObjHashCode() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

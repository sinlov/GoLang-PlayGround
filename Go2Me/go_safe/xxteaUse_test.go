package go_safe

import (
	"reflect"
	"testing"
)

func TestXByteEncode(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := XByteEncode(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("XByteEncode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXByteDecode(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := XByteDecode(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("XByteDecode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXHexStringEncode(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name:"TestXHexStringEncode",
			args:args{
				data:`{"info":{"result":"0","errorinfo":"\u64cd\u4f5c\u6210\u529f"},"data":{"uid":"2518383@MINIGAME","session":"du48foaifko9vc6ogu373lc2q4","sex":"0","nickname":"S0ZaUzE5MzgwMTI=","status":"0","regionver":"20120704","mail":"tianran5@kuaifazs.com","mobile":"","username":"","mark":"0","account_mark":"1","solidify":"1","initialpwd":"","lastserver":[],"serverlist":[],"roleaccount":"2518383"}}`,
			},
			want:"ab1571a520e7984f7c87678db1a49e2be9a1b5cf79da18fa1e1828130e99cc51e6ced93c47845dcfe0e8201b233482e612898f75ec0af0c964ee121572fb4f038fa4155aef1871a0478c9e81fa991ed383ddca79275354802644af11df1b54302d8f8ab19958c0b2ebbadcb9cf9d7e272ccb3e1055326193b9c4da033418a189ea1bac8ff480da2772bfeb93077ef2c91f445ef66002c6d1da1cd3e07902320f016c26c009b2bb90fcdd95670739c92219fee9839c98be19ebd1ecc3b346c88489dfe021b9382f04770c5a4ad950a080f5046c68402b5f940857702b6d19baf7a5baf6e78446a258a5cfc5def23073f535c941db34f2542cfdc2e323707654dc5b27752974df07fd447a816a7fb10b3a886f0788ccde36de52f8b66bdd788ea999881fc9272f11d6f0c6912781ae0ab14ef0f7a72be5c9eeeb43b19a6c0cb293b15861a6fd55b3810f90f0e28375da8572f1f5f0d2da9e78b90df4a619a828dc787ac4b1e247d659c61a39092e4aab854da7c95e715d18d5b54b65344fcd27550fd10b9ed9110191",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := XHexStringEncode(tt.args.data); got != tt.want {
				t.Errorf("XHexStringEncode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXHexStringDecode(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name:"TestXHexStringDecode",
			args:args{
				data:"ab1571a520e7984f7c87678db1a49e2be9a1b5cf79da18fa1e1828130e99cc51e6ced93c47845dcfe0e8201b233482e612898f75ec0af0c964ee121572fb4f038fa4155aef1871a0478c9e81fa991ed383ddca79275354802644af11df1b54302d8f8ab19958c0b2ebbadcb9cf9d7e272ccb3e1055326193b9c4da033418a189ea1bac8ff480da2772bfeb93077ef2c91f445ef66002c6d1da1cd3e07902320f016c26c009b2bb90fcdd95670739c92219fee9839c98be19ebd1ecc3b346c88489dfe021b9382f04770c5a4ad950a080f5046c68402b5f940857702b6d19baf7a5baf6e78446a258a5cfc5def23073f535c941db34f2542cfdc2e323707654dc5b27752974df07fd447a816a7fb10b3a886f0788ccde36de52f8b66bdd788ea999881fc9272f11d6f0c6912781ae0ab14ef0f7a72be5c9eeeb43b19a6c0cb293b15861a6fd55b3810f90f0e28375da8572f1f5f0d2da9e78b90df4a619a828dc787ac4b1e247d659c61a39092e4aab854da7c95e715d18d5b54b65344fcd27550fd10b9ed9110191",
			},
			want:`{"info":{"result":"0","errorinfo":"\u64cd\u4f5c\u6210\u529f"},"data":{"uid":"2518383@MINIGAME","session":"du48foaifko9vc6ogu373lc2q4","sex":"0","nickname":"S0ZaUzE5MzgwMTI=","status":"0","regionver":"20120704","mail":"tianran5@kuaifazs.com","mobile":"","username":"","mark":"0","account_mark":"1","solidify":"1","initialpwd":"","lastserver":[],"serverlist":[],"roleaccount":"2518383"}}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := XHexStringDecode(tt.args.data); got != tt.want {
				t.Errorf("XHexStringDecode() = %v, want %v", got, tt.want)
			}
		})
	}
}

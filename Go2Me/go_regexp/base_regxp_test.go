package go_regexp

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"regexp"
	"testing"
)

func Test_password(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "password-right",
			args: args{
				text: `AAAAdd111`,
			},
		},
		{
			name: "password-right",
			args: args{
				text: `_Aa111111`,
			},
		},
		{
			name: "password-error",
			args: args{
				text: `http://my.com`,
			},
		},
		{
			name: "password-error",
			args: args{
				text: `123456`,
			},
		},
		{
			name: "password-error",
			args: args{
				text: `aaaaaa`,
			},
		},
	}
	for _, tt := range tests {
		fmt.Printf("test name: %v ", tt.name)
		password(tt.args.text)
	}
}

func TestNoAscii(t *testing.T) {
	convey.Convey("mock TestNoAscii", t, func() {
		// mock
		text := "我的daw123zxc21你的"
		convey.Convey("do TestNoAscii", func() {
			// do
			asscii, err := regexp.Compile("[[:^ascii:]]+")
			if err != nil {
				t.Fatalf("regexp error %v", err)
			}
			convey.Convey("verify TestNoAscii", func() {
				// verify
				res := asscii.ReplaceAllString(text, "")
				fmt.Printf("res -> %v\n", res)
				convey.So(res, convey.ShouldNotEqual, "")
			})
		})
	})
}

func Test_email(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "email",
			args: args{
				text: `123@my.com`,
			},
		},
		{
			name: "not email",
			args: args{
				text: `http://my.com`,
			},
		},
	}
	for _, tt := range tests {
		email(tt.args.text)
	}
}

func Test_chat(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "chat",
			args: args{
				text: `Hello 世界！123 Go.`,
			},
		},
	}
	for _, tt := range tests {
		chat(tt.args.text)
	}
}

func Test_innerStringFind(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "chat",
			args: args{
				text: `Hello 世界！123 Go.`,
			},
		},
	}
	for _, tt := range tests {
		innerStringFind(tt.args.text)
	}
}

func Test_punctuationMark(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "chat",
			args: args{
				text: `Hello 世界！123 Go.`,
			},
		},
	}
	for _, tt := range tests {
		punctuationMark(tt.args.text)
	}
}

func Test_replaceString(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "chat",
			args: args{
				text: `Hello 世界！123 Go.`,
			},
		},
	}
	for _, tt := range tests {
		replaceString(tt.args.text)
	}
}

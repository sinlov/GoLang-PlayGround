package go_regexp

import "testing"

func Test_web(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name:"chat",
			args:args{
				text: `123@kf.com`,
			},
		},
	}
	for _, tt := range tests {
		web(tt.args.text)
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
			name:"chat",
			args:args{
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
			name:"chat",
			args:args{
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
			name:"chat",
			args:args{
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
			name:"chat",
			args:args{
				text: `Hello 世界！123 Go.`,
			},
		},
	}
	for _, tt := range tests {
		replaceString(tt.args.text)
	}
}

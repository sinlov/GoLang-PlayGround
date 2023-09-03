package OptionMode

type Option struct {
	A string
	B string
	C int
}

type OptionFunc func(*Option)

var (
	defaultOption = setDefaultOption()
)

func setDefaultOption() *Option {
	return &Option{
		A: "A",
		B: "B",
		C: 100,
	}
}

func NewOption(opts ...OptionFunc) (opt *Option) {
	opt = defaultOption
	for _, o := range opts {
		o(opt)
	}
	defaultOption = setDefaultOption()
	return
}

//nolint:golint,unused
func newOptionABS(a, b string, c int) *Option {
	return &Option{
		A: a,
		B: b,
		C: c,
	}
}

func WithA(a string) OptionFunc {
	return func(o *Option) {
		o.A = a
	}
}

func WithB(b string) OptionFunc {
	return func(o *Option) {
		o.B = b
	}
}

func WithC(c int) OptionFunc {
	return func(o *Option) {
		o.C = c
	}
}

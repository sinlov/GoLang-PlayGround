package go_reflect

type FooStruct struct {
	Pub string
	prv int
}

type BazStruct struct {
	Pub string
	prv int
}

type BarStruct struct {
	FooStruct FooStruct
	bazStruct BazStruct

	PFooStruct *FooStruct
	pBazStruct *BazStruct

	Pub string
	prv int
}

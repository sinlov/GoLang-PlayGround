.PHONY: test check clean build dist all

install:
	go get -u github.com/mailru/easyjson/...
	go get -u github.com/pquerna/ffjson
	go get -u github.com/ugorji/go/codec/codecgen

jsonGen:
	easyjson --all ebook.go
	ffjson fbook.go

help:
	@echo "need https://github.com/mailru/easyjson"
	@echo "need https://github.com/pquerna/ffjson"
	@echo "need https://github.com/ugorji/go/codec/codecgen"
	@echo "$$ make install     ~>      install gen tools"
	@echo "$$ make jsonGen     ~>      gen json file"
	@echo ""
	@echo "-- more info see Makefile --"
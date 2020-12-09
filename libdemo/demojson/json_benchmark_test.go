package demojson

import (
	"bufio"
	"bytes"
	"encoding/json"
	"strings"
	"testing"

	"github.com/buger/jsonparser"
	jsoniter "github.com/json-iterator/go"
	"github.com/mailru/easyjson"
	"github.com/pquerna/ffjson/ffjson"
	"github.com/ugorji/go/codec"
)

// run bench
// go test -bench=. *

func BenchmarkMarshalStdJson(b *testing.B) {
	book := Book{
		BookId: 12125924,
		Title:  "人类简史-从动物到上帝",
		Author: "尤瓦尔·赫拉利",
		Price:  40.8,
		Hot:    true,
		Weight: 100,
	}

	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(&book)
		if err != nil {
			b.Errorf("BenchmarkMarshalStdJson err: %v", err)
			return
		}
	}
}

func BenchmarkMarshalJsonIterator(b *testing.B) {
	book := Book{
		BookId: 12125924,
		Title:  "人类简史-从动物到上帝",
		Author: "尤瓦尔·赫拉利",
		Price:  40.8,
		Hot:    true,
		Weight: 100,
	}

	var jsonIterator = jsoniter.ConfigCompatibleWithStandardLibrary
	for i := 0; i < b.N; i++ {
		_, err := jsonIterator.Marshal(&book)
		if err != nil {
			b.Errorf("BenchmarkMarshalJsonIterator err: %v", err)
			return
		}
	}
}

func BenchmarkMarshalFfjson(b *testing.B) {
	book := FBook{
		BookId: 12125924,
		Title:  "人类简史-从动物到上帝",
		Author: "尤瓦尔·赫拉利",
		Price:  40.8,
		Hot:    true,
		Weight: 100,
	}

	for i := 0; i < b.N; i++ {
		_, err := ffjson.Marshal(&book)
		if err != nil {
			b.Errorf("BenchmarkMarshalFfjson err: %v", err)
			return
		}
	}
}

func BenchmarkMarshalEasyjson(b *testing.B) {
	book := EBook{
		BookId: 12125924,
		Title:  "人类简史-从动物到上帝",
		Author: "尤瓦尔·赫拉利",
		Price:  40.8,
		Hot:    true,
		Weight: 100,
	}

	for i := 0; i < b.N; i++ {
		_, err := easyjson.Marshal(&book)
		if err != nil {
			b.Errorf("BenchmarkMarshalEasyjson err: %v", err)
			return
		}

	}
}

func BenchmarkMarshalCodecJson(b *testing.B) {
	book := Book{
		BookId: 12125924,
		Title:  "人类简史-从动物到上帝",
		Author: "尤瓦尔·赫拉利",
		Price:  40.8,
		Hot:    true,
		Weight: 100,
	}

	buf := make([]byte, 0, 1024)
	jsonHandler := &codec.JsonHandle{}
	encoder := codec.NewEncoderBytes(&buf, jsonHandler)
	for i := 0; i < b.N; i++ {
		err := encoder.Encode(&book)
		if err != nil {
			b.Errorf("BenchmarkMarshalCodecJson err: %v", err)
			return
		}

	}
}

func BenchmarkMarshalCodecJsonWithBufio(b *testing.B) {
	book := Book{
		BookId: 12125924,
		Title:  "人类简史-从动物到上帝",
		Author: "尤瓦尔·赫拉利",
		Price:  40.8,
		Hot:    true,
		Weight: 100,
	}

	jsonHandler := &codec.JsonHandle{}
	for i := 0; i < b.N; i++ {
		buf := bytes.NewBuffer(make([]byte, 0, 1024))
		writer := bufio.NewWriter(buf)
		encoder := codec.NewEncoder(writer, jsonHandler)
		err := encoder.Encode(&book)
		if err != nil {
			b.Errorf("BenchmarkMarshalCodecJsonWithBufio encoder.Encode err: %v", err)
			return
		}
		err = writer.Flush()
		if err != nil {
			b.Errorf("BenchmarkMarshalCodecJsonWithBufio writer.Flush err: %v", err)
			return
		}
	}
}

func BenchmarkUnMarshalStdJson(b *testing.B) {
	data := []byte(`{"id":12125925,"title":"未来简史-从智人到智神","author":"尤瓦尔·赫拉利","price":40.8,"hot":true}`)
	var book Book

	for i := 0; i < b.N; i++ {
		err := json.Unmarshal(data, &book)
		if err != nil {
			b.Errorf("BenchmarkUnMarshalStdJson err: %v", err)
			return
		}
	}
}

func BenchmarkUnMarshalJsonIterator(b *testing.B) {
	data := []byte(`{"id":12125925,"title":"未来简史-从智人到智神","author":"尤瓦尔·赫拉利","price":40.8,"hot":true}`)
	var book Book

	var jsonIterator = jsoniter.ConfigCompatibleWithStandardLibrary
	for i := 0; i < b.N; i++ {
		err := jsonIterator.Unmarshal(data, &book)
		if err != nil {
			b.Errorf("BenchmarkUnMarshalJsonIterator err: %v", err)
			return
		}
	}
}

func BenchmarkUnMarshalFfjson(b *testing.B) {
	data := []byte(`{"id":12125925,"title":"未来简史-从智人到智神","author":"尤瓦尔·赫拉利","price":40.8,"hot":true}`)
	var book FBook

	for i := 0; i < b.N; i++ {
		err := ffjson.Unmarshal(data, &book)
		if err != nil {
			b.Errorf("BenchmarkUnMarshalFfjson err: %v", err)
			return
		}
	}
}

func BenchmarkUnMarshalEasyjson(b *testing.B) {
	data := []byte(`{"id":12125925,"title":"未来简史-从智人到智神","author":"尤瓦尔·赫拉利","price":40.8,"hot":true}`)
	var book EBook

	for i := 0; i < b.N; i++ {
		err := easyjson.Unmarshal(data, &book)
		if err != nil {
			b.Errorf("BenchmarkUnMarshalEasyjson err: %v", err)
			return
		}
	}
}

func BenchmarkUnMarshalCodecJson(b *testing.B) {
	data := []byte(`{"id":12125925,"title":"未来简史-从智人到智神","author":"尤瓦尔·赫拉利","price":40.8,"hot":true}`)
	var book Book

	jsonHandler := &codec.JsonHandle{}
	decoder := codec.NewDecoderBytes(data, jsonHandler)
	if decoder == nil {
		b.Logf("BenchmarkUnMarshalCodecJson decoder is nil")
		return
	}
	for i := 0; i < b.N; i++ {
		err := decoder.Decode(&book)
		if err != nil {
			b.Logf("BenchmarkUnMarshalCodecJson err: %v", err)
			return
		}
	}
}

func BenchmarkUnMarshalCodecJsonWithBufio(b *testing.B) {
	data := []byte(`{"id":12125925,"title":"未来简史-从智人到智神","author":"尤瓦尔·赫拉利","price":40.8,"hot":true}`)
	var book Book

	jsonHandler := &codec.JsonHandle{}
	for i := 0; i < b.N; i++ {
		reader := bufio.NewReader(strings.NewReader(string(data)))
		decoder := codec.NewDecoder(reader, jsonHandler)
		err := decoder.Decode(&book)
		if err != nil {
			b.Errorf("BenchmarkUnMarshalCodecJsonWithBufio err: %v", err)
			return
		}
	}
}

func BenchmarkUnMarshalJsonparser(b *testing.B) {
	data := []byte(`{"id":12125925,"title":"未来简史-从智人到智神","author":"尤瓦尔·赫拉利","price":40.8,"hot":true}`)
	var book Book

	for i := 0; i < b.N; i++ {
		id, err := jsonparser.GetInt(data, "id")
		if err != nil {
			b.Errorf("BenchmarkUnMarshalJsonparser GetInt id err: %v", err)
			return
		}
		book.BookId = id
		book.Title, err = jsonparser.GetString(data, "title")
		if err != nil {
			b.Errorf("BenchmarkUnMarshalJsonparser GetString title err: %v", err)
			return
		}
		book.Author, err = jsonparser.GetString(data, "author")
		if err != nil {
			b.Errorf("BenchmarkUnMarshalJsonparser GetString author err: %v", err)
			return
		}
		book.Price, err = jsonparser.GetFloat(data, "price")
		if err != nil {
			b.Errorf("BenchmarkUnMarshalJsonparser GetFloat price err: %v", err)
			return
		}
		book.Hot, err = jsonparser.GetBoolean(data, "hot")
		if err != nil {
			b.Errorf("BenchmarkUnMarshalJsonparser GetBoolean hot err: %v", err)
			return
		}
	}
}

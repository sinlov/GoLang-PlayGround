package go_cid

import (
	"github.com/ipfs/go-cid"
	mh "github.com/multiformats/go-multihash"
	"testing"
)

func Test_cid_by_test(t *testing.T) {
	// Create a cid manually by specifying the 'prefix' parameters
	pref := cid.Prefix{
		Version:  0,
		Codec:    cid.DagProtobuf,
		MhType:   mh.SHA2_256,
		MhLength: -1, // default length
	}

	// And then feed it some data
	c, err := pref.Sum([]byte("beep boop"))
	if err != nil {
		t.Fatalf("err, %v", err)
	}
	t.Log(c.String())
}

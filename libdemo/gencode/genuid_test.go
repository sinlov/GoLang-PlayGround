package gencode

import (
	"testing"
)

func TestGenCodeFull(t *testing.T) {
	t.Logf("github.com/lithammer/shortuuid: %s\n", genShortUUID())

	t.Logf("github.com/google/uuid:         %s\n", genUUID())
	t.Logf("github.com/rs/xid:              %s\n", genXid())
	t.Logf("github.com/segmentio/ksuid:     %s\n", genKsuid())
	t.Logf("github.com/kjk/betterguid:      %s\n", genBetterGUID())
	t.Logf("github.com/oklog/ulid:          %s\n", genUlid())

	t.Logf("github.com/chilts/sid:          %s\n", genSid())

	// Note: this is base16, could shorten by encoding as base62 string
	sonyFlake, err := genSonyflake()
	if err != nil {
		t.Errorf("genSonyflake err: %v", err)
	}
	t.Logf("github.com/sony/sonyflake:      %d\n", sonyFlake)

	flake, err := genSnowFlake()
	if err != nil {
		t.Errorf("genSnowFlake err: %v", err)
	}
	snowFlake := flake.Generate()
	t.Logf("github.com/bwmarrin/snowflake Int64   ID:    %d\n", snowFlake)
	t.Logf("github.com/bwmarrin/snowflake String  ID:    %s\n", snowFlake)
	t.Logf("github.com/bwmarrin/snowflake Base2   ID:    %s\n", snowFlake.Base2())
	t.Logf("github.com/bwmarrin/snowflake Base64  ID:    %s\n", snowFlake.Base64())
	t.Logf("github.com/bwmarrin/snowflake Int64   ID:    %d\n", snowFlake.Int64())
	t.Logf("github.com/bwmarrin/snowflake Time:  %d\n", snowFlake.Time())
	t.Logf("github.com/bwmarrin/snowflake Node:  %d\n", snowFlake.Node())
	t.Logf("github.com/bwmarrin/snowflake Step:  %d\n", snowFlake.Step())

	t.Logf("github.com/satori/go.uuid:      %s\n", genUUIDv4().String())
}

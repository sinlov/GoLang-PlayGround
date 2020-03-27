package gencode

import (
	"math/rand"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/chilts/sid"
	guuid "github.com/google/uuid"
	"github.com/kjk/betterguid"
	"github.com/lithammer/shortuuid"
	"github.com/oklog/ulid"
	"github.com/rs/xid"
	"github.com/satori/go.uuid"
	"github.com/segmentio/ksuid"
	"github.com/sony/sonyflake"
)

func genShortUUID() string {
	return shortuuid.New()
}

func genUUID() guuid.UUID {
	return guuid.New()
}

func genXid() xid.ID {
	return xid.New()
}

func genKsuid() ksuid.KSUID {
	return ksuid.New()
}

func genBetterGUID() string {
	return betterguid.New()
}

func genUlid() string {
	t := time.Now().UTC()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	id := ulid.MustNew(ulid.Timestamp(t), entropy)
	return id.String()
}

func genSonyflake() (uint64, error) {
	flake := sonyflake.NewSonyflake(sonyflake.Settings{})
	return flake.NextID()
}

func genSnowFlake() (*snowflake.Node, error) {
	return snowflake.NewNode(1)
}

func genSid() string {
	return sid.Id()
}

func genUUIDv4() uuid.UUID {
	return uuid.NewV4()
}

// Mostly a port of https://github.com/Mediagone/small-uid
package uid

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const max20BitInt = 524288

type UID struct {
	uid int64
}

// Returns a randomly-initialized UID.
func New() *UID {
	return initialize()
}

// Returns a randomly-initialized UID that is guaranteed to not already
// exist in the passed-in map.
func NewWithExisting(existing map[int64]bool) *UID {
	uid := initialize()
	for existing[uid.uid] {
		uid = initialize()
	}
	existing[uid.uid] = true

	return uid
}

func initialize() *UID {
	timeMs := time.Now().UnixMilli()
	rand := rand.Intn(max20BitInt)
	uidStr := fmt.Sprintf("%d%06d", timeMs, rand)

	uid, err := strconv.Atoi(uidStr)
	if err != nil {
		return &UID{uid: -1}
	}

	return &UID{uid: int64(uid)}
}

func (u UID) Int() int64 {
	return u.uid
}

func (u UID) Hex() string {
	return strconv.FormatInt(u.uid, 16)
}

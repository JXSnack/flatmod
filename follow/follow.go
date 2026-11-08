package follow

import (
	"crypto/md5"
	"encoding/hex"
)

type Info struct {
	Config Config
	URL    string
	Branch string
}

func (i Info) Checksum() string {
	nameSum := md5.Sum([]byte(i.URL + i.Branch))
	return hex.EncodeToString(nameSum[:])
}

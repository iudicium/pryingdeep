package favicon

import (
	"bytes"
	"encoding/base64"
	"fmt"

	"github.com/twmb/murmur3"
)

func StandBase64(braw []byte) []byte {
	b64enc := base64.StdEncoding.EncodeToString(braw)
	var buffer bytes.Buffer
	for i := 0; i < len(b64enc); i++ {
		ch := b64enc[i]
		buffer.WriteByte(ch)
		if (i+1)%76 == 0 {
			buffer.WriteByte('\n')
		}
	}

	buffer.WriteByte('\n')
	return buffer.Bytes()
}

func Mmh3Hash32(raw []byte) string {
	var h32 = murmur3.New32()
	h32.Write(raw)
	return fmt.Sprintf("%d", int32(h32.Sum32()))
}

func IconHash(content []byte) string {
	return Mmh3Hash32(StandBase64(content))
}

package password

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"time"
)

// Generate (t time.Time, nm string, s []byte) generates a time dependent password for the specified
// name (e.g. a hostname). The time is truncated to the hour so that the password will remain for for the hour.
func Generate(t time.Time, nm string, s []byte) string {
	layout := "2006-01-02 15"
	ts := t.Format(layout)
	salt := string(s)
	h := md5.New()
	io.WriteString(h, salt)
	io.WriteString(h, ts)
	io.WriteString(h, nm)
	pwdbytes := h.Sum(nil)
	pwdstr := hex.EncodeToString(pwdbytes)
	return pwdstr
}

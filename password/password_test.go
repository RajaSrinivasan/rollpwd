package password

import (
	"log"
	"testing"
	"time"

	"../salt"
)

func TestGenerate(t *testing.T) {
	n := time.Now().UTC()
	s := salt.Generate("admin")
	pwd := Generate(n, "gitlab.com", s)
	layout := "2006-01-02 15"
	ts := n.Format(layout)
	log.Printf("Timestamp: %s Password: %s\n", ts, pwd)
}

func testWithDrift(t time.Time, driftmins int) {
	s := salt.Generate("admin")

	tahead := t.Add(time.Duration(driftmins) * time.Minute)
	tbehind := t.Add(time.Duration(-driftmins) * time.Minute)

	pwd := Generate(t, "gitlab.com", s)
	pwdahead := Generate(tahead, "gitlab.com", s)
	pwdbehind := Generate(tbehind, "gitlab.com", s)
	layout := "2006-01-02 15"
	ts := t.Format(layout)
	log.Printf("Timestamp: %s\n> Password: %s\n> Ahead:    %s\n> Behind:   %s\n", ts, pwd, pwdahead, pwdbehind)

}
func TestGenerateDrift(t *testing.T) {

	driftmins := 31

	log.Println("Current time and ahead fall in the same hour")
	n := time.Date(2019, 11, 01, 2, 28, 0, 0, time.UTC)
	testWithDrift(n, driftmins)

	log.Println("Current time and behind fall in the same hour")
	n = time.Date(2019, 11, 01, 2, 32, 0, 0, time.UTC)
	testWithDrift(n, driftmins)
}

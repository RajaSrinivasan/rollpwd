package main

import (
	"log"
	"os"
	"strings"
	"time"

	"./password"
	"./salt"
	"github.com/spf13/pflag"
)

var verbose bool
var username string
var hostname string

const driftmins = 30

var authenticate string

func echoArguments() {
	log.Printf("User %s for host %s\n", username, hostname)
	if len(authenticate) > 0 {
		log.Printf("Will authenticate the password %s\n", authenticate)
	} else {
		log.Printf("Will generate a password")
	}
}
func processCommandLine() {
	pflag.BoolVar(&verbose, "verbose", false, "high verbosity")
	pflag.StringVar(&authenticate, "authenticate", "", "authenticate this password")
	pflag.Parse()

	if len(pflag.Args()) != 1 {
		log.Println("A user id in the form username@hostname is required.")
		os.Exit(1)
	}
	userid := pflag.Arg(0)
	usps := strings.Split(userid, "@")
	username = usps[0]
	hostname = usps[1]

	if verbose {
		echoArguments()
	}

}

func main() {

	processCommandLine()
	usalt := salt.Generate(username)
	nowtime := time.Now()
	pwdnow := password.Generate(nowtime, hostname, usalt)
	if len(authenticate) == 0 {
		log.Printf("%s\n", pwdnow)
		os.Exit(0)
	}

	// Check the password for this hour
	if pwdnow == authenticate {
		log.Printf("Authenticated\n")
		os.Exit(0)
	}

	// In case this host's clock is behind and hence the real time is ahead check
	// the password against the presumably real time
	tahead := nowtime.Add(time.Duration(driftmins) * time.Minute)
	pwdahead := password.Generate(tahead, hostname, usalt)
	if pwdahead == authenticate {
		log.Printf("Authenticated")
		os.Exit(0)
	}

	// Alternatively the host's clock has drifted ahead and the real time presumably is
	// behind.
	tbehind := nowtime.Add(time.Duration(-driftmins) * time.Minute)
	pwdbehind := password.Generate(tbehind, hostname, usalt)
	if pwdbehind == authenticate {
		log.Printf("Authenticated\n")
		os.Exit(0)
	}
	log.Printf("Not authenticated\n")
	os.Exit(1)

}

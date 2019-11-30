SOURCES=$(wildcard *.go)
EXEC=rollpwd
BINARIES=../bin

all: $(SOURCES)
	GOOS=darwin GOARCH=amd64 GOARM=6 go build -o $(BINARIES)/mac/$(EXEC) $(SOURCES)
	GOOS=windows GOARCH=amd64 GOARM=6 go build -o $(BINARIES)/win64/$(EXEC).exe $(SOURCES)	
	GOOS=linux GOARCH=amd64 GOARM=6 go build -o $(BINARIES)/linux64/$(EXEC) $(SOURCES)
	zip rollpwd.zip -r $(BINARIES)

clean:
	$(RM) $(BINARIES)/mac/$(EXEC)
	$(RM) $(BINARIES)/win64/$(EXEC)
	$(RM) $(BINARIES)/linux64/$(EXEC)

dependencies:
	go get -u -v github.com/spf13/pflag
include $(GOROOT)/src/Make.inc

TARG=mp3
CGOFILES=\
	music.go

include $(GOROOT)/src/Make.pkg

main: install main.go
	$(GC) main.go
	$(LD) -o $@ main.$O

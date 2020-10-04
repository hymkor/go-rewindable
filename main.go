package rewindable

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
)

type Reader struct {
	sourceReader  io.Reader
	pastBuffer    []byte
	tmpBuffer     bytes.Buffer
	currentReader io.Reader
}

func (this *Reader) Read(b []byte) (int, error) {
	return this.currentReader.Read(b)
}

func NewReader(sourceReader io.Reader) *Reader {
	this := &Reader{}
	this.sourceReader = sourceReader
	this.currentReader = io.TeeReader(this.sourceReader, &this.tmpBuffer)
	return this
}

func (this *Reader) tmp2past() {
	if this.tmpBuffer.Len() > 0 {
		this.pastBuffer = append(this.pastBuffer, this.tmpBuffer.Bytes()...)
		this.tmpBuffer.Reset()
	}
}

func (this *Reader) Rewind() {
	this.tmp2past()
	if this.pastBuffer == nil && len(this.pastBuffer) <= 0 {
		this.currentReader = io.TeeReader(this.sourceReader, &this.tmpBuffer)
	} else {
		this.currentReader = io.MultiReader(
			bytes.NewReader(this.pastBuffer),
			io.TeeReader(this.sourceReader, &this.tmpBuffer))
	}
}

func (this *Reader) seekStart(pos int64) (int64, error) {
	alreadyReadBytes := int64(len(this.pastBuffer) + this.tmpBuffer.Len())
	if pos < alreadyReadBytes {
		this.tmp2past()
		r1 := bytes.NewReader(this.pastBuffer)
		r2 := io.TeeReader(this.sourceReader, &this.tmpBuffer)
		move, err := r1.Seek(pos, io.SeekStart)
		this.currentReader = io.MultiReader(r1, r2)
		return move, err
	} else if pos > alreadyReadBytes {
		move, err := io.CopyN(ioutil.Discard, this.currentReader, pos-alreadyReadBytes)
		return alreadyReadBytes + move, err
	} else {
		return alreadyReadBytes, nil
	}
}

func (this *Reader) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	default:
		return 0, fmt.Errorf("not support: %d", whence)
	case io.SeekStart:
		return this.seekStart(offset)
	}
}

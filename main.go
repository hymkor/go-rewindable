package rewindable

import (
	"bytes"
	"io"
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

func (this *Reader) Rewind() {
	if this.tmpBuffer.Len() > 0 {
		this.pastBuffer = append(this.pastBuffer, this.tmpBuffer.Bytes()...)
		this.tmpBuffer.Reset()
	}
	if this.pastBuffer == nil && len(this.pastBuffer) <= 0 {
		this.currentReader = io.TeeReader(this.sourceReader, &this.tmpBuffer)
	} else {
		this.currentReader = io.MultiReader(
			bytes.NewReader(this.pastBuffer),
			io.TeeReader(this.sourceReader, &this.tmpBuffer))
	}
}

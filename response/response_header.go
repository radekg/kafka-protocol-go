package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/pkg/errors"
	"github.com/radekg/kafka-protocol-go/decoder"
	"github.com/radekg/kafka-protocol-go/encoder"
	protoerrors "github.com/radekg/kafka-protocol-go/errors"
	"github.com/radekg/kafka-protocol-go/request"
)

type ResponseHeader struct {
	Length        int32
	CorrelationID int32
}

func (r *ResponseHeader) Decode(pd decoder.DDef) (err error) {
	r.Length, err = pd.GetInt32()
	if err != nil {
		return err
	}
	if r.Length <= 4 {
		return protoerrors.PacketDecodingError{fmt.Sprintf("message of length %d too small", r.Length)}
	}

	r.CorrelationID, err = pd.GetInt32()
	return err
}

func (r *ResponseHeader) Encode(pe encoder.EDef) (err error) {
	pe.PutInt32(r.Length)
	pe.PutInt32(r.CorrelationID)
	return nil
}

type ResponseHeaderTaggedFields struct {
	version int16
}

type ByteReader struct {
	io.Reader
}

func (r ByteReader) ReadByte() (byte, error) {
	b := make([]byte, 1)
	n, err := r.Reader.Read(b)
	if err != nil {
		return 0, err
	}
	if n != 1 {
		return 0, errors.New("should read one byte")
	}
	return b[0], nil
}

func (r *ResponseHeaderTaggedFields) MaybeRead(reader io.Reader) ([]byte, error) {
	if r.version >= 1 {
		var bufferRead bytes.Buffer
		reader = io.TeeReader(reader, &bufferRead)

		byteReader := ByteReader{Reader: reader}
		numTaggedFields, err := binary.ReadUvarint(byteReader)
		if err != nil {
			return nil, errors.Wrap(err, "error while reading tagged field numTaggedFields")
		}
		for i := 0; i < int(numTaggedFields); i++ {
			// read tag
			if _, err := binary.ReadUvarint(byteReader); err != nil {
				return nil, errors.Wrap(err, "error while reading tagged field tag")
			}
			// number of data bytes
			size, err := binary.ReadUvarint(byteReader)
			if err != nil {
				return nil, errors.Wrap(err, "error while reading tagged field size")
			}
			if size < 0 {
				return nil, errors.New("negative size of tagged field data")
			} else if size == 0 {
				continue
			} else {
				if _, err := io.CopyN(ioutil.Discard, reader, int64(size)); err != nil {
					return nil, errors.Wrap(err, "error while reading tagged field data")
				}
			}
		}
		return bufferRead.Bytes(), nil
	}
	return make([]byte, 0), nil
}

func NewResponseHeaderTaggedFields(req *request.RequestKeyVersion) (*ResponseHeaderTaggedFields, error) {
	version := req.ResponseHeaderVersion()
	if version < 0 {
		return nil, protoerrors.PacketDecodingError{fmt.Sprintf("ResponseHeaderVersion: unknown ApiKey %d", req.ApiKey)}
	}
	return &ResponseHeaderTaggedFields{version: version}, nil
}

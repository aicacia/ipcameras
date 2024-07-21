package format

import (
	"bufio"
	"encoding/binary"
	"io"
	"os"

	"github.com/deepch/vdk/av"
)

type Demuxer struct {
	idx          int8
	codec        av.CodecData
	file         *os.File
	headerOffset int64
	scanner      *bufio.Scanner
}

const maxCapacity = 1024 * 1024 * 8

func NewDemuxer(file *os.File, idx int8) (*Demuxer, error) {
	var codec av.CodecData
	scanner := NewDelimScanner(file, Delim)
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	var headerOffset int64
	if scanner.Scan() {
		bytes := scanner.Bytes()
		err := FromBytes(bytes, &codec)
		if err != nil {
			return nil, err
		}
		headerOffset = int64(len(bytes))
	} else {
		return nil, scanner.Err()
	}

	return &Demuxer{
		idx:          idx,
		file:         file,
		codec:        codec,
		scanner:      scanner,
		headerOffset: headerOffset,
	}, nil
}

func (d *Demuxer) Codec() av.CodecData {
	return d.codec
}

func (d *Demuxer) SeekToTime(seekMicroseconds int64) (*av.Packet, error) {
	d.file.Seek(d.headerOffset, io.SeekStart)
	for d.scanner.Scan() {
		scannerBytes := d.scanner.Bytes()

		if len(scannerBytes) > 0 {
			// skip non key frames
			if scannerBytes[0] != 1 {
				continue
			}
			timestamp := int64(binary.BigEndian.Uint64(scannerBytes[9:17]))

			if seekMicroseconds >= timestamp {
				return ReadPacket(d.idx, scannerBytes), nil
			}
		}
	}
	return nil, io.EOF
}

func (d *Demuxer) ReadPacket(direction int8) (*av.Packet, error) {
	if d.scanner.Scan() {
		scannerBytes := d.scanner.Bytes()

		if len(scannerBytes) > 0 {
			return ReadPacket(d.idx, scannerBytes), nil
		}
	}
	return nil, io.EOF
}

func (d *Demuxer) Close() (err error) {
	return d.file.Close()
}

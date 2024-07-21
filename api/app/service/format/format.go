package format

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"io"
	"time"

	"github.com/deepch/vdk/av"
	"github.com/deepch/vdk/codec"
	"github.com/deepch/vdk/codec/h264parser"
	"github.com/deepch/vdk/codec/h265parser"
	"github.com/deepch/vdk/format"
)

var Delim = []byte{'\n', '\n', '\n', '\n'}

func ToBytes(e any) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(e)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil

}

func FromBytes(b []byte, e any) error {
	var buf bytes.Buffer
	_, err := buf.Write(b)
	if err != nil {
		return err
	}
	dec := gob.NewDecoder(&buf)
	return dec.Decode(e)
}

func NewDelimScanner(r io.Reader, delim []byte) *bufio.Scanner {
	scanner := bufio.NewScanner(r)
	scanner.Split(createScanLines(delim))
	return scanner
}

func createScanLines(delim []byte) bufio.SplitFunc {
	return func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}
		for i := 0; i+len(delim) <= len(data); {
			j := i + bytes.IndexByte(data[i:], delim[0])
			if j < i {
				break
			}
			if bytes.Equal(data[j+1:j+len(delim)], delim[1:]) {
				return j + len(delim), data[0:j], nil
			}
			i = j + 1
		}
		if atEOF {
			return len(data), data, nil
		}
		return 0, nil, nil
	}
}

func SetPacketTime(packet *av.Packet) *av.Packet {
	packet.Time = time.Duration(time.Now().UTC().UnixNano())
	return packet
}

func TimeFromDuration(d time.Duration) time.Time {
	return time.UnixMicro(d.Microseconds()).UTC()
}

func GetPacketTime(packet *av.Packet) time.Time {
	return TimeFromDuration(packet.Time)
}

func WritePacket(w io.Writer, pkt *av.Packet) error {
	bytes := make([]byte, 17)

	if pkt.IsKeyFrame {
		bytes[0] = 1
	} else {
		bytes[0] = 0
	}

	binary.BigEndian.PutUint64(bytes[1:9], uint64(pkt.Duration))
	binary.BigEndian.PutUint64(bytes[9:17], uint64(pkt.Time))

	if _, err := w.Write(bytes); err != nil {
		return err
	}
	if _, err := w.Write(pkt.Data); err != nil {
		return err
	}
	_, err := w.Write(Delim)
	return err
}

func ReadPacket(idx int8, bytes []byte) *av.Packet {
	isKeyFrame := false
	if bytes[0] == 1 {
		isKeyFrame = true
	}
	duration := time.Duration(int64(binary.BigEndian.Uint64(bytes[1:9])))
	timestamp := int64(binary.BigEndian.Uint64(bytes[9:17]))

	return &av.Packet{
		IsKeyFrame:      isKeyFrame,
		Idx:             idx,
		CompositionTime: 0,
		Time:            time.Duration(timestamp),
		Duration:        duration,
		Data:            bytes[17:],
	}
}

func init() {
	format.RegisterAll()

	gob.Register(h264parser.CodecData{})
	gob.Register(h265parser.CodecData{})
	gob.Register(codec.PCMUCodecData{})
}

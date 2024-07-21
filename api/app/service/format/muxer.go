package format

import (
	"fmt"
	"os"

	"github.com/deepch/vdk/av"
)

type Muxer struct {
	folder string
	files  []*os.File
}

func NewMuxer(folder string) (*Muxer, error) {
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		err := os.MkdirAll(folder, os.ModePerm)
		if err != nil {
			return nil, err
		}
	}
	return &Muxer{folder: folder, files: nil}, nil
}

func (element *Muxer) WriteHeader(codecs []av.CodecData) error {
	files := make([]*os.File, len(codecs))

	for idx, codec := range codecs {
		file, err := os.Create(element.folder + fmt.Sprintf("/%d", idx))
		if err != nil {
			return err
		}
		codecBytes, err := ToBytes(codec)
		if err != nil {
			return err
		}
		if _, err := file.Write(append(codecBytes, Delim...)); err != nil {
			return err
		}
		files[idx] = file
	}
	element.files = files
	return nil

}

func (element *Muxer) WritePacket(pkt *av.Packet) (err error) {
	file := element.files[pkt.Idx]
	return WritePacket(file, pkt)
}

func (element *Muxer) Close() (err error) {
	for _, file := range element.files {
		err = file.Close()
	}
	return err
}

package huffman_compression

import (
	"bytes"
	"encoding/binary"
	"io"
	"io/ioutil"
)

type decoder struct {
	codes map[int]string
	r     io.Reader
	w     io.Writer
}

func (e *decoder) decode() (err error) {
	err = e.readHuffmanHeader()
	if err != nil {
		return err
	}
	//datas,err := e.readHuffmanDatas()
	//if err != nil {
	//	return err
	//}
	return
}

func (e *decoder) readHuffmanHeader() error {
	bts, err := ioutil.ReadAll(e.r)
	if err != nil {
		return err
	}

	buf := bytes.NewBuffer(bts)
	header := new(huffmanTreeHeader)
	err = binary.Read(buf, binary.BigEndian, &header.eLen)
	if err != nil {
		return err
	}
	for i := 0; i < int(header.eLen); i++ {
		elem := huffmanTreeHeaderElem{}
		err := binary.Read(buf, binary.BigEndian, &elem.k)
		if err != nil {
			return err
		}
		err = binary.Read(buf, binary.BigEndian, &elem.vsLen)
		if err != nil {
			return err
		}
		elem.vs = make([]byte, elem.vsLen)
		err = binary.Read(buf, binary.BigEndian, &elem.vs)
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *decoder) readHuffmanDatas() (data []byte, err error) {
	return nil, nil
}

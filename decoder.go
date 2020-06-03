package huffman_compression

import (
	"bytes"
	"encoding/binary"
	"io"
	"io/ioutil"
	"strconv"
)

type decoder struct {
	codes map[int]string
	r     io.Reader
	bs []byte
	w     io.Writer
}

func (e *decoder) decode() (err error) {
	err = e.readHuffmanHeader()
	if err != nil {
		return err
	}
	err = e.readHuffmanDatas()
	if err != nil {
		return err
	}
	return
}

func (e *decoder) readHuffmanHeader() error {
	e.codes = map[int]string{}
	bts, err := ioutil.ReadAll(e.r)
	if err != nil {
		return err
	}
	len := 0
	buf := bytes.NewBuffer(bts)
	header := new(huffmanTreeHeader)
	err = binary.Read(buf, binary.BigEndian, &header.eLen)
	if err != nil {
		return err
	}
	len += 16
	for i := 0; i < int(header.eLen); i++ {
		elem := huffmanTreeHeaderElem{}
		err := binary.Read(buf, binary.BigEndian, &elem.k)
		if err != nil {
			return err
		}
		len += 8
		err = binary.Read(buf, binary.BigEndian, &elem.vsLen)
		if err != nil {
			return err
		}
		len += 8
		elem.vs = make([]byte, elem.vsLen)
		err = binary.Read(buf, binary.BigEndian, &elem.vs)
		if err != nil {
			return err
		}
		len = len + int(elem.vsLen*8)

		header.elem = append(header.elem, elem)
		e.codes[int(elem.k)] = string(elem.vs)
	}
	e.bs = bts[len+1:]
	return nil
}

func (e *decoder) readHuffmanDatas() (err error) {
	s:=""
	for _,b:=range e.bs{
		s+=strconv.FormatUint(uint64(b),2)
	}

	return nil
}

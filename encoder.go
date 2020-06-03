package huffman_compression

import (
	"bytes"
	"encoding/binary"
	"errors"
	"github.com/godaner/func/tree"
	huffmantree "github.com/godaner/func/tree/huffman"
	"github.com/icza/bitio"
	"io"
	"io/ioutil"
	"strconv"
)

type encoder struct {
	huffmanTree tree.HuffmanTree
	codes       map[int]string
	r           io.Reader
	w           io.Writer
}

func (e *encoder) encode() (err error) {

	datas, err := ioutil.ReadAll(e.r)
	if err != nil {
		return err
	}
	err = e.buildHuffmanTree(datas)
	if err != nil {
		return err
	}
	err = e.writeHuffmanHeader()
	if err != nil {
		return err
	}
	err = e.writeHuffmanDatas(datas)
	if err != nil {
		return err
	}
	return
}

// writeHuffmanHeader
func (e *encoder) writeHuffmanHeader() (err error) {
	header := e.buildHuffmanTreeHeader()
	buf := new(bytes.Buffer)
	err = binary.Write(buf, binary.BigEndian, header.eLen)
	if err != nil {
		return err
	}
	for _, elem := range header.elem {
		err = binary.Write(buf, binary.BigEndian, elem.k)
		if err != nil {
			return err
		}
		err = binary.Write(buf, binary.BigEndian, elem.vsLen)
		if err != nil {
			return err
		}
		err = binary.Write(buf, binary.BigEndian, elem.vs)
		if err != nil {
			return err
		}
	}
	_, err = e.w.Write(buf.Bytes())
	if err != nil {
		return nil
	}
	return nil
}

func (e *encoder) buildHuffmanTreeHeader() huffmanTreeHeader {
	elems := make([]huffmanTreeHeaderElem, 0)
	for signal, code := range e.codes {
		vs := []byte{}
		for _, c := range code {
			vs = append(vs, byte(c))
		}
		elems = append(elems, huffmanTreeHeaderElem{
			k:     byte(signal),
			vsLen: byte(len(vs)),
			vs:    vs,
		})
	}
	return huffmanTreeHeader{
		eLen: uint16(len(elems)),
		elem: elems,
	}
}

// buildHuffmanTree
func (e *encoder) buildHuffmanTree(datas []byte) (err error) {
	wdMap := map[int]huffmantree.WeightData{}
	for _, s := range datas {
		d := int(s)
		r, ok := wdMap[d]
		if !ok {
			r = huffmantree.WeightData{
				Data:   d,
				Weight: 0,
			}
		}
		r.Weight++
		wdMap[d] = r
	}
	wds := make([]huffmantree.WeightData, 0)
	for _, wd := range wdMap {
		wds = append(wds, wd)
	}
	if len(wds) == 0 {
		return errors.New("nil data")
	}
	e.huffmanTree = huffmantree.Build(wds...)
	if e.huffmanTree == nil {
		return errors.New("nil tree")
	}
	e.codes = e.huffmanTree.Codes()
	return nil
}

func (e *encoder) writeHuffmanDatas(datas []byte) error {
	bw := bitio.NewWriter(e.w)
	defer bw.Close()
	for _, d := range datas {
		code := e.codes[int(d)]
		r, err := strconv.ParseUint(code, 2, 64)
		if err != nil {
			return err
		}
		err = bw.WriteBits(r, uint8(len(code)))
		if err != nil {
			return err
		}
	}

	return nil
}

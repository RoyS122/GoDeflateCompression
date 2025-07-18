package GoDeflateCompression

import (
	"bufio"
	"bytes"
	"fmt"
)

func FullCompression(s string) ([]byte, []byte, int, bool) {
	var LZCompressed []byte
	var useLZ77 bool
	var ser_tree bytes.Buffer

	if len(s) < 100000 {
		LZCompressed = serializeLZTuples(lz77Compression(s, len(s)/2))
		if len(LZCompressed) < len(s) {
			useLZ77 = true
		}
	} else {
		LZCompressed = serializeLZTuples(lz77Compression(s, 4096))
		if len(LZCompressed) < len(s) {
			useLZ77 = true
		}
	}

	if useLZ77 {
		data, tree, count := compressTextIntoBinary(LZCompressed)

		writer := bufio.NewWriter(&ser_tree)
		serializeTree(tree, writer)
		writer.Flush()
		return data, ser_tree.Bytes(), count, useLZ77
	}
	data, tree, count := compressTextIntoBinary([]byte(s))

	writer := bufio.NewWriter(&ser_tree)
	serializeTree(tree, writer)
	writer.Flush()
	return data, ser_tree.Bytes(), count, useLZ77

}

func FullDecompression(bin []byte, tree []byte, totalChars int, usedLZ bool) (final string) {
	treeReader := bufio.NewReader(bytes.NewReader(tree))

	deserialized_tree, err := deserializeTree(treeReader)
	if err != nil {
		fmt.Printf("Erreur désérialisation arbre : %v\n", err)
		return ""
	}
	decompressedBinary := decompress(bin, deserialized_tree, totalChars)

	if usedLZ {
		// Décompression LZ après Huffman
		tuples := deserializeLZTuples(decompressedBinary)
		final = lz77Decompression(tuples)
	} else {
		final = string(decompressedBinary)
	}
	return final
}

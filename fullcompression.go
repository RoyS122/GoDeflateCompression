package main

func FullCompression(s string) ([]byte, *Node, int, bool) {
	var LZCompressed []byte
	var useLZ77 bool

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
		return data, tree, count, true
	}
	data, tree, count := compressTextIntoBinary([]byte(s))
	return data, tree, count, false
}

func FullDecompression(bin []byte, tree *Node, totalChars int, usedLZ bool) (final string) {
	decompressedBinary := decompress(bin, tree, totalChars)

	if usedLZ {
		// Décompression LZ après Huffman
		tuples := deserializeLZTuples(decompressedBinary)
		final = lz77Decompression(tuples)
	} else {
		final = string(decompressedBinary)
	}
	return final
}

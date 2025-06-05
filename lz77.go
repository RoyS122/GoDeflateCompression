package GoDeflateCompression

import (
	"bytes"
	"encoding/binary"
)

type LZTuples struct {
	Offset   uint16
	Length   uint8
	NextChar byte
}

func lz77Compression(str string, windowSize int) (result []LZTuples) {
	// taille de la fenêtre de recherche
	i := 0

	for i < len(str) {
		matchOffset := 0
		matchLength := 0
		var nextChar byte

		start := i - windowSize
		if start < 0 {
			start = 0
		}
		window := str[start:i]

		// Recherche de la correspondance la plus longue dans la fenêtre
		for j := 0; j < len(window); j++ {
			k := 0
			for i+k < len(str) && j+k < len(window) && str[i+k] == window[j+k] {
				k++
			}
			if k > matchLength {
				matchLength = k
				matchOffset = len(window) - j
			}
		}

		if i+matchLength < len(str) {
			nextChar = str[i+matchLength]
		} else {
			nextChar = 0 // indique la fin de la chaîne
		}

		// Format: (offset,length,nextChar)
		result = append(result, LZTuples{uint16(matchOffset), uint8(matchLength), nextChar})

		i += matchLength + 1
	}

	return result
}

func serializeLZTuples(tuples []LZTuples) []byte {
	buf := new(bytes.Buffer)
	for _, t := range tuples {
		binary.Write(buf, binary.BigEndian, t.Offset)
		binary.Write(buf, binary.BigEndian, t.Length)
		binary.Write(buf, binary.BigEndian, t.NextChar)
	}
	return buf.Bytes()
}

func deserializeLZTuples(data []byte) []LZTuples {
	var tuples []LZTuples
	buf := bytes.NewReader(data)

	for buf.Len() >= 4 {
		var offset uint16
		var length uint8
		var nextChar byte

		binary.Read(buf, binary.BigEndian, &offset)
		binary.Read(buf, binary.BigEndian, &length)
		binary.Read(buf, binary.BigEndian, &nextChar)

		tuples = append(tuples, LZTuples{offset, length, nextChar})
	}
	return tuples
}

func lz77Decompression(compressed []LZTuples) string {
	var result []byte

	for _, tuple := range compressed {
		start := len(result) - int(tuple.Offset)
		if int(tuple.Offset) > len(result) {
			continue // ou log erreur
		}
		for i := 0; i < int(tuple.Length); i++ {
			result = append(result, result[start+i])
		}
		if tuple.NextChar != 0 { // 0 indique fin de texte
			result = append(result, tuple.NextChar)
		}
	}

	return string(result)
}

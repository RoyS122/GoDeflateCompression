package GoDeflateCompression

import (
	"testing"
)

func TestCompressDecompress(t *testing.T) {
	// Texte original à compresser
	// original := []byte("hello hello hello compression test qjqqdshsssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssdqvdqsjkdhnqsldqlskdhjkqsbjdqjskdkjqskhdkjbqsjkdhkqsgdqhsdhqnskjdbqsdnjkqbkdhqsdybq jdbjqsdlkqshkjdgqsdoiqgsdkjqsoihdhvbqsjkdjqsvdjnqslkdbhvzbaohugzduzhhqisjdlqbsjdvbjksbdjqsdoih")
	og := "IL FAUT QUE J'ecrive ce putain de test à la main sinon ces mongoles n'y arrive pas"
	a, b, c, d := FullCompression(og)

	decompressed := FullDecompression(a, b, c, d)

	// Vérifie que le texte décompressé est identique à l'original
	if og != decompressed {
		t.Errorf("texte décompressé différent de l'original:\noriginal:    %s\ndécompressé: %s",
			string(og), string(decompressed))
	}
}

// func TestTreeSerializationSymmetry(t *testing.T) {
// 	data := []byte("hello world, hello compression")

// 	_, root, _ := compressTextIntoBinary(data)

// 	var buf bytes.Buffer
// 	writer := bufio.NewWriter(&buf)

// 	err := serializeTree(root, writer)
// 	if err != nil {
// 		t.Fatalf("Erreur lors de la sérialisation : %v", err)
// 	}
// 	writer.Flush()

// 	deserializedRoot, err := deserializeTree(bufio.NewReader(&buf))
// 	if err != nil {
// 		t.Fatalf("Erreur lors de la désérialisation : %v", err)
// 	}

// 	compressed, _, length := compressTextIntoBinary(data)
// 	result := decompress(compressed, deserializedRoot, length)

// 	if string(result) != string(data) {
// 		t.Errorf("La décompression après désérialisation a échoué.\nAttendu : %s\nReçu : %s", string(data), string(result))
// 	}
// }

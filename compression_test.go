package GoDeflateCompression

import (
	"bufio"
	"bytes"
	"testing"
)

func TestCompressDecompress(t *testing.T) {
	// Texte original à compresser
	original := []byte("hello hello hello compression test")

	// Compression
	compressed, tree, originalLen := compressTextIntoBinary(original)

	// Sérialisation de l'arbre (pour simuler un vrai scénario)
	var buf bytes.Buffer
	err := serializeTree(tree, &buf)
	if err != nil {
		t.Fatalf("erreur lors de la sérialisation de l'arbre: %v", err)
	}

	// Désérialisation
	treeReader := bufio.NewReader(&buf)
	deserializedTree, err := deserializeTree(treeReader)
	if err != nil {
		t.Fatalf("erreur lors de la désérialisation de l'arbre: %v", err)
	}

	// Décompression
	decompressed := decompress(compressed, deserializedTree, originalLen)

	// Vérifie que le texte décompressé est identique à l'original
	if !bytes.Equal(original, decompressed) {
		t.Errorf("texte décompressé différent de l'original:\noriginal:    %s\ndécompressé: %s",
			string(original), string(decompressed))
	}
}

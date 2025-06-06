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


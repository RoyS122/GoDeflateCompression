package main

import (
	"bufio"
	"fmt"
	"io"
)

type Node struct {
	b     byte
	freq  int
	left  *Node
	right *Node
}

type PriorityQueue []*Node

type Code struct {
	bits  uint32 // Le code Huffman binaire (ex: 101 -> bits = 0b101)
	depth int    // Le nombre de bits significatifs (ex: 3)
}

func buildHuffmanTree(freq map[byte]int) *Node {
	pq := PriorityQueue{}
	for i, k := range freq {
		var inserted bool
		for ind, node := range pq {
			if node.freq > k {
				continue
			}
			inserted = true
			var n Node
			n.freq = k
			n.b = i
			pq = append(pq[:ind], append([]*Node{&n}, pq[ind:]...)...)

			break
		}
		if !inserted {
			var n Node
			n.freq = k
			n.b = i
			pq = append(pq, &n)
		}
	}

	for len(pq) > 1 {
		// Prend les deux plus petits
		left := pq[0]
		right := pq[1]
		pq = pq[2:]

		// Crée un nouveau nœud parent
		parent := &Node{
			freq:  left.freq + right.freq,
			left:  left,
			right: right,
		}

		// Réinsère le nouveau nœud dans la file triée
		inserted := false
		for i, node := range pq {
			if parent.freq < node.freq {
				pq = append(pq[:i], append([]*Node{parent}, pq[i:]...)...)
				inserted = true
				break
			}
		}
		if !inserted {
			pq = append(pq, parent)
		}
	}

	// La racine de l’arbre est le seul élément restant
	if len(pq) == 1 {
		return pq[0]
	}
	return nil

}

func createFreqMap(data []byte) map[byte]int {
	m := make(map[byte]int)
	for _, k := range data {

		m[k] += 1
	}
	return m
}

func buildCodeTable(node *Node) map[byte]Code {
	codes := make(map[byte]Code)

	var walk func(n *Node, currentBits uint32, depth int)
	walk = func(n *Node, currentBits uint32, depth int) {
		if n == nil {
			return
		}

		// Si c'est une feuille, on ajoute le code dans la map
		if n.left == nil && n.right == nil {
			codes[n.b] = Code{
				bits:  currentBits,
				depth: depth,
			}
			return
		}

		// Parcours à gauche : ajoute un 0
		walk(n.left, currentBits<<1, depth+1)

		// Parcours à droite : ajoute un 1
		walk(n.right, (currentBits<<1)|1, depth+1)
	}

	// Lancement du parcours
	walk(node, 0, 0)
	return codes
}

func compressTextIntoBinary(s []byte) ([]byte, *Node, int) {
	freqMap := createFreqMap(s)
	t := buildHuffmanTree(freqMap)
	table := buildCodeTable(t)

	var bytes []byte
	var currentByte byte = 0
	var bitCount int = 0

	for _, r := range s {
		code, ok := table[r]
		if !ok {
			continue // ou panic(fmt.Sprintf("caractère %c non trouvé", r))
		}

		for i := code.depth - 1; i >= 0; i-- {
			bit := (code.bits >> i) & 1
			currentByte = (currentByte << 1) | byte(bit)
			bitCount++

			if bitCount == 8 {
				bytes = append(bytes, currentByte)
				currentByte = 0
				bitCount = 0
			}
		}
	}

	// Si le dernier byte n'est pas complet, on le remplit avec des 0 (padding à droite)
	if bitCount > 0 {
		currentByte <<= (8 - bitCount)
		bytes = append(bytes, currentByte)
	}

	return bytes, t, len(s)
}

func decompress(binaryData []byte, root *Node, totalChars int) []byte {
	var result []byte
	current := root
	bitCount := 0
	charsDecoded := 0

	for _, b := range binaryData {
		for i := 7; i >= 0; i-- { // on lit bit par bit du MSB au LSB
			bit := (b >> i) & 1
			if bit == 0 {
				current = current.left
			} else {
				current = current.right
			}

			if current.left == nil && current.right == nil {
				// Feuille atteinte
				result = append(result, current.b)
				charsDecoded++
				if charsDecoded == totalChars {
					return result // Tous les caractères décodés
				}
				current = root // reset à la racine
			}
			bitCount++
		}
	}
	return result
}

func serializeTree(node *Node, w io.Writer) error {
	if node == nil {
		return nil
	}
	if node.left == nil && node.right == nil {
		// feuille : écrire 1 + rune
		if _, err := w.Write([]byte{1}); err != nil {
			return err
		}
		// écrire le rune (ex: UTF-8)
		_, err := w.Write([]byte(string(node.b)))
		return err
	} else {
		// interne : écrire 0
		if _, err := w.Write([]byte{0}); err != nil {
			return err
		}
		if err := serializeTree(node.left, w); err != nil {
			return err
		}
		return serializeTree(node.right, w)
	}
}

func deserializeTree(r *bufio.Reader) (*Node, error) {
	flag, err := r.ReadByte()
	if err != nil {
		return nil, err
	}

	if flag == 1 {
		cb, err := r.ReadByte()
		if err != nil {
			return nil, err
		}
		return &Node{b: cb}, nil
	} else if flag == 0 {
		left, err := deserializeTree(r)
		if err != nil {
			return nil, err
		}
		right, err := deserializeTree(r)
		if err != nil {
			return nil, err
		}
		return &Node{left: left, right: right}, nil
	} else {
		return nil, fmt.Errorf("valeur inattendue dans le flux: %v", flag)
	}
}

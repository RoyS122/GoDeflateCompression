# GoDeflateCompression
A pure Go implementation of DEFLATE compression (LZ77 + Huffman coding) optimized for string data.

## Features
- Dual-stage compression (LZ77 followed by Huffman encoding)
- Automatic cost analysis (skips LZ77 when ineffective)
- Zero external dependencies
- MIT Licensed

## Installation
```bash
go get github.com/RoyS122/GoDeflateCompression
```

## Usage 
```Go
package main

import (
    "fmt"
    deflate "github.com/RoyS122/GoDeflateCompression"
)

func main() {
    text := "Your sample text with repeated patterns like abcabcabc"
    
    // Compression
    binaryValue, tree, charCount, usedLZ := deflate.FullCompression(text)
    
    // Decompression
    decompressed := deflate.FullDecompression(binaryValue, tree, charCount, usedLZ)
    
    fmt.Println("Original:", text)
    fmt.Println("Decompressed:", decompressed)
}

```
*Here BinaryValue is the value of your data (compressed)*
*Tree is the value of huffman tree*
*TotalCharCount is the original length of your text*
*UsedLZ is a boolean who indicate if LZ77 was used (is not used whe it's not worth)*

## Warning
This package was developped by my own to fit to my project I put it to open source because the algorithm is public and it's only my implementation of it, it could be update or not, there is no certainty


## Contributing
Pull requests are welcome. For major changes, please open an issue first.

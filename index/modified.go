package index

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/cdvelop/vanify/common"
	"golang.org/x/net/html"
)

func (h *handler) UpdateFileOnDisk(filePath string) error {
	const e = "UpdateFileOnDisk Html "
	if filePath == "" {
		return nil
	}

	fmt.Println("Compilando HTML..." + filePath)

	time.Sleep(10 * time.Millisecond) // Esperar antes de intentar leer el archivo de nuevo

	//1- read file content from filePath
	content, err := os.ReadFile(filePath)
	if err != nil {
		return errors.New(e + err.Error())
	}

	doc, err := html.Parse(bytes.NewReader(content))

	fmt.Println("DOC NODE:", doc.Data)

	traverseNode(doc, 0)
	// h.UpdateFileContentInMemory(filePath, content)

	// if err := h.Minify(&h.buf); err != nil {
	// 	return errors.New(e + err.Error())
	// }

	// if err := common.FileWrite(h.FilePath, h.buf); err != nil {
	// 	return errors.New(e + err.Error())
	// }

	return errors.New("TEST")
	// return nil
}

func traverseNode(n *html.Node, depth int) {
	indent := strings.Repeat("  ", depth)

	switch n.Type {
	case html.ElementNode:
		fmt.Printf("%s<%s>\n", indent, n.Data)
	case html.TextNode:
		text := strings.TrimSpace(n.Data)
		if text != "" {
			fmt.Printf("%s%s\n", indent, text)
		}
	case html.CommentNode:
		fmt.Printf("%s<!-- %s -->\n", indent, n.Data)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		traverseNode(c, depth+1)
	}
}

func (h *handler) FileWasModifiedOnDiskOLD(filePath string) error {
	const e = "UpdateFileOnDisk Html "
	if filePath == "" {
		return nil
	}

	fmt.Println("Compilando HTML..." + filePath)

	time.Sleep(10 * time.Millisecond) // Esperar antes de intentar leer el archivo de nuevo

	//1- read file content from filePath
	content, err := os.ReadFile(filePath)
	if err != nil {
		return errors.New(e + err.Error())
	}

	h.UpdateFileContentInMemory(filePath, content)

	if err := h.Minify(&h.buf); err != nil {
		return errors.New(e + err.Error())
	}

	if err := common.FileWrite(h.FilePath, h.buf); err != nil {
		return errors.New(e + err.Error())
	}

	return nil
}

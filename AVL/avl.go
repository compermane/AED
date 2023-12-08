package AVL

import (
	"math"
)

type Node struct {
	filhoEsq *Node
	filhoDir *Node
	chave    float64
	altura   int64
}

func createNode(v float64) *Node {
	return &Node{nil, nil, v, 0}
}

func insertValue(v float64, raiz *Node) *Node {
	if raiz == nil {
		return createNode(v)
	}

	if v < raiz.chave {
		raiz.filhoEsq = insertValue(v, raiz.filhoEsq)
	} else if v > raiz.chave {
		raiz.filhoDir = insertValue(v, raiz.filhoDir)
	} else {
		return raiz
	}

	// Atualizar o fator de balanco para cada noh e balancear a arvore
	raiz.altura = int64(math.Max(float64(raiz.filhoEsq.getHeight()), float64(raiz.filhoDir.getHeight()))) + 1

	balance := raiz.getBalance()
	if balance > 1 && v < raiz.filhoEsq.chave {
		return rightRotate(raiz)
	}
	if balance < -1 && v > raiz.filhoDir.chave {
		return leftRotate(raiz)
	}
	if balance > 1 && v > raiz.filhoEsq.chave {
		raiz.filhoEsq = leftRotate(raiz.filhoEsq)
		return rightRotate(raiz)
	}
	if balance < -1 && v < raiz.filhoDir.chave {
		raiz.filhoDir = rightRotate(raiz.filhoDir)
		return leftRotate(raiz)
	}

	return raiz
}

func deleteValue(v float64, raiz *Node) *Node {
	if raiz == nil {
		return raiz
	}

	if v < raiz.chave {
		raiz.filhoEsq = deleteValue(v, raiz.filhoEsq)
	} else if v > raiz.chave {
		raiz.filhoDir = deleteValue(v, raiz.filhoDir)
	} else {
		if raiz.filhoEsq != nil && raiz.filhoDir != nil {
			menorValorDireito := findSmallest(raiz.filhoDir)
			raiz.chave = menorValorDireito.chave
			raiz.filhoDir = deleteValue(menorValorDireito.chave, raiz.filhoDir)
		} else if raiz.filhoEsq != nil {
			raiz = raiz.filhoEsq
		} else if raiz.filhoDir != nil {
			raiz = raiz.filhoDir
		} else {
			raiz = nil
			return raiz
		}

	}

	return rebalanceTree(raiz)
}

func rebalanceTree(raiz *Node) *Node {
	if raiz == nil {
		return raiz
	}

	raiz.altura = 1 + int64(math.Max(float64(raiz.filhoEsq.getHeight()), float64(raiz.filhoDir.getHeight())))

	fator := raiz.getBalance()

	if fator == -2 {
		if raiz.filhoDir.filhoEsq.getHeight() > raiz.filhoDir.filhoDir.getHeight() {
			raiz.filhoDir = rightRotate(raiz.filhoDir)
		}

		return leftRotate(raiz)
	} else if fator == 2 {
		if raiz.filhoEsq.filhoDir.getHeight() > raiz.filhoEsq.filhoEsq.getHeight() {
			raiz.filhoEsq = leftRotate(raiz.filhoEsq)
		}
		return rightRotate(raiz)
	}

	return raiz
}

func findSmallest(raiz *Node) *Node {
	if raiz.filhoEsq != nil {
		return findSmallest(raiz.filhoEsq)
	}

	return raiz
}

func (raiz *Node) getHeight() int64 {
	if raiz == nil {
		return 0
	}
	return raiz.altura
}

func (raiz *Node) getBalance() int64 {
	if raiz == nil {
		return 0
	}
	return raiz.filhoEsq.getHeight() - raiz.filhoDir.getHeight()
}

func rightRotate(raiz *Node) *Node {
	x := raiz.filhoEsq
	y := x.filhoDir

	x.filhoDir = raiz
	raiz.filhoEsq = y

	// Atualizacao de alturas
	raiz.altura = int64(math.Max(float64(raiz.filhoEsq.getHeight()), float64(raiz.filhoDir.getHeight()))) + 1
	x.altura = int64(math.Max(float64(x.filhoEsq.getHeight()), float64(x.filhoDir.getHeight()))) + 1

	return x
}

func leftRotate(raiz *Node) *Node {
	x := raiz.filhoDir
	y := x.filhoEsq

	x.filhoEsq = raiz
	raiz.filhoDir = y

	// Atualizacao de alturas
	raiz.altura = int64(math.Max(float64(raiz.filhoEsq.getHeight()), float64(raiz.filhoDir.getHeight()))) + 1
	x.altura = int64(math.Max(float64(x.filhoEsq.getHeight()), float64(x.filhoDir.getHeight()))) + 1

	return x
}

func (raiz *Node) inorder() []float64 {
	if raiz == nil {
		return nil
	}

	esq := raiz.filhoEsq.inorder()
	dir := raiz.filhoDir.inorder()

	out := make([]float64, 0)

	out = append(out, esq...)
	out = append(out, raiz.chave)
	out = append(out, dir...)

	return out
}

func (raiz *Node) preorder() []float64 {
	if raiz == nil {
		return nil
	}
	out := make([]float64, 0)
	out = append(out, raiz.chave)

	esq := raiz.filhoEsq.preorder()
	dir := raiz.filhoDir.preorder()

	out = append(out, esq...)
	out = append(out, dir...)

	return out
}

func (raiz *Node) postorder() []float64 {
	if raiz == nil {
		return nil
	}

	dir := raiz.filhoDir.postorder()
	esq := raiz.filhoEsq.postorder()
	out := make([]float64, 0)

	out = append(out, dir...)
	out = append(out, esq...)
	out = append(out, raiz.chave)

	return out
}

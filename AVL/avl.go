package AVL

import (
	"fmt"
	"math"
)

type Node struct {
	filhoEsq *Node
	filhoDir *Node
	chave    float64
	altura   int64
}

func CreateNode(v float64) *Node {
	fmt.Println("BRUH")
	return &Node{nil, nil, v, 0}
}

func InsertValue(v float64, raiz *Node) *Node {
	if raiz == nil {
		return CreateNode(v)
	}

	if v < raiz.chave {
		raiz.filhoEsq = InsertValue(v, raiz.filhoEsq)
	} else if v > raiz.chave {
		raiz.filhoDir = InsertValue(v, raiz.filhoDir)
	} else {
		return raiz
	}

	// Atualizar o fator de balanco para cada noh e balancear a arvore
	raiz.altura = int64(math.Max(float64(raiz.filhoEsq.GetHeight()), float64(raiz.filhoDir.GetHeight()))) + 1

	balance := raiz.GetBalance()
	if balance > 1 && v < raiz.filhoEsq.chave {
		return RightRotate(raiz)
	}
	if balance < -1 && v > raiz.filhoDir.chave {
		return LeftRotate(raiz)
	}
	if balance > 1 && v > raiz.filhoEsq.chave {
		raiz.filhoEsq = LeftRotate(raiz.filhoEsq)
		return RightRotate(raiz)
	}
	if balance < -1 && v < raiz.filhoDir.chave {
		raiz.filhoDir = RightRotate(raiz.filhoDir)
		return LeftRotate(raiz)
	}

	return raiz
}

func DeleteValue(v float64, raiz *Node) *Node {
	if raiz == nil {
		return raiz
	}

	if v < raiz.chave {
		raiz.filhoEsq = DeleteValue(v, raiz.filhoEsq)
	} else if v > raiz.chave {
		raiz.filhoDir = DeleteValue(v, raiz.filhoDir)
	} else {
		if raiz.filhoEsq != nil && raiz.filhoDir != nil {
			menorValorDireito := FindSmallest(raiz.filhoDir)
			raiz.chave = menorValorDireito.chave
			raiz.filhoDir = DeleteValue(menorValorDireito.chave, raiz.filhoDir)
		} else if raiz.filhoEsq != nil {
			raiz = raiz.filhoEsq
		} else if raiz.filhoDir != nil {
			raiz = raiz.filhoDir
		} else {
			raiz = nil
			return raiz
		}

	}

	return RebalanceTree(raiz)
}

func RebalanceTree(raiz *Node) *Node {
	if raiz == nil {
		return raiz
	}

	raiz.altura = 1 + int64(math.Max(float64(raiz.filhoEsq.GetHeight()), float64(raiz.filhoDir.GetHeight())))

	fator := raiz.GetBalance()

	if fator == -2 {
		if raiz.filhoDir.filhoEsq.GetHeight() > raiz.filhoDir.filhoDir.GetHeight() {
			raiz.filhoDir = RightRotate(raiz.filhoDir)
		}

		return LeftRotate(raiz)
	} else if fator == 2 {
		if raiz.filhoEsq.filhoDir.GetHeight() > raiz.filhoEsq.filhoEsq.GetHeight() {
			raiz.filhoEsq = LeftRotate(raiz.filhoEsq)
		}
		return RightRotate(raiz)
	}

	return raiz
}

func FindSmallest(raiz *Node) *Node {
	if raiz.filhoEsq != nil {
		return FindSmallest(raiz.filhoEsq)
	}

	return raiz
}

func (raiz *Node) GetHeight() int64 {
	if raiz == nil {
		return 0
	}
	return raiz.altura
}

func (raiz *Node) GetBalance() int64 {
	if raiz == nil {
		return 0
	}
	return raiz.filhoEsq.GetHeight() - raiz.filhoDir.GetHeight()
}

func RightRotate(raiz *Node) *Node {
	x := raiz.filhoEsq
	y := x.filhoDir

	x.filhoDir = raiz
	raiz.filhoEsq = y

	// Atualizacao de alturas
	raiz.altura = int64(math.Max(float64(raiz.filhoEsq.GetHeight()), float64(raiz.filhoDir.GetHeight()))) + 1
	x.altura = int64(math.Max(float64(x.filhoEsq.GetHeight()), float64(x.filhoDir.GetHeight()))) + 1

	return x
}

func LeftRotate(raiz *Node) *Node {
	x := raiz.filhoDir
	y := x.filhoEsq

	x.filhoEsq = raiz
	raiz.filhoDir = y

	// Atualizacao de alturas
	raiz.altura = int64(math.Max(float64(raiz.filhoEsq.GetHeight()), float64(raiz.filhoDir.GetHeight()))) + 1
	x.altura = int64(math.Max(float64(x.filhoEsq.GetHeight()), float64(x.filhoDir.GetHeight()))) + 1

	return x
}

func (raiz *Node) Inorder() []float64 {
	if raiz == nil {
		return nil
	}

	esq := raiz.filhoEsq.Inorder()
	dir := raiz.filhoDir.Inorder()

	out := make([]float64, 0)

	out = append(out, esq...)
	out = append(out, raiz.chave)
	out = append(out, dir...)

	return out
}

func (raiz *Node) Preorder() []float64 {
	if raiz == nil {
		return nil
	}
	out := make([]float64, 0)
	out = append(out, raiz.chave)

	esq := raiz.filhoEsq.Preorder()
	dir := raiz.filhoDir.Preorder()

	out = append(out, esq...)
	out = append(out, dir...)

	return out
}

func (raiz *Node) Postorder() []float64 {
	if raiz == nil {
		return nil
	}

	dir := raiz.filhoDir.Postorder()
	esq := raiz.filhoEsq.Postorder()
	out := make([]float64, 0)

	out = append(out, dir...)
	out = append(out, esq...)
	out = append(out, raiz.chave)

	return out
}

package Hash

import (
	"encoding/binary"
	"fmt"
)

const (
	LOAD_FACTOR_LIMIT 	float64 = 0.5
	STD_LIMIT 			int 	= 16
)

// Uma hash table consiste basicamente de um array de ponteiros para seus itens (buckets).
// Podemos também armazenar seu tamanho
type HashTable struct {
	lenght int64
	buckets [][]Item 
}

// Um item de uma hash table consiste em seu valor e sua chave associada
type Item struct {
	key PreHashable
	value interface{}
}

// Em go, interfaces servem para generalizar métodos para determinados tipos
// No caso, essa interface implementa duas funções para os tipos IntKey e StringKey:
// - Uma para representar chaves inteiras como um array de bytes
// - E outra para conferir se duas chaves são iguais
type PreHashable interface {
	HashBytes() []byte
	Equal(PreHashable) bool
}

type IntKey int
type StringKey string

func (i IntKey) HashBytes() []byte {
	buf := make([]byte, binary.MaxVarintLen64)	// Cria um array de bytes com o valor máximo de bytes que um int64 pode ter
	n := binary.PutVarint(buf, int64(i))		// Converte a chave inteira para bytes, armazena em buf e retorna a quantidade de bytes

	return buf[:n]								// Retorna os n bytes codificados
}

func (s StringKey) HashBytes() []byte {			// Para chaves string é mais fácil, basta converter o valor da string para um
	return []byte(s)							// array de bytes
}

func (i IntKey) Equal(other PreHashable) bool {
	v, ok := other.(IntKey)

	return ok && i == v
}

func (s StringKey) Equal(other PreHashable) bool {
	v, ok := other.(StringKey)

	return ok && s == v
}

func hashFunction(key PreHashable, limit int64) (hash uint64) {
	// Definindo uma função hash baseada no algoritmo Fowler-Noll-Vo (FNV)
	const (
		offset uint64 = 0xcbf29ce484222325
		prime uint64 = 0x00000100000001b3
	)

	hash = offset

	for _, b := range key.HashBytes() {
		hash ^= uint64(b)
		hash *= prime
	}

	return uint64(hash % uint64(limit))
}

// Cria uma nova hash table de acordo com um número size de buckets
func CreateHashTable() *HashTable {
	return CreateHashTableVar(STD_LIMIT)
}

func CreateHashTableVar(size int) *HashTable {
	if size <= 0 {
		return &HashTable{
			buckets: make([][]Item, STD_LIMIT),
		}
	}
	return &HashTable{
		buckets: make([][]Item, size),
	}
}

func (table *HashTable) loadFactor() float64 {
	// O load factor (fator de carga) de uma hash table é a razão entre a quantidade de chaves
	// armarzenadas e a quantidade de buckets. 
	return float64(table.lenght) / float64(len(table.buckets))
}

func (table *HashTable) expandTable() error {
	newTable := make([][]Item, len(table.buckets) * 2)

	for _, bucket := range table.buckets {
		for _, e := range bucket {
			newHash := hashFunction(e.key, int64(len(newTable)))
			newTable[newHash] = append(newTable[newHash], Item{e.key, e.value})
		}
	}

	table.buckets = newTable
	return nil
}

func (table *HashTable) Add(key PreHashable, value interface{}) {
	hash := hashFunction(key, int64(len(table.buckets)))

	// Checa se a chave a ser inserida já está na tabela. 
	// Se sim, sobrescreve.
	for i, e := range table.buckets[hash] {
		if e.key == key {
			table.buckets[hash][i].value = value
			return
		}
	}

	// Se não, adiciona a chave
	table.buckets[hash] = append(table.buckets[hash], Item{key, value})
	table.lenght += 1		

	if table.loadFactor() > LOAD_FACTOR_LIMIT {
		table.expandTable()
	}
}

func (table *HashTable) Get(key PreHashable) (interface{}, bool) {
	hash := hashFunction(key, int64(len(table.buckets)))

	for _, v := range table.buckets[hash] {
		if v.key == key {
			return v.value, true
		}
	}

	return nil, false
}

func (table *HashTable) Delete(key PreHashable) error {
	hash := hashFunction(key, int64(len(table.buckets)))

	_, ok := table.Get(key) 
	// Caso em que a chave não existe
	if !ok {
		return nil
	}

	for i, e := range table.buckets[hash] {
		if e.key == key {
			current := table.buckets[hash]
			current[i] = current[len(current) - 1]
			current = current[:len(current) - 1]

			table.lenght -= 1
			table.buckets[hash] = current

			return nil
		}
	}

	return fmt.Errorf("Key error: key %v does not exist", key)
}

func (table *HashTable) PrintTable() {
	for _, bucket := range table.buckets {
		for _, e := range bucket {
			fmt.Printf("%v: %v\n", e.key, e.value)
		}
	}
}
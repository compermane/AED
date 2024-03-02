package Hash

import (
	"testing"
)

type testCase struct {
	stringKeys 	[]StringKey
	intKeys 	[]IntKey
	values		[]interface{}
}

var unexistingStringKeys = []StringKey{"unexisting", "string", "keys"}
var unexistingIntKeys = []IntKey{-90, 70000, 5677}
var tableSizes = []int64{0, -1, 16, 32, 90000}
var testCases = []testCase {
	{	
		stringKeys: []StringKey{"foo", "bar", "teste", "golang"},
		intKeys: []IntKey{-1, 10, 20, 1000000000},
		values: []interface{}{"varios", "valores", 10, -1},
	},
	{
		stringKeys: []StringKey{},
		intKeys: []IntKey{},
		values: []interface{}{},
	},
}

func assertTableSize(t *testing.T, table *HashTable, size int64) {
	if size <= 0 {
		if int64(len(table.buckets)) != 16 {
			t.Errorf("Table size error: expected %v got %v", 16, int64(len(table.buckets)))
		}
	} else if int64(len(table.buckets)) != size {
		t.Errorf("Table size error: expected %v got %v", size, int64(len(table.buckets)))
	}
}

func assertNoItems(t *testing.T, table *HashTable) {
	for _, testCase := range testCases {
		for _, key := range testCase.stringKeys {
			_, ok := table.Get(key)
			if ok != false {
				t.Errorf("Key error: key %v found on empty table", key)
			}
		}

		for _, key := range testCase.intKeys {
			_, ok := table.Get(key)
			if ok != false {
				t.Errorf("Key error: key %v found on empty table", key)
			}
		}
	}
}

func assertKeyIsAdded(t *testing.T, table *HashTable, key PreHashable, value interface{}) {
	v, ok := table.Get(key)

	if v != value || !ok {
		if v != value {
			t.Errorf("Value error: expected %v got %v", v, value)
		} else {
			t.Errorf(".Get error: key not found")
		}
	}
}

func assertKeyNotAdded(t *testing.T, table *HashTable, key PreHashable, value interface{}) {
	v, ok := table.Get(key)

	if v != nil || ok {
		if v != nil {
			t.Errorf("Value error: expected nil, got %v", v)
		} else {
			t.Errorf("Unexistent key found: %v", key)
		}
	}
}

func assertKeyIsDeleted(t *testing.T, table *HashTable, key PreHashable, tableLen int64) {
	err := table.Delete(key)

	if err != nil {
		t.Errorf("Deletion error: could not delete key %v", key)
	} else if table.lenght != tableLen - 1 {
		t.Errorf("Lenght error: table lenght not updated")
	}
}

func assertTableExpansion(t *testing.T, table *HashTable, initialSize int) {
	if initialSize == 0 {
		return
	}

	if len(table.buckets) != initialSize * 2 {
		t.Errorf("Expansion error: table did not expand, expected %v got %v", 2 * initialSize, len(table.buckets))
	}
}

func TestCreateTable(t *testing.T) {
	for _, size := range tableSizes {
		table := CreateHashTableVar(int(size))
		assertTableSize(t, table, size)
		assertNoItems(t, table)
	}
}

func TestAddKey(t *testing.T) {
	for _, testCase := range testCases {
		table := CreateHashTable()
		for i := range testCase.intKeys {
			table.Add(testCase.intKeys[i], testCase.values[i])
			assertKeyIsAdded(t, table, testCase.intKeys[i], testCase.values[i])
		}
		for i := range testCase.stringKeys {
			table.Add(testCase.stringKeys[i], testCase.values[i])
			assertKeyIsAdded(t, table, testCase.stringKeys[i], testCase.values[i])
		}
		for _, v := range unexistingIntKeys {
			assertKeyNotAdded(t, table, v, "i do not exist")
		}
		for _, v := range unexistingStringKeys {
			assertKeyNotAdded(t, table, v, "shit")
		}
	}
}

func TestDeleteKey(t *testing.T) {
	for _, testCase := range testCases {
		table := CreateHashTable()
		for i := range testCase.intKeys {
			table.Add(testCase.intKeys[i], testCase.values[i])
		}
		for _, key := range testCase.intKeys {
			assertKeyIsDeleted(t, table, key, table.lenght)
		}
		for i := range testCase.stringKeys {
			table.Add(testCase.stringKeys[i], testCase.values[i])
		}
		for _, key := range testCase.stringKeys {
			assertKeyIsDeleted(t, table, key, table.lenght)
		}
	}

}

func TestExpansion(t *testing.T) {
	for _, testCase := range testCases {
		initialSize := len(testCase.intKeys) + len(testCase.stringKeys)
		table := CreateHashTableVar(initialSize)
		for i := range testCase.intKeys {
			table.Add(testCase.intKeys[i], testCase.values[i])
			table.Add(testCase.stringKeys[i], testCase.values[i])
		}

		assertTableExpansion(t, table, initialSize)
	}
}
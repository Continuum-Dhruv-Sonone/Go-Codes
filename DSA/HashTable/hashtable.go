package hashtable

//Record is used to store key, value and link to next Record
type Record struct {
	Key   int
	Value int
	Next  *Record
}

//HashTable ...
type HashTable struct {
	ht   []*Record
	size int
}

//NewHashTable ...
func NewHashTable(tableSize int) *HashTable {
	return &HashTable{
		ht:   make([]*Record, tableSize),
		size: 0,
	}
}

func hash(key, size int) int {
	return key % size
}

//Insert into hashtable
func (ht *HashTable) Insert(key, value int) {
	hashKey := hash(key, ht.size)
	r := &Record{key, value, nil}

	if ht.ht[hashKey] == nil {
		ht.ht[hashKey] = r
		return
	}

	prev := &Record{0, 0, nil}

	iterator := ht.ht[hashKey]

	for iterator != nil {
		if iterator.Key == key {
			iterator.Value = value
			return
		}

		prev = iterator
		iterator = iterator.Next
	}

	prev.Next = r

	ht.size = ht.size + 1
}

//Delete key from hashtable
func (ht *HashTable) Delete(key int) {
	hashKey := hash(key, ht.size)

	if ht.ht[hashKey] == nil {
		return
	}

	iterator := ht.ht[hashKey]

	if iterator.Key == key {
		ht.ht[hashKey] = iterator.Next
		return
	}

	prev := iterator
	iterator = iterator.Next

	for iterator != nil {

		if iterator.Key == key {
			prev.Next = iterator.Next
			ht.size = ht.size - 1
			return
		}

		prev = iterator
		iterator = iterator.Next
	}
}

//Get return value for specific key
func (ht HashTable) Get(key int) (bool, int) {
	hashKey := hash(key, ht.size)

	if ht.ht[hashKey] == nil {
		return false, 0
	}

	iterator := ht.ht[hashKey]

	for iterator != nil {
		if iterator.Key == key {
			return true, iterator.Value
		}

		iterator = iterator.Next
	}

	return false, 0
}

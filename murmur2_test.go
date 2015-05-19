package murmur

import (
	"fmt"
	"hash"
	"hash/fnv"
	"testing"
)

func printv(v uint32) {
	fmt.Printf("%#v %d\n", v, v)
}

// Sorry, no real tests here. Just some manual checks that things look right.
// Feel free to submit a pull request with some better testing :)

func TestStuff(t *testing.T) {
	printv(MurmurHash2([]byte("foo"), 123))       // == 1412061192
	printv(MurmurHash2([]byte("zztop"), 123))     // == 1878194508
	printv(MurmurHash2([]byte("foobarbaz"), 234)) // == 1777016281
	printv(MurmurHash2([]byte("blam"), 777))      // == 1668928339

	println("-----------------")

	var s uint32
	var h hash.Hash32

	h = New32(123)
	h.Write([]byte("zztop"))
	h.Write([]byte("zztop"))
	h.Write([]byte("zztop"))
	h.Write([]byte("zztop!"))

	s = h.Sum32()
	printv(s)

	s = h.Sum32()
	printv(s) // should be the same as above

	println("-----------------")

	h.Reset()
	h.Write([]byte("foo"))
	s = h.Sum32()
	printv(s)

	h.Reset()
	h.Write([]byte("foo"))
	s = h.Sum32()
	printv(s) // should be the same as above

	println("-----------------")
}

// -----------------------------------------------------------------------------

var sampleBytes = []byte("hardly a good test, but hey.")

func BenchmarkMurmurHash2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MurmurHash2(sampleBytes, 42)
	}
}

func BenchmarkMurmurHash2A(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MurmurHash2A(sampleBytes, 42)
	}
}

func BenchmarkMurmurHash64A(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MurmurHash64A(sampleBytes, 42)
	}
}

func BenchmarkHash32_Murmur2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		h := New32(42)
		h.Write(sampleBytes)
		h.Write(sampleBytes)
		h.Write(sampleBytes)
		h.Sum32()
	}
}

// Benchmark "hash/fnv" to get a comparison on speed.

func BenchmarkHash32_FNV1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		h := fnv.New32()
		h.Write(sampleBytes)
		h.Write(sampleBytes)
		h.Write(sampleBytes)
		h.Sum32()
	}
}

func BenchmarkHash32_FNV1a(b *testing.B) {
	for i := 0; i < b.N; i++ {
		h := fnv.New32a()
		h.Write(sampleBytes)
		h.Write(sampleBytes)
		h.Write(sampleBytes)
		h.Sum32()
	}
}

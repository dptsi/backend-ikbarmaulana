package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Ikbar")
	}

	// bisa juga menggunakan sub dan table benchmark dengan cara hampir sama
}

func TestMain(m *testing.M) {
	//before
	fmt.Println("Sebelum unit test")
	m.Run() // ekseskusi semua unit test

	//after
	fmt.Println("setelah unit test")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Ikbar")

	if result != "Hello Ikbar" {
		//panic("Result isn't Hello Ikbar")
		//t.Fail() // melanjutkan code selanjutnya
		//t.FailNow() // berhenti saat ini juga

		//t.Error("harusnya gagitu") // memberi pesan dan melanjutkan code selanjutnya
		t.Fatal("harusnya ga gitu") // memberi pesan dan mengakhiri test
	}

	//fmt.Println("dieksekusi")
}

func TestHelloWorldUsingAssertion(t *testing.T) { // assertion == fail
	result := HelloWorld("Ikbar")
	assert.Equal(t, "Hello Ikbar", result, "harusnya ga gitu")
	fmt.Println("assertion done")
}

func TestHelloWorldUsingRequire(t *testing.T) { // require == fatal
	result := HelloWorld("Ikbar")
	require.Equal(t, "Hello Ikbar", result, "Harusnya ga gitu")
	fmt.Println("require done")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Diskip aja")
	}

	fmt.Println("lololoo")
}

func TestSubTest(t *testing.T) {
	t.Run("Ikbar", func(t *testing.T) {
		result := HelloWorld("Ikbar")
		require.Equal(t, "Hello Ikbar", result, "is must be Ikbar")
	})
	t.Run("Maulana", func(t *testing.T) {
		result := HelloWorld("Maulana")
		require.Equal(t, "Hello Maulana", result, "is must be Maulana")
	})
}

func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "HelloWorld(Ikbar)",
			request:  "Ikbar",
			expected: "Hello Ikbar",
		},
		{
			name:     "HelloWorld(Maulana)",
			request:  "Maulana",
			expected: "Hello Maulana",
		},
		{
			name:     "HelloWorld(Alexander)",
			request:  "Alexander",
			expected: "Hello Alexander",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result, "hohoh salah")
		})
	}
}

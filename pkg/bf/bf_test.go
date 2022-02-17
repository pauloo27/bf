package bf_test

import (
	"testing"

	"github.com/Pauloo27/bf/pkg/bf"
	"github.com/stretchr/testify/require"
)

// TODO: some edge cases

func TestMemoryNavigation(t *testing.T) {
	t.Skip()
	t.Run("empty program", func(t *testing.T) {
		pgr := bf.NewProgram("", "")
		_, err := pgr.Run()
		require.Nil(t, err)
		require.Equal(t, 0, pgr.MemoryIndex)
	})

	t.Run("> operator", func(t *testing.T) {
		pgr := bf.NewProgram(">>>", "")
		_, err := pgr.Run()
		require.Nil(t, err)
		require.Equal(t, 3, pgr.MemoryIndex)
	})

	t.Run("< operator", func(t *testing.T) {
		pgr := bf.NewProgram(">>>", "")
		_, err := pgr.Run()
		require.Nil(t, err)
		require.Equal(t, 3, pgr.MemoryIndex)
	})

	t.Run("< and > operators", func(t *testing.T) {
		pgr := bf.NewProgram(">>><<", "")
		_, err := pgr.Run()
		require.Nil(t, err)
		require.Equal(t, 1, pgr.MemoryIndex)
	})
}

func TestMemoryUpdate(t *testing.T) {
	t.Skip()
	t.Run("+ operator", func(t *testing.T) {
		pgr := bf.NewProgram("+++++", "")
		_, err := pgr.Run()
		require.Nil(t, err)
		require.Equal(t, 0, pgr.MemoryIndex)
		require.Equal(t, byte(5), pgr.Memory[0])
	})

	t.Run("- operator", func(t *testing.T) {
		pgr := bf.NewProgram("----", "")
		_, err := pgr.Run()
		require.Nil(t, err)
		require.Equal(t, 0, pgr.MemoryIndex)
		require.Equal(t, byte(252), pgr.Memory[0]) // underflow
	})

	t.Run("+ and - operators", func(t *testing.T) {
		pgr := bf.NewProgram("+++++----", "")
		_, err := pgr.Run()
		require.Nil(t, err)
		require.Equal(t, 0, pgr.MemoryIndex)
		require.Equal(t, byte(1), pgr.Memory[0])
	})
}

func TestInputRead(t *testing.T) {
	t.Skip()
	t.Run("read from input", func(t *testing.T) {
		pgr := bf.NewProgram(">,", " ")
		_, err := pgr.Run()
		require.Nil(t, err)
		require.Equal(t, 1, pgr.MemoryIndex)
		require.Equal(t, byte(32), pgr.Memory[1])
	})

	t.Run("read from input multiple times", func(t *testing.T) {
		pgr := bf.NewProgram(">,>,", " A")
		_, err := pgr.Run()
		require.Nil(t, err)
		require.Equal(t, 2, pgr.MemoryIndex)
		require.Equal(t, byte(32), pgr.Memory[1])
		require.Equal(t, byte(65), pgr.Memory[2])
	})
}

func TestWriteOutput(t *testing.T) {
	t.Skip()
	t.Run("write to output", func(t *testing.T) {
		pgr := bf.NewProgram(">,>,.<.", " A")
		out, err := pgr.Run()
		require.Nil(t, err)
		require.Equal(t, 1, pgr.MemoryIndex)
		require.Equal(t, byte(32), pgr.Memory[1])
		require.Equal(t, byte(65), pgr.Memory[2])

		require.Equal(t, byte(65), out[0])
		require.Equal(t, byte(32), out[1])
	})
}

func TestLoop(t *testing.T) {
	t.Run("loop 3 times", func(t *testing.T) {
		pgr := bf.NewProgram("+++[>+<-].>.", "")
		out, err := pgr.Run()
		require.Nil(t, err)
		require.Equal(t, 1, pgr.MemoryIndex)

		require.Equal(t, byte(0), out[0])
		require.Equal(t, byte(3), out[1])
	})
}

func TestHelloMom(t *testing.T) {
	t.Run("show hello mom", func(t *testing.T) {
		pgr := bf.NewProgram(`
			+++++++[>++++++++++<-]>++.		H
			---.													E
			+++++++.											L
			.															L
			+++.													O
			<+++++[>----------<-]>+++.		SPACE
			<+++++[>++++++++++<-]>-----.	M
			++.														O
			--.														M
		`, "")
		out, err := pgr.Run()
		require.Nil(t, err)
		require.Equal(t, 1, pgr.MemoryIndex)

		require.Equal(t, byte('H'), out[0])
		require.Equal(t, byte('E'), out[1])
		require.Equal(t, byte('L'), out[2])
		require.Equal(t, byte('L'), out[3])
		require.Equal(t, byte('O'), out[4])
		require.Equal(t, byte(' '), out[5])
		require.Equal(t, byte('M'), out[6])
		require.Equal(t, byte('O'), out[7])
		require.Equal(t, byte('M'), out[8])
	})
}

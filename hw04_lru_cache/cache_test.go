package hw04lrucache

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Run("empty cache", func(t *testing.T) {
		c := NewCache(10)

		_, ok := c.Get("aaa")
		require.False(t, ok)

		_, ok = c.Get("bbb")
		require.False(t, ok)
	})

	t.Run("simple", func(t *testing.T) {
		c := NewCache(5)

		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200)
		require.False(t, wasInCache)

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 200, val)

		wasInCache = c.Set("aaa", 300)
		require.True(t, wasInCache)

		val, ok = c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 300, val)

		val, ok = c.Get("ccc")
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("purge logic", func(t *testing.T) {
		c := NewCache(3)
		c.Set("a", 100)
		c.Set("b", 200)
		c.Set("c", 300)
		c.Set("d", 400)

		val, ok := c.Get("a")
		require.False(t, ok)
		require.Nil(t, val)

		val, ok = c.Get("d")
		require.True(t, ok)
		require.Equal(t, 400, val)
	})

	t.Run("purge logic long time", func(t *testing.T) {
		c := NewCache(3)
		c.Set("a", 100) // [a]
		c.Set("b", 200) // [b, a]
		c.Set("c", 300) // [c, b, a]

		val, ok := c.Get("a") // [a, c, b]
		require.True(t, ok)
		require.Equal(t, 100, val)

		val, ok = c.Get("b") // [b, a, c]
		require.True(t, ok)
		require.Equal(t, 200, val)

		val, ok = c.Get("c") // [c, b, a]
		require.True(t, ok)
		require.Equal(t, 300, val)

		wasInCache := c.Set("b", 500) // [b, c, a]
		require.True(t, wasInCache)

		wasInCache = c.Set("c", 600) // [c, b, a]
		require.True(t, wasInCache)

		val, ok = c.Get("b") // [b, c, a]
		require.True(t, ok)
		require.Equal(t, 500, val)

		c.Set("d", 400) // [d, b, c]

		val, ok = c.Get("d") // [d, b, c]
		require.True(t, ok)
		require.Equal(t, 400, val)

		val, ok = c.Get("a") // [d, b, c]
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("clear cache", func(t *testing.T) {
		c := NewCache(3)
		c.Set("a", 100) // [a]
		c.Set("b", 200) // [b, a]
		c.Set("c", 300) // [c, b, a]

		val, ok := c.Get("a") // [a, c, b]
		require.True(t, ok)
		require.Equal(t, 100, val)

		val, ok = c.Get("b") // [b, a, c]
		require.True(t, ok)
		require.Equal(t, 200, val)

		c.Clear() // []

		_, ok = c.Get("a") // []
		require.False(t, ok)

		_, ok = c.Get("b") // []
		require.False(t, ok)
	})
}

func TestCacheMultithreading(t *testing.T) {
	t.Skip() // Remove me if task with asterisk completed.

	c := NewCache(10)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Set(Key(strconv.Itoa(i)), i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Get(Key(strconv.Itoa(rand.Intn(1_000_000))))
		}
	}()

	wg.Wait()
}

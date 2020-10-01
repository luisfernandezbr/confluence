package html

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmoji(t *testing.T) {
	assert := assert.New(t)
	buf, err := Parse(`:)`)
	assert.NoError(err)
	assert.Equal("😃", buf)
	buf, err = Parse(`:(`)
	assert.NoError(err)
	assert.Equal("🙁", buf)
	buf, err = Parse(`:P`)
	assert.NoError(err)
	assert.Equal("😛", buf)
	buf, err = Parse(`:D`)
	assert.NoError(err)
	assert.Equal("😁", buf)
	buf, err = Parse(`;)`)
	assert.NoError(err)
	assert.Equal("😉", buf)
	buf, err = Parse(`(y)`)
	assert.NoError(err)
	assert.Equal("👍", buf)
	buf, err = Parse(`(n)`)
	assert.NoError(err)
	assert.Equal("👎", buf)
	buf, err = Parse(`(i)`)
	assert.NoError(err)
	assert.Equal("ℹ️", buf)
	buf, err = Parse(`(/)`)
	assert.NoError(err)
	assert.Equal("✅", buf)
	buf, err = Parse(`(x)`)
	assert.NoError(err)
	assert.Equal("❗️", buf)
	buf, err = Parse(`(!)`)
	assert.NoError(err)
	assert.Equal("⚠️", buf)
	buf, err = Parse(`(+)`)
	assert.NoError(err)
	assert.Equal("➕", buf)
	buf, err = Parse(`(-)`)
	assert.NoError(err)
	assert.Equal("➖", buf)
	buf, err = Parse(`(?)`)
	assert.NoError(err)
	assert.Equal("❓", buf)
}

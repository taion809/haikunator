package haikunator

import (
	"math/rand"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestDefaultReturnsTwoWordsAndInt(t *testing.T) {
	h := NewHaikunator()

	haiku := h.Haikunate()

	parts := strings.Split(haiku, "-")

	if len(parts) < 3 {
		t.Errorf("Generated haiku [%s] is less than 3 parts: %v\n", haiku, parts)
	}

	_, err := strconv.ParseInt(parts[2], 10, 64)
	if err != nil {
		t.Error("Third part is not integer: ", parts[2])
	}
}

func TestNonDefaultReturnsTwoWordsAndInt(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	h := Haikunator{r: r, delim: ".", token: 1000}

	haiku := h.Haikunate()

	parts := strings.Split(haiku, ".")

	if len(parts) < 3 {
		t.Errorf("Generated haiku [%s] is less than 3 parts: %v\n", haiku, parts)
	}

	token, err := strconv.ParseInt(parts[2], 10, 64)
	if err != nil {
		t.Error("Third part is not integer: ", parts[2])
	}

	if token < 0 || token > 1000 {
		t.Error("Generated token is outside of bounds 0-1000", token)
	}
}

func TestTokenHaikunateBetweenZeroAndMax(t *testing.T) {
	h := NewHaikunator()

	haiku := h.TokenHaikunate(1000)

	parts := strings.Split(haiku, "-")

	if len(parts) < 3 {
		t.Errorf("Generated haiku [%s] is less than 3 parts: %v\n", haiku, parts)
	}

	token, err := strconv.ParseInt(parts[2], 10, 64)
	if err != nil {
		t.Error("Third part is not integer: ", parts[2])
	}

	if token < 0 || token > 1000 {
		t.Error("Generated token is outside of bounds 0-1000", token)
	}
}

func TestZeroTokenGeneratesNoTokenHaiku(t *testing.T) {
	h := NewHaikunator()

	haiku := h.TokenHaikunate(0)

	parts := strings.Split(haiku, "-")

	if len(parts) != 2 {
		t.Errorf("Generated haiku has invalid number of parts, expected 2: %v\n", parts)
	}
}

func TestSpaceDelimHaikuHasCorrectDelimAndNoToken(t *testing.T) {
	h := NewHaikunator()

	haiku := h.DelimHaikunate(" ")

	parts := strings.Split(haiku, " ")

	if len(parts) != 2 {
		t.Errorf("Generated haiku has invalid number of parts, expected 2: %v\n", parts)
	}
}

func TestTokenDelimHaikuHasDelimAndToken(t *testing.T) {
	h := NewHaikunator()

	haiku := h.TokenDelimHaikunate(1000, ".")

	parts := strings.Split(haiku, ".")

	if len(parts) != 3 {
		t.Errorf("Generated haiku has invalid number of parts, expected 3: %v\n", parts)
	}

	token, err := strconv.ParseInt(parts[2], 10, 64)
	if err != nil {
		t.Error("Third part is not integer: ", parts[2])
	}

	if token < 0 || token > 1000 {
		t.Error("Generated token is outside of bounds 0-1000", token)
	}
}

func TestZeroTokenDelimHaikuHasDelimAndNoToken(t *testing.T) {
	h := NewHaikunator()

	haiku := h.TokenDelimHaikunate(0, " ")

	parts := strings.Split(haiku, " ")

	if len(parts) != 2 {
		t.Errorf("Generated haiku has invalid number of parts, expected 2: %v\n", parts)
	}
}

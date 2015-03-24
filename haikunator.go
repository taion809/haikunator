package haikunator

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var (
	adjectives = []string{"autumn", "hidden", "bitter", "misty", "silent", "empty", "dry", "dark", "summer",
		"icy", "delicate", "quiet", "white", "cool", "spring", "winter", "patient",
		"twilight", "dawn", "crimson", "wispy", "weathered", "blue", "billowing",
		"broken", "cold", "damp", "falling", "frosty", "green", "long", "late", "lingering",
		"bold", "little", "morning", "muddy", "old", "red", "rough", "still", "small",
		"sparkling", "throbbing", "shy", "wandering", "withered", "wild", "black",
		"young", "holy", "solitary", "fragrant", "aged", "snowy", "proud", "floral",
		"restless", "divine", "polished", "ancient", "purple", "lively", "nameless"}
	nouns = []string{"waterfall", "river", "breeze", "moon", "rain", "wind", "sea", "morning",
		"snow", "lake", "sunset", "pine", "shadow", "leaf", "dawn", "glitter", "forest",
		"hill", "cloud", "meadow", "sun", "glade", "bird", "brook", "butterfly",
		"bush", "dew", "dust", "field", "fire", "flower", "firefly", "feather", "grass",
		"haze", "mountain", "night", "pond", "darkness", "snowflake", "silence",
		"sound", "sky", "shape", "surf", "thunder", "violet", "water", "wildflower",
		"wave", "water", "resonance", "sun", "wood", "dream", "cherry", "tree", "fog",
		"frost", "voice", "paper", "frog", "smoke", "star"}
)

type Haikunator struct {
	r     *rand.Rand
	delim string
	token int64
}

func NewHaikunator() Haikunator {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	h := Haikunator{r: r, delim: "-", token: 9999}

	return h
}

func (h *Haikunator) haikunate(token, delim string) string {
	haiku := fmt.Sprintf("%s%s%s", adjectives[h.r.Intn(len(adjectives))], delim, nouns[h.r.Intn(len(nouns))])

	if len(token) > 0 {
		haiku = fmt.Sprintf("%s%s%s", haiku, delim, token)
	}

	return haiku
}

func (h *Haikunator) Haikunate() string {
	tokenString := strconv.FormatInt(h.r.Int63n(h.token), 10)

	return h.haikunate(tokenString, h.delim)
}

func (h *Haikunator) TokenHaikunate(token int64) string {
	tokenString := ""

	if token > 0 {
		tokenString = strconv.FormatInt(h.r.Int63n(token), 10)
	}

	return h.haikunate(tokenString, h.delim)
}

func (h *Haikunator) DelimHaikunate(delim string) string {
	tokenString := ""
	return h.haikunate(tokenString, delim)
}

func (h *Haikunator) TokenDelimHaikunate(token int64, delim string) string {
	tokenString := ""

	if token > 0 {
		tokenString = strconv.FormatInt(h.r.Int63n(token), 10)
	}

	return h.haikunate(tokenString, delim)
}

func main() {
	h := NewHaikunator()

	name := h.Haikunate()
	fmt.Println("Haiku: ", name)
}

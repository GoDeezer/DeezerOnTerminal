package format

import (
	"fmt"

	"github.com/godeezer/lib/deezer"
)

func StringLimitLength(s string, n int) string {
	if len(s) > n {
		return s[n:]
	}
	for i := 0; i < n-len(s); i++ {
		s += " "
	}
	return s
}

func FormatSongs(songs []deezer.Song, col int) []string {
	ret := make([]string, len(songs))
	for i, s := range songs {
		ret[i] = fmt.Sprintf("%s %s", StringLimitLength(s.Title, col/3), StringLimitLength(s.ArtistName, col/3))
	}
	return ret
}

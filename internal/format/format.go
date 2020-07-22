package format

import (
	"fmt"

	"github.com/godeezer/lib/deezer"
)

func StringLimitLength(s string, n int) string {
	if len(s) >= n {
		return s[0:n]
	}
	ret := s
	for i := 0; i < n-len(s); i++ {
		ret += " "
	}
	return ret
}

func FormatSongs(songs []deezer.Song, col int) []string {
	ret := make([]string, len(songs))
	for i, s := range songs {
		if s.ExplicitContent.CoverStatus == 1 && s.ExplicitContent.LyricsStatus == 1 || true {
			ret[i] = StringLimitLength(fmt.Sprintf("%s - %s",
				StringLimitLength(s.Title, col/3),
				StringLimitLength(s.ArtistName, col/3)), col-12) + "[ Explicit ](mod:reverse)"
		} else {
			ret[i] = StringLimitLength(fmt.Sprintf("%s - %s",
				StringLimitLength(s.Title, col/3),
				StringLimitLength(s.ArtistName, col/3)), col)
		}
	}
	return ret
}

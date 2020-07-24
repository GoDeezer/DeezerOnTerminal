package format

import (
	"fmt"

	"unicode/utf8"

	"github.com/godeezer/lib/deezer"
)

func StringLimitLength(s string, n int) string {
	if len(s) >= n {
		return s[0:n]
	}
	ret := s
	for i := 0; i < n-utf8.RuneCountInString(s); i++ {
		ret += " "
	}
	return ret
}

func FormatSongs(songs []deezer.Song, col int) []string {
	ret := make([]string, len(songs))
	for i, s := range songs {
		if s.ExplicitContent.CoverStatus == 1 && s.ExplicitContent.LyricsStatus == 1 || true {
			ret[i] = StringLimitLength(fmt.Sprintf("%s - %s - %s",
				StringLimitLength(s.Title, col/4),
				StringLimitLength(s.ArtistName, col/4),
				StringLimitLength(s.AlbumTitle, col/4)), col-12) + "[ Explicit ](mod:reverse)"
		} else {
			ret[i] = StringLimitLength(fmt.Sprintf("%s - %s - %s",
				StringLimitLength(s.Title, col/4),
				StringLimitLength(s.ArtistName, col/4),
				StringLimitLength(s.AlbumTitle, col/4)), col)
		}
	}
	return ret
}

func FormatAlbums(albums []deezer.Album, col int) []string {
	ret := make([]string, len(albums))
	for i, a := range albums {
		if a.ExplicitContent.CoverStatus == 1 && a.ExplicitContent.LyricsStatus == 1 {
			ret[i] = StringLimitLength(fmt.Sprintf("%s - %s",
				StringLimitLength(a.Title, col/3),
				StringLimitLength(a.ArtistName, col/3)), col-12) + "[ Explicit ](mod:reverse)"
		} else {
			ret[i] = StringLimitLength(fmt.Sprintf("%s - %s",
				StringLimitLength(a.Title, col/3),
				StringLimitLength(a.ArtistName, col/3)), col)
		}
	}
	return ret
}

func FormatArtists(artists []deezer.Artist, col int) []string {
	ret := make([]string, len(artists))
	for i, a := range artists {
		ret[i] = StringLimitLength(a.Name, col-12)
	}
	return ret
}

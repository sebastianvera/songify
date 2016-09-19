package spotify

import (
	"fmt"
	"os/exec"
	"strings"
)

const (
	cmdStart = `if application "Spotify" is running then tell application "Spotify" to`
	cmdTmpl  = "%s %s"

	currentSongID        = "get ID of current track"
	currentSongNameCmd   = "get name of current track"
	currentArtistNameCmd = "get artist of current track"

	currentTrackErr = "(-1728)"
	connectionErr   = "(-609)"
)

// Track ...
type Track struct {
	ID     string `json:"id"`
	Artist string `json:"artist"`
	Name   string `json:"name"`
}

// GetCurrentTrack ...
func GetCurrentTrack() (*Track, error) {
	artist, err := execCmd(currentArtistNameCmd)
	if err != nil {
		return nil, err
	}

	name, err := execCmd(currentSongNameCmd)
	if err != nil {
		return nil, err
	}

	id, err := execCmd(currentSongID)
	if err != nil {
		return nil, err
	}

	return &Track{Artist: string(artist), Name: string(name), ID: string(id)}, nil
}

func execCmd(cmd string) (string, error) {
	cmd = buildCommand(cmd)
	out, err := exec.Command("osascript", "-e", cmd).CombinedOutput()
	if err != nil {
		eo := string(out)
		// TODO: better error handling when closing or opening spotify
		if strings.Contains(eo, connectionErr) || strings.Contains(eo, currentTrackErr) {
			return "", nil
		}

		return "", err
	}

	return strings.TrimSpace(string(out)), nil
}

func buildCommand(cmd string) string {
	return fmt.Sprintf(cmdTmpl, cmdStart, cmd)
}

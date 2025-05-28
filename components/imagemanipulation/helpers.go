package imagemanipulation

import (
	"bytes"
	"io"
	"net/http"

	"github.com/bwmarrin/discordgo"
	"gopkg.in/gographics/imagick.v3/imagick"
)

func ReadFromHttp(url string) (*imagick.MagickWand, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	content, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	wand := imagick.NewMagickWand()
	if err := wand.ReadImageBlob(content); err != nil {
		return nil, err
	}
	return wand, nil
}

func WriteWandToReader(wand *imagick.MagickWand) (io.Reader, error) {
	wand.SetImageFormat("png")

	blob, err := wand.GetImageBlob()
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(blob), nil
}

func OptionDataToWand(inter *discordgo.Interaction, narg int) (*imagick.MagickWand, error) {
	fileID := inter.ApplicationCommandData().Options[0].Options[narg].Value.(string)
	atta := inter.ApplicationCommandData().Resolved.Attachments[fileID]
	return ReadFromHttp(atta.ProxyURL)
}

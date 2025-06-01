package imagemanipulation

import (
	"bytes"
	"errors"
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

func OptionDataToWand(inter *discordgo.Interaction, sublayers int, namearg string) (*imagick.MagickWand, error) {
	var opt *discordgo.ApplicationCommandInteractionDataOption
	if sublayers == 0 {
		opt = inter.ApplicationCommandData().GetOption(namearg)
	} else {
		current := inter.ApplicationCommandData().Options[0]
		for range sublayers - 1 {
			if current == nil || len(current.Options) < 1 {
				return nil, errors.New("cannot unwrap all layers")
			}
			current = current.Options[0]
		}
		opt = current.GetOption(namearg)
	}

	if opt == nil {
		return nil, errors.New("option was not specified")
	}
	fileID := opt.Value.(string)
	atta := inter.ApplicationCommandData().Resolved.Attachments[fileID]
	return ReadFromHttp(atta.ProxyURL)
}

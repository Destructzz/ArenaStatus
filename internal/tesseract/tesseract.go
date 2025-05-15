package tesseract

import (
	"bytes"
	"image"
	"image/png"
	"os/exec"
)

type Tesseract struct{
	Language string
}

func New(language string) *Tesseract{
	return &Tesseract{
		Language: language,
	}
}

func (t *Tesseract) Recognize(img *image.RGBA) (string, error) {
	var buf, out bytes.Buffer

	if err := png.Encode(&buf, img); err != nil{
		return "", err
	}

	cmd := exec.Command("tesseract", "stdin", "stdout", "-l", t.Language)

	cmd.Stdin = &buf
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return "", err
	}

	return out.String(), nil
}
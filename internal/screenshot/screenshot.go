package screenshot

import (
	"ArenaStatus/internal/screenshot/config"
	"ArenaStatus/internal/screenshot/plate"
	"image"
	"log"
	"github.com/kbinani/screenshot"
)

type ScreenshotManager struct {
	config *config.MyConfig
}

func New() *ScreenshotManager{
	return &ScreenshotManager{
		config: config.New(),
	}
}

func (s *ScreenshotManager) screen() (*image.RGBA, error){
	bounds := screenshot.GetDisplayBounds(0)

	img, err := screenshot.CaptureRect(bounds)

	if err != nil{
		return nil, err
	}

	return img, nil
}

func (s *ScreenshotManager)subImage(img *image.RGBA, plate plate.Plate) *image.RGBA{
	sub := img.SubImage(
		image.Rect(
			plate.X0, plate.Y0, plate.X1, plate.Y1,
		),
	)
	rgba, ok := sub.(*image.RGBA)

	if !ok {
		return nil
	}

	return rgba
}

func (s *ScreenshotManager) NickNames() ([]*image.RGBA){
	screen, err := s.screen()

	if err != nil{
		log.Println("Screenshot was not captured", err)
		return nil
	}

	if s.config == nil {
		log.Fatalln("config object was not initialize")
	}

	var imgs []*image.RGBA

	for _, item := range s.config.Plates(){
		imgs = append(imgs, s.subImage(screen, item))
	}
	
	return imgs
}
package application

import (
	filterpipe "ArenaStatus/internal/filter-pipe"
	"ArenaStatus/internal/screenshot"
	"ArenaStatus/internal/tesseract"
	"fmt"
	"image"
	"log"
	"sync"
	"time"

	"github.com/lxn/win"
)

type Application struct {
	tesseract *tesseract.Tesseract
	screenshot *screenshot.ScreenshotManager
	filter *filterpipe.FilterPipe
}

func New() *Application{
	return &Application{
		tesseract: tesseract.New("eng+rus"),
		screenshot: screenshot.New(),
		filter: filterpipe.New(),
	}
}

func testRecognize(img *image.RGBA, tesseract *tesseract.Tesseract) {
	const rounds = 100

	filter := filterpipe.New()
	results := make([]string, 0, rounds)

	for i := 0; i < rounds; i++ {
		nick, err := tesseract.Recognize(img)
		if err != nil {
			log.Println("Unable to recognize nickname:", err)
			continue // пропускаем неудачную попытку
		}
		normalized := filter.Fileter(nick)
		results = append(results, normalized)
	}

	if len(results) == 0 {
		return
	}

	// Сравниваем с первым как эталон
	ref := results[0]
	var matched int
	for _, val := range results {
		if val == ref {
			matched++
		}
	}

	log.Println("occuracy", rounds / matched)
}


func (a *Application)Run(){
	for{
		if uint16(win.GetKeyState(win.VK_RIGHT))&0x8000 != 0 {
			var wg sync.WaitGroup
			img := a.screenshot.NickNames()
			fmt.Println("-------------------------------")
			for _, item := range img{
				nick, err := a.tesseract.Recognize(item)

				wg.Add(1)

				ptr := item
				go func(copy *image.RGBA){
					defer wg.Done()
					testRecognize(copy, a.tesseract)
				}(ptr)

				if err != nil{
					fmt.Println("Recognize err", err)
				}

				nick = a.filter.Fileter(nick)

				fmt.Println(nick)
			}
			wg.Wait()
			fmt.Println("-------------------------------")
			time.Sleep(time.Second * 3)
		}
		time.Sleep(time.Millisecond)
	}
}
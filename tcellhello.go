package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
	"math/rand"
	"log"

	"github.com/gdamore/tcell"
	"github.com/urfave/cli"
)

func ScreenInit() tcell.Screen {
	rand.Seed(time.Now().Unix())
	s, err := tcell.NewScreen()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	err = s.Init()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	return s
}

func main() {

	app := cli.NewApp()

	app.Name = "tcellhello"
	app.Usage = "Hello World for tcell"
	app.Version = "0.0.0"

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name: "debug",
		},
	}

	start := time.Now()

	app.Action = func(c *cli.Context) error {
		s := ScreenInit()
		defer s.Fini()

		eventChan := make(chan tcell.Event)
		go func() {
			for {
				event := s.PollEvent()
				eventChan <- event
			}
		}()

		sigChan := make(chan os.Signal)
		signal.Notify(sigChan, os.Interrupt)
		signal.Notify(sigChan, os.Kill)

		//ticker := time.NewTicker(10 * time.Millisecond)
		ticker := time.NewTicker(1)
		defer ticker.Stop()

		width, height := 10, 16
		lx := 0
		ly := 0
		go func() {
			//for ;; {
			for _ = range ticker.C {

				x := rand.Int() % width
				y := rand.Int() % height

				st := tcell.StyleDefault.Foreground(tcell.Color(rand.Int() % s.Colors()))
				//st := tcell.StyleDefault

				gl := 'X'
				if (rand.Int() & 1) != 0 {
					gl = ' '
				}

				//s.SetCell(lx, ly, tcell.StyleDefault, ' ')
				s.SetCell(x, y, st, gl)
				lx = x
				ly = y

				s.Show()
			}
		}()

		done := false
		for !done {
			select {
			case event := <-eventChan:
				//log.Printf("tcell Event\n")
				switch ev := event.(type) {
				case *tcell.EventKey:
					switch  ev.Key() {
					case tcell.KeyCtrlZ, tcell.KeyCtrlC:
						done = true
						continue
					case tcell.KeyRune:
						switch ev.Rune() {
						case 'q':
							done = true
							continue
						case 'c':
							s.Clear()
						}
					}

				case *tcell.EventResize:
					//log.Printf("tcell Resize Event\n")
					width, height = ev.Size()

				case *tcell.EventError:
					log.Fatalf("tcell Error: %v\n", ev.Error())
					done = true
					continue
				}
			case _ = <-sigChan:
				done = true
				continue
			}
		}
		ticker.Stop()
		fmt.Println("Elapsed", time.Since(start))

		return nil 
	}


	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

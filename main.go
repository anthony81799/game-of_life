package main

import (
	"flag"
	"runtime"
	"time"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

const (
	width  = 500
	height = 500
)

var (
	rows      = 20
	columns   = 20
	seed      = time.Now().UnixNano()
	threshold = 0.2
	fps       = 20
)

func init() {
	flag.IntVar(&columns, "columns", columns, "Sets the number of columns.")
	flag.IntVar(&rows, "rows", rows, "Sets the number of columns.")
	flag.Int64Var(&seed, "seed", seed, "Sets the starting seed of the game, used to randomize the initial state.")
	flag.Float64Var(&threshold, "threshold", threshold, "A percentage between 0 and 1 used in conjunction with the -seed to determine if a cell starts alive. For example, 0.15 means each cell has a 15% chance of starting alive.")
	flag.IntVar(&fps, "fps", fps, "Sets the frames-per-second, used set the speed of the simulation.")
	flag.Parse()
}

func main() {
	runtime.LockOSThread()

	window := initGlfw()
	defer glfw.Terminate()

	program := initOpenGL()

	cells := makeCells(seed, threshold)
	for !window.ShouldClose() {
		t := time.Now()
		aliveCells := 0
		changedCells := 0
		for x := range cells {
			for _, c := range cells[x] {
				c.checkState(cells)
				if c.alive {
					aliveCells++
				}
				if c.stateChanged {
					changedCells++
				}
			}
		}
		if aliveCells == 0 || changedCells == 0 {
			glfw.GetCurrentContext().SetShouldClose(true)
		}
		draw(cells, window, program)
		time.Sleep(time.Second/time.Duration(fps) - time.Since(t))
	}
}

func draw(cells [][]*cell, window *glfw.Window, program uint32) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(program)

	for x := range cells {
		for _, c := range cells[x] {
			c.draw()
		}
	}

	glfw.PollEvents()
	window.SwapBuffers()
}

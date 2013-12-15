package main

import (
  "log"

  "github.com/go-gl/gl"
  glfw "github.com/go-gl/glfw3"
)

const (
  Width  = 800
  Height = 640
  Title  = "GLFW3 play!!"
)

var (
  width, height    int
  centerX, centerY float32
)

func onError(err glfw.ErrorCode, desc string) {
  log.Fatalf("%v: %v\n", err, desc)
}

func main() {
  glfw.SetErrorCallback(onError)

  if !glfw.Init() {
    panic("Can't init glfw!")
  }
  defer glfw.Terminate()

  window, err := glfw.CreateWindow(Width, Height, Title, nil, nil)
  if err != nil {
    panic(err)
  }
  defer window.Destroy()

  window.MakeContextCurrent()

  // Set callbacks
  window.SetKeyCallback(onKey)
  window.SetFramebufferSizeCallback(framebufferSizeCallback)

  // Get window size
  updateDimensions(window)

  // Main loop
  for !window.ShouldClose() {
    drawScene(window)
    window.SwapBuffers()
    // SwapBuffers doesn't call PollEvents() in glfw3
    // This polls for new events (e.g. input)
    glfw.PollEvents()
    // Unlike PollEvents(), WaitEvents() will block (sleep) until an event is received (good for games)
    // which is good if you only want to render on input (good for visual editors).
    // glfw.WaitEvents()
  }
}

func updateDimensions(window *glfw.Window) {
  width, height = window.GetFramebufferSize()
  centerX = float32(width) / 2
  centerY = float32(height) / 2
}

func framebufferSizeCallback(window *glfw.Window, width, height int) {
  ratio := float64(width) / float64(height)

  updateDimensions(window)

  gl.Viewport(0, 0, width, height)
  gl.Clear(gl.COLOR_BUFFER_BIT)
  gl.MatrixMode(gl.PROJECTION)
  gl.LoadIdentity()
  gl.Ortho(-ratio, ratio, -1.0, 1.0, 1.0, -1.0)

  gl.MatrixMode(gl.MODELVIEW)
  gl.LoadIdentity()
}

func onKey(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
  if action == glfw.Press {
    switch key {
    case glfw.KeyEscape:
      window.SetShouldClose(true)
    }
  }
}

func drawScene(window *glfw.Window) {
  var mouseX, mouseY float64

  gl.Clear(gl.COLOR_BUFFER_BIT)
  gl.LoadIdentity()

  // Read timer since glfw.Init()
  time := glfw.GetTime()

  mouseX, mouseY = window.GetCursorPosition()
  mouseRatioX, mouseRatioY := float32(mouseX)/float32(width), float32(mouseY)/float32(height)
  gl.Translatef(mouseRatioX-(centerX/float32(width)), -mouseRatioY+(centerY/float32(height)), 0)
  gl.Rotatef(float32(time)*50.0, -0.2, 0.0, 1.0)
  /* gl.Translatef(-0.9, 0, 0)*/

  gl.Begin(gl.TRIANGLES)
  gl.Color3f(1.0, 0.0, 0.0)
  gl.Vertex3f(-0.6, -0.4, 0.0)
  gl.Color3f(0.0, 1.0, 0.0)
  gl.Vertex3f(0.6, -0.4, 0.0)
  gl.Color3f(0.0, 0.0, 1.0)
  gl.Vertex3f(0.0, 0.6, 0.0)
  gl.End()
}

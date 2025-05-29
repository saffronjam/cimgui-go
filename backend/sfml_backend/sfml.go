package sfml_backend

// #cgo amd64,linux LDFLAGS: -L${SRCDIR}/../../lib/linux/x64 -lsfml-audio -lsfml-graphics -lsfml-network -lsfml-system -lsfml-window -lcsfml-graphics -lcsfml-window -lcsfml-system -lcsfml-audio -lcsfml-network
// #cgo CPPFLAGS: -DCIMGUI_GO_USE_SFML
// #cgo CXXFLAGS: -std=c++17 -I${SRCDIR}/../../thirdparty/SFML/include -I${SRCDIR}/../../cwrappers/imgui
// #cgo linux LDFLAGS: -lstdc++
// #include <stdlib.h>
// #include <stdint.h>
// #include "sfml_backend.h"
import "C"

import (
	"errors"
	"time"
	"unsafe"
)

type SfmlBackend struct {
}

func NewSfmlBackend() *SfmlBackend {
	return &SfmlBackend{}
}

func (b *SfmlBackend) Init(cWindow unsafe.Pointer, loadDefaultFont bool) error {
	cFont := C._Bool(loadDefaultFont)
	if !C.igInit(cWindow, cFont) {
		return errors.New("failed to initialize SFML")
	}
	return nil
}

func (b *SfmlBackend) Shutdown(sfWindow unsafe.Pointer) {
	C.igShutdown(sfWindow)
}

func (b *SfmlBackend) NewFrame(sfWindow unsafe.Pointer, dt time.Duration) {
	cDt := C.uint64_t(dt.Microseconds())
	C.igUpdate(sfWindow, cDt)
}

func (b *SfmlBackend) Render(sfWindow unsafe.Pointer) {
	C.igRender(sfWindow)
}

func (b *SfmlBackend) ProcessEvent(sfWindow unsafe.Pointer, event unsafe.Pointer) {
	C.igProcessEvent(sfWindow, event)
}

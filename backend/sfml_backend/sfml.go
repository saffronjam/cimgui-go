package sfml_backend

// #cgo amd64,linux LDFLAGS: -L${SRCDIR}/../../lib/linux/x64 -lsfml-audio -lsfml-graphics -lsfml-network -lsfml-system -lsfml-window -lcsfml-graphics -lcsfml-window -lcsfml-system -lcsfml-audio -lcsfml-network
// #cgo CPPFLAGS: -DCIMGUI_GO_USE_SFML
// #cgo CXXFLAGS: -std=c++17 -I${SRCDIR}/../../thirdparty/SFML/include
// #cgo linux LDFLAGS: -lstdc++
// #include <stdlib.h>
// #include <stdint.h>
// #include "sfml_backend.h"
import "C"

import (
	"errors"
	"time"
)

type SfmlBackend struct {
}

func NewSfmlBackend() *SfmlBackend {
	return &SfmlBackend{}
}

func (b *SfmlBackend) Init(sfWindow uintptr, loadDefaultFont bool) error {
	if C.igInit(sfWindow, loadDefaultFont) < 0 {
		return errors.New("failed to initialize SFML")
	}
	return nil
}

func (b *SfmlBackend) Shutdown(sfWindow uintptr) {
	C.igShutdown(sfWindow)
}

func (b *SfmlBackend) NewFrame(sfWindow uintptr, dt time.Duration) {
	C.igUpdate(sfWindow, C.double(dt.Microseconds()))
}

func (b *SfmlBackend) Render(sfWindow uintptr) {
	C.igRender(sfWindow)
}

func (b *SfmlBackend) ProcessEvent(sfWindow uintptr, event uintptr) {
	C.igProcessEvent(sfWindow, event)
}

func (b *SfmlBackend) UpdateFontTexture(sfWindow uintptr) error {
	if C.igUpdateFontTexture(sfWindow) < 0 {
		return errors.New("failed to update font texture")
	}

	return nil
}

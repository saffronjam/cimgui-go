package sdlbackend

// #cgo amd64,linux LDFLAGS: -L${SRCDIR}/../../lib/linux/x64 -lsfml-audio -lsfml-graphics -lsfml-network -lsfml-system -lsfml-window -lcsfml-graphics -lcsfml-window -lcsfml-system -lcsfml-audio -lcsfml-network
// #cgo CPPFLAGS: -DCIMGUI_GO_USE_SFML
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
		return errors.New("failed to initialize SDL: " + C.GoString(C.SDL_GetError()))
	}
	return nil
}

func (b *SfmlBackend) Shutdown(sfWindow uintptr) {
	C.igShutdown(sfWindow)
}

func (b *SfmlBackend) NewFrame(sfWindow uintptr, dt time.Duration) error {
	if C.igUpdate(sfWindow, C.double(dt.Microseconds())) < 0 {
		return errors.New("failed to start new frame in SFML")
	}
	return nil
}

func (b *SfmlBackend) Render(sfWindow uintptr) error {
	if C.igRender(sfWindow) < 0 {
		return errors.New("failed to render SFML frame")
	}
	return nil
}

func (b *SfmlBackend) ProcessEvent(sfWindow uintptr, event uintptr) error {
	if C.igProcessEvent(sfWindow, event) < 0 {
		return errors.New("failed to process SFML event")
	}
	return nil
}

func (b *SfmlBackend) UpdateFontTexture(sfWindow uintptr) bool {
	return C.igUpdateFontTexture(sfWindow) != 0
}

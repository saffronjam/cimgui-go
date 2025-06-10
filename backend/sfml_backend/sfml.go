package sfml_backend

// #cgo amd64,linux LDFLAGS: -L${SRCDIR}/../../lib/linux/x64 -lcsfml-system -lcsfml-window -lcsfml-graphics -lcsfml-audio -lcsfml-network -lsfml-audio -lsfml-network -lsfml-graphics -lsfml-window -lsfml-system -lX11 -ldl -lGL -lXrandr -lstdc++ -ludev -lfreetype
// #cgo CPPFLAGS: -DCIMGUI_GO_USE_SFML
// #cgo CXXFLAGS: -I${SRCDIR}/../../thirdparty/SFML/include -I${SRCDIR}/../../cwrappers/imgui
// #include <stdlib.h>
// #include <stdint.h>
// #include "sfml_backend.h"
import "C"

import (
	"errors"
	_ "github.com/saffronjam/cimgui-go/imgui"
	"log"
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
	if !C.igSfmlInit(cWindow, cFont) {
		return errors.New("failed to initialize SFML")
	}
	return nil
}

func (b *SfmlBackend) Shutdown(sfWindow unsafe.Pointer) {
	C.igSfmlShutdown(sfWindow)
}

func (b *SfmlBackend) NewFrame(sfWindow unsafe.Pointer, dt time.Duration) {
	cDt := C.uint64_t(dt.Microseconds())
	C.igSfmlUpdate(sfWindow, cDt)
}

func (b *SfmlBackend) Render(sfWindow unsafe.Pointer) {
	C.igSfmlRender(sfWindow)
}

func (b *SfmlBackend) ProcessEvent(sfWindow unsafe.Pointer, event unsafe.Pointer) {
	C.igSfmlProcessEvent(sfWindow, event)
}

func (b *SfmlBackend) UpdateFontTexture() error {
	if !C.igSfmlUpdateFontTexture() {
		return errors.New("failed to update font texture")
	}
	return nil
}

func main() {
	// This is just a placeholder to ensure the package can be built.
	// The actual usage will be in the GUI package or wherever needed.

	log.Println("hello from SFML backend package")
}

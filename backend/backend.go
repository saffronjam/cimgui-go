package backend

// extern void loopCallback();
// extern void beforeRender();
// extern void afterRender();
// extern void afterCreateContext();
// extern void beforeDestroyContext();
// extern void dropCallback(void*, int, char**);
// extern void keyCallback(void*, int, int, int, int);
// extern void sizeCallback(void*, int, int);
// #include "../imgui/extra_types.h"
// #include "../imgui/wrapper.h"
// #include "../imgui/typedefs.h"
import "C"

import (
	"errors"
	"fmt"
	"image"
	"unsafe"

	"github.com/saffronjam/cimgui-go/imgui"
)

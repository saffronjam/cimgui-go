package main

// #cgo amd64,linux LDFLAGS: -L${SRCDIR}/../../lib/linux/x64 -l:cimgui.a -lsfml-audio -lsfml-graphics -lsfml-network -lsfml-window -lsfml-system -lcsfml-graphics -lcsfml-window -lcsfml-system -lcsfml-audio -lcsfml-network -lX11 -ldl -lGL -lXrandr -lstdc++ -ludev
// #cgo CPPFLAGS: -DCIMGUI_GO_USE_SFML
// #cgo CXXFLAGS: -I${SRCDIR}/../../thirdparty/SFML/include -I${SRCDIR}/../../cwrappers/imgui
// #include <stdlib.h>
// #include <stdint.h>
// #include "test.h"
import "C"

import "log"

func main() {
	log.Println("started with calling just_for_testing():", C.just_for_testing())
}

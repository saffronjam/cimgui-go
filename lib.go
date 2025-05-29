//go:build required
// +build required

package imgui

import (
	_ "github.com/saffronjam/cimgui-go/lib/linux/x64"
	_ "github.com/saffronjam/cimgui-go/lib/macos/arm64"
	_ "github.com/saffronjam/cimgui-go/lib/macos/x64"
	_ "github.com/saffronjam/cimgui-go/lib/windows/x64"
)

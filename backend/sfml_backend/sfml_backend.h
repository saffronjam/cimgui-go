#ifndef IMGUI_SFML_H
#define IMGUI_SFML_H

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

extern bool igInit(void *window, bool loadDefaultFont);
extern void igProcessEvent(void *window, void *event);
extern void igUpdate(void *window, uint64_t dtUs);
extern void igRender(void *window);
extern bool igUpdateFontTexture();
extern void igShutdown(void *window);

#ifdef __cplusplus
}
#endif

#endif // # IMGUI_SFML_H

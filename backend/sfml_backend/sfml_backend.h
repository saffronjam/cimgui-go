#ifndef IMGUI_SFML_H
#define IMGUI_SFML_H

#include <stdint.h>

#ifdef __cplusplus
extern "C"
{
#endif

    extern bool igSfmlInit(void *window, bool loadDefaultFont);
    extern void igSfmlProcessEvent(void *window, void *event);
    extern void igSfmlUpdate(void *window, uint64_t dtUs);
    extern void igSfmlRender(void *window);
    extern bool igSfmlUpdateFontTexture();
    extern void igSfmlShutdown(void *window);

#ifdef __cplusplus
}
#endif

#endif // # IMGUI_SFML_H

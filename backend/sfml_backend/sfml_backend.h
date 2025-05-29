#ifndef IMGUI_SFML_H
#define IMGUI_SFML_H

extern bool igInit(void *window, bool loadDefaultFont);
extern void igProcessEvent(void *window, void *event);
extern void igUpdate(void *window, u_int64_t dtUs);
extern void igRender(void *window);
extern bool igUpdateFontTexture();
extern void igShutdown(void *window);

#endif // # IMGUI_SFML_H

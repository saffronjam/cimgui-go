#ifndef IMGUI_SFML_H
#define IMGUI_SFML_H

namespace sf
{
    class Event;
    class RenderTexture;
    class RenderWindow;
} // namespace sf

extern bool igInit(sf::RenderWindow *window, bool loadDefaultFont = true);
extern void igProcessEvent(sf::RenderWindow *window, sf::Event *event);
extern void igUpdate(sf::RenderWindow *window, u_int64_t dtUs);
extern void igRender(sf::RenderWindow *window);
extern bool igUpdateFontTexture();
extern void igShutdown(sf::RenderWindow *window);

#endif // # IMGUI_SFML_H

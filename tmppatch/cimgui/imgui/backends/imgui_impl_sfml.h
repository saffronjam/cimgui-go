// dear imgui: Platform Backend for SFML

#pragma once
#include "imgui.h" // IMGUI_IMPL_API
#include "stdint.h"
#ifndef IMGUI_DISABLE

namespace sf
{
    class Event;
    class RenderWindow;
    class Time;
} // namespace sf

IMGUI_IMPL_API bool ImGui_ImplSFML_Init(sf::RenderWindow *window, bool load_default_font = true);
IMGUI_IMPL_API void ImGui_ImplSFML_Shutdown(sf::RenderWindow *window);
IMGUI_IMPL_API void ImGui_ImplSFML_ProcessEvent(sf::RenderWindow *window, sf::Event *event);
IMGUI_IMPL_API void ImGui_ImplSFML_Update(sf::RenderWindow *window, sf::Time dt);
IMGUI_IMPL_API void ImGui_ImplSFML_Render(sf::RenderWindow *window);
IMGUI_IMPL_API bool ImGui_ImplSFML_UpdateFontTexture();

#ifdef __cplusplus
extern "C"
{
#endif
    bool cimgui_ImGui_ImplSFML_Init(void *window, bool load_default_font);
    void cimgui_ImGui_ImplSFML_Shutdown(void *window);
    void cimgui_ImGui_ImplSFML_ProcessEvent(void *window, void *event);
    void cimgui_ImGui_ImplSFML_Update(void *window, uint64_t dtUs);
    void cimgui_ImGui_ImplSFML_Render(void *window);
    bool cimgui_UpdateFontTexture();
#ifdef __cplusplus
}
#endif

#endif // #ifndef IMGUI_DISABLE

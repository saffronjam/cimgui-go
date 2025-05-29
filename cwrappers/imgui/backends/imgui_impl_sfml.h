// dear imgui: Platform Backend for SFML

#pragma once
#include "imgui.h" // IMGUI_IMPL_API
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

#endif // #ifndef IMGUI_DISABLE

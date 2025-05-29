#include "sfml_backend.h"

#include "../../cwrappers/imgui/backends/imgui_impl_sfml.h"

#include "../../thirdparty/SFML/include/SFML/Graphics/RenderWindow.hpp"

bool igInit(void *window, bool loadDefaultFont)
{
    auto *sfmlWindow = static_cast<sf::RenderWindow *>(window);
    return ImGui_ImplSFML_Init(sfmlWindow, loadDefaultFont);
}

void igProcessEvent(void *window, void *event)
{
    auto *sfmlWindow = static_cast<sf::RenderWindow *>(window);
    auto *sfmlEvent = static_cast<sf::Event *>(event);
    ImGui_ImplSFML_ProcessEvent(sfmlWindow, sfmlEvent);
}

void igUpdate(void *window, u_int64_t dt)
{
    auto *sfmlWindow = static_cast<sf::RenderWindow *>(window);
    ImGui_ImplSFML_Update(sfmlWindow, sf::microseconds(dt));
}

void igRender(void *window)
{
    auto *sfmlWindow = static_cast<sf::RenderWindow *>(window);
    ImGui_ImplSFML_Render(sfmlWindow);
}

bool igUpdateFontTexture()
{
    return ImGui_ImplSFML_UpdateFontTexture();
}

void igShutdown(void *window)
{
    auto *sfmlWindow = static_cast<sf::RenderWindow *>(window);
    ImGui_ImplSFML_Shutdown(sfmlWindow);
}
#include "sfml_backend.h"

#include "../../cwrappers/imgui/backends/imgui_impl_sfml.h"

#include "../../thirdparty/SFML/include/SFML/Graphics/RenderWindow.hpp"

void igInit(sf::RenderWindow *window, bool loadDefaultFont)
{
    ImGui_ImplSFML_Init(window, loadDefaultFont);
}

void igProcessEvent(sf::RenderWindow *window, sf::Event *event)
{
    ImGui_ImplSFML_ProcessEvent(window, event);
}

void igUpdate(sf::RenderWindow *window, u_int64_t dt)
{
    ImGui_ImplSFML_Update(window, sf::microseconds(dt));
}

void igRender(sf::RenderWindow *window)
{
    ImGui_ImplSFML_Render(window);
}

bool igUpdateFontTexture()
{
    return ImGui_ImplSFML_UpdateFontTexture();
}

void igShutdown(sf::RenderWindow *window)
{
    ImGui_ImplSFML_Shutdown(window);
}
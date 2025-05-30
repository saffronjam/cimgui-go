#include "sfml_backend.h"

#include "../../cwrappers/imgui/backends/imgui_impl_sfml.h"

#include "../../thirdparty/SFML/include/SFML/Graphics/RenderWindow.hpp"

#ifdef __cplusplus
extern "C"
{
#endif

    bool igSfmlInit(void *window, bool loadDefaultFont)
    {
        return cimgui_ImGui_ImplSFML_Init(window, loadDefaultFont);
    }

    void igSfmlProcessEvent(void *window, void *event)
    {
        cimgui_ImGui_ImplSFML_ProcessEvent(window, event);
    }

    void igSfmlUpdate(void *window, uint64_t dt)
    {
        cimgui_ImGui_ImplSFML_Update(window, dt);
    }

    void igSfmlRender(void *window)
    {
        cimgui_ImGui_ImplSFML_Render(window);
    }

    bool igSfmlUpdateFontTexture()
    {
        return cimgui_UpdateFontTexture();
    }

    void igSfmlShutdown(void *window)
    {
        cimgui_ImGui_ImplSFML_Shutdown(window);
    }

#ifdef __cplusplus
}
#endif
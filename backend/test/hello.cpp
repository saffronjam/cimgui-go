
#include "../../cwrappers/imgui/backends/imgui_impl_sfml.h"
#include <stdio.h>

#ifdef __cplusplus
extern "C"
{
#endif

    bool just_for_testing()
    {
        return cimgui_UpdateFontTexture();
    }

#ifdef __cplusplus
}
#endif

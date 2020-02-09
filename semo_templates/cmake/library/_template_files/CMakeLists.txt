#
# Dependencies
#


#
# Sources
#
file(GLOB HEADERS_PUB_ROOT "include/{{ .lib_name }}/*.h")
file(GLOB HEADERS_PUB_SUB "include/{{ .lib_name }}/**/*.h")
set(HEADERS_PUB ${HEADERS_PUB_ROOT} ${HEADERS_PUB_SUB})

set(SOURCES
    ${HEADERS_PUB}
    src/{{ .lib_name }}.cpp
)

source_group(TREE ${CMAKE_CURRENT_SOURCE_DIR} FILES ${SOURCES})

#
# Target
#
add_library({{ .lib_name }} ${SOURCES})
target_include_directories({{ .lib_name }}
    PUBLIC
        include
    PRIVATE
        src
)
# target_link_libraries({{ .lib_name }}
#     PUBLIC
#         some_lib
# )

#
# Install
#
install(
    TARGETS {{ .lib_name }}
    RUNTIME DESTINATION bin
    LIBRARY DESTINATION lib
    ARCHIVE DESTINATION lib
)
install(DIRECTORY include/{{ .lib_name }} DESTINATION include)

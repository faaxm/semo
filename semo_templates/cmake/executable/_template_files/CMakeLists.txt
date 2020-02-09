#
# Dependencies
#

#
# Sources
#
set(SOURCES
    src/main.cpp
)

source_group(TREE ${CMAKE_CURRENT_SOURCE_DIR} FILES ${SOURCES})


#
# Target
#
add_executable({{ .app_name }} ${SOURCES})
#target_link_libraries({{ .app_name }}
#    PRIVATE
#        some_lib
#)

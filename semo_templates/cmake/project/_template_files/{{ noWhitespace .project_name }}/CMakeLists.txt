cmake_minimum_required(VERSION {{ .min_cmake_version }})
project({{ .project_name }} LANGUAGES {{ .languages }})

# option({{ toUpper (noWhitespace .project_name) }}_BUILD_EXAMPLES "Build {{ .project_name }} examples." ON)
# option({{ toUpper (noWhitespace .project_name) }}_BUILD_TESTS "Build {{ .project_name }} unit tests." OFF)

set(CMAKE_MODULE_PATH "${CMAKE_SOURCE_DIR}/cmake")
{{if .cxx_standard}}set(CMAKE_CXX_STANDARD {{ .cxx_standard }}){{end}}

# add_subdirectory(lib)

# if({{ toUpper (noWhitespace .project_name) }}_BUILD_EXAMPLES)
#     add_subdirectory(examples)
# endif()
#
# if({{ toUpper (noWhitespace .project_name) }}_BUILD_TESTS)
#     enable_testing()
#     add_subdirectory(lib/tests)
# endif()

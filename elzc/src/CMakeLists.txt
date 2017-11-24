project(elzc)

add_subdirectory(parser/)
include_directories(parser/)

add_library(elzc_lib
    AstBuilder.cc)
target_link_libraries(elzc_lib
    parser)
add_executable(elzc main.cc)
target_link_libraries(elzc
    elzc_lib)
install(
    TARGETS elzc
    RUNTIME DESTINATION bin
    )

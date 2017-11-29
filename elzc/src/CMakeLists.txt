project(elzc)

add_subdirectory(parser/)
include_directories(parser/)

find_package(LLVM REQUIRED CONFIG)

llvm_map_components_to_libnames(llvm_libs support core irreader)


add_library(elzc_lib
    AstBuilder.cc)
target_link_libraries(elzc_lib
    parser
    ${llvm_libs})
add_executable(elzc main.cc)
target_link_libraries(elzc
    elzc_lib)
install(
    TARGETS elzc
    RUNTIME DESTINATION bin
    )

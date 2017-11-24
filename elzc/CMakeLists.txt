cmake_minimum_required(VERSION 3.9.0)
project(elzc)

include(gtest.cmake) # Embedded googletest(By submodule), so we can use it anywhere.

add_compile_options(-std=c++14) # using c++ 14 std to compile this project

add_subdirectory(src/)
include_directories(src/parser)

file(GLOB_RECURSE TEST_CODES src/*_test.cc)
enable_testing()
add_library(AllTests ${TEST_CODES})
#add_executable(tests ${TEST_CODES})
target_link_libraries(tests
    AllTests
    elzc_lib # elz compiler without main function
    gtest)
add_test(
    NAME testelzcompiler
    COMMAND tests
    )

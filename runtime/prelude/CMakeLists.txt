cmake_minimum_required(VERSION 3.15)
project(prelude)

set(CMAKE_CXX_STANDARD 14)

include_directories(.)


add_library(prelude
        stdio.cpp
        string.cpp
        string.hpp)
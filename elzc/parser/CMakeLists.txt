project(parser)

find_package(Threads)
add_library(parser
    ElzBaseListener.cpp
    ElzParser.cpp
    ElzLexer.cpp
    )
target_link_libraries(parser
    antlr4-runtime
    ${CMAKE_THREAD_LIBS_INIT})

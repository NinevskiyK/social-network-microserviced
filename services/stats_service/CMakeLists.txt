project(stats_service)
set(CMAKE_CXX_STANDARD 20)
set(CMAKE_CXX_COMPILER "/usr/bin/clang++-14")
add_executable(stats_service server.cpp)
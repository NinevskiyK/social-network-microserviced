from ubuntu:22.04

RUN apt-get update && apt-get install libboost-all-dev python3 python3-pip clang++-14 -y
RUN pip3 install cmake
COPY CMakeLists.txt server.cpp /app/
WORKDIR /app/
RUN mkdir build && cd build && cmake .. && make -j 100

ENTRYPOINT [ "/app/build/stats_service" ]
CMD ["0.0.0.0", "8188"]
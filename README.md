### I/O performance tests

Contains tests to read and write small files into the disk and measure the time it was taken to do so.
This aims to mimic the behavior of vm reading configuration during cold start.
Note that both tests don't run in concurrent fashion. Concurrency will be added as a follow up.

Usage:
1. docker build -f Dockerfile . -t $DOCKERREGISTRY/REPO.
2. make your pod use the docker image.
3. run the binary and see output: `./read_small_files`, `./write_small_files`

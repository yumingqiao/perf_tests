### Contains tests to read and write small files into the disk and measure the time it was taken to do so.
1. docker build -f Dockerfile . -t $DOCKERREGISTRY/REPO.
2. make your pod use the docker image.
3. run the binary and see output: `./read_small_files`, `./write_small_files`


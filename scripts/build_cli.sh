#!/bin/bash

# Create build directory if it doesn't exist
BUILD_DIR="build"
ENTRY_FILE="./main.go"

if [ -d "${BUILD_DIR}" ];then
  rm -rf ${BUILD_DIR}
fi

mkdir -p "${BUILD_DIR}"

# Build for all platforms and architectures
for GOOS in darwin linux windows; do
  for GOARCH in 386 amd64 arm arm64; do
    output_name="${BUILD_DIR}/xmchn-mng-${GOOS}_${GOARCH}"
    if [ $GOOS = "windows" ]; then
      output_name="$output_name.exe"
    fi
    env GOOS=$GOOS GOARCH=$GOARCH go build -o $output_name ${ENTRY_FILE} && \
      # Create checksum for compiled binary
      shasum -a 256 $output_name > $output_name.sha256 && \
      # Archive compiled binary and checksum
      tar -czvf $output_name.tar.gz $output_name $output_name.sha256 && \
      # Clean up temporary files
      rm $output_name
  done
done
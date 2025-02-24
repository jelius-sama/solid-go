#!/bin/bash

# Define paths
SRC_FILE="./cmd/main.go"
BUILD_OUTPUT="./bin/solid-go"

# Ensure the bin directory exists
mkdir -p ./bin

# Handle commands
case "$1" in
  build)
    echo "🛠️ Building Frontend..."
    cd frontend || exit
    bun run build
    echo "✅ Frontend build complete"

    echo "🔨 Building Backend..."
    cd .. || exit
    GOOS=linux GOARCH=amd64 go build -ldflags "-X 'solid-go/utils.ENV=prod' -X 'solid-go/utils.PORT=5000'" -o "$BUILD_OUTPUT" "$SRC_FILE"
    echo "✅ Backend build complete: $BUILD_OUTPUT"
    ;;
  dev)
    echo "🛠️ Starting Frontend and Backend server..."
    air & (cd frontend && bun run dev)  & wait
    ;;
  start)
    echo "🚀 Starting server..."
    ./"$BUILD_OUTPUT"
    ;;
  clean)
    echo "🧹 Cleaning build artifacts..."
    cd frontend || exit
    rm -rf dist
    
    cd .. || exit
    rm -f "$BUILD_OUTPUT"
    echo "✅ Clean complete"
    ;;
  *)
    echo "❌ Unknown command: $1"
    echo "Usage: $0 {build|dev|start|clean}"
    exit 1
    ;;
esac

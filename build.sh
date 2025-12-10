#!/bin/bash

APP_NAME="happy_new_year"
OUTPUT_DIR="builds"

echo "🧹 Cleaning up old builds..."
rm -rf $OUTPUT_DIR
mkdir -p $OUTPUT_DIR

echo "🚀 Building for Windows (x64)..."
GOEXPERIMENT=greenteagc GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o $OUTPUT_DIR/${APP_NAME}-windows-amd64.exe .

echo "🚀 Building for Linux (x64)..."
GOEXPERIMENT=greenteagc GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o $OUTPUT_DIR/${APP_NAME}-linux-amd64 .

echo "🚀 Building for macOS (Intel)..."
GOEXPERIMENT=greenteagc GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o $OUTPUT_DIR/${APP_NAME}-macos-intel .

echo "🚀 Building for macOS (Apple Silicon M1/M2)..."
GOEXPERIMENT=greenteagc GOOS=darwin GOARCH=arm64 go build -ldflags "-s -w" -o $OUTPUT_DIR/${APP_NAME}-macos-arm64 .


echo "✅ Done! All files are in the '$OUTPUT_DIR' folder."
ls -lh $OUTPUT_DIR

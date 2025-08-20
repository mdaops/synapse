#!/bin/bash
set -e

if [ ! -d ".syn" ]; then
    echo "Error: .syn directory does not exist"
    exit 1
fi

if [ ! -f "crossplane.yaml" ]; then
    echo "Error: crossplane.yaml not found in current directory"
    exit 1
fi

if [ ! -d "apis" ]; then
    echo "Error: apis directory not found in current directory"
    exit 1
fi

echo "Copying crossplane.yaml to .syn/"
cp crossplane.yaml .syn/

echo "Copying apis/ directory to .syn/"
cp -r apis .syn/

echo "✅ Successfully prepared .syn directory"
echo "Contents of .syn:"
ls -la .syn/

echo ""
echo "Building crossplane package..."
cd .syn

crossplane xpkg build --package-root=. --package-file=../tenancy-config.xpkg

cd ..

echo "✅ Package built successfully: tenancy-config.xpkg"

# Push the package
echo ""
echo "Pushing package to registry..."

# Default registry and tag - you can modify these
REGISTRY="ghcr.io/mdaops"
PACKAGE_NAME="synapse-tenancy"
VERSION=${1:-"latest"}  # Use first argument as version, default to "latest"

FULL_PACKAGE_NAME="${REGISTRY}/${PACKAGE_NAME}:${VERSION}"

crossplane xpkg push ${FULL_PACKAGE_NAME} -f tenancy-config.xpkg

echo "✅ Package pushed successfully to: ${FULL_PACKAGE_NAME}"
echo ""
echo "To install the package:"
echo "crossplane xpkg install configuration ${FULL_PACKAGE_NAME}"

rm -rf .syn/*

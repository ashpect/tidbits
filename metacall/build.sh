#!/bin/bash

# Install Homebrew if not already installed
if ! command -v brew &> /dev/null; then
    /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
fi

# Update brew
brew update

# Install necessary packages
brew install pkg-config openssl libtool libuv boost libxml2 rapidjson swagger-codegen ruby dotnet wget

# Download and install Xcode command line tools
if test ! $(xcode-select -p); then
    echo "Installing Xcode command line tools..."
    sudo xcode-select --install
	sudo xcode-select --switch /Library/Developer/CommandLineTools
fi
echo "Installed Xcode CLI Tools"

mkdir metacall
LOC="$PWD/metacall"

# Build Metacall
build_meta() {
	cd "$LOC" || error "cd $LOC failed"
	echo "Building MetaCall"

	# Export compiler options
	export SDKROOT=$(xcrun --show-sdk-path)
	export MACOSX_DEPLOYMENT_TARGET=''
	export CC=$(xcrun --find clang)
	export CXX=$(xcrun --find clang++)

	# Clone Metacall repository
	git clone https://github.com/metacall/core.git
	cd core
	echo "Cloned metacall repository!"

	# Checkout the latest release tag
	latest_tag=$(git describe --tags `git rev-list --tags --max-count=1`)
	git checkout $latest_tag

	# Create build folder
	mkdir -p "$LOC/core/build"
	cd "$LOC/core/build" || error "cd $LOC/core/build failed"

	# Configure
	cmake -Wno-dev \
		-DCMAKE_BUILD_TYPE=Release \
		-DOPTION_BUILD_SECURITY=OFF \
		-DOPTION_FORK_SAFE=OFF \
		-DOPTION_BUILD_SCRIPTS=OFF \
		-DOPTION_BUILD_TESTS=OFF \
		-DOPTION_BUILD_EXAMPLES=OFF \
		-DOPTION_BUILD_LOADERS_PY=ON \
		-DOPTION_BUILD_LOADERS_NODE=OFF \
		-DOPTION_BUILD_LOADERS_CS=OFF \
		-DOPTION_BUILD_LOADERS_RB=OFF \
		-DOPTION_BUILD_LOADERS_TS=OFF \
		-DOPTION_BUILD_PORTS=ON \
		-DOPTION_BUILD_PORTS_PY=ON \
		-DOPTION_BUILD_PORTS_NODE=OFF \
		-DCMAKE_INSTALL_PREFIX="$LOC" \
		-G "Unix Makefiles" .. || error "Cmake configuration failed."

	# Build
	cmake --build . --target install --config Debug || error "Cmake build target install failed."
}

build_meta

# Set environment variables
export PKG_CONFIG_PATH=/opt/homebrew/bin/pkg-config
export LD_LIBRARY_PATH=/usr/local/lib
export DYLD_LIBRARY_PATH=/usr/local/lib

# Install dependencies in Metacall folder
mkdir lib include
for dep in openssl libtool libuv boost libxml2 rapidjson swagger-codegen ruby dotnet wget; do
  pkg-config --cflags $dep | xargs -I{} cp -v {} include/
  pkg-config --libs $dep | xargs -I{} cp -v {} lib/
done
echo "Dependencies installed in metacall folder."

# Update library links
install_name_tool -id ./libmetacall.dylib lib/libmetacall.dylib
for lib in lib/lib*.dylib; do
  install_name_tool -change /usr/local/lib/libssl.1.0.0.dylib lib/libssl.1.0.0.dylib $lib
  install_name_tool -change /usr/local/lib/libcrypto.1.0.0.dylib lib/libcrypto.1.0.0.dylib $lib
  install_name_tool -change /usr/local/opt/libuv/lib/libuv.1.dylib lib/libuv.1.dylib $lib
  install_name_tool -change /usr/local/opt/boost/lib/libboost_filesystem.dylib lib/libboost_filesystem.dylib $lib
  install_name_tool -change /usr/local/opt/boost/lib/libboost_system.dylib lib/libboost_system.dylib $lib
  install_name_tool -change /usr/local/opt/rapidjson/lib/librapidjson.dylib lib/librapidjson.dylib $lib
  install_name_tool -change /usr/local/opt/ruby/lib/libruby.2.7.dylib lib/libruby.2.7.dylib $lib
done
echo "Updated Library links"

# Install brew-pkg
brew tap timsutton/formulae
brew install brew-pkg

# Wrap up and compress the executable into tar file
# brew pkg --identifier com.metacall.metacall --root=/usr/local/Cellar/metacall/$latest_tag .

echo "Package created successfully!"

# Prepare the release for the above package and use this link to create a recipe
# brew create https://github.com/metacall/core/releases/download/$latest_tag/metacall-$latest_tag-runtime.tar.gz

# Prepare the metacall formula in brew
# brew tap brillard1/homebrew-distributable # format <user/repo>
# brew install metacall

echo "Formula created for metacall distributable macos."

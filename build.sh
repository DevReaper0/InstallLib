#!/bin/bash

# rm -r bin

cd src

# windows linux darwin android aix dragonfly freebsd hurd illumos ios js nacl netbsd openbsd plan9 solaris zos
#gooses=( windows linux darwin android freebsd openbsd solaris )
gooses=( linux )
gooses=($(printf -- '%s\n' "${gooses[@]}" | sort))

# This list is based on:
# https://github.com/golang/go/blob/master/src/go/build/syslist.go
# 386 amd64 amd64p32 arm armbe arm64 arm64be ppc64 ppc64le loong64 mips mipsle mips64 mips64le mips64p32 mips64p32le ppc riscv riscv64 s390 s390x sparc sparc64 wasm
#goarches=( 386 amd64 arm arm64 ppc64 s390x wasm )
goarches=( amd64 )
goarches=($(printf -- '%s\n' "${goarches[@]}" | sort))

for goos in "${gooses[@]}"
do
    for goarch in "${goarches[@]}"
    do
        echo "Building for $goos/$goarch..."
        GOOS="$goos"
        GOARCH="$goarch"
        filepath="../bin/$goos/InstallLib-$goos-$goarch"
        if [[ "$filepath" == *"windows"* ]]; then
          filepath+=".exe"
        fi
        go build -o "$filepath" .
    done
    cp ../install.sh "../bin/$goos"
    cp -r ../fonts "../bin/$goos"
done


# platforms=( $(echo $(go tool dist list) | tr '\n' ' ') )
# for platform in "${platforms[@]}"
# do
#     platform_split=(${platform//\// })
#     GOOS=${platform_split[0]}
#     GOARCH=${platform_split[1]}
#     echo "Building for $GOOS/$GOARCH..."
#     filepath="../bin/$GOOS/InstallLib-$GOOS-$GOARCH"
#     if [ $GOOS = "windows" ]; then
#         filepath+='.exe'
#     fi    
#     echo $filepath

#     env GOOS=$GOOS GOARCH=$GOARCH go build -o "$filepath" .
#     if [ $? -ne 0 ]; then
#            echo 'An error has occurred! Aborting the script execution...'
#         exit 1
#     fi
# done

cd ..

echo "Finished building."

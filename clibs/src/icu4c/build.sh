#!/bin/bash

# Provide a all lowercased OS name for everyone that sources us.
if [[ "x`uname`" == "xMINGW32_NT-6.1" ]]; then
	export OS_NAME="windows"
else
	export OS_NAME=`uname | tr '[:upper:]' '[:lower:]'`
fi

if [ `uname -m` == "x86_64" ]; then
	export ARCH=amd64
else
	export ARCH=i386
fi

if [[ $1 == "jenkins" ]]; then
  export CLIBS_HOME="$HOME/userContent/clibs/$OS_NAME-$ARCH"
else
	[ -z $CLIBS_HOME ] && [ ! -z $MOOV_HOME ] && export CLIBS_HOME=$MOOV_HOME/clibs
fi

if [ -z $CLIBS_HOME ]; then
	echo "Please set CLIBS_HOME or MOOV_HOME before running this script."
	exit 1
fi

if [[ "x`uname`" == xMINGW32_NT* ]]; then
	CLIBS_HOME=$(echo "$CLIBS_HOME" | sed 's/\\/\//g' | sed -r 's/(^[^\/]):/\/\1/')
fi

if [[ "`uname`" == Darwin* ]]; then
	export LIBTOOL=`which glibtool`
	export LIBTOOLIZE=`which glibtoolize`
fi

cd source
[ ! -d $CLIBS_HOME ] && mkdir -p $CLIBS_HOME

# make an output folder just for this lib
echo "Placing output in $CLIBS_HOME/output/icu4c"
rm -rf "$CLIBS_HOME/output/icu4c"
mkdir -p "$CLIBS_HOME/output/icu4c"

./configure --prefix="$CLIBS_HOME/output/icu4c"
if [[ "x`uname`" == xMINGW32_NT* ]]; then
	#mkdir -p lib
	#cd stubdata && make install && cd ..
	#cd common && make install && cd ..
	#cd i18n && make install && cd ..
	make install -i
	cp $CLIBS_HOME/output/icu4c/lib/*.dll $CLIBS_HOME/output/icu4c/bin # ICU installs DLLs in the wrong place
else
	make install
fi

# empty the dumping ground and re-copy all the latest clib outputs into it
echo "Removing clib output folders under $CLIBS_HOME"
rm -rf "$CLIBS_HOME/bin"
rm -rf "$CLIBS_HOME/include"
rm -rf "$CLIBS_HOME/lib"
rm -rf "$CLIBS_HOME/sbin"
rm -rf "$CLIBS_HOME/share"
for f in "$CLIBS_HOME/output/*/*"
do
	echo "Moving $f to $CLIBS_HOME"
  cp -R $f $CLIBS_HOME
done

#!/bin/bash
set -e

cp /home/gradle/build/dependencies-licences.gradle /home/gradle/project/

# Download dependencies
#exec gradle compileJava &&
exec gradle "$@"
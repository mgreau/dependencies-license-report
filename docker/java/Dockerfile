FROM gradle:4.1.0-jdk8-alpine

ENV GRADLE_FILE = "/home/gradle/build/dependencies-licences.gradle"

COPY ./config /home/gradle/build

# script to be able to copy the gradle file to the mount project
COPY docker-entrypoint.sh /usr/bin/docker-entrypoint

WORKDIR "/home/gradle/project/"

ENTRYPOINT [ "/usr/bin/docker-entrypoint" ]

CMD [ "-b", "dependencies-licences.gradle", ":downloadLicenses", "--stacktrace", "-debug" ]
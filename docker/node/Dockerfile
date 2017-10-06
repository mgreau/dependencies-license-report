FROM node:6.11.3-slim

# git is required by npm
RUN apt-get update && apt-get install -y --no-install-recommends \
		git \
    && npm i npm-license-crawler -g \
	&& rm -rf /var/lib/apt/lists/*

USER node

WORKDIR /home/node

RUN mkdir -p /home/node/reports

ENTRYPOINT [ "npm-license-crawler" ]

CMD [ "--dependencies", "--json", "/home/node/reports/dependencies-licenses.json"]

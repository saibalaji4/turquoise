# Turquoise

This is a sample blue/green application. Based on the version you release, it will
show it under the /color endpoint. Use the Makefile to build the image and tag
it based on the version.

## Run
1. Go to [joatmon08/turquoise](https://hub.docker.com/r/joatmon08/turquoise) on Docker Hub for version information.
1. `docker run -d -e TURQUOISE_APP_PORT=8080 joatmon08/turquoise:<tag>`

## API
| Endpoint       | Description           |
| :------------- |:-------------|
| /     | shows the app name and healthy status |
| /color     | shows the version |

## Updates
1. Change the version under "version" file.
1. Issue `make golang-build` to build the binary and container.
1. Issue `make push` to push it to Docker Hub.
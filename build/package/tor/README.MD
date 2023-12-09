# Tor container

## Build

*Directory that the docker container is in:*

- ```sh
    docker build -t tor .
  ```

*Root directory:*

- ```sh
    docker build -t tor -f build/package/tor/Dockerfile .
  ```

## Run

This will publish the docker container on localhost, so that the containers can communicate. Adjust port accordingly.

```sh
docker run --network host --publish 9050:9050 tor
```
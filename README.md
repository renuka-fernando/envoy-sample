# 1. Envoy Sample
Envoy sample demonstration

## 1.1. Prerequisite

Go Lang 1.18
[Envoy binary](https://www.envoyproxy.io/docs/envoy/latest/start/install)
[Backend Server - Echo](https://github.com/cake-baker/request-info)

## 1.2. Static Sample

Update [envoy-static.yaml](envoy-static.yaml) file. Execute following.

-   Envoy
    ```sh
    envoy -c envoy-static.yaml
    ```

    or with docker (please update upstream address URLs as `host.docker.internal`)

    ```sh
    docker run --rm -it -v $(pwd)/envoy-static.yaml:/etc/envoy/envoy-static.yaml -p 18080:18080 -p 19000:19000 envoyproxy/envoy:v1.20.2 -c /etc/envoy/envoy-static.yaml
    ```

-   Ext authz Server - gRPC
    ```sh
    cd ext-auth-server
    go run main/main.go
    ```

-   Start echo backend.
    ```sh
    cd ~/git/request-info/
    go run main.go -pretty
    ```

    or with docker

    ```sh
    docker run --rm -it -p 8080:8080 cakebakery/request-info:v1 -pretty
    ```

-   Invoke backend through envoy.
    ```sh
    curl http://localhost:18080/foo\?hello\=world \
        -d '{"data":"hellllo world"}' \
        -H 'Authorization: Bearer token1' \
        -H "foo: bar" -H "foo: baz" -H "X-Forwarded-For: 192.168.7.7" -v
    ```

## 1.3. Dynamic Sample - with xDS Server

-   xDS Server - gRPC
    ```sh
    cd xds-server
    go run main/main.go -debug
    ```

-   Ext authz Server - gRPC
    ```sh
    cd ext-auth-server
    go run main/main.go
    ```

-   Envoy
    ```sh
    envoy -c envoy-dynamic.yaml
    ```

    or with docker (please update upstream address URLs as `host.docker.internal`)

    ```sh
    docker run --rm -it -v $(pwd)/envoy-dynamic.yaml:/etc/envoy/envoy-dynamic.yaml -p 18080:18080 -p 19000:19000 envoyproxy/envoy:v1.20.2 -c /etc/envoy/envoy-dynamic.yaml
    ```

-   Start echo backend.
    ```sh
    cd ~/git/request-info/
    go run main.go -pretty
    ```

    or with docker

    ```sh
    docker run --rm -it -p 8080:8080 cakebakery/request-info:v1 -pretty
    ```

-   cURL
    ```sh
    curl http://localhost:18080/foo -d '{"data":"hellllo world"}' -H 'Authorization: Bearer token1' -v
    ```

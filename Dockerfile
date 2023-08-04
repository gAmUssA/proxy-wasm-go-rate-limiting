ARG TC_KONG_IMAGE
FROM ${TC_KONG_IMAGE:-kong/kong-gateway-dev:3.4.0.0-rc.1}

RUN mkdir -p /usr/local/kong/wasm

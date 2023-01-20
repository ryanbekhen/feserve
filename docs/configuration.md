# Configuration

English | [Indonesia](configuration-id.md)

```yaml
# configuration example
version: 1
host: 0.0.0.0
port: 8000
headers: 
  X-Custom-Header: "hi"
allowOrigins: "https://gofiber.io, https://gofiber.net"
timezone: Asia/Jakarta
publicDir: public
proxyHeader: CF-Connecting-IP
routes:
  - path: /
    file: index.html
  - path: /about
    file: about.html
  - path: /myjs
    file: myjavascript.js
```

## Version

The version here is not the application version but the configuration version.

## Host

If we set the `host` configuration with a certain IP then the application can only be accessed with that IP. For example, if we set `host` with IP `192.168.1.1`, then when accessing it we have to use the url <http://192.168.1.1:8000>. When we access it via <http://127.0.0.1:8000> it will get an error `ERR_CONNECTION_REFUSED` in your browser. If we don't set it, by default the application can be accessed via various IPs on your computer.

## Port

If we do not set the `port` configuration then by default the application runs on port `8000`.

## Headers

This configuration to do custom header on response.

## Allow Origins

To enable cors, set the configuration with the origin you want by simply writing `allowOrigins: "https://example.com"`. If you want multiple origins, separate them separately using commas (`,`) like this `allowOrigins: "https://example.com, https://example.net"`. By default cors is empty.

## Timezone

This configuration is to make the time format on the log with the time zone that we specify. For example, we set `timezone` with `Asia/Jakarta` timezone, then the log format will be like this.

```shell
[2023-01-17T17:34:33+07:00] -  - 200 GET / Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36
```

If not set, it will default to `UTC` timezone.

## Public Directory

In the written configuration `publicDir`, this configuration is to specify the location of the `public` directory. For example, when we use React.js, the build directory is named `build`, we can simply set `publicDir` with the name of the build directory. If not set by default, the directory is `public`.

## Proxy Header

This configuration is used to get the user's IP if the host is running behind a Load Balancer. For example, if we use Cloudflare, usually the user's IP is set to a header with the name `CF-Connecting-IP` then in the `proxyHeader` configuration it is set to `proxyHeader: CF-Connecting-IP`. By default the value is empty.

## Routes

By default if we do not set the `routes` configuration, Feserve will use the following configuration.

```yaml
routes:
  - path: *
    file: index.html
```

With the configuration above, every url and parameter will be responded with an `index.html` file, perfect if we use SPA.

You can also manipulate the url as follows.

```yaml
routes:
  - path: /
    file: index.html
  - path: /about
    file: about.html
  - path: /myjs
    file: myjavascript.js
```

With the configuration above when we access `/about` the file that is responded is `about.html` which in fact we can also access it with `/about.html`. Here we also see from the configuration above that if we access `/myjs` we will get a javascript response file of `myjavascript.js`.

If you want the feserve to be a load balancer with a config like the following:

```yaml
routes:
  - path: "*"
    file: index.html
  - path: /payment
    balancer:
      - http://localhost:3000
      - http://localhost:3001
```

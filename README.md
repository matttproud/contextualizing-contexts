This repository contains code samples to demonstrate vagaries around Go's
context API.

## Structure

* `{grpc,http}-in-series`: shows the behavior of a function with *manual
  deadline* management juxtaposed with the context's native cancellation
  mechanisms as implemented respectively through gRPC and HTTP server libraries.

* `{grpc,http}-branch`: shows the behavior of the same function from
  `{grpc,http}-in-series` run in a separate goroutine from the respective server
  libraries to connect how the server's dispatch functions handle context
  lifetime.

* `{grpc,http}-notify`: shows the lifetime behaviors of the context from the
  respective server libraries to establish when lifetime begins and ends in
  isolation.

* `{grpc,http}-deadline`: reveals what deadline the server library attaches to
  the context (if any).

  `http-deadline-synth` demonstrates a deadline with the HTTP stack using
  a homegrown `X-MTP-Deadline` request header (uses `http.TimeFormat` values).

  ```
  $ curl -H "X-MTP-Deadline: $(LC_ALL=C TZ=GMT date -d 'now + 1 minutes' '+%a, %d %b %Y %T %Z')" localhost:8080
  ```

The server samples serve on `localhost:8080`:

* Clients can exercise the HTTP servers with curl:

  ```
  $ curl localhost:8080
  ```

* Clients can exercise the gRPC servers the gRPC CLI ([`grpc_cli`]):

  ```
  $ grpc_cli call localhost:8080 proto.Test.Exercise ''
  ```

Try exercising the gRPC servers with varying values of the `grpc_cli`'s
`--timeout` flag.  There is [no default timeout/deadline].

[`grpc_cli`]: https://github.com/grpc/grpc/blob/master/doc/command_line_tool.md
[no default timeout/deadline]: https://grpc.io/docs/guides/deadlines/#deadlines-on-the-client
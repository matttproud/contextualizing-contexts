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

The server samples serve to `:8080`:

* Clients can exercise the HTTP serverswith curl:

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
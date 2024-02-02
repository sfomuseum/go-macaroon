# go-macaroon

Code for testing Macaroons and the superfly/macaroon package.

## Important

This is not production code. It's not even useful package code. It is _example_ code to test using the parts of the [superfly/macaroon](https://pkg.go.dev/github.com/superfly/macaroon) package that aren't (don't seem to be) specific to the Fly.io service.

## Motivation

If you don't know what "macaroons" are then you should start by reading the following documents:

* [Macaroons Escalated Quickly](https://fly.io/blog/macaroons-escalated-quickly/)
* [Integral Principles of the Structural Dynamics of Macaroons](https://github.com/superfly/macaroon/blob/main/macaroon-thought.md)
* [Third Party Discharge Protocol](https://github.com/superfly/macaroon/tree/main/tp)

This package exposes tools and code to test using the parts of the [superfly/macaroon](https://pkg.go.dev/github.com/superfly/macaroon) package that aren't (don't seem to be) specific to the Fly.io service. It is being released as-is in the spirit of generousity and because it might be helpful or interesting to others. Suggestions, bug-fixes and (gentle) cluebats are welcome.

Importantly, the `superfly/macaroon` package comes with the following warning:

> We don't think you should use any of this code; it's shrink-wrapped around some peculiar details of our production network, and the data model is Fly-specific. But if it's an interesting read, that's great too.

So, buyer beware. I have gotten basic custom (non-Fly.io) caveat creation and testing both locally and with third-party discharges working so that bodes well for the idea of common code not bound to any one service but it's early days.

## Tools

### hello-world

Demonstrate basic custom (non-Fly.io) caveat creation and access validation locally.

```
$> go run cmd/hello-world/main.go 
2024/02/02 12:02:36 Expires 2024-02-02 12:02:46 -0800 PST
2024/02/02 12:02:36 Sleeping...
2024/02/02 12:02:45 OK alice
2024/02/02 12:02:45 OK bob
2024/02/02 12:02:45 DENY doug
2024/02/02 12:02:45 Sleep again...
2024/02/02 12:02:47 Failed to validate bob, unauthorized: token only valid until 2024-02-02 12:02:46 -0800 PST; unauthorized: token only valid until 2024-02-02 12:02:46 -0800 PST
```

### tp-discharge

Demonstrate basic custom (non-Fly.io) caveat creation and access validation locally and with a (simulated) third-party discharge exchange.

```
$> go run cmd/tp-discharge/main.go 
2024/02/02 12:03:09 Sleeping 2 seconds to simulate exchanging macaroon
2024/02/02 12:03:11 Sleeping 2 seconds to simulate sending discharge ticket
2024/02/02 12:03:13 Sleeping 2 seconds to simulate receiving discharge token
2024/02/02 12:03:15 All caveats validate
2024/02/02 12:03:15 Sleeping 4 seconds to trigger expiration
2024/02/02 12:03:19 Macaroon expired, unauthorized: token only valid until 2024-02-02 12:03:19 -0800 PST; unauthorized: token only valid until 2024-02-02 12:03:19 -0800 PST
```

## See also

* https://pkg.go.dev/github.com/superfly/macaroon
* https://fly.io/blog/macaroons-escalated-quickly/
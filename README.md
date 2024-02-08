# go-macaroon

Go package for working with Macaroons and the [superfly/macaroon](https://pkg.go.dev/github.com/superfly/macaroon) package.

## Documentation

Documentation is incomplete. Consult the tests and in particular [caveats/caveats_test.go](caveats/caveats_test.go) to get started.

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

### new-macaroon

```
$> ./bin/new-macaroon -h
Generate a new Macaroon token and emit it as a base64-encoded string.
Usage:
	 ./bin/new-macaroon [options]
  -duration string
    	A valid ISO8061 duration time used to set the time-to-live for the Macaroon token. (default "PT10M")
  -location string
    	The 'location' string to associate with the Macaroon token. (default "sfomuseum.org")
  -signing-key-uri string
    	A valid sfomuseum/runtimevar URI that contains the signing key for the Macaroon token.
  -urlescape
    	A boolean flag to URL escape the final base64-encoded Macaroon token.
```

For example:

```
$> ./bin/new-macaroon -signing-key-uri file:///usr/local/sfomuseum/go-macaroon/fixtures/signing.key
lJPAxBDUE507S463PTwmomfDUVIlwq1zZm9tdXNldW0ub3JnkgSSzmXFQLPOZcVDC8Qgp8rB2CYGZ0o6El7wOQtnfcgMB80FvT3Vv2If5Pj6hss=
```

See also:

* https://github.com/sfomuseum/runtimevar

### tp-ensure-account

Append a "ensure account" caveat to discharge with a third-party to a Macaroon token and emit it as a base64-encoded string.

```
$> ./bin/tp-ensure-account -h
Append a "ensure account" caveat to discharge with a third-party to a Macaroon token and emit it as a base64-encoded string.
Usage:
	 ./bin/tp-ensure-account [options]
  -account-id int
    	The account ID to assign to the third-party caveat.
  -encryption-key-uri string
    	A valid sfomuseum/runtimevar URI that contains the shared encryption key for the third-party caveat.
  -location string
    	The 'location' string to associate with the Macaroon token. (default "")
  -macaroon string
    	A base64-encoded string containing the Macaroon token to update. If the value is '-' then data will be read from STDIN.
  -urlescape
    	A boolean flag to URL unescape the base64-encoded Macaroon token being updated.
  -urlunescape
    	A boolean flag to URL escape the final base64-encoded Macaroon token.
```

For example (reading data from `STDIN`):

```
$> ./bin/new-macaroon \
	-signing-key-uri file:///usr/local/sfomuseum/go-macaroon/fixtures/signing.key \
	-duration PT30S \
	| \
	./bin/tp-ensure-account -encryption-key-uri file:///usr/local/sfomuseum/go-macaroon/fixtures/3p-encryption.key \
	-location example.com \
	-macaroon -

lJPAxBA2ZfumuM+3ye8ksRwEfMqKwq1zZm9tdXNldW0ub3JnlASSzmXFQwfOZcVDJQuTq2V4YW1wbGUuY29txDxfN+zomqLyS1MH2w5OkYRKPFnz7d6UwLBbbYiOLdd9qIhlnzz+B9HuuzIYQcqjlEjDpFo+2kzNTALYxrzETp0uDb6u3XHQ5nx4xDM8I6zNOr9skiIUpLWlGQWc32rL56ILzSBSfIUvVkngOEuyG/bY58s2KQltpggV4IsGNwE4IhyjWB+zns8fW8C09MQghGB+25sWgzazuqI0sLNAd4in5pUves6nT179GjL1HY0=
```

See also

* https://github.com/sfomuseum/runtimevar

### signing-key

Generate a random base64-encoded Macaroon signing key.

```
$> ./bin/signing-key -h
Generate a random base64-encoded Macaroon signing key.
Usage:
	 ./bin/signing-key
```

For example:

```
$> ./bin/signing-key 
SLDGG+DU9wAXQz5l1bYmymEhyRpEyK1f4wrXs58N4iw=
```

### encryption-key

Generate a random base64-encoded Macaroon encyption key.

```
$> ./bin/encryption-key -h
Generate a random base64-encoded Macaroon encyption key.
Usage:
	 ./bin/encryption-key
```

For example:

```
$> ./bin/encryption-key 
hMonc06ho508zB+Hn3N70jg7kkKvJu3IrDZxJxcgSrQ=
```

## See also

* https://pkg.go.dev/github.com/superfly/macaroon
* https://fly.io/blog/macaroons-escalated-quickly/
* https://github.com/sfomuseum/runtimevar
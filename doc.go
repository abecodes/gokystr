/*
package gokystr is a simple implementation of a keystore for private keys.

A new keystore instance is created of a folder on the filesystem.
The keystore adds private keyfiles present in the folder on initialization
automatically to the internal store, if they are parseable.

The name of the keyfile is it's internal ID.

By default private keys will be recognized via the fileextension ".pem".
*/
package gokystr

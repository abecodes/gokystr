# :package: Gokystr

**Gokystr** is a simple implementation of a filebased keystore for private keys.

On initialization the keystore reads the private keys from a folder and keeps them in memory.

Additional added keys can be written back to the file system, if enabled.

The ID for the key equals the filename minus the fileextension.

## :computer: Example

```go
package main

import "github.com/abecodes/gokystr"

func main() {
	kystr, err := gokystr.New("/folder/with/keyfiles")

	// Retrieving a key
	prvKey, err := kystr.PrivateKey("keyname")
	pubKey, err := kystr.PublicKey("keyname")

	// Adding a key
	err := kystr.Add("keyname", *rsa.PrivateKey)

	// Removing a key
	err := kystr.Remove("keyid")
}

```

## :clipboard: Options

| Option | Info | Default |
| --- | --- | --- |
| AcceptEmpty | If true, the keystore will not throw if the passed folder does not contain any keys | false |
| NoWrite | If true, additional added keys will not be written to the filesystem | false |
| SetKeyExt | Change the fileextension used to identify keyfiles | .pem |


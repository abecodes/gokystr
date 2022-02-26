package gokystr

type option func(k *Keystore) error

//revive:disable:unexported-return
// SetKeyExt changes the default "pem" extension used to identify keyfiles
func SetKeyExt(ext string) option {
	return func(k *Keystore) error {
		if !validateExt(ext) {
			return ErrWrongExtFormat
		}
		k.keyExt = ext

		return nil
	}
}

// AcceptEmpty prevents the keystore from throwing when using a empty dir
func AcceptEmpty() option {
	return func(k *Keystore) error {
		k.acceptEmpty = true

		return nil
	}
}

// NoWrite will prevent the keystore from writing keyfiles
func NoWrite() option {
	return func(k *Keystore) error {
		k.write = false

		return nil
	}
}

//revive:enable:unexported-return

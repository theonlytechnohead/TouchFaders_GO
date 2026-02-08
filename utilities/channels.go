package utilities

/* https://stackoverflow.com/a/75962269/8705144 */
func ReadChannel[T any](channel <-chan T) (value T) {
	select {
	case value, _ := <-channel:
		return value
	default:
		var zeroT T
		return zeroT
	}
}

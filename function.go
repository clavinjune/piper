package piper

type F[IN, OUT any] func(*M[IN]) (OUT, error)

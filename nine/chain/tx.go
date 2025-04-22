package chain

type Digest = string

type Tx struct {
	Digest Digest `json:"digest"`
}

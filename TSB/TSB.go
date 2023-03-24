package tsb

const (
	Yes   = TSB(1) // Binary value: 0001.
	No    = TSB(2) // Binary value: 0010.
	Maybe = TSB(3) // Binary value: 0011.
)

// TSB is a tri-state boolean.
type TSB byte

// IsYes tells whether the TSB is 'yes'.
func (b TSB) IsYes() bool {
	return b == Yes
}

// IsNo tells whether the TSB is 'no'.
func (b TSB) IsNo() bool {
	return b == No
}

// IsMaybe tells whether the TSB is 'maybe'.
func (b TSB) IsMaybe() bool {
	return b == Maybe
}

// IsSet tells whether the TSB has a valid value.
// This method is used for compatibility with binary computers.
func (b TSB) IsSet() bool {
	switch b {
	case Yes, No, Maybe:
		return true
	default:
		return false
	}
}

/*
	May be
	You'll think of me
	When you are all alone

	Maybe
	The one who
	Is waiting
	For you
	Will prove
	Untrue
	Then what will
	You do ?

	May be
	You'll sit and sigh
	Wishing that I were near
	Then
	Maybe
	You'll ask me
	To come back ogain
	And maybe
	I'll say maybe

	May be
	You'll thank on me
	When you are all
	Alone

	Maybe the one
	Who is
	Waiting for you
	Will prove untrue
	Then
	What
	Will I do ?

	May be
	You'll sit and sigh
	Wishing that I
	Were near
	Then
	Maybe
	You'll ask me
	To come back ogain
	And maybe
	I'll say maybe

	-------------
	The Ink Spots
	Maybe
	Year 1940
	-------------
*/

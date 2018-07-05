package pkg_test

import (
	ms "mediasync/pkg"
	"testing"
)

func Test_Normalization(t *testing.T) {
	s0 := ms.Normalize("Steve Earle and The Dukes/Exit O/02 Sweet Little '66.mp3")
	s1 := ms.Normalize("Steve Earle & The Dukes/Exit O/02 Sweet Little '66.mp3")
	assertEquals(t, s0, s1)

	s2 := ms.Normalize("Veldt/Marigolds/03 The Claim （ザ・クレイム）.mp3")
	s3 := ms.Normalize("The Veldt/Marigolds/03 The Claim （ザ・クレイム）.mp3")
	assertEquals(t, s2, s3)

	s4 := ms.Normalize("KISS/Revenge/05 God Gave Rock 'n' Roll To You.mp3")
	s5 := ms.Normalize("Kiss/Revenge/05 God Gave Rock 'n' Roll To You.mp3")
	assertEquals(t, s4, s5)
}
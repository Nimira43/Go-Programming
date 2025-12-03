package main

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expect deck length of 52, but got", len(d))
	}
}
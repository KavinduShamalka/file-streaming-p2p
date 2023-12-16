package main

import (
	"bytes"
	"testing"
)

func TestPathTransformFunc(t *testing.T) {
	key := "momsbestpicture"
	pathKey := CASPathTransformFun(key)
	expectedOriginalKey := "6804429f74181a63c50c3d81d733a12f14a353ff"
	expectedPathName := "68044/29f74/181a6/3c50c/3d81d/733a1/2f14a/353ff"
	if pathKey.Pathname != expectedPathName {
		t.Errorf("have %s want %s", pathKey.Pathname, expectedPathName)
	}

	if pathKey.Original != expectedPathName {
		t.Errorf("have %s want %s", pathKey.Original, expectedOriginalKey)
	}
}

func TestStore(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFun,
	}

	s := NewStore(opts)

	data := bytes.NewReader([]byte("some jpg bytes"))

	if err := s.writeStream("myspecialpicture", data); err != nil {
		t.Error(err)
	}
}

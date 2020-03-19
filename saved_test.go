package instago

import (
	"fmt"
	"os"
	"testing"
)

func ExampleGetSavedPosts(t *testing.T) {
	mgr, err := NewInstagramApiManager("auth.json")
	if err != nil {
		t.Error(err)
		return
	}

	items, err := mgr.GetSavedPosts(20)
	if err != nil {
		t.Error(err)
		return
	}

	for _, item := range items {
		fmt.Println(item.GetPostUrl())
		fmt.Println(item.SavedCollectionIds)
	}
}

func ExampleGetSavedCollection(t *testing.T) {
	mgr, err := NewInstagramApiManager("auth.json")
	if err != nil {
		t.Error(err)
		return
	}

	items, err := mgr.GetSavedCollection(os.Getenv("COLLECTION_ID"))
	if err != nil {
		t.Error(err)
		return
	}

	for _, item := range items {
		fmt.Println(item.GetPostUrl())
		fmt.Println(item.SavedCollectionIds)
	}
}

func ExampleGetSavedCollectionList(t *testing.T) {
	mgr, err := NewInstagramApiManager("auth.json")
	if err != nil {
		t.Error(err)
		return
	}

	collections, err := mgr.GetSavedCollectionList()
	if err != nil {
		t.Error(err)
		return
	}

	for _, collection := range collections {
		fmt.Println(collection)
	}
}

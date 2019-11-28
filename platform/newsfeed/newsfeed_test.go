package newsfeed

import (
	"testing"
)

func TestAdd(t *testing.T) {
	feed := New()
	feed.Add(Item{Title: "An Item!", Post: "With Body"})

	if len(feed.Items) == 0 {
		t.Errorf("Item was not added")
	}
}

func TestGetAll(t *testing.T) {
	feed := New()
	feed.Add(Item{Title: "An Item!", Post: "With Body"})

	result := feed.GetAll()

	if len(result) != 1 {
		t.Errorf("Item was not added")
	}

}

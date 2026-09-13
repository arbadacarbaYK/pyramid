package grasp

import (
	"testing"

	"fiatjaf.com/nostr"
)

func TestStatusRootETagPrefersRootMarker(t *testing.T) {
	event := nostr.Event{
		Tags: nostr.Tags{
			{"p", "owner"},
			{"e", "reply-id", "", "reply"},
			{"e", "root-id", "", "root"},
		},
	}
	got := statusRootETag(event)
	if got == nil || got[1] != "root-id" {
		t.Fatalf("expected root-id, got %v", got)
	}
}

func TestStatusRootETagFallsBackToFirstE(t *testing.T) {
	event := nostr.Event{
		Tags: nostr.Tags{
			{"e", "only-id"},
			{"p", "owner"},
		},
	}
	got := statusRootETag(event)
	if got == nil || got[1] != "only-id" {
		t.Fatalf("expected only-id, got %v", got)
	}
}

func TestStatusRootETagNilWhenMissing(t *testing.T) {
	if got := statusRootETag(nostr.Event{Tags: nostr.Tags{{"p", "owner"}}}); got != nil {
		t.Fatalf("expected nil, got %v", got)
	}
}

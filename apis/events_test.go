package apis

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/ubuntu/tutorial-deployment/paths"

	"os"
)

func TestNewEvents(t *testing.T) {
	testCases := []struct {
		eventsDir string

		wantEvents Events
		wantErr    bool
	}{
		{"testdata/events/valid",
			Events{"event-1": event{Name: "Event 1", Logo: "img/event1.jpg", Description: "This workshop is taking place at Event 1."},
				"event-2": event{Name: "Event 2", Logo: "event2.jpg", Description: "This workshop is taking place at Event 2."},
			},
			false},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("create event for: %+v", tc.eventsDir), func(t *testing.T) {
			// Setup/Teardown
			p := paths.New()
			p.MetaData = tc.eventsDir

			// Test
			e, err := NewEvents()
			if (err != nil) != tc.wantErr {
				t.Errorf("NewEvents() error = %v, wantErr %v", err, tc.wantErr)
			}
			if err != nil {
				return
			}

			if !reflect.DeepEqual(*e, tc.wantEvents) {
				t.Errorf("Generated events: got %+v; want %+v", *e, tc.wantEvents)
			}
		})
	}
}

func tempDir(t *testing.T) (string, func()) {
	path, err := ioutil.TempDir("", "tutorial-test")
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	return path, func() {
		if err := os.RemoveAll(path); err != nil {
			t.Fatalf("err: %s", err)
		}
	}
}

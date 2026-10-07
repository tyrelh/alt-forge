package main

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGenerator(t *testing.T) {
	for _, ancestry := range ancestries {
		names, ok := generate(ancestry, 6)
		if !ok || len(names) != 6 {
			t.Fatalf("invalid batch for %s", ancestry)
		}
		seen := map[Name]bool{}
		for _, n := range names {
			if n.First == "" || n.Last == "" || seen[n] {
				t.Fatalf("empty or duplicate name: %+v", n)
			}
			seen[n] = true
		}
	}
	if _, ok := generate("Murloc", 6); ok {
		t.Fatal("accepted unknown ancestry")
	}
}

func TestUndeadNamePool(t *testing.T) {
	pool, ok := pools["Undead"]
	if !ok {
		t.Fatal("missing Undead name pool")
	}
	for _, part := range []struct {
		name   string
		values []string
		min    int
	}{
		{"first names", pool.first, 16},
		{"surname prefixes", pool.prefix, 10},
		{"surname suffixes", pool.suffix, 10},
	} {
		if len(part.values) < part.min {
			t.Errorf("Undead %s: got %d entries, want at least %d", part.name, len(part.values), part.min)
		}
		seen := make(map[string]bool)
		for _, value := range part.values {
			if value == "" || seen[value] {
				t.Errorf("Undead %s: empty or duplicate entry %q", part.name, value)
			}
			seen[value] = true
		}
	}
}

func TestPages(t *testing.T) {
	for _, tc := range []struct {
		path     string
		partial  bool
		status   int
		selected string
	}{
		{"/", false, 200, ""}, {"/?ancestry=Orc", true, 200, ""}, {"/?ancestry=Night+elf", false, 200, ""},
		{"/?ancestry=Undead", false, 200, "Undead"}, {"/?ancestry=Undead", true, 200, "Undead"},
		{"/?ancestry=Goblin", false, 200, "Goblin"}, {"/?ancestry=Goblin", true, 200, "Goblin"},
		{"/?ancestry=unknown", true, 400, ""}, {"/missing", false, 404, ""}, {"/static/htmx.min.js", false, 200, ""},
	} {
		r := httptest.NewRequest("GET", tc.path, nil)
		if tc.partial {
			r.Header.Set("HX-Request", "true")
		}
		w := httptest.NewRecorder()
		handler().ServeHTTP(w, r)
		if w.Code != tc.status {
			t.Fatalf("%s: got %d", tc.path, w.Code)
		}
		if tc.status == 200 && !strings.HasPrefix(tc.path, "/static/") {
			body := w.Body.String()
			if strings.Contains(body, "<!doctype html>") == tc.partial {
				t.Fatal("wrong page/fragment response")
			}
			if strings.Count(body, `class="name-card"`) != 6 {
				t.Fatal("expected six cards")
			}
			if tc.selected != "" {
				if !strings.Contains(body, `<span>`+tc.selected+` · 6 names</span>`) {
					t.Fatalf("%s: wrong selected ancestry in results", tc.path)
				}
				if !tc.partial && !strings.Contains(body, `<option selected>`+tc.selected+`</option>`) {
					t.Fatalf("%s: selected ancestry missing from form", tc.path)
				}
			}
		}
	}
}

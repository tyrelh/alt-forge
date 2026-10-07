package main

import "testing"

func TestGoblinNamePool(t *testing.T) {
	found := false
	for _, a := range ancestries {
		if a == "Goblin" {
			found = true
		}
	}
	if !found {
		t.Fatal("Goblin missing from ancestries")
	}
	pool, ok := pools["Goblin"]
	if !ok {
		t.Fatal("missing Goblin name pool")
	}
	for name, values := range map[string][]string{"first": pool.first, "prefix": pool.prefix, "suffix": pool.suffix} {
		if len(values) == 0 {
			t.Errorf("Goblin %s: empty", name)
		}
		seen := map[string]bool{}
		for _, v := range values {
			if v == "" || seen[v] {
				t.Errorf("Goblin %s: empty or duplicate entry %q", name, v)
			}
			seen[v] = true
		}
	}
}

func TestGoblinGenerate(t *testing.T) {
	names, ok := generate("Goblin", 6)
	if !ok || len(names) != 6 {
		t.Fatal("invalid Goblin batch")
	}
	seen := map[Name]bool{}
	for _, n := range names {
		if n.First == "" || n.Last == "" || seen[n] {
			t.Fatalf("empty or duplicate name: %+v", n)
		}
		seen[n] = true
	}
}

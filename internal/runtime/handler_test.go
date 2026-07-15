package runtime

import "testing"

func TestAccept(t *testing.T) { if !Accept("a:1") { t.Fatal("expected accepted event") } }
func TestRejectEmpty(t *testing.T) { if Accept("  ") { t.Fatal("expected empty event rejection") } }
func TestAcceptTrimmed(t *testing.T) { if !Accept(" a:1 ") { t.Fatal("expected trimmed event") } }
func TestAcceptBatch(t *testing.T) { if !AcceptBatch([]string{"a:1","b:2"}) { t.Fatal("expected valid batch") }; if AcceptBatch([]string{"a:1",""}) { t.Fatal("expected invalid batch") } }

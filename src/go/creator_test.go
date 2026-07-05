package common

import "testing"

func TestHTAddNewClassToIdx(t *testing.T) {
	t.Run("appends value to last content", func(t *testing.T) {
		idx := &ClassIdx{
			Content:    []ClassContent{{Value: []ClassContentValue{{ID: "existing"}}}},
			LastUpdate: []string{""},
		}

		HTAddNewClassToIdx(idx, "newFile")

		if len(idx.Content) != 1 {
			t.Fatalf("expected 1 content, got %d", len(idx.Content))
		}
		vals := idx.Content[0].Value
		if len(vals) != 2 {
			t.Fatalf("expected 2 values, got %d", len(vals))
		}
		if vals[1].ID != "newFile" {
			t.Errorf("expected ID newFile, got %s", vals[1].ID)
		}
	})

	t.Run("no-op when content is empty", func(t *testing.T) {
		idx := &ClassIdx{}
		HTAddNewClassToIdx(idx, "newFile")
	})
}

func TestHTAddNewFamilyToIdx(t *testing.T) {
	t.Run("appends value to last content", func(t *testing.T) {
		idx := &IdxFamily{
			Contents:   []IdxFamilyContent{{Value: []IdxFamilyValue{{ID: "existing"}}}},
			LastUpdate: []string{""},
		}

		HTAddNewFamilyToIdx(idx, "newFile", "en")

		if len(idx.Contents) != 1 {
			t.Fatalf("expected 1 content, got %d", len(idx.Contents))
		}
		vals := idx.Contents[0].Value
		if len(vals) != 2 {
			t.Fatalf("expected 2 values, got %d", len(vals))
		}
		if vals[1].ID != "newFile" {
			t.Errorf("expected ID newFile, got %s", vals[1].ID)
		}
		if vals[1].GEDCOM != "gedcom/newFile_en.ged" {
			t.Errorf("expected GEDCOM gedcom/newFile_en.ged, got %s", vals[1].GEDCOM)
		}
		if vals[1].CSV != "csv/newFile_en.csv" {
			t.Errorf("expected CSV csv/newFile_en.csv, got %s", vals[1].CSV)
		}
	})

	t.Run("no-op when content is empty", func(t *testing.T) {
		idx := &IdxFamily{}
		HTAddNewFamilyToIdx(idx, "newFile", "en")
	})
}

func TestHTSetDefaultClassTemplateValues(t *testing.T) {
	fp := &ClassTemplateFile{
		Sources:    []string{""},
		Scripts:    []string{""},
		Index:      []string{""},
		LastUpdate: []string{""},
		Authors:    []string{""},
		Reviewers:  []string{""},
	}

	HTSetDefaultClassTemplateValues(fp, "myfile", "history")

	if fp.Title != "" {
		t.Errorf("expected empty title, got %q", fp.Title)
	}
	if fp.Header != "" {
		t.Errorf("expected empty header, got %q", fp.Header)
	}
	if fp.Sources[0] != "myfile" {
		t.Errorf("expected source myfile, got %q", fp.Sources[0])
	}
	if fp.Scripts[0] != "myfile" {
		t.Errorf("expected script myfile, got %q", fp.Scripts[0])
	}
	if fp.Index[0] != "history" {
		t.Errorf("expected index history, got %q", fp.Index[0])
	}
	if fp.Type != "class" {
		t.Errorf("expected type class, got %q", fp.Type)
	}
	if fp.Version != 2 {
		t.Errorf("expected version 2, got %d", fp.Version)
	}
}

func TestHTNewFamilySetDefaultValues(t *testing.T) {
	family := &Family{
		Sources:    []string{""},
		Scripts:    []string{""},
		LastUpdate: []string{""},
	}

	HTNewFamilySetDefaultValues(family, "en", "myfamily")

	if family.Title != "" {
		t.Errorf("expected empty title, got %q", family.Title)
	}
	if family.Header != "" {
		t.Errorf("expected empty header, got %q", family.Header)
	}
	if family.Sources[0] != "myfamily" {
		t.Errorf("expected source myfamily, got %q", family.Sources[0])
	}
	if family.Scripts[0] != "myfamily" {
		t.Errorf("expected script myfamily, got %q", family.Scripts[0])
	}
	if family.GEDCOM != "gedcom/myfamily_en.ged" {
		t.Errorf("expected GEDCOM gedcom/myfamily_en.ged, got %q", family.GEDCOM)
	}
	if family.CSV != "csv/myfamily_en.csv" {
		t.Errorf("expected CSV csv/myfamily_en.csv, got %q", family.CSV)
	}
}

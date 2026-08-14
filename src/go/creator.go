// SPDX-License-Identifier: GPL-3.0-or-later

package common

// HTAddNewClassToIdx appends a new class content value to the last content
// entry of the class index and updates the timestamp.
func HTAddNewClassToIdx(index *ClassIdx, newFile string) {
	HTAddNewClassToIdxAfter(index, newFile, "")
}

// HTAddNewClassToIdxAfter inserts a new class content value into the value
// vector right after the entry with the given ID. If afterID is empty or is
// not found in any content entry, the value is appended to the last content
// entry. The timestamp is updated.
func HTAddNewClassToIdxAfter(index *ClassIdx, newFile string, afterID string) {
	newValue := ClassContentValue{ID: newFile, Name: "", Desc: ""}

	if afterID != "" {
		for i := range index.Content {
			content := &index.Content[i]
			for j := range content.Value {
				if content.Value[j].ID == afterID {
					content.Value = append(content.Value, ClassContentValue{})
					copy(content.Value[j+2:], content.Value[j+1:])
					content.Value[j+1] = newValue
					index.LastUpdate[0] = HTUpdateTimestamp()
					return
				}
			}
		}
	}

	lastContent := len(index.Content) - 1
	if lastContent < 0 {
		return
	}

	content := &index.Content[lastContent]
	content.Value = append(content.Value, newValue)

	index.LastUpdate[0] = HTUpdateTimestamp()
}

// HTSetDefaultClassTemplateValues sets the default values on a class template
// file struct. The className parameter is the index name (e.g. "history",
// "physics") that the new class belongs to.
func HTSetDefaultClassTemplateValues(fp *ClassTemplateFile, newFile string, className string) {
	fp.Title = ""
	fp.Header = ""
	fp.Sources[0] = newFile
	fp.Scripts[0] = newFile
	fp.Index[0] = className
	fp.LastUpdate[0] = HTUpdateTimestamp()
	fp.Authors[0] = ""
	fp.Reviewers[0] = ""
	fp.Type = "class"
	fp.Version = 2
}

// HTNewFamilySetDefaultValues sets the default values on a family struct.
func HTNewFamilySetDefaultValues(family *Family, lang string, fileName string) {
	family.Title = ""
	family.Header = ""
	family.Sources[0] = fileName
	family.Scripts[0] = fileName
	family.LastUpdate[0] = HTUpdateTimestamp()
	family.GEDCOM = "gedcom/" + fileName + "_" + lang + ".ged"
	family.CSV = "csv/" + fileName + "_" + lang + ".csv"
}

// HTAddNewFamilyToIdx appends a new family content value to the last content
// entry of the family index and updates the timestamp.
func HTAddNewFamilyToIdx(index *IdxFamily, newFile string, lang string) {
	HTAddNewFamilyToIdxAfter(index, newFile, lang, "")
}

// HTAddNewFamilyToIdxAfter inserts a new family content value into the value
// vector right after the entry with the given ID. If afterID is empty or is
// not found in any content entry, the value is appended to the last content
// entry. The timestamp is updated.
func HTAddNewFamilyToIdxAfter(index *IdxFamily, newFile string, lang string, afterID string) {
	newValue := IdxFamilyValue{ID: newFile, GEDCOM: "gedcom/" + newFile + "_" + lang + ".ged", CSV: "csv/" + newFile + "_" + lang + ".csv"}

	if afterID != "" {
		for i := range index.Contents {
			content := &index.Contents[i]
			for j := range content.Value {
				if content.Value[j].ID == afterID {
					content.Value = append(content.Value, IdxFamilyValue{})
					copy(content.Value[j+2:], content.Value[j+1:])
					content.Value[j+1] = newValue
					index.LastUpdate[0] = HTUpdateTimestamp()
					return
				}
			}
		}
	}

	lastContent := len(index.Contents) - 1
	if lastContent < 0 {
		return
	}

	content := &index.Contents[lastContent]
	content.Value = append(content.Value, newValue)

	index.LastUpdate[0] = HTUpdateTimestamp()
}

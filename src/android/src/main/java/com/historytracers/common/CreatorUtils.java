package com.historytracers.common;

public final class CreatorUtils {
    private CreatorUtils() {}

    public static void addNewClassToIdx(ClassIdx index, String newFile) {
        int lastContent = index.content.size() - 1;
        if (lastContent < 0) return;

        ClassContent content = index.content.get(lastContent);
        ClassContentValue newValue = new ClassContentValue();
        newValue.id = newFile;
        content.value.add(newValue);

        index.lastUpdate.set(0, TimestampUtils.updateTimestamp());
    }

    public static void setDefaultClassTemplateValues(ClassTemplateFile fp, String newFile, String className) {
        fp.sources.set(0, newFile);
        fp.scripts.set(0, newFile);
        fp.index.set(0, className);
        fp.lastUpdate.set(0, TimestampUtils.updateTimestamp());
        fp.type = "class";
        fp.version = 2;
    }

    public static void newFamilySetDefaultValues(Family family, String lang, String fileName) {
        family.sources.set(0, fileName);
        family.scripts.set(0, fileName);
        family.lastUpdate.set(0, TimestampUtils.updateTimestamp());
        family.gedcom = "gedcom/" + fileName + "_" + lang + ".ged";
        family.csv = "csv/" + fileName + "_" + lang + ".csv";
    }

    public static void addNewFamilyToIdx(IdxFamily index, String newFile, String lang) {
        int lastContent = index.contents.size() - 1;
        if (lastContent < 0) return;

        IdxFamilyContent content = index.contents.get(lastContent);
        IdxFamilyValue newValue = new IdxFamilyValue();
        newValue.id = newFile;
        newValue.gedcom = "gedcom/" + newFile + "_" + lang + ".ged";
        newValue.csv = "csv/" + newFile + "_" + lang + ".csv";
        content.value.add(newValue);

        index.lastUpdate.set(0, TimestampUtils.updateTimestamp());
    }
}

package com.historytracers.common;

import com.google.gson.annotations.SerializedName;

import java.util.List;

public class FamilyPersonMarriage {
    @SerializedName("type")
    public String marriageType;

    public String id;

    @SerializedName("gedcom_id")
    public String gedcomId;

    public boolean official;

    @SerializedName("family_id")
    public String familyId;

    @SerializedName("external_family_file")
    public boolean externalFile;

    public String name;
    public List<HTText> history;

    @SerializedName("date_time")
    public FamilyPersonEvent dateTime;
}

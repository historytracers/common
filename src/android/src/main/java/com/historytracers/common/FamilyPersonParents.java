package com.historytracers.common;

import com.google.gson.annotations.SerializedName;

public class FamilyPersonParents {
    @SerializedName("type")
    public String parentType;

    @SerializedName("father_external_family_file")
    public boolean fatherExternalFile;

    @SerializedName("father_family")
    public String fatherFamily;

    @SerializedName("father_id")
    public String fatherID;

    @SerializedName("father_name")
    public String fatherName;

    @SerializedName("mother_external_family_file")
    public boolean motherExternalFile;

    @SerializedName("mother_family")
    public String motherFamily;

    @SerializedName("mother_id")
    public String motherID;

    @SerializedName("mother_name")
    public String motherName;
}

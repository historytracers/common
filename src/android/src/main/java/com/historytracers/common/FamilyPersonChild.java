package com.historytracers.common;

import com.google.gson.annotations.SerializedName;

import java.util.List;

public class FamilyPersonChild {
    @SerializedName("type")
    public String childType;

    public String id;

    @SerializedName("marriage_id")
    public String marriageId;

    public String name;

    @SerializedName("family_id")
    public String familyId;

    @SerializedName("external_family_file")
    public boolean externalFile;

    @SerializedName("add_link")
    public boolean addLink;

    public List<HTText> history;

    @SerializedName("adopted_child")
    public boolean adoptedChild;
}

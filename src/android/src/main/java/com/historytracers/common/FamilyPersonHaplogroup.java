package com.historytracers.common;

import com.google.gson.annotations.SerializedName;

import java.util.List;

public class FamilyPersonHaplogroup {
    @SerializedName("type")
    public String haplogroupType;

    public String haplogroup;
    public List<HTSource> sources;
}

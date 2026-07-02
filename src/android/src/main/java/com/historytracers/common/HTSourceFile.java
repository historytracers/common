package com.historytracers.common;

import com.google.gson.annotations.SerializedName;

import java.util.List;

public class HTSourceFile {
    public List<String> license;

    @SerializedName("last_update")
    public List<String> lastUpdate;

    public int version;

    @SerializedName("type")
    public String fileType;

    public List<HTSourceElement> primary_sources;
    public List<HTSourceElement> reference_sources;
    public List<HTSourceElement> religious_sources;
    public List<HTSourceElement> social_media_sources;
}

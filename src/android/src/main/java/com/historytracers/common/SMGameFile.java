package com.historytracers.common;

import com.google.gson.annotations.SerializedName;

import java.util.List;

public class SMGameFile {
    public List<String> sources;
    public List<String> license;

    @SerializedName("last_update")
    public List<String> lastUpdate;

    public String authors;
    public String reviewers;
    public int version;

    @SerializedName("type")
    public String type;

    public List<SMGameContent> content;

    @SerializedName("levels")
    public List<SMGameLevel> levels;

    @SerializedName("date_time")
    public List<HTDate> dateTime;
}

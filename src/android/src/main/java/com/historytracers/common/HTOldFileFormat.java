package com.historytracers.common;

import com.google.gson.annotations.SerializedName;

import java.util.List;

public class HTOldFileFormat {
    public String title;
    public String header;
    public List<String> license;
    public List<String> sources;

    @SerializedName("last_update")
    public List<String> lastUpdate;

    public List<HTAudio> audio;
    public List<HTCommonContent> contents;

    @SerializedName("date_time")
    public List<HTDate> dateTime;
}

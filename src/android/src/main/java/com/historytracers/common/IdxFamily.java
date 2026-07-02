package com.historytracers.common;

import com.google.gson.annotations.SerializedName;

import java.util.List;

public class IdxFamily {
    public String title;
    public String header;
    public List<String> license;
    public List<String> sources;

    @SerializedName("last_update")
    public List<String> lastUpdate;

    public List<HTAudio> audio;
    public String gedcom;
    public String csv;

    public List<IdxFamilyContent> contents;

    @SerializedName("date_time")
    public List<HTDate> dateTime;
}

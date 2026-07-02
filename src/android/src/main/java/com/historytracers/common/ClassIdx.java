package com.historytracers.common;

import com.google.gson.annotations.SerializedName;

import java.util.List;

public class ClassIdx {
    public String title;
    public String header;
    public List<HTAudio> audio;

    @SerializedName("last_update")
    public List<String> lastUpdate;

    public List<String> sources;
    public List<String> scripts;
    public List<String> license;
    public int version;

    @SerializedName("type")
    public String type;

    public List<ClassContent> content;

    @SerializedName("date_time")
    public List<HTDate> dateTime;
}

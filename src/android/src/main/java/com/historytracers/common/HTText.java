package com.historytracers.common;

import com.google.gson.annotations.SerializedName;

import java.util.List;

public class HTText {
    public String text;
    public List<HTSource> source;

    @SerializedName("date_time")
    public List<HTDate> fillDates;

    public boolean isTable;
    public String imgdesc;
    public String format;

    @SerializedName("PostMention")
    public String postMention;
}

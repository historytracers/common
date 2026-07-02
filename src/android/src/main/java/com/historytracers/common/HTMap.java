package com.historytracers.common;

import com.google.gson.annotations.SerializedName;

import java.util.List;

public class HTMap {
    public String text;
    public String img;
    public int order;

    @SerializedName("date_time")
    public List<HTDate> dateTime;
}

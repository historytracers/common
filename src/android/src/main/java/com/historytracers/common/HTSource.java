package com.historytracers.common;

import com.google.gson.annotations.SerializedName;

public class HTSource {
    @SerializedName("type")
    public int sourceType;

    public String uuid;
    public String text;
    public String page;

    @SerializedName("date_time")
    public HTDate dateTime;
}

package com.historytracers.common;

import com.google.gson.annotations.SerializedName;

import java.util.List;

public class HTGameDesc {
    public String imagePath;
    public String imageDesc;

    @SerializedName("date_time")
    public List<HTDate> dateTime;
}

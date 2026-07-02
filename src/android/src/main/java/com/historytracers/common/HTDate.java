package com.historytracers.common;

import com.google.gson.annotations.SerializedName;

public class HTDate {
    @SerializedName("type")
    public String dateType;

    public String year;
    public String month;
    public String day;
}

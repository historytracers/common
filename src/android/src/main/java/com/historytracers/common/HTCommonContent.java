package com.historytracers.common;

import com.google.gson.annotations.SerializedName;

import java.util.List;

public class HTCommonContent {
    public String id;
    public String desc;
    public String target;
    public String page;

    @SerializedName("value_type")
    public String valueType;

    @SerializedName("html_value")
    public String htmlValue;

    public List<IdxFamilyValue> value;

    @SerializedName("date_time")
    public List<HTDate> fillDates;
}

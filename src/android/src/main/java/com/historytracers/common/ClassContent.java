package com.historytracers.common;

import com.google.gson.annotations.SerializedName;

import java.util.List;

public class ClassContent {
    public String id;
    public String target;
    public String page;

    @SerializedName("value_type")
    public String valueType;

    @SerializedName("html_value")
    public String htmlValue;

    public List<ClassContentValue> value;

    @SerializedName("date_time")
    public List<HTDate> dateTime;
}

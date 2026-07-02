package com.historytracers.common;

import com.google.gson.annotations.SerializedName;

public class HTSourceElement {
    public String id;
    public String citation;

    @SerializedName("date_time")
    public String date;

    @SerializedName("published")
    public String publishDate;

    public String url;
}

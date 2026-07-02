package com.historytracers.common;

import com.google.gson.annotations.SerializedName;

import java.util.List;

public class ClassContentValue {
    @SerializedName("family_id")
    public String familyId;

    @SerializedName("person_id")
    public String personId;

    public String id;
    public String name;
    public String desc;

    @SerializedName("date_time")
    public List<HTDate> dateTime;
}

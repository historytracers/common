package com.historytracers.common;

import com.google.gson.annotations.SerializedName;

import java.util.List;

public class FamilyPersonEvent {
    public List<HTDate> date;
    public String address;

    @SerializedName("city_id")
    public String cityId;

    public String city;

    @SerializedName("state_id")
    public String stateId;

    public String state;
    public String pc;

    @SerializedName("country_id")
    public String countryId;

    public String country;
    public List<HTSource> sources;
}

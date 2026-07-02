package com.historytracers.common;

import com.google.gson.annotations.SerializedName;

import java.util.List;

public class Family {
    public String title;
    public String header;
    public List<String> sources;
    public List<String> scripts;
    public List<HTAudio> audio;
    public List<String> index;
    public List<String> license;

    @SerializedName("last_update")
    public List<String> lastUpdate;

    public String authors;
    public String reviewers;

    @SerializedName("documentsInfo")
    public List<String> documentsInfo;

    @SerializedName("periodOfTime")
    public List<String> periodOfTime;

    public List<HTMap> maps;

    @SerializedName("common")
    public List<HTText> commonText;

    public List<String> prerequisites;
    public String gedcom;
    public String csv;
    public int version;
    public boolean editing;

    @SerializedName("type")
    public String type;

    public List<FamilyBody> families;

    @SerializedName("exercise_v2")
    public List<HTExercise> exercises;

    @SerializedName("date_time")
    public List<HTDate> dateTime;
}

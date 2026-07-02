package com.historytracers.common;

import com.google.gson.annotations.SerializedName;

import java.util.List;

public class ClassTemplateFile {
    public String title;
    public String header;
    public List<String> sources;
    public List<String> scripts;
    public List<HTAudio> audio;
    public List<String> index;
    public List<String> license;

    @SerializedName("last_update")
    public List<String> lastUpdate;

    public List<String> authors;
    public List<String> reviewers;

    @SerializedName("type")
    public String type;

    public int version;
    public boolean editing;

    public List<ClassTemplateContent> content;

    @SerializedName("exercise_v2")
    public List<HTExercise> exercises;

    @SerializedName("game_v2")
    public List<HTGameDesc> gameV2;

    @SerializedName("date_time")
    public List<HTDate> dateTime;
}

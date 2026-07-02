package com.historytracers.common;

import com.google.gson.annotations.SerializedName;

import java.util.List;

public class SMGameContent {
    public String id;
    public String audio;
    public List<HTText> text;

    public Object desc;
    public String next;
    public String prev;

    @SerializedName("jumpTo")
    public String jumpTo;

    public Object answer;
    public int score;

    @SerializedName("position")
    public String position;
}

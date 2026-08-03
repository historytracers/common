// SPDX-License-Identifier: GPL-3.0-or-later

package com.historytracers.common;

import com.google.gson.annotations.SerializedName;

import java.util.List;

public class SMGameContent {
    public String id;
    public String audio;
    public String smile;

    @SerializedName("source_menu")
    public List<HTSource> sourceMenu;

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

// SPDX-License-Identifier: GPL-3.0-or-later

package com.historytracers.common;

import com.google.gson.annotations.SerializedName;

import java.util.List;

public class HTExercise {
    public String question;

    @SerializedName("yesNoAnswer")
    public String yesNoAnswer;

    public String additionalInfo;

    public List<HTSource> source;
}

package com.historytracers.common;

import com.google.gson.annotations.SerializedName;

public class HTExercise {
    public String question;

    @SerializedName("yesNoAnswer")
    public String yesNoAnswer;

    public String additionalInfo;
}

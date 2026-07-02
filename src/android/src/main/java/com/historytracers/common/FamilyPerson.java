package com.historytracers.common;

import com.google.gson.annotations.SerializedName;

import java.util.List;

public class FamilyPerson {
    public String id;
    public String name;
    public String surname;
    public String patronymic;

    @SerializedName("fullname")
    public String fullName;

    public String sex;
    public String gender;

    @SerializedName("is_real")
    public boolean isReal;

    public List<FamilyPersonHaplogroup> haplogroup;
    public List<HTText> history;
    public List<FamilyPersonParents> parents;
    public List<FamilyPersonEvent> birth;
    public List<FamilyPersonEvent> baptism;
    public List<FamilyPersonMarriage> marriages;
    public List<FamilyPersonMarriage> divorced;
    public List<FamilyPersonChild> children;
    public List<FamilyPersonEvent> death;
}

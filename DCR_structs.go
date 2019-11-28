package main

type DCR struct {
    Events Event `xml:"events"`
    Constraints Constraint `xml:"constraints"`
    Markings Marking `xml:"markings"`
}

type Event struct {
    Names []string `xml:"event"`
}

type Constraint struct {
    Conditions []string `xml:"conditions>condition"`
    Responses []string `xml:"responses>response"`
    Corresponses []string `xml:"corresponses>corresponse"`
    Excludes []string `xml:"excludes>exclude"`
    Includes []string `xml:"includes>include"`
    Milestones []string `xml:"milestones>milestone"`
    Spawns []string `xml:"spawns>spawn"`
}

type Marking struct {
    Executed []string `xml:"executed>event"`
    Included []string `xml:"included>event"`
    Pending []string `xml:"pending>event"`
}

type Constraint_map struct {
     constraint_map map[string][]string
}

type DCR_graph struct{
    nodes []string

    conditions_for Constraint_map
    milestones_for Constraint_map
    responses_to   Constraint_map
    excludes_to    Constraint_map
    includes_to    Constraint_map

    marking Marking
}

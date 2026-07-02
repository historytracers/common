package com.historytracers.common;

public class HTConfigBase {
    public int port;
    public String srcPath;
    public String contentPath;
    public String logPath;

    public HTConfigBase(int port, String srcPath, String contentPath, String logPath) {
        this.port = port;
        this.srcPath = srcPath;
        this.contentPath = contentPath;
        this.logPath = logPath;
    }
}

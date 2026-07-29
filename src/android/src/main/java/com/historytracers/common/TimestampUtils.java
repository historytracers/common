// SPDX-License-Identifier: GPL-3.0-or-later

package com.historytracers.common;

public final class TimestampUtils {
    private TimestampUtils() {}

    public static String updateTimestamp() {
        return String.valueOf(System.currentTimeMillis() / 1000L);
    }
}

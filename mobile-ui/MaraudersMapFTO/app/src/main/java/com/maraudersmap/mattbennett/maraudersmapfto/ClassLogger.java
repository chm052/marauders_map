package com.maraudersmap.mattbennett.maraudersmapfto;

import android.util.Log;

/**
 * Created by matt.bennett on 12/05/2016.
 */
public class ClassLogger {
    public static void debug(Object tag, String message) {
        Log.d(tag.getClass().getCanonicalName(), message);
    }
}

package com.maraudersmap.mattbennett.maraudersmapfto;

import android.location.Location;
import android.support.design.widget.Snackbar;
import android.view.View;

/**
 * Created by matt.bennett on 12/05/2016.
 */
public class PostLocationListener implements View.OnClickListener
{
    private LocationFinder locationFinder;

    public PostLocationListener(LocationFinder locationFinder) {
        this.locationFinder = locationFinder;
    }

    @Override
    public void onClick(View view) {
        try {
            Location loc = locationFinder.myLocation();
            snackbar(view, String.format("Found your location: <%s>", locationText(loc)));
        } catch (Exception e) {
            snackbar(view, e.getMessage());
        }
    }

    private String locationText(Location loc) {
        return String.format("lat: <%f> long: <%f>", loc.getLatitude(), loc.getLongitude());
    }

    private void snackbar(View view, String text) {
        Snackbar.make(view, text, Snackbar.LENGTH_LONG)
                .setAction("Action", null)
                .show();
    }
}

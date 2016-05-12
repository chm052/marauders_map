package com.maraudersmap.mattbennett.maraudersmapfto;

import android.location.Location;
import android.support.design.widget.Snackbar;
import android.view.View;

import com.android.volley.Response;

/**
 * Created by matt.bennett on 12/05/2016.
 */
public class PostLocationListener implements View.OnClickListener
{
    private LocationFinder locationFinder;
    private MaraudersApiClient maraudersApiClient;

    public PostLocationListener(LocationFinder locationFinder, MaraudersApiClient maraudersApiClient) {
        this.locationFinder = locationFinder;
        this.maraudersApiClient = maraudersApiClient;
    }

    @Override
    public void onClick(final View view) {
        try {
            Location loc = locationFinder.myLocation();
            ClassLogger.debug(this, "Sending location to MM");
            //snackbar(view, String.format("Found your location: <%s>", locationText(loc)));
            maraudersApiClient.postLocationToMarauders(loc, new Response.Listener<String>() {
                @Override
                public void onResponse(String response) {
                    ClassLogger.debug(this, "It worked! Gonna pop up a snacky bar");
                    snackbar(view, response);
                }
            });
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

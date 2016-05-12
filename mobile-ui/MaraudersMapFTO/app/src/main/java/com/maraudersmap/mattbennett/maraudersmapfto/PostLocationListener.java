package com.maraudersmap.mattbennett.maraudersmapfto;

import android.graphics.Color;
import android.location.Location;
import android.support.annotation.ColorInt;
import android.support.design.widget.Snackbar;
import android.view.View;
import android.widget.Button;


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
            String result = maraudersApiClient.postLocationToMarauders(loc);
            if(result.contains("Opening food truck")) {
                snackbar(view, "Opened your truck!");
                view.setBackgroundColor(Color.RED);
                ((Button) view).setText("I'm Closed!");
            } else {
                snackbar(view, "There seems to have been an error!");
            }
        } catch (Exception e) {
            //ClassLogger.debug(this, e.getMessage());
            e.printStackTrace();
            snackbar(view, e.getMessage());
        }
    }

    private void snackbar(View view, String text) {
        Snackbar.make(view, text, Snackbar.LENGTH_LONG)
                .setAction("Action", null)
                .show();
    }
}

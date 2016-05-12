package com.maraudersmap.mattbennett.maraudersmapfto;

import android.location.Location;

/**
 * Created by matt.bennett on 12/05/2016.
 */
public class MaraudersApiClient {
    private String MaraudersBaseUrl = "http://localhost:9001/";
    private String OpenFormat = "api/trucks/open/%d?lat=%f&long=%f";

    public MaraudersApiClient(String url) {
        MaraudersBaseUrl = url;
    }

    public MaraudersApiClient() {
    }

    public String postLocationToMaruaders(Location location) {


        return "Saved succesfully!";
    }

    private String openUrl(Location location, int id) {
        return String.format(OpenFormat, id, location.getLatitude(), location.getLongitude());
    }
}

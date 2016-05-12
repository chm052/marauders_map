package com.maraudersmap.mattbennett.maraudersmapfto;

import android.Manifest;
import android.content.Context;
import android.content.pm.PackageManager;
import android.location.Location;
import android.location.LocationManager;
import android.support.v4.content.ContextCompat;
import android.util.Log;

import java.util.List;

/**
 * Created by matt.bennett on 12/05/2016.
 */
public class LocationFinder {
    private LocationManager locationManager;
    private Context context;

    public LocationFinder(Context context) {
        this.context = context;
        locationManager = (LocationManager) context.getSystemService(Context.LOCATION_SERVICE);
    }

    public Location myLocation() throws Exception {
        List<String> providers = locationManager.getProviders(false);

        if(providers.size() > 0) {
            Log.d("tag!", String.format("here's the locations %s", providers.toString()));

            if (ContextCompat.checkSelfPermission(context, Manifest.permission.ACCESS_FINE_LOCATION) == PackageManager.PERMISSION_GRANTED) {
                return locationManager.getLastKnownLocation(providers.get(0));
            }
            throw new Exception("I don't have the permissions!");
        }
        throw new Exception("could not find providers");
    }
}

package com.maraudersmap.mattbennett.maraudersmapfto;

import android.Manifest;
import android.app.Activity;
import android.content.pm.PackageManager;
import android.location.Location;
import android.os.Bundle;
import android.support.annotation.NonNull;
import android.support.annotation.Nullable;
import android.util.Log;

import com.google.android.gms.common.ConnectionResult;
import com.google.android.gms.common.api.GoogleApiClient;
import com.google.android.gms.location.LocationServices;

/**
 * Created by matt.bennett on 12/05/2016.
 */
public class LocationFinder implements GoogleApiClient.ConnectionCallbacks, GoogleApiClient.OnConnectionFailedListener {
    private GoogleApiClient apiClient;
    private Activity context;
    private boolean connected = false;

    final int REQUEST_LOCATION_CODE = 420;

    public LocationFinder(Activity context) {
        this.context = context;

        // Create an instance of GoogleAPIClient.
        Log.d("PostActivity", "Instatiating google api client");
        if (this.apiClient == null) {
            this.apiClient = new GoogleApiClient.Builder(context)
                    .addConnectionCallbacks(this)
                    .addOnConnectionFailedListener(this)
                    .addApi(LocationServices.API)
                    .build();
        }
    }

    public Location myLocation() throws Exception {
        if(connected) {
            Location currentLocation = null;

            if (context.checkSelfPermission(Manifest.permission.ACCESS_FINE_LOCATION) != PackageManager.PERMISSION_GRANTED) {
               requestPermission(Manifest.permission.ACCESS_FINE_LOCATION);
            }

            while (currentLocation == null) {
                currentLocation = LocationServices.FusedLocationApi.getLastLocation(this.apiClient);
            }
            return currentLocation;
        }

        //TODO: This should throw an error and you should wait on a connection
        return null;
    }

    public void connect() {
        this.apiClient.connect();
    }

    public void disconnect() {
        this.apiClient.disconnect();
    }

    private void requestPermission(String... permission) {
        context.requestPermissions(permission, RequestCode.LOCATION);
    }

    @Override
    public void onConnected(@Nullable Bundle bundle) {
        ClassLogger.debug(this, "Have connected");
        connected = true;
    }

    @Override
    public void onConnectionSuspended(int i) {
        ClassLogger.debug(this, "Have suspended connected");
        connected = false;
    }

    @Override
    public void onConnectionFailed(@NonNull ConnectionResult connectionResult) {
        ClassLogger.debug(this, "Everything has gone to shit :(");
        connected = false;
    }
}

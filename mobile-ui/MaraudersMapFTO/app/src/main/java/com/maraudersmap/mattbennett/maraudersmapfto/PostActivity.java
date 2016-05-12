package com.maraudersmap.mattbennett.maraudersmapfto;

import android.Manifest;
import android.content.pm.PackageManager;
import android.location.Location;
import android.os.Bundle;
import android.support.annotation.NonNull;
import android.support.annotation.Nullable;
import android.support.design.widget.Snackbar;
import android.support.v7.app.AppCompatActivity;
import android.support.v7.widget.Toolbar;
import android.util.Log;
import android.view.View;
import android.view.Menu;
import android.view.MenuItem;
import android.widget.Button;

import com.google.android.gms.common.ConnectionResult;
import com.google.android.gms.common.api.GoogleApiClient;
import com.google.android.gms.location.LocationServices;

public class PostActivity extends AppCompatActivity
    implements GoogleApiClient.ConnectionCallbacks, GoogleApiClient.OnConnectionFailedListener {
    private Location currentLocation;
    private GoogleApiClient googleApiClient;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);

        // Create an instance of GoogleAPIClient.
        Log.d("PostActivity", "Instatiating google api client");
        if (this.googleApiClient == null) {
            this.googleApiClient = new GoogleApiClient.Builder(this)
                    .addConnectionCallbacks(this)
                    .addOnConnectionFailedListener(this)
                    .addApi(LocationServices.API)
                    .build();
        }

        setContentView(R.layout.activity_post);
        Toolbar toolbar = (Toolbar) findViewById(R.id.toolbar);
        setSupportActionBar(toolbar);

//        FloatingActionButton fab = (FloatingActionButton) findViewById(R.id.fab);
//        fab.setOnClickListener(new View.OnClickListener() {
//            @Override
//            public void onClick(View view) {
//                Snackbar.make(view, "Replace with your own action", Snackbar.LENGTH_LONG)
//                        .setAction("Action", null).show();
//            }
//        });

        Button button = (Button) findViewById(R.id.open_button);
        button.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View view) {
                LocationFinder location = new LocationFinder(view.getContext());
                try {
                    Location loc = currentLocation;
                    System.out.println("Got location " + currentLocation);
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
        });
    }

    @Override
    public boolean onCreateOptionsMenu(Menu menu) {
        // Inflate the menu; this adds items to the action bar if it is present.
        getMenuInflater().inflate(R.menu.menu_post, menu);
        return true;
    }

    protected void onStart() {
        Log.d("PostActivity", "starting connection");
        this.googleApiClient.connect();
        super.onStart();
    }

    protected void onStop() {
        this.googleApiClient.disconnect();
        super.onStop();
    }

    @Override
    public boolean onOptionsItemSelected(MenuItem item) {
        // Handle action bar item clicks here. The action bar will
        // automatically handle clicks on the Home/Up button, so long
        // as you specify a parent activity in AndroidManifest.xml.
        int id = item.getItemId();

        //noinspection SimplifiableIfStatement
        if (id == R.id.action_settings) {
            return true;
        }

        return super.onOptionsItemSelected(item);
    }

    @Override
    public void onConnected(@Nullable Bundle bundle) {
        Log.d("PostActivity", "Do I have permisssions/S?S?S?S?/");
        Log.d("PostActivity", "check: " + checkSelfPermission(Manifest.permission.ACCESS_FINE_LOCATION) + " permgranted: " + PackageManager.PERMISSION_GRANTED);
        if(checkSelfPermission(Manifest.permission.ACCESS_FINE_LOCATION) == PackageManager.PERMISSION_GRANTED) {
            Log.d("PostActivity", "Finding current location");
            while (currentLocation == null) {
                currentLocation = LocationServices.FusedLocationApi.getLastLocation(this.googleApiClient);
                Log.d("PostActivity", "Found current location " + currentLocation);
            }
        } else {
            Log.d("PostActivity", "NOPE");
            requestPermissions(new String[] { Manifest.permission.ACCESS_FINE_LOCATION }, 0);
        }
    }

    @Override
    public void onRequestPermissionsResult(int requestCode,
                                           String permissions[], int[] grantResults) {
        Log.d("PostActivity", "doing the thiiiiiing: permission request came back");
        switch (requestCode) {
            case 0: {
                // If request is cancelled, the result arrays are empty.
                if (grantResults.length > 0 && grantResults[0] == PackageManager.PERMISSION_GRANTED) {
                    Log.d("PostActivity", "doing the thiiiiiing: permission request came back AND IT'S GOOD!");
                    // permission was granted, yay! Do the
                    // location-related task you need to do.
                    if(checkSelfPermission(Manifest.permission.ACCESS_FINE_LOCATION) == PackageManager.PERMISSION_GRANTED)
                        currentLocation = LocationServices.FusedLocationApi.getLastLocation(this.googleApiClient);

                } else {
                    Log.d("PostActivity", "FUuuuuuuuuuuuuuuuck you!");
                    // permission denied, boo! Disable the
                    // functionality that depends on this permission.
                }
                return;
            }
        }
    }

    @Override
    public void onConnectionSuspended(int i) {
        Log.d("PostActivity", "I was suspended!!" + i);
    }

    @Override
    public void onConnectionFailed(@NonNull ConnectionResult connectionResult) {
        Log.d("PostActivity", "something went wrong");
    }
}

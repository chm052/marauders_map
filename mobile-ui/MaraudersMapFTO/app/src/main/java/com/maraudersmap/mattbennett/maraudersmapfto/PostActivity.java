package com.maraudersmap.mattbennett.maraudersmapfto;

import android.content.pm.PackageManager;
import android.os.Bundle;
import android.support.v7.app.AppCompatActivity;
import android.support.v7.widget.Toolbar;
import android.util.Log;
import android.view.Menu;
import android.view.MenuItem;
import android.widget.Button;

public class PostActivity extends AppCompatActivity {
    private LocationFinder locationFinder;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        locationFinder = new LocationFinder(this);

        setContentView(R.layout.activity_post);
        Toolbar toolbar = (Toolbar) findViewById(R.id.toolbar);
        setSupportActionBar(toolbar);

        Button button = (Button) findViewById(R.id.open_button);
        button.setOnClickListener(new PostLocationListener(locationFinder));
    }

    @Override
    public boolean onCreateOptionsMenu(Menu menu) {
        // Inflate the menu; this adds items to the action bar if it is present.
        getMenuInflater().inflate(R.menu.menu_post, menu);
        return true;
    }

    protected void onStart() {
        Log.d("PostActivity", "starting connection");
        this.locationFinder.connect();
        super.onStart();
    }

    protected void onStop() {
        this.locationFinder.disconnect();
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
    public void onRequestPermissionsResult(int requestCode, String permissions[], int[] grantResults) {
        Log.d("PostActivity", "doing the thiiiiiing: permission request came back");
        switch (requestCode) {
            case RequestCode.LOCATION: {
                // If request is cancelled, the result arrays are empty.
                if (grantResults.length > 0 && grantResults[0] == PackageManager.PERMISSION_GRANTED) {
                    Log.d("PostActivity", "doing the thiiiiiing: permission request came back AND IT'S GOOD!");
                    // permission was granted, yay! Do the
                    // location-related task you need to do.
                } else {
                    Log.d("PostActivity", "FUuuuuuuuuuuuuuuuck you!");
                    // permission denied, boo! Disable the
                    // functionality that depends on this permission.
                }
                return;
            }
        }
    }
}

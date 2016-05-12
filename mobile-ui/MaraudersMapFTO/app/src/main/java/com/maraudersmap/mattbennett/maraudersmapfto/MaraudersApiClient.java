package com.maraudersmap.mattbennett.maraudersmapfto;

import android.content.Context;
import android.location.Location;

import com.android.volley.Request;
import com.android.volley.RequestQueue;
import com.android.volley.Response;
import com.android.volley.VolleyError;
import com.android.volley.toolbox.StringRequest;
import com.android.volley.toolbox.Volley;

/**
 * Created by matt.bennett on 12/05/2016.
 */
public class MaraudersApiClient {
    private String maraudersBaseUrl;
    private String OpenFormat = "trucks/open/%d?lat=%f&long=%f";
    private RequestQueue requestQueue;

   public MaraudersApiClient(String maraudersBaseUrl, Context context) {
       requestQueue = Volley.newRequestQueue(context);
       this.maraudersBaseUrl = maraudersBaseUrl;
   }

    public void postLocationToMarauders(Location location, Response.Listener<String> listener) {
        //call endpoint
        String url = maraudersBaseUrl + openUrl(location, 1);
        // Request a string response from the provided URL.
        StringRequest stringRequest = new StringRequest(Request.Method.POST, url, listener,
            new Response.ErrorListener() {
                @Override
                public void onErrorResponse(VolleyError error) {
                    ClassLogger.debug(this, "there was an error: " + error.getMessage());
                }
        });
        // Add the request to the RequestQueue.
        requestQueue.add(stringRequest);
        ClassLogger.debug(this, "request queued");
    }

    private String openUrl(Location location, int id) {
        return String.format(OpenFormat, id, location.getLatitude(), location.getLongitude());
    }
}

package com.maraudersmap.mattbennett.maraudersmapfto;

import android.location.Location;
import android.os.StrictMode;

import java.io.IOException;

import okhttp3.MediaType;
import okhttp3.OkHttpClient;
import okhttp3.Request;
import okhttp3.RequestBody;
import okhttp3.Response;
import okio.BufferedSink;

/**
 * Created by matt.bennett on 12/05/2016.
 */
public class MaraudersApiClient {
    private String maraudersBaseUrl;
    private String OpenFormat = "trucks/open/%d?lat=%f&lon=%f";
    private OkHttpClient okHttpClient = new OkHttpClient();

   public MaraudersApiClient(String maraudersBaseUrl) {
       this.maraudersBaseUrl = maraudersBaseUrl;
   }

    public String postLocationToMarauders(Location location) {
        //call endpoint
        String url = maraudersBaseUrl + openUrl(location, 1);
        okhttp3.Request request = new Request.Builder()
                .url(url)
                .method("POST", new RequestBody() {
                    @Override
                    public MediaType contentType() {
                        return MediaType.parse("application/json");
                    }

                    @Override
                    public void writeTo(BufferedSink sink) throws IOException {

                    }
                })
                .build();
        ClassLogger.debug(this, String.format("sending request to <%s>", url));

        try {
            enableStrictMode();
            Response res = okHttpClient.newCall(request).execute();
            ClassLogger.debug(this, String.format("response body: <%s>", res.body()));
            return res.body().string();
        } catch (IOException e) {
            ClassLogger.debug(this, "EVERYTHING HAS GONE TO SHIT");
        }
        return ":(";
    }

    private void enableStrictMode() {
        StrictMode.ThreadPolicy policy = new StrictMode.ThreadPolicy.Builder().permitAll().build();
        StrictMode.setThreadPolicy(policy);
    }

    private String openUrl(Location location, int id) {
        return String.format(OpenFormat, id, location.getLatitude(), location.getLongitude());
    }
}

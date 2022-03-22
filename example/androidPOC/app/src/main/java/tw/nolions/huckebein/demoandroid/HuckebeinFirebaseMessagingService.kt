package tw.nolions.huckebein.demoandroid

import android.util.Log
import com.google.firebase.messaging.FirebaseMessagingService

class HuckebeinFirebaseMessagingService: FirebaseMessagingService() {

    override fun onNewToken(token: String) {
        Log.d("HuckebeinAndroid", "Refreshed token: $token")
    }
}
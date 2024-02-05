<template>
    <div v-if="isOffline" class="v-banner">
        <v-slide-y-transition mode="out-in" appear>
            <v-row>
                <v-col class="d-flex justify-center">
                    <p class="d-none d-sm-flex">"Hello there! It seems like you're currently offline. To access online
                        services and
                        get the most out of our
                        platform,Please check your internet connection and connect online.
                    </p>
                    <p class="d-flex d-sm-none">Please check your internet connection and connect online.</p>
                </v-col>
            </v-row>
        </v-slide-y-transition>
    </div>
</template>
  
<script>
export default {
    data() {
        return {
            isOffline: !navigator.onLine,
        };
    },
    created() {
        window.addEventListener('online', this.updateOnlineStatus);
        window.addEventListener('offline', this.updateOnlineStatus);
    },
    methods: {
        updateOnlineStatus() {
            this.isOffline = !navigator.onLine;
        },
    },
};
</script>
  
<style scoped>
.v-banner {
    /* Your banner styles here */
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    background-color: #f44336;
    /* Red color for error */
    color: white;
    padding: 10px;
    text-align: center;
    z-index: 1000;
    /* Ensure it appears above other page content */
}
</style>
  

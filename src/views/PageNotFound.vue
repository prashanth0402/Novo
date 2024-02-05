<template>
    <div class="container">
        <v-row class="d-flex justify-center flex-column text-center mt-10">
            <v-col class="d-flex flex-column align-center justify-center">
                <!-- <span class="text-h1 font-weight-black">404</span>
                <span class="text-h4">Page Not Found</span> -->
                <v-img src="https://flattrade.s3.ap-south-1.amazonaws.com/promo/Oops! 404 Error with a broken robot.gif"
                    max-width="500" min-width="500" min-height="50" contain
                    lazy-src="https://flattrade.s3.ap-south-1.amazonaws.com/promo/Oops! 404 Error with a broken robot.gif"></v-img>
            </v-col>
            <v-col>
                <p class="caption">Sorry, we were unable to locate this page.You can find more things on our <a
                        href="/dashboard">Homepage</a> ,
                    so don't worry.</p>
            </v-col>
        </v-row>
    </div>
</template>

<script>
export default {
    data: () => ({
        path: "/"
    }),
    // computed: {
    //     router: {
    //         get() {
    //             if (this.$globalData.links.length > 0) {
    //                 return this.$globalData.links[0].path
    //             } else {
    //                 return this.path
    //             }
    //         }
    //     }
    // }
}
</script>

<style scoped>
a {
    text-decoration: none;
}

p a:hover {
    color: #197219;
    text-decoration: underline;
}
</style>
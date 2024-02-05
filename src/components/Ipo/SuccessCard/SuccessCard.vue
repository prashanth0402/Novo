<template>
    <div>
        <v-snackbar class="mt-15" rounded="0"  :timeout="-1" color="green lighten-5" elevation="1" width="200px" v-model="Alert" top right>
            <v-layout>
                <v-flex>
                    <v-icon color="green darken-2" size="20">mdi-check-all</v-icon>
                </v-flex>

                <v-flex>
                    <div class="caption pl-2" style="color: #000000">
                        IPO application submitted. <br />
                        The UPI collect request from the bank might be delayed. Accept it to
                        apply whenever you receive the collect request on your UPI app
                        (might take up to end of the day).
                    </div>
                </v-flex>
                <v-flex>
                    <v-btn icon @click="closeCard">
                        <v-icon class="red--text">mdi-close</v-icon>
                    </v-btn>
                </v-flex>
            </v-layout>
        </v-snackbar>


    </div>
</template>

<script>
export default {
    props: {
        Alert: Boolean,
    },
    watch: {
        Alert: function () {
            if (this.Alert == true) {
                setTimeout(() => {
                    this.closeCard()
                }, 10000);
            }
        }
    },
    methods: {
        closeCard() {
            this.$emit("timeout2");
        },
    },
};
</script>

<style scoped>
.card {
    position: absolute;
    top: 0;
    right: 0;
    z-index: 1;
}
</style>

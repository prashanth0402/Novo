<template>
    <div>
        <v-layout>
            <v-flex>
                <v-tabs active-class="black--text">
                    <v-tab @click="Show('I')" class="text-capitalize " >
                        <span v-if="InvestCount <= 0 || InvestCount == undefined">Invest</span>
                        <v-badge v-else color="blue lighten-3" :content="InvestCount">Invest </v-badge>
                    </v-tab>
                    <v-tab @click="Show('O')" class="text-capitalize ">
                        <span v-if="OrderCount <= 0 || OrderCount == undefined">Order</span>
                        <v-badge v-else color="blue lighten-3" :content="OrderCount">Order</v-badge>
                    </v-tab>
                </v-tabs>
                <v-divider></v-divider>
            </v-flex>
        </v-layout>
    </div>
</template>

<script>
export default {
    name: "NCBTAB",
    methods: {
        Show(flag) {
            if (flag == "I") {
                this.$emit('show', flag);
            } else {
                this.$emit('show', flag);
            }
        },
    },
    props: {
        InvestCount: Number,
        OrderCount: Number,
    },
}
</script>

<style scoped></style>
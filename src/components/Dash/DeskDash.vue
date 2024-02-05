<template>
    <div>
        <v-row class="pa-5">
            <v-col v-for="segment, idx in segmentDetail" :key="idx" cols="12" xl="4" lg="4" md="6" sm="6" xs="12">
                <DashCard :segment="segment" />
            </v-col>
        </v-row>
    </div>
</template>

<script>
import DashCard from './DashCard.vue';

export default {
    props: {
        segmentDetail: Array
    },
    components: { DashCard }
}
</script>

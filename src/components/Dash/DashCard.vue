<template>
    <div>
        <v-hover v-slot="{ hover }">
            <v-card :class="`elevation-${hover ? 5 : 0}`" class="mx-auto pa-6 transition-swing rounded-lg"
                :outlined="!hover" @click="changeRoute(segment.path)" :color="segment.color">
                <v-expand-transition>
                    <v-card-title class=" transition-fast-in-fast-out hover_card_text v-card--reveal "
                        style="height: 100%;">
                        <!-- <v-layout v-if="hover" class="d-flex flex-column">
                            <v-flex>
                                <span v-if="segment.available > 0">
                                    Current available <strong>{{ segment.name }}</strong> : {{
                                        segment.available }}
                                </span>
                                <span v-else class="text--disabled"> No {{ segment.name }}s are open for sale
                                    currently.</span>
                            </v-flex>
                            <v-flex>
                                <span v-if="segment.preApply > 0">Preapply open for {{ segment.preApply
                                }} {{ segment.segment }}</span>
                                <span v-else class="text--disabled">No Preapply {{ segment.name }}'s are available</span>
                            </v-flex>
                            <v-flex>
                                <span v-if="segment.preApply > 0">Upcoming {{ segment.segment }}'s : {{ segment.upcoming
                                }} </span>
                                <span v-else class="text--disabled">Upcoming {{ segment.name }}'s are unavailable</span>
                            </v-flex>
                        </v-layout> -->
                        <v-layout>
                            <v-img :lazy-src="segment.image" :srcset="segment.image" :src="segment.image" height="50"
                                width="50" contain class="mr-2"></v-img>
                            <span class="mt-2">{{ segment.fullname }}</span>
                        </v-layout>
                    </v-card-title>

                </v-expand-transition>
            </v-card>
        </v-hover>
    </div>
</template>

<script>
export default {
    data: () => ({
        Titles: [],
        hovered: false
    }),
    methods: {
        changeRoute(route) {
            if (this.$route.path != route) {
                this.$router.push(route);
            }
        },
    },
    props: {
        segment: Object
    }
}
</script>


<style scoped>
.hover_card_text {
    font-size: 14px;
}

.v-responsive {
    flex: none;
}

.slide {
    transform: translateX(30%);
    transition: all 0.5s;
}

.v-card:hover {
    color: #000;
    cursor: pointer;
}
</style>
<template>
    <v-container>
        <v-slide-y-transition mode="out-in" appear>
            <v-breadcrumbs :items="items">
                <template v-slot:divider>
                    <v-icon>mdi-chevron-right</v-icon>
                </template>
            </v-breadcrumbs>
        </v-slide-y-transition>
        <v-slide-x-transition mode="out-in" appear>
            <MasterControlMobile :IpoMasterData="IpoMasterData" :SgbMasterData="SgbMasterData" :NcbMasterData="NcbMasterData" @GetAllMaster="$emit('GetAllMaster')" v-if="this.$vuetify.breakpoint.width < 600">
            </MasterControlMobile>
            <MasterControlDesk :IpoMasterData="IpoMasterData" :SgbMasterData="SgbMasterData" :NcbMasterData="NcbMasterData" @GetAllMaster="$emit('GetAllMaster')" v-else></MasterControlDesk>
        </v-slide-x-transition>
    </v-container>
</template>
<script>
import MasterControlDesk from './MasterControlComponents/MasterControlDesk.vue';
import MasterControlMobile from './MasterControlComponents/MasterControlMobile.vue'
export default {

    components: {
        MasterControlDesk, MasterControlMobile
    },
    data: () => ({
        items: [
            {
                text: 'Meta Data Setup',
                disabled: true,
            },
            {
                text: 'MasterControl',
                disabled: false,
                href: '/mastercontrol',
            },
        ],
    }),
    props: {
        IpoMasterData: [], SgbMasterData:[],NcbMasterData:[]

    },
    methods:{
        GetAllMaster(){}
    }
};

</script>
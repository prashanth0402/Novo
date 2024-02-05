<template>
    <div>
        <LookUpMob v-if="this.$vuetify.breakpoint.width < 600" :Header="Header" @GetHeaders="GetHeaders"/>
        <LookUpDesk v-else  :Header="Header" @GetHeaders="GetHeaders"/>
    </div>
</template>

<script>
import LookUpDesk from './LookUpComp.vue/LookUpDesk.vue';
import LookUpMob from './LookUpComp.vue/LookUpMob.vue';
export default {
    props: {
        Header: Array
    },
    components: {
        LookUpDesk,
        LookUpMob
    },
    methods:{
        GetHeaders(){
            this.$emit("GetHeaders")
        }
    }
}
</script>


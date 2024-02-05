<template>
    <div style="width:100%" >
        <!-- Mobile Stepper -->
        <v-stepper vertical v-if="this.$vuetify.breakpoint.name == 'xs'">
            <v-stepper-step :complete="true" step="" complete-icon=" mdi-checkbox-blank-circle">
                <span> Offer start</span>
                <small class="small">{{ NewStruct.startDate }}</small>
            </v-stepper-step>

            <v-stepper-content step="" :style="currentDate >= NewStruct.endDate && NewStruct.endDate != '' ?
                'border-left:3px solid #1E88E5 !important' : undefined">
            </v-stepper-content>

            <v-stepper-step step="" complete-icon=" mdi-checkbox-blank-circle"
                :complete="currentDate >= NewStruct.endDate && NewStruct.endDate != '' ? true : false">
                <span>Offer end</span>
                <small class="small">{{ NewStruct.endDate }}</small>
            </v-stepper-step>

            <v-stepper-content step="" :style="currentDate >= NewStruct.allotment && NewStruct.allotment != '' ?
                'border-left:3px solid #1E88E5 !important' : undefined">
            </v-stepper-content>

            <v-stepper-step step="" complete-icon=" mdi-checkbox-blank-circle"
                :complete="currentDate >= NewStruct.allotment && NewStruct.allotment != '' ? true : false">
                <span>Allotment</span>
                <small class="small">{{ NewStruct.allotment }}</small>
            </v-stepper-step>

            <v-stepper-content step="" :style="currentDate >= NewStruct.refund && NewStruct.refund != '' ?
                'border-left:3px solid #1E88E5 !important' : undefined">
            </v-stepper-content>

            <v-stepper-step step="" complete-icon=" mdi-checkbox-blank-circle"
                :complete="currentDate >= NewStruct.refund && NewStruct.refund != '' ? true : false">
                <span>Refund</span>
                <small class="small">{{ NewStruct.refund }}</small>
            </v-stepper-step>
            <v-stepper-content step="" :style="currentDate >= NewStruct.demat && NewStruct.demat != '' ?
                'border-left:3px solid #1E88E5 !important' : undefined">
            </v-stepper-content>

            <v-stepper-step step="" complete-icon=" mdi-checkbox-blank-circle"
                :complete="currentDate >= NewStruct.demat && NewStruct.demat != '' ? true : false">
                <span> Demat</span>
                <small class="small">{{ NewStruct.demat }}</small>
            </v-stepper-step>

            <v-stepper-content step="" :style="currentDate >= NewStruct.listing && NewStruct.listing != '' ?
                'border-left:3px solid #1E88E5 !important' : undefined">
            </v-stepper-content>

            <v-stepper-step step="" complete-icon=" mdi-checkbox-blank-circle"
                :complete="currentDate >= NewStruct.listing && NewStruct.listing != '' ? true : false">
                <span>Listing</span>
                <small class="small">{{ NewStruct.listing }}</small>
            </v-stepper-step>
        </v-stepper>
        <!-- desktop view -->
        <v-stepper alt-labels elevation="0" style="border-bottom: 1px solid #F0F0F0;" v-else>
            <!-- <v-stepper-header :class="StepStruct == undefined ? 'd-flex flex-column' : 'd-flex'"> -->
            <v-stepper-header class="d-flex">
                <v-stepper-step step="" complete-icon=" mdi-checkbox-blank-circle"
                    :complete="currentDate >= NewStruct.startDate && NewStruct.startDate != '' ? true : false">
                    <span> Offer Start</span>
                    <small class="small">{{ NewStruct.startDate }}</small>
                </v-stepper-step>
                <v-divider
                    :color="currentDate >= NewStruct.endDate && NewStruct.endDate != '' ? '#1E88E5' : undefined"></v-divider>

                <v-stepper-step step="" complete-icon=" mdi-checkbox-blank-circle"
                    :complete="currentDate >= NewStruct.endDate && NewStruct.endDate != '' ? true : false">
                    <span>Offer End</span>
                    <small class="small">{{ NewStruct.endDate }}</small>
                </v-stepper-step>

                <v-divider
                    :color="currentDate >= NewStruct.allotment && NewStruct.allotment != '' ? '#1E88E5' : undefined"></v-divider>

                <v-stepper-step step="" complete-icon=" mdi-checkbox-blank-circle"
                    :complete="currentDate >= NewStruct.allotment && NewStruct.allotment != '' ? true : false">
                    <span>Allotment</span>
                    <small class="small">{{ NewStruct.allotment }}</small>
                </v-stepper-step>

                <v-divider
                    :color="currentDate >= NewStruct.refund && NewStruct.refund != '' ? '#1E88E5' : undefined"></v-divider>

                <v-stepper-step step="" complete-icon=" mdi-checkbox-blank-circle"
                    :complete="currentDate >= NewStruct.refund && NewStruct.refund != '' ? true : false">
                    <span>Refund</span>
                    <small class="small">{{ NewStruct.refund }}</small>
                </v-stepper-step>

                <v-divider
                    :color="currentDate >= NewStruct.demat && NewStruct.demat != '' ? '#1E88E5' : undefined"></v-divider>

                <v-stepper-step step="" complete-icon=" mdi-checkbox-blank-circle"
                    :complete="currentDate >= NewStruct.demat && NewStruct.demat != '' ? true : false">
                    <span> Demat</span>
                    <small class="small">{{ NewStruct.demat }}</small>
                </v-stepper-step>

                <v-divider
                    :color="currentDate >= NewStruct.listing && NewStruct.listing != '' ? '#1E88E5' : undefined"></v-divider>

                <v-stepper-step step="" complete-icon=" mdi-checkbox-blank-circle"
                    :complete="currentDate >= NewStruct.listing && NewStruct.listing != '' ? true : false">
                    <span>Listing</span>
                    <small class="small">{{ NewStruct.listing }}</small>
                </v-stepper-step>
            </v-stepper-header>
        </v-stepper>
    </div>
</template>
<script>
export default {
    data() {
        return {
            e6:null,
            currentDate:'',
            show:true
        }
    },
    props: {
        StepStruct: {},
        Item: {}
    },
    created() {
        this.currentDate = new Date().toISOString().slice(0, 10) ;

    
    },
    // mounted(){
    //     console.log(this.NewStruct,"newStruct")
    //     if(this.NewStruct==null||this.NewStruct==undefined){
    //         this.show=false
    //     }
    // },
    
    computed: {
        NewStruct: {
            get() {
                if (this.Item != null) {
                    return this.Item
                }else  {
                    return this.StepStruct
                }
                
            }
        }
    }
}
</script>

<style scoped>
.v-stepper--alt-labels .v-stepper__step {
    flex-direction: column;
    justify-content: flex-start;
    align-items: center;
    flex-basis: 128px;
}

.v-stepper--alt-labels .v-stepper__step {
    flex-basis: 128px;
}

span {
    font-size: 15px;
}

.small {
    font-size: 12px;
    margin-top: 5px;
}

::v-deep .v-stepper__label {
    font-size: 8px;
    margin-top: 5px;
}

::v-deep .v-stepper--alt-labels .v-stepper__step__step {
    margin-bottom: 0px;
}

.v-divider {
    border: 2px solid;
}

::v-deep .v-stepper--alt-labels .v-stepper__header .v-divider {
    margin: 35px -50px 0;
}

/* Mobile Stepper */
::v-deep .v-stepper__content {
    margin: -18px -14px -24px 34px !important;
    /* border-left: 3px solid #1E88E5 !important; */
    border-left: 3px solid #F0F0F0 !important;
    height: 50px !important;
}

::v-deep .v-stepper__wrapper {
    display: none !important;
}

::v-deep .v-sheet.v-stepper:not(.v-sheet--outlined) {
    box-shadow: none;
}
</style>
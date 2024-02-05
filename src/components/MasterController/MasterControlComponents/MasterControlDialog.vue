<template>
    <div>
        <v-dialog v-model="ControlDialog" max-width='600' persistent :fullscreen="this.$vuetify.breakpoint.width <= 600"
            :transition="this.$vuetify.breakpoint.width <= 600 ? 'dialog-bottom-transition' : undefined">
            <v-card elevation="0">
                <v-card-title>
                    <span class="text-h5">{{ FormTitle }}&nbsp;&nbsp;{{ item.symbol }}</span>
                </v-card-title>
                <v-card-text class="text-h6 d-flex justify-center">
                    <span v-if="item.softDelete == 'Y'">
                        Are you Sure Want to Enable {{ item.symbol }} Symbol ?
                    </span>
                    <span v-else>
                        Are you Sure Want to Disable {{ item.symbol }} Symbol ?
                    </span>
                </v-card-text>
                <v-card-actions>
                    <v-spacer></v-spacer>
                    <v-btn color="blue darken-1" text @click="closeDialog()">
                        No
                    </v-btn>
                    <v-btn color="blue darken-1" text @click="SetMasterControl()" >
                        Yes
                    </v-btn>
                </v-card-actions>
            </v-card>
        </v-dialog>
    </div>
</template>
<script>
import EventServices from '@/services/EventServices';
export default {
    props: {
        item: {},
        ControlDialog: Boolean,
        CurrentTittle:String
    },
    computed: {
        FormTitle() {
            return 'Edit'
        },
    },
    methods: {
        closeDialog() {
            this.$emit("closeonly")
        },

        SetMasterControl() {
            this.item.softDelete = this.item.softDelete == "Y" ? "N" :"Y"
            
            EventServices.SetMasterControl(this.item,this.CurrentTittle)
                .then((response) => {
                    if (response.data.status == 'S') {
                        this.MessageBar('S', response.data.errMsg)
                        //   this.$refs.form.resetValidation()
                        this.$emit("closeDialog")
                    } else {
                        this.MessageBar('E', response.data.errMsg)
                    }
                })
                .catch((error) => {
                    console.log("Error :", error)
                })
        },

    },
};
</script>
      
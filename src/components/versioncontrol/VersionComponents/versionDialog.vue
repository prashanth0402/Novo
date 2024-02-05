<template>
  <div>
    <v-dialog v-model="VersionDialog" max-width='1100' persistent :fullscreen="this.$vuetify.breakpoint.width <= 600"
      :transition="this.$vuetify.breakpoint.width <= 600 ? 'dialog-bottom-transition' : undefined">
      <v-card elevation="0">
        <v-card-title>
          <span class="text-h5">{{ FormTitle }}</span>
        </v-card-title>
        <v-card-text class="mt-6">
          <v-slide-x-transition mode="out-in" appear>
            <v-container>
              <v-layout>
                <v-form ref="form" v-model="valid" lazy-validation>
                  <v-row>
                    <v-col cols="12" lg="3" sm="3" md="3">
                      <v-text-field :disabled="true" v-model="VersionManger.os" label="Device Name" outlined
                        dense></v-text-field>
                    </v-col>
                    <v-col cols="12" lg="3" sm="3" md="3">
                      <v-autocomplete v-model="VersionManger.forceUpdate" :items="updateOptions" outlined
                        label="Force Update" dense item-value="Status" item-text="Desc" :rules="statusrules">
                      </v-autocomplete>
                    </v-col>
                    <v-col cols="12" lg="3" sm="3" md="3">
                      <v-text-field v-model="VersionManger.version" label="version" outlined dense
                      :disabled="VersionManger.id == 0 ? false:true"  :rules="nameRules"  @keypress="VersionCharacters"></v-text-field>
                    </v-col>
                    <v-col cols="12" lg="3" sm="3" md="3">
                      <v-autocomplete v-model="VersionManger.appStatus" :items="statusOptions" outlined label="Status"
                        dense item-value="Status" item-text="Desc" :rules="statusrules"> </v-autocomplete>
                    </v-col>
                  </v-row>
                </v-form>
              </v-layout>
            </v-container>
          </v-slide-x-transition>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="blue darken-1" text @click="closeDialog()">
            Cancel
          </v-btn>
          <v-btn color="blue darken-1" text @click="SetVersion()" :disabled="issave">
            Save
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </div>
</template>
<script>
import EventServices from '@/services/EventServices';
export default {
  data() {
    return {
      valid: true,
      statusOptions: [{ Status: 'Y', Desc: 'Active' }, { Status: 'N', Desc: 'InActive' }],
      updateOptions: [{ Status: 'Y', Desc: 'Yes' }, { Status: 'N', Desc: 'No' }],
      nameRules: [v => !!v || 'Version is required',],
      statusrules: [v => !!v || 'This field is required. Please choose a value from the dropdown',]
    };
  },
  props: {
    VersionMangerCopy: {},
    VersionManger: {},
    VersionDialog: Boolean,
  },
  computed: {
    FormTitle() {
      return this.VersionManger.id == 0 ? 'Add Version' : 'Edit Version'

    },
    issave() {

      return JSON.stringify(this.VersionManger) == JSON.stringify(this.VersionMangerCopy)
        || this.VersionManger.os == ""
        || this.VersionManger.forceUpdate == null
        || this.VersionManger.forceUpdate == ""
        || this.VersionManger.version == ""
        || this.VersionManger.appStatus == null
        || this.VersionManger.appStatus == ""

    }
  },
 
  methods: {
    closeDialog() {
      this.$emit("closeOnly")
      this.$refs.form.resetValidation()

    },
    VersionCharacters($event) {
    let keyCode = $event.keyCode ? $event.keyCode : $event.which;

    // Allow numbers (0-9), dot (.), and minus sign (-)
    if (
        (keyCode < 48 || keyCode > 57) &&
        keyCode !== 46 
    ) {
        $event.preventDefault();
    }
},
    SetVersion() {
      EventServices.SetVersion(this.VersionManger)
        .then((response) => {
          if (response.data.status == 'S') {
            this.MessageBar('S', response.data.errMsg)
            this.$refs.form.resetValidation()
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
    
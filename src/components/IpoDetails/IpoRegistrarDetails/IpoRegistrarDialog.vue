<template>
    <div>
      <v-dialog v-model="RegistrarDialog" max-width="600" persistent
        :fullscreen="this.$vuetify.breakpoint.width <= 600"
        :transition="this.$vuetify.breakpoint.width <= 600 ? 'dialog-bottom-transition' : undefined">
        <v-card elevation="0">
          <v-card-title class="">
            <span class="text-h5">{{ FormTitle }}</span>
          </v-card-title>
          <v-card-text class="mt-6">
            <v-slide-x-transition mode="out-in" appear>
              <v-container>
                <v-layout>
                  <v-form ref="form" lazy-validation>
                    <v-row >
                      <v-col cols="12" lg="6" sm="6" md="6">
                        <v-text-field  v-model="Registrar.registrarName" label="Registrar Name" outlined dense
                          :rules="nameRules"></v-text-field>
                      </v-col>
                      <v-col lg="6" cols="12" sm="6" md="6">
                        <v-text-field  v-model="Registrar.registrarLink" label="Registrar Link" outlined dense
                          :rules="linkrules"></v-text-field>
                      </v-col>
                    </v-row>
                  </v-form>
                </v-layout>
              </v-container>
            </v-slide-x-transition>
          </v-card-text>
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn color="blue darken-1" text @click="closeDialog">
              Cancel
            </v-btn>
            <v-btn color="blue darken-1" text @click="SetRegistrar" :disabled="issave">
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
        nameRules: [v => !!v || 'Name is required',],
        linkrules: [
            (v) => !!v || "Address is required",
            // (v) =>
            //     /^https:\/\/[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/.test(v) ||
            //     "Domain Name must be valid (e.g., https://example.com)",
        ],
  
      };
    },
    props: {
      Registrar:{},
      Registrarcopy:{},
      RegistrarDialog: Boolean,
    },
    computed: {
      FormTitle() {
        return this.Registrar.id == 0 ? 'Add Registrar' : 'Edit Registrar'
      },
      issave() {
        return JSON.stringify(this.Registrar) == JSON.stringify(this.Registrarcopy) || this.Registrar.registrarName == "" || this.Registrar.registrarLink == ""
         },
      },
    methods: {
      closeDialog() {
        this.$refs.form.resetValidation()
        this.$emit("closeOnly")
      },
      SetRegistrar() {

        EventServices.SetRegistrar(this.Registrar)
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
    
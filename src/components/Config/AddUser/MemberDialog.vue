<template>
  <div>
    <v-dialog v-model="dialog" max-width="800px" persistent :fullscreen="this.$vuetify.breakpoint.width < 600"
      :hide-overlay="this.$vuetify.breakpoint.width < 600"
      :transition="this.$vuetify.breakpoint.width < 600 ? 'dialog-bottom-transition' : undefined">
      <v-card>
        <v-card-title>
          <span class="text-h5">{{ formTitle }}</span>
        </v-card-title>
        <v-card-text>
          <v-form ref="form" lazy-validation>
          <v-container>
              <v-row>
  
                <v-col cols="12" sm="6" xs="6" md="4">
                  <v-text-field v-model="detail.clientId" label="ClientId" :rules="Requeird"
            ></v-text-field>
                </v-col>
                <v-col cols="12" sm="6" xs="6" md="4">
                  <v-autocomplete v-model="detail.roleId" :items="RoleLisrArr" label="Role" :rules="Requeird" item-text="roleName"
                    item-value="roleId"></v-autocomplete>
                </v-col>
                <v-col cols="12" sm="6" xs="6" md="4">
                  <v-autocomplete v-model="detail.status" :items="statusArr" item-text="text" item-value="value"
                    label="Type"></v-autocomplete>
                </v-col>
              </v-row>
            </v-container>
          </v-form>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="error" text @click="close"> Cancel </v-btn>
          <v-btn color="blue darken-1" text @click="save" :disabled="isSame"> Save </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </div>
</template>

<script>
import EventServices from '../../../services/EventServices';
export default {
  data() {
    return {
      statusArr: [
        { text: "Active", value: "Y" },
        { text: "InActive", value: "N" }
      ],
      Requeird: [(value) => !!value || "Required."],
      detail: {
        id:0,
        clientId: "",
        roleId: 0,
        status: ""
      }
    }
  },
  computed: {
    formTitle() {
      return this.editedIndex === -1 ? "Add Role" : "Edit Role";
    },
    RoleLisrArr() {
      if (this.roleArr != null){

        return this.roleArr.filter(item => item.roleName != 'BrokerSuperAdmin' && item.roleName != 'BrokerAdmin');
      }else{
        return []
      }
    },
    isSame: {
      get() {
        if (this.editedItem != undefined) {
            return this.editedItem.clientId == ""  || this.editedItem.roleId == 0
            || this.editedItem.status == "" || this.editedItem.roleId == null
            || this.editedItem.status == null || JSON.stringify(this.editedItemcopy) == JSON.stringify(this.editedItem)
        } else {
          return false
        }
      }
    }
  },
  props: {
    editedItemcopy:{},
    editedItem: Object,
    editedIndex: Number,
    dialog: Boolean,
    roleArr: Array
  },
  methods: {
    save() {
      if (this.$refs.form.validate()) {
        if (
          this.editedItem.clientid != "" &&
          this.editedItem.role != 0 &&
          this.editedItem.status != ""
        ) {
          // console.log("Requested to save", this.editedItem);
          this.$globalData.overlay = true;
          EventServices.Adduser(this.editedItem)
            .then((response) => {
              this.$globalData.overlay = false;
              if (response.data.status == "S") {
                this.$emit("reInitialize")
                this.MessageBar("S", "Details updated successfully");
              } else {
                this.MessageBar("E", response.data.errMsg);
              }
            })
            .catch((error) => {
              this.$globalData.overlay = false;
              this.MessageBar("E", error);
            });
          this.close();
        } else {
          this.MessageBar("E", "Fill all the details");
        }
      }
    },
    close() {
      // console.log("Closed")
      this.$emit("closeDialog");
      this.$refs.form.resetValidation()
    },
  },
  watch: {
    dialog: function (bool) {
      if (bool == true) {
        this.detail = this.editedItem
      }
    }
  }

}
</script>

<style scoped></style>
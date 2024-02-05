<template>
  <div>
    <v-dialog v-model="NewDialog" :max-width="RoleDialog ? '600' : '800'" persistent
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
                <v-form>
                  <!-- Role Master Dialog -->
                  <v-row v-if="RoleDialog">
                    <v-col cols="12" lg="6" sm="6" md="6">
                      <v-text-field :disabled="Newrole" v-model="RoleMaster.RoleName" label="Role Name" outlined dense
                        :rules="nameRules"></v-text-field>
                    </v-col>
                    <v-col lg="6" cols="12" sm="6" md="6">
                      <v-autocomplete v-model="RoleMaster.Status" :items="statusOptions" outlined label="Status" dense
                        item-value="Status" item-text="Desc" :rules="statusrules"> </v-autocomplete>
                    </v-col>
                  </v-row>
                  <!-- Task Master Dialog -->
                  <v-row v-if="TaskDialog">
                    <v-col lg="3" cols="12" sm="3" md="3">
                      <v-text-field v-model="TaskMaster.RouterName" :rules="nameRules" label="Task Name" outlined
                        dense></v-text-field>
                    </v-col>
                    <v-col lg="3" cols="12" sm="3" md="3">
                      <v-text-field v-model="TaskMaster.Router"  outlined
                        label="Task Link" dense></v-text-field>
                    </v-col>
                    <v-col lg="3" cols="12" sm="3" md="3">
                      <v-autocomplete v-model="TaskMaster.Status" :items="statusOptions" outlined label="Status" dense
                        item-value="Status" item-text="Desc" :rules="statusrules"></v-autocomplete>
                    </v-col>
                    <v-col lg="3" cols="12" sm="3" md="3">
                      <v-autocomplete v-if="isParent == false" v-model="TaskMaster.ParentId" :items="RouterListArr"
                        outlined label="Parent Menu" :rules="statusrules" dense item-text='routerName'
                        item-value="taskId"></v-autocomplete>
                      <v-checkbox v-model="isParent" outlined label="Is Parent ?" dense></v-checkbox>
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
          <v-btn color="blue darken-1" text @click="save" :disabled="issave || isDataChanged">
            Save
          </v-btn>
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
      isParent: false,
      statusOptions: [{ Status: 'Y', Desc: 'Active' }, { Status: 'N', Desc: 'InActive' }],
      nameRules: [v => !!v || 'Name is required',],
      statusrules: [v => !!v || 'This field is required. Please choose a value from the dropdown',],
      linkrules: [v => !!v || 'Link is Requierd',]


    };
  },
  props: {
    TaskMasterCopy: {},
    RoleMastercopy: {},
    TaskListArr: Array,
    TaskMaster: {},
    Newrole: Boolean,
    NewDialog: Boolean,
    RoleMaster: {},
    RoleDialog: Boolean,
    TaskDialog: Boolean
  },
  watch: {
    TaskMaster: {
      immediate: true,
      deep: true,
      handler() {
        this.isParent = this.TaskMaster.ParentId == 0 ? true : false
      }
    }
  },
  computed: {
    RouterListArr() {
      return this.TaskListArr.filter(item => item.taskId !== this.TaskMaster.TaskId && item.router == "");
    },
    FormTitle() {
      if (this.RoleDialog == true) {

        return this.RoleMaster.Id == 0 ? 'Add Role' : 'Edit Role'
      } else {

        return this.TaskMaster.TaskId == 0 ? 'Add Task' : 'Edit Task'
      }
    },
    issave() {
      if (this.TaskDialog) {
        if (this.isParent) {
          return this.TaskMaster.RouterName == null  || this.TaskMaster.Status == null
        } else {
          return this.TaskMaster.RouterName == null || this.TaskMaster.Status == null || this.TaskMaster.ParentId == 0 || this.TaskMaster.ParentId == null
        }
      } else {
        return this.RoleMaster.RoleName == "" || this.RoleMaster.Status == null || this.RoleMaster.RoleName == null
      }
    },
    isDataChanged() {
      if (this.RoleDialog) {
        if (JSON.stringify(this.RoleMastercopy) == JSON.stringify(this.RoleMaster)) {
          return true
        } else {
          return false
        }
      } else {
        if (JSON.stringify(this.TaskMasterCopy) == JSON.stringify(this.TaskMaster)) {
          return true
        } else {
          return false
        }
      }
    }
  },
  methods: {
    // ... (existing methods) ...
    openAddRoleDialog() {
      // Set dialogAddRole to true to open the dialog
      this.dialogAddRole = true;
    },
    closeDialog() {
      this.$emit("closeOnly")
      // Close the dialog and reset editedItem
    },
    SetTask() {
      EventServices.SetTask(this.TaskMaster)
        .then((response) => {
          if (response.data.status == 'S') {
            this.MessageBar('S', response.data.errMsg)
            this.$emit("closeDialog")
          } else {
            this.MessageBar('E', response.data.errMsg)
          }
        })
        .catch((error) => {
          this.MessageBar('E', error)
        })
    },
    SetRole() {
      EventServices.SetRole(this.RoleMaster)
        .then((response) => {
          if (response.data.status == 'S') {
            this.MessageBar('S', response.data.errMsg)
            this.$emit("closeDialog")
          } else {
            this.MessageBar('E', response.data.errMsg)
          }
        })
        .catch((error) => {
          this.MessageBar('E', error)

        })
    },
    save() {
      this.isParent == true ? this.TaskMaster.ParentId = 0 : undefined
      this.RoleDialog == true ? this.SetRole() : this.SetTask()
    },
  },
};
</script>
  
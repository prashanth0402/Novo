<template>
    <div>
        <v-dialog v-model="RoleTaskDialog" max-width="800" persistent :fullscreen="this.$vuetify.breakpoint.width <= 600"
            :transition="this.$vuetify.breakpoint.width <= 600 ? 'dialog-bottom-transition' : undefined">
            <v-card>
                <v-card-title class="">
                    <span class="text-h5">{{ formTitle }}</span>
                </v-card-title>
                <v-card-text class="mt-6">
                    <v-slide-x-transition mode="out-in" appear>
                        <v-container>
                            <v-form ref="form">
                                <v-row>
                                    <v-col lg="4" cols="12" sm="4" md="4">
                                        <v-autocomplete v-model="RoleTask.RoleId" :items="rolename" outlined label="Role"
                                            dense item-text='roleName' item-value="roleId"
                                            :rules="DropDown"></v-autocomplete>
                                    </v-col>
                                    <v-col lg="4" cols="12" sm="4" md="4">
                                        <v-autocomplete v-model="RoleTask.TaskId" :items="taskname" outlined label="Task"
                                            dense item-text='routerName' item-value="taskId"
                                            :rules="DropDown"></v-autocomplete>
                                    </v-col>

                                    <v-col lg="4" cols="12" sm="4" md="4">
                                        <v-autocomplete v-model="RoleTask.Status" :items="statusOptions" outlined
                                            label="Status" dense item-value="Status" item-text="Desc"
                                            :rules="DropDown"></v-autocomplete>
                                    </v-col>
                                </v-row>
                            </v-form>
                        </v-container>
                    </v-slide-x-transition>
                </v-card-text>
                <v-card-actions>
                    <v-spacer></v-spacer>
                    <v-btn color="red darken-1" text @click="closeOnly">
                        Cancel
                    </v-btn>
                    <v-btn color="blue darken-1" text :disabled="issave" @click="SetRoleTaskMaster()">
                        Save
                    </v-btn>
                </v-card-actions>
            </v-card>
        </v-dialog>
    </div>
</template>
<script>
import EventServices from '../../../services/EventServices'
export default {
    data() {
        return {
            DropDown: [v => !!v || 'This field is required. Please choose a value from the dropdown',],
        }
    },
    props: {
        RoleTaskcopy:{},
        rolename: Array,
        taskname: Array,
        statusOptions: Array,
        RoleTaskDialog: Boolean,
        RoleTask: {},
    },
    computed: {
        formTitle() {
            return this.RoleTask.TaskRoleId === 0 ? 'New Task' : 'Edit Task'
        },
        issave() {
            return this.RoleTask.RoleId == null || this.RoleTask.TaskId == null || this.RoleTask.Status == null ||
            JSON.stringify(this.RoleTask) == JSON.stringify (this.RoleTaskcopy)
        }
    },
    methods: {
        SetRoleTaskMaster() {
            // console.log(this.RoleTask,"this.RoleTask")
            EventServices.SetRoleTaskMaster(this.RoleTask)
                .then((response) => {
                    if (response.data.status == 'S') {
                        this.MessageBar('S', response.data.errMsg)
                        this.closeRoleTask();
                    } else {
                        this.MessageBar('E', response.data.errMsg)
                    }
                })
                .catch((error) => {
                    this.MessageBar('E', error)
                })
        },
        closeRoleTask() {
            this.$emit("CloseRoleTask");
            this.$refs.form.resetValidation();
        },
        closeOnly() {
            this.$emit("closeOnly");
            this.$refs.form.resetValidation();

        }


    }
}
</script>


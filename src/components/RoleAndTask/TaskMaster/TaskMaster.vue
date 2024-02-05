<template>
    <div>
        <!-- <v-toolbar-title class="black--text subtitle-1 text--primary mt-3">ROLE TASK MASTER</v-toolbar-title> -->
        <v-layout row wrap class="mb-5" style="padding-left: 12px;">
            <v-flex xs12 lg9 class="mt-6 d-flex justify-left">
                <v-icon left color="blue darken-4" medium>
                    mdi-account-check
                </v-icon>
                <span class="text-subtitle-1">
                    Role Task Master</span>
            </v-flex>
        </v-layout>
        <v-card class="grey lighten-2 elevation-0">
            <v-slide-x-transition mode="out-in" appear>
                <v-data-table :headers="headers" :items="roletask" :items-per-page="10" item-key="Id" :search="search"
                    fixed-header :loading="loading" class="mt-5">
                    <template v-slot:top>
                        <v-card-title>
                            <v-text-field v-model="search" append-icon="mdi-magnify" label="Search" single-line hide-details
                                class="rounded"></v-text-field>
                            <v-divider class="mx-2" inset vertical></v-divider>
                            <v-btn color="primary" text small class="mt-2 text-capitalize" @click="OpenRoleTask">
                                + Add
                            </v-btn>
                        </v-card-title>
                    </template>
                    <template v-slot:item.actions="{ item }">
                        <v-hover v-slot="{ hover }">
                            <v-btn small icon :class="hover ? 'secondary white--text' : 'blue lighten-4 primary--text'">
                                <v-icon small @click="editItem(item)">
                                    mdi-pencil
                                </v-icon>
                            </v-btn>
                        </v-hover>
                    </template>
                </v-data-table>
            </v-slide-x-transition>
        </v-card>
        <TaskMasterDialog :rolename="rolename" :taskname="taskname" :statusOptions="statusOptions"
            :RoleTaskDialog="RoleTaskDialog" @CloseRoleTask="CloseRoleTask" @closeOnly="closeOnly" :RoleTask="RoleTask"
            :RoleTaskcopy="RoleTaskcopy">
        </TaskMasterDialog>
    </div>
</template>
  
<script>
import TaskMasterDialog from './TaskMasterDialog.vue'
import EventServices from '../../../services/EventServices'
export default {
    components: {
        TaskMasterDialog
    },
    data() {
        return {
            loading: false,
            RoleTask: {
                TaskRoleId: 0,
                RoleId: null,
                TaskId: null,
                Status: null
            },
            DefaultRoleTask: {
                TaskRoleId: 0,
                RoleId: null,
                TaskId: null,
                Status: null

            },
            RoleTaskcopy:{},
            RoleTaskDialog: false,
            search: "",
            statusOptions: [{ Status: 'Y', Desc: 'Active' }, { Status: 'N', Desc: 'InActive' }],
            rolename: [],
            taskname: [],
            roletask: [],
            headers: [
                { text: 'ID', align: 'start', sortable: false, value: 'id' },
                { text: 'Role', value: 'roleName' },
                { text: 'Task', value: 'routerName' },
                { text: 'Enabled', value: 'type' },
                { text: 'CreatedDate', value: 'createdDate' },
                { text: 'CreatedBy', value: 'createdBy' },
                { text: 'UpdatedDate', value: 'updatedDate' },
                { text: 'UpdatedBy', value: 'updatedBy' },
                { text: "Actions", value: "actions", sortable: false },
            ],
        }
    },
    methods: {
        OpenRoleTask() {
            this.RoleTaskDialog = true
        },
        CloseRoleTask() {
            this.RoleTaskDialog = false
            this.close()
            this.GetRoleTaskMaster()
        },
        closeOnly() {
            this.RoleTaskDialog = false
            this.close()
        },
        GetRoleTaskMaster() {
            this.loading = true
            EventServices.GetRoleTaskMaster()
                .then((response) => {
                    if (response.data.status == 'S') {
                        this.roletask = response.data.roleTaskArr
                        this.rolename = response.data.roleConnector
                        this.taskname = response.data.taskConnector
                        this.loading = false
                    } else {
                        this.MessageBar('E', response.data.errMsg)
                    }
                })
                .catch((error) => {
                    this.MessageBar('E', error)
                })
        },

        editItem(item) {
            this.RoleTask.TaskRoleId = item.id
            this.RoleTask.RoleId = item.roleId
            this.RoleTask.TaskId = item.taskId
            this.RoleTask.Status = item.type
            this.RoleTaskcopy ={...this.RoleTask}
            this.RoleTaskDialog = true
        },


        close() {
            this.$nextTick(() => {
                this.RoleTask = Object.assign({}, this.DefaultRoleTask)
                this.RoleTaskcopy = Object.assign({}, this.DefaultRoleTask)
            })
        },
    },
    mounted() {
        this.GetRoleTaskMaster()
    }
}
</script>
<template>
    <div>
        <!-- <v-toolbar-title class="black--text subtitle-1 text--primary ml-5">Role Task Master</v-toolbar-title> -->
        <v-layout row wrap class="mb-5" style="padding-left: 12px;">
            <v-flex xs12 lg9 class="mt-6 d-flex justify-left">
                <v-icon left color="blue darken-4" medium>
                    mdi-account-check
                </v-icon>
                <span class="text-subtitle-1">
                    Role Task Master</span>
            </v-flex>
        </v-layout>
        <v-card elevation="0">
            <v-slide-x-transition mode="out-in" appear>
                <v-data-table :headers="TaskMasterMobile" :items="roletask" :search="search" :items-per-page="10"
                    :footer-props="{ 'items-per-page-options': [5] }" :loading="loading"
                    loading-text="Loading... Please wait" no-data-text="No Records available">
                    <template v-slot:top>
                            <v-toolbar flat>
                                <v-text-field v-model="search" append-icon="mdi-magnify" label="Search" single-line hide-details
                    class="rounded"></v-text-field>
                <v-divider class="mx-2" inset vertical></v-divider>
                <v-btn color="primary" text small class="mt-2 text-capitalize" @click="OpenRoleTask">
                    + Add
                </v-btn>
                            </v-toolbar>
                        </template>
                    <template v-slot:item.roleName="{ item }">
                        <v-row wrap class="d-flex text">
                            <v-col cols="6" class="d-flex flex-column text-left">
                                <span class="text--disabled">Role Name</span>
                                <span>{{ item.roleName }}</span>
                            </v-col>
                            <v-col class="d-flex flex-column text-right" cols="6">
                                <span class="text--disabled">Task Name</span>
                                <span>{{ item.routerName }}</span>
                            </v-col>
                        </v-row>
                        <v-row wrap class="text">
                            <v-col class="d-flex text-left" cols="6">
                                <v-layout>
                                    <v-flex class="d-flex flex-column">
                                        <span class="text--disabled">CreatedBy</span>
                                        <span>{{ item.createdBy }}</span>
                                    </v-flex>
                                </v-layout>
                            </v-col>
                            <v-col class="d-flex flex-column text-right" cols="6">
                                <span class="text--disabled">UpdatedBy</span>
                                <span>{{ item.updatedBy }}</span>
                            </v-col>
                        </v-row>
                        <v-row wrap class="text">
                            <v-col class="d-flex text-left" cols="6">
                                <v-layout>
                                    <v-flex class="d-flex flex-column ">
                                        <span class="text--disabled">CreatedDate</span>
                                        <span> {{ item.createdDate }}</span>
                                    </v-flex>
                                </v-layout>
                            </v-col>
                            <v-col class="d-flex flex-column text-right" cols="6">
                                <span class="text--disabled">UpdatedDate</span>
                                <span>{{ item.updatedDate }}</span>
                            </v-col>
                        </v-row>
                        <v-row class="text ">
                            <v-col class="d-flex flex-column text-left" cols="6">
                                <span class="text--disabled">Status</span>
                                <span> {{ item.type == 'Y' ? 'Active' : 'InActive' }}</span>
                            </v-col>
                            <v-col class="text-right align-end" cols="6">
                                <v-hover v-slot="{ hover }">
                                    <v-btn small width="50" elevation="0"
                                        :class="hover ? 'secondary mt-2' : 'text-capitalize blue lighten-4 primary--text mt-2'">
                                        <span @click="editItem(item)">Edit</span>
                                    </v-btn>
                                </v-hover>
                            </v-col>
                        </v-row>
                    </template>
                </v-data-table>
            </v-slide-x-transition>
        </v-card>
        <TaskMasterDialog :rolename="rolename" :taskname="taskname" :statusOptions="statusOptions"
            :RoleTaskDialog="RoleTaskDialog" @CloseRoleTask="CloseRoleTask" @closeOnly="closeOnly" :RoleTask="RoleTask" :RoleTaskcopy="RoleTaskcopy">
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
            TaskMasterMobile: [{ text: "", sortable: false, align: "center", value: "roleName" },],
            // headers: [
            //     { text: 'ID', align: 'start', sortable: false, value: 'id' },
            //     { text: 'Role', value: 'roleName' },
            //     { text: 'Task', value: 'routerName' },
            //     { text: 'Enabled', value: 'type' },
            //     { text: 'CreatedDate', value: 'createdDate' },
            //     { text: 'CreatedBy', value: 'createdBy' },
            //     { text: 'UpdatedDate', value: 'updatedDate' },
            //     { text: 'UpdatedBy', value: 'updatedBy' },
            //     { text: "Actions", value: "actions", sortable: false },
            // ],
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
<style scoped>
::v-deep .v-data-table__mobile-row__header {
    padding-right: 0px !important;
}

.text {
    font-size: .9em;
}

.v-card__title {
    padding: 4px;
}

::v-deep .v-data-table__mobile-row__cell {
    width: 100% !important;
}

::v-deep .v-data-table__mobile-row__cell table tr td:first-child {
    text-align: left;
}

::v-deep .v-data-table__mobile-row__cell table tr td:last-child {
    text-align: right;
}

::v-deep .v-data-footer {
    justify-content: end;
    /* padding: 0;
  padding-left: 200px; */
}

::v-deep .v-data-table__mobile-row__cell table tr {
    padding-bottom: 3px;
}

.row {
    margin-top: 2px;
    margin-bottom: 2px;
}

::v-deep .v-data-table>.v-data-table__wrapper .v-data-table__mobile-row {
    height: initial;
    min-height: 2px;
}

::v-deep .v-data-table>.v-data-table__wrapper>table>tbody>tr>td {
    padding: 16px !important;
}

.col-xl,
.col-xl-auto,
.col-xl-12,
.col-xl-11,
.col-xl-10,
.col-xl-9,
.col-xl-8,
.col-xl-7,
.col-xl-6,
.col-xl-5,
.col-xl-4,
.col-xl-3,
.col-xl-2,
.col-xl-1,
.col-lg,
.col-lg-auto,
.col-lg-12,
.col-lg-11,
.col-lg-10,
.col-lg-9,
.col-lg-8,
.col-lg-7,
.col-lg-6,
.col-lg-5,
.col-lg-4,
.col-lg-3,
.col-lg-2,
.col-lg-1,
.col-md,
.col-md-auto,
.col-md-12,
.col-md-11,
.col-md-10,
.col-md-9,
.col-md-8,
.col-md-7,
.col-md-6,
.col-md-5,
.col-md-4,
.col-md-3,
.col-md-2,
.col-md-1,
.col-sm,
.col-sm-auto,
.col-sm-12,
.col-sm-11,
.col-sm-10,
.col-sm-9,
.col-sm-8,
.col-sm-7,
.col-sm-6,
.col-sm-5,
.col-sm-4,
.col-sm-3,
.col-sm-2,
.col-sm-1,
.col,
.col-auto,
.col-12,
.col-11,
.col-10,
.col-9,
.col-8,
.col-7,
.col-6,
.col-5,
.col-4,
.col-3,
.col-2,
.col-1 {
    width: 100%;
    padding: 0px !important;
}
</style>
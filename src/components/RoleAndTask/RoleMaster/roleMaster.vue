<template>
    <v-card elevation="0">
        <v-tabs v-model="activeTab" left>
            <v-tab>Role</v-tab>
            <v-tab>Task</v-tab>
        </v-tabs>
        <v-window v-model="activeTab" class="elevation-0 mt-9">
            <v-window-item v-for="(tab, index) in tabs" :key="index">
                <v-card v-if="activeTab === index" elevation="0">
                    <!-- <v-toolbar-title class="black--text subtitle-1 text--primary ml-5">{{ currentTitle }}</v-toolbar-title> -->
                    <v-layout row wrap class="mb-5" style="padding-left: 12px;">
                        <v-flex xs12 lg9 class="mt-6 d-flex justify-left">
                            <v-icon left color="blue darken-4" medium>
                                {{ currentIcon }}
                            </v-icon>
                            <span class="text-subtitle-1 text-capitalize">
                                {{ currentTitle }}
                            </span>
                        </v-flex>
                    </v-layout>
                    <v-card elevation="0">
                        <v-toolbar flat>
                            <v-text-field v-model="search" append-icon="mdi-magnify" label="Search" single-line hide-details
                                class="rounded px-2"></v-text-field>
                            <v-divider class="mx-4" inset vertical></v-divider>
                            <v-btn color="primary" text small class="mb-2 text-capitalize" @click="openNewItemDialog()">
                                + Add
                            </v-btn>
                        </v-toolbar>
                        <v-data-table :search="search" :headers="currentHeaders" :items="currentItems" :items-per-page="10" fixed-header
                            :loading="loading" >
                            <template v-slot:item.type="{ item }">
                                {{ item.type == 'Y' ? 'Active' : 'InActive' }}
                            </template>
                            <template v-slot:item.parentId="{ item }">
                                {{ item.parentId == 0 ? 'NA' : item.parentId }}
                            </template>
                            <template v-slot:item.actions="{ item }">
                                <v-hover v-slot="{ hover }">
                                    <v-btn small icon
                                        :class="hover ? 'secondary white--text' : 'blue lighten-4 primary--text'">
                                        <v-icon small @click="editItem(item)">
                                            mdi-pencil
                                        </v-icon>
                                    </v-btn>
                                </v-hover>
                            </template>
                        </v-data-table>
                    </v-card>
                </v-card>
            </v-window-item>
        </v-window>
        <RoleMasterDialog :NewDialog="NewDialog" :Newrole="Newrole" :RoleDialog="RoleDialog" :RoleMaster="RoleMaster"
            :TaskMaster="TaskMaster" :TaskDialog="TaskDialog" @closeDialog="closeDialog" @closeOnly="closeOnly" :RoleMastercopy="RoleMastercopy"
            :TaskListArr="TaskListArr" :TaskMasterCopy="TaskMasterCopy">
        </RoleMasterDialog>
    </v-card>
</template>
<script>
import EventServices from '../../../services/EventServices';
import RoleMasterDialog from './roleTaskDialog.vue';
export default {
    components: {
        RoleMasterDialog
    },
    data() {
        return {
            loading: false,
            DefaultTask: {
                TaskId: 0,
                RouterName: null,
                Router: null,
                ParentId: 0,
                Status: null,
            },
            TaskMaster: {
                TaskId: 0,
                RouterName: null,
                Router: null,
                ParentId: 0,
                Status: null,
            },
            DefaultRole: {
                Id: 0,
                RoleName: null,
                Status: null,
            },
            RoleMaster: {
                Id: 0,
                RoleName: null,
                Status: null,
            },
            TaskMasterCopy:{},
            RoleMastercopy:{},
            TaskDialog: false,
            Newrole: false,
            NewDialog: false,
            RoleDialog: false,
            search: "",
            activeTab: 0,
            isAddingNewItem: true,
            tabs: ['ROLEMASTER', 'TASKMASTER'],
            titleIcon: ['mdi-account-edit', 'mdi-marker-check'],
            RoleTableArr: [],
            TaskTableArr: [],
            TaskListArr: [],
            roleHeaders: [
                { text: 'ID', align: 'start', sortable: false, value: 'roleId' },
                { text: 'RoleName', value: 'roleName' },
                { text: 'Status', value: 'type' },
                { text: 'CreatedBy', value: 'createdBy' },
                { text: 'CreatedDate', value: 'createdDate' },
                { text: 'UpdatedBy', value: 'updatedBy' },
                { text: 'UpdatedDate', value: 'updatedDate' },
                { text: "Actions", value: "actions", sortable: false },
            ],

            taskHeaders: [
                { text: 'ID', align: 'start', sortable: false, value: 'taskId' },
                { text: 'TaskName', value: 'routerName' },
                { text: 'TaskLink', value: 'router' },
                { text: 'ParentMenu', value: 'parentId' },
                { text: 'Status', value: 'type' },
                { text: 'Application', value: 'application' },
                { text: 'CreatedBy', value: 'createdBy' },
                { text: 'CreatedDate', value: 'createdDate' },
                { text: 'UpdatedBy', value: 'updatedBy' },
                { text: 'UpdatedDate', value: 'updatedDate' },
                { text: "Actions", value: "actions", sortable: false },
            ],


        }
    },
    computed: {
        currentHeaders() {
            return this.activeTab === 0 ? this.roleHeaders : this.taskHeaders;
        },
        currentItems() {
            return this.activeTab === 0 ? this.RoleTableArr : this.TaskTableArr;
        },
        currentTitle() {
            return this.tabs[this.activeTab];
        },
        currentIcon() {
            return this.titleIcon[this.activeTab];
        }
    },

    methods: {
        closeDialog() {
            this.close()
            this.TaskDialog = false
            this.RoleDialog = false
            this.NewDialog = false
            this.GetRoleTask()
        },
        closeOnly() {
            this.close()
            this.TaskDialog = false
            this.RoleDialog = false
            this.NewDialog = false
        },
        GetRoleTask() {
            this.loading = true
            EventServices.GetRoleTask()
                .then((response) => {
                    if (response.data.status == 'S') {
                        this.RoleTableArr = response.data.roleMasterArr
                        this.TaskTableArr = response.data.taskMasterArr
                        this.TaskListArr = response.data.taskListArr
                        this.loading = false
                    } else {
                        this.MessageBar('E', response.data.errMsg)
                    }
                })
                .catch((error) => {
                    this.MessageBar('E', error)
                })
        },
        openNewItemDialog() {
            if (this.currentTitle == 'ROLEMASTER') {
                this.NewDialog = true;
                this.RoleDialog = true
                this.Newrole = false
            } else if (this.currentTitle == 'TASKMASTER') {
                this.NewDialog = true;
                this.TaskDialog = true
            }
        },
        close() {
            this.$nextTick(() => {
                this.RoleMaster = Object.assign({}, this.DefaultRole)
                this.TaskMaster = Object.assign({}, this.DefaultTask)
                this.RoleMastercopy = Object.assign({}, this.DefaultRole)
                this.TaskMasterCopy = Object.assign({}, this.DefaultTask)

                // this.editedIndex = -1
            })
        },
        editItem(item) {
            if (this.currentTitle == 'ROLEMASTER') {
                this.RoleMaster.Id = item.roleId
                this.RoleMaster.RoleName = item.roleName
                this.RoleMaster.Status = item.type
                 this.RoleMastercopy = { ...this.RoleMaster };
                this.Newrole = true
                this.NewDialog = true;
                this.RoleDialog = true
            } else if (this.currentTitle == 'TASKMASTER') {
                this.TaskMaster.TaskId = item.taskId
                this.TaskMaster.Status = item.type
                this.TaskMaster.Router = item.router
                this.TaskMaster.ParentId = item.parentId
                this.TaskMaster.RouterName = item.routerName
                this.TaskMasterCopy =  { ...this.TaskMaster };
                this.TaskDialog = true
                this.NewDialog = true
            }
        },
    },
    mounted() {
        this.GetRoleTask()
    }

};

</script>
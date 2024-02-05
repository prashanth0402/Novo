<template>
    <div>
        <v-tabs v-model="activeTab" left>
            <v-tab>Role</v-tab>
            <v-tab>Task</v-tab>
        </v-tabs>
        <v-window v-model="activeTab" class="elevation-0 mt-9">
            <v-window-item v-for="(tab, index) in tabs" :key="index">
                <v-card v-if="activeTab === index" elevation="0">
                    <!-- <v-toolbar-title class="black--text subtitle-1 text--primary ml-5"></v-toolbar-title> -->
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
                    <v-data-table :headers="currentHeaders" :items="currentItems" :search="search" :items-per-page="10"
                        :footer-props="{ 'items-per-page-options': [10], pageText: '', }" :loading="loading" 
                        loading-text="Loading... Please wait" no-data-text="No Records available">
                        <template v-slot:top>
                            <v-toolbar flat>
                                <v-text-field v-model="search" append-icon="mdi-magnify" label="Search" single-line hide-details
                                class="rounded"></v-text-field>
                                <v-divider class="mx-2" inset vertical></v-divider>
                                <v-btn color="primary" text small class="mt-2 text-capitalize" @click="openNewItemDialog">
                                    + Add
                                </v-btn>
                            </v-toolbar>
                        </template>
                        <template v-slot:item.roleName="{ item }" v-if="activeTab == 0">
                            <v-row wrap class="d-flex text">
                                <v-col cols="6" class="d-flex flex-column text-left">
                                    <span class="text--disabled">Role Name</span>
                                    <span>{{ item.roleName }}</span>
                                </v-col>
                                <v-col class="d-flex flex-column text-right" cols="6">
                                    <span class="text--disabled">Status</span>
                                    <span> {{ item.type == 'Y' ? 'Active' : 'InActive' }}</span>
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
                                        <v-flex class="d-flex flex-column">

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
                            <v-row wrap class="text">
                                <v-col class="text-center" cols="12">
                                    <v-hover v-slot="{ hover }">
                                        <v-btn small width="80"
                                            :class="hover ? 'secondary' : 'text-capitalize blue lighten-4 primary--text elevation-0'">
                                            <span @click="editItem(item)">Edit</span>
                                        </v-btn>
                                    </v-hover>
                                </v-col>
                            </v-row>
                        </template>
                        <template v-slot:item.routerName="{ item }" v-else>
                            <v-row wrap class="d-flex text">
                                <v-col cols="6" class="d-flex flex-column text-left">
                                    <span class="text--disabled">Task Name</span>
                                    <span>{{ item.routerName }}</span>
                                </v-col>
                                <v-col class="d-flex flex-column text-right" cols="6">
                                    <span class="text--disabled">Task Link</span>
                                    <span>{{ item.router }}</span>
                                </v-col>
                            </v-row>
                            <v-row wrap class="d-flex text">
                                <v-col cols="6" class="d-flex flex-column text-left">
                                    <span class="text--disabled">Parent Menu</span>
                                    <span> {{ item.parentId == 0 ? 'NA' : item.parentId }}</span>
                                </v-col>
                                <v-col class="d-flex flex-column text-right" cols="6">
                                    <span class="text--disabled">Status</span>
                                    <span> {{ item.type == 'Y' ? 'Active' : 'InActive' }}</span>
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
                                        <v-flex class="d-flex flex-column">
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
                            <v-row wrap class="text">
                                <v-col class="text-center" cols="12">
                                    <v-hover v-slot="{ hover }">
                                        <v-btn small width="80"
                                            :class="hover ? 'secondary' : 'text-capitalize blue lighten-4 primary--text elevation-0'">
                                            <span @click="editItem(item)">Edit</span>
                                        </v-btn>
                                    </v-hover>
                                </v-col>
                            </v-row>
                        </template>
                    </v-data-table>
                </v-card>
            </v-window-item>
        </v-window>
        <roleTaskDialog :NewDialog="NewDialog" :Newrole="Newrole" :RoleDialog="RoleDialog" :RoleMaster="RoleMaster"
            :TaskMaster="TaskMaster" :TaskDialog="TaskDialog" @closeDialog="closeDialog" @closeOnly="closeOnly"
            :TaskListArr="TaskListArr" :RoleMastercopy="RoleMastercopy" :TaskMasterCopy="TaskMasterCopy">
        </roleTaskDialog>
    </div>
</template>
<script>
import EventServices from '../../../services/EventServices';
import roleTaskDialog from './roleTaskDialog.vue';
export default {
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


            roleMasterMobile: [{ text: "", sortable: false, align: "center", value: "roleName"},],
            TaskMasterMobile: [{ text: "", sortable: false, align: "center", value: "routerName"},],
        }
    },
    components: {
        roleTaskDialog
    },
    computed: {
    //     searchItems(){
    //        return this.search !="" ? this.filteredItems : this.currentItems
    //     },
    //     filteredItems() {
    //   const search = this.search.toLowerCase();
    //   return this.currentItems.filter((item) =>
    //     item.roleName.toLowerCase().includes(search) ||
    //     item.type.toLowerCase().includes(search) ||
    //     item.createdBy.toLowerCase().includes(search) ||
    //     item.updatedBy.toLowerCase().includes(search) ||
    //     item.createdDate.toLowerCase().includes(search) ||
    //     item.updatedDate.toLowerCase().includes(search)
    //   )
    // },
    currentHeaders() {
            return this.activeTab === 0 ? this.roleMasterMobile : this.TaskMasterMobile;
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
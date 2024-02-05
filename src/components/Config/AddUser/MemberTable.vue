<template>
    <div>
        <v-layout row wrap class="mb-3 mx-1 d-flex">
            <v-flex xs12 lg9 class="d-flex justify-start">
                <v-icon left color="blue darken-4" medium>
                    mdi-account-multiple-plus
                </v-icon>
                <span class="text-subtitle-1">Manage User</span>
            </v-flex>
            <!-- <v-flex class="d-flex justify-end align-center">
                <v-btn @click="manualFetch" small width="150" class="caption text-capitalize elevation-0" color="F0F0F0">
                    <v-icon size="15" class="primary--text mr-1">mdi-cached</v-icon>manualFetch</v-btn>
            </v-flex> -->
        </v-layout>
        <v-card elevation="0">
            <v-data-table :headers="this.$vuetify.breakpoint.name == 'xs' ? headerMobile : headerDesk" :items="MemberList"
                :search="search" :loading="loading" loading-text="Loading... Please wait"
                no-data-text="No Records available"
                :footer-props="hidetab == true ? { 'items-per-page-options': [5, 10, 15, -1] } : { 'items-per-page-options': [5] }">
                <template v-slot:top>
                    <v-toolbar flat>
                        <v-text-field v-model="search" append-icon="mdi-magnify" label="Search" single-line
                            hide-details></v-text-field>
                        <v-divider class="mx-4" inset vertical></v-divider>
                        <v-btn color="primary" text small class="mb-2 text-capitalize" @click="add">
                            + Add
                        </v-btn>
                    </v-toolbar>
                </template>
                <!-- <template v-slot:item.clientId="{ item }">
                    <span>{{  + 1 }}</span>
                </template> -->
                <template v-slot:item.status="{ item }" v-if="hidetab">
                    <span>{{ item.status == "Y" ? "Active" : "Inactive" }}</span>
                </template>
                <template v-slot:item.action="{ item }" v-if="hidetab">
                    <v-layout>
                        <v-flex class="d-flex align-center">
                            <v-hover v-slot="{ hover }">
                                <v-btn small icon :class="hover ? 'secondary white--text' : 'blue lighten-4 primary--text'"
                                    @click="editItem(item)" elevation="0">
                                    <v-icon small>mdi-pencil</v-icon>
                                </v-btn>
                            </v-hover>
                        </v-flex>
                    </v-layout>
                </template>
                <template v-if="!hidetab" v-slot:item.clientId="{ item }">
                    <v-row wrap class="d-flex text">
                        <v-col cols="6" class="d-flex flex-column text-left">
                            <span class="text--disabled">ClientId</span>
                            <span>{{ item.clientId }}</span>
                            <!-- <span class="text--disabled">Unit</span>
                            <span>{{ item.unit }}</span> -->
                        </v-col>
                        <v-col class="d-flex flex-column text-right text" cols="6">
                            <span class="text--disabled">Role</span>
                            <span>{{ item.roleName }}</span>
                        </v-col>
                    </v-row>
                    <v-row wrap class="text d-flex align-center">
                        <v-col class="d-flex text-left" cols="6">
                            <v-layout>
                                <v-flex class="d-flex flex-column">
                                    <span class="text--disabled">Status</span>
                                    <span>{{ item.status }}</span>
                                </v-flex>
                            </v-layout>
                        </v-col>
                        <v-col class="d-flex justify-end  text-left" cols="6">
                            <v-hover v-slot="{ hover }">
                                <v-btn small :class="hover ? 'secondary white--text' : 'blue lighten-4 primary--text'"
                                    @click="editItem(item)" elevation="0" width="100">
                                    <span class="text-capitalize">Edit</span>
                                </v-btn>
                            </v-hover>
                        </v-col>
                    </v-row>
                </template>
            </v-data-table>
        </v-card>
        <MemberDialog :editedIndex="editedIndex" :editedItem="editedItem" :dialog="dialog" :roleArr="Role"
            @closeDialog="close" @reInitialize="initialize()" :editedItemcopy="editedItemcopy"/>
    </div>
</template>
<script>
import MemberDialog from './MemberDialog.vue';
import EventServices from "@/services/EventServices";
export default {
    components: {
        MemberDialog
    },
    data: () => ({
        search: "",
        dialog: false,
        dialogDelete: false,
        checkBox: false,
        editedIndex: -1,
        editedItemcopy:{},
        editedItem: {
            id:0,
            roleId: 0,
            clientId:null,
            status: null,
        },
        defaultItem: {
            id:0,
            roleId: 0,
            clientId: null,
            status: null,
        },
        headerDesk: [
            { text: "ClientId", value: "clientId" },
            { text: "Role", value: "roleName" },
            { text: "Status", value: "status" },
            { text: "Actions", value: "action", sortable: false, editable: false }
        ],
        headerMobile: [{ text: "", value: "clientId", sortable: false, editable: false }],
        Role: [],
        MemberList: [],
        loading: false

    }),
    mounted() {
        this.initialize();
    },
    methods: {
        initialize() {
            this.loading = true;
            EventServices.GetMemberUser()
                .then((response) => {
                    this.loading = false;
                    if (response.data.status == "S") {
                        this.MemberList = response.data.memberListArr;
                        this.Role = response.data.roleListArr
                    } else {
                        this.loading = false;
                        this.MessageBar("E", response.data.errMsg);
                    }
                })
                .catch((error) => {
                    this.loading = false;
                    this.MessageBar("E", error);
                });

        },

        editItem(item) {
            this.editedIndex = this.MemberList.indexOf(item);
            this.editedItem = Object.assign({}, item);
            this.editedItemcopy =  Object.assign({}, item);
            this.dialog = true;
        },
        add() {
            this.dialog = true;
        },

        close() {
            this.dialog = false;
            // this.$nextTick(() => {
                this.editedItem = Object.assign({}, this.defaultItem);
                this.editedIndex = -1;
                
            // });
        },
    },
    computed: {
        hidetab: {
            get() {
                if (this.$vuetify.breakpoint.name == 'xs') {
                    return false
                } else {
                    return true
                }
            }
        },
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
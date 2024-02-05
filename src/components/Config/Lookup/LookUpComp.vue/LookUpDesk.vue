<template>
    <v-card elevation="0">
        <v-tabs v-model="activeTab" left>
            <v-tab>Header</v-tab>
            <v-tab>Detail</v-tab>
        </v-tabs>
        <v-window v-model="activeTab" class="elevation-0 mt-9">
            <v-window-item v-for="(tab, index) in tabs" :key="index">
                <v-card v-if="activeTab === index" elevation="0">
                    <v-layout row wrap class="mb-5" style="padding-left: 12px;">
                        <v-flex xs12 lg9 class="mt-6 d-flex justify-left align-center">
                            <v-icon left color="blue darken-4" medium>
                                {{ currentIcon }}
                            </v-icon>
                            <span class="text-subtitle-1 text-capitalize">
                                {{ currentTitle }}
                            </span>
                            <span class="caption ml-5 primary--text" v-if="activeTab == 0"> ( Double Click the row to get
                                it's Detail )</span>
                        </v-flex>
                    </v-layout>
                    <v-card elevation="0">
                        <v-data-table :headers="currentHeaders" :items="currentItems" :items-per-page="10" fixed-header
                            :loading="loading" @dblclick:row="FetchDetail" :search="search">
                            <template v-slot:top>
                                <v-toolbar flat>
                            <v-text-field v-model="search" append-icon="mdi-magnify" label="Search" single-line hide-details
                                class="rounded px-2"></v-text-field>
                            <v-divider class="mx-4" inset vertical></v-divider>
                            <v-btn color="primary" text small class="mb-2 text-capitalize" @click="openNewItemDialog()"
                                :disabled="currentTitle == 'Detail' && details.length <= 0 && DetailItem.headerId == '' ? true : false">
                                + Add
                            </v-btn>
                        </v-toolbar>
                            </template>
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
        <LookUpDialog :NewDialog="NewDialog" :HeaderDialog="HeaderDialog" :HeaderItem="HeaderItem" :DetailItem="DetailItem"
            :DetailDialog="DetailDialog" @closeDialog="closeDialog" @closeOnly="closeOnly" @DetailsDialog="DetailsDialog"
            :DetailItemcopy="DetailItemcopy" :HeaderItemcopy="HeaderItemcopy">
        </LookUpDialog>
    </v-card>
</template>
<script>
import EventServices from '../../../../services/EventServices';
import LookUpDialog from './LookUpDialog.vue';
export default {
    props: {
        Header: Array
    },
    components: {
        LookUpDialog
    },
    data() {
        return {
            loading: false,
            defaulHead: {
                id: "",
                user: "",
                code: "",
                description: "",
            },
            HeaderItem: {
                id: "",
                user: "",
                code: "",
                description: "",
            },
            defaultDetail: {
                id: "",
                headerId: "",
                code: "",
                description: "",
                attribute: "",
            },
            DetailItem: {
                id: "",
                headerId: "",
                code: "",
                description: "",
                attribute: "",
            },
            HeaderItemcopy:{},
            DetailItemcopy:{},
            NewDialog: false,
            DetailDialog: false,
            HeaderDialog: false,
            search: "",
            activeTab: 0,
            isAddingNewItem: true,
            tabs: ['Header', 'Detail'],
            titleIcon: ['mdi-account-edit', 'mdi-text-box-edit-outline'],
            Headers: [
                // { text: "Header ID", value: "headerID", class: "blue lighten-5", width: '5px' },
                { text: 'ID', align: 'start', sortable: false, value: 'id' },
                { text: "Code", value: "code" },
                { text: "Description", value: "description" },
                { text: "Created By", value: "createdBy" },
                { text: "Created Date", value: "createdDate" },
                { text: "Updated By", value: "updatedBy" },
                { text: "Updated Date", value: "updatedDate" },
                { text: 'Actions', value: 'actions', sortable: false },
            ],

            Details: [
                { text: 'ID', align: 'start', sortable: false, value: 'id' },
                { text: 'Code', value: 'code' },
                { text: "Description", value: "description" },
                { text: "Attribute", value: "attribute" },
                { text: "Created By", value: "createdBy" },
                { text: "Created Date", value: "createdDate" },
                { text: "Updated By", value: "updatedBy" },
                { text: "Updated Date", value: "updatedDate" },
                { text: "Actions", value: "actions", sortable: false },
            ],
            details: []
        }
    },
    computed: {
        currentHeaders() {
            return this.activeTab === 0 ? this.Headers : this.Details;
        },
        currentItems() {
            return this.activeTab === 0 ? this.Header : this.details;
        },
        currentTitle() {
            return this.tabs[this.activeTab];
        },
        currentIcon() {
            return this.titleIcon[this.activeTab];
        }
    },

    methods: {
        DetailsDialog(HeaderId){
            this.DetailDialog = false
            this.HeaderDialog = false
            this.NewDialog = false
            this.FetchLookUp(HeaderId)
            this.close()

        },
        closeDialog() {
            this.close()
            this.DetailDialog = false
            this.HeaderDialog = false
            this.NewDialog = false
            // this.$globalData.overlay =false
            
            this.GetHeaders()
        },
        closeOnly() {
            this.close()
            this.DetailDialog = false
            this.HeaderDialog = false
            this.NewDialog = false
            // this.$globalData.overlay =false

        },
        GetHeaders(){
this.$emit('GetHeaders')
        },

        openNewItemDialog() {
            if (this.currentTitle == 'Header') {
                this.HeaderDialog = true
                this.NewDialog = true;
            } else if (this.currentTitle == 'Detail') {
                this.DetailDialog = true
                this.NewDialog = true;
            }
        },
        close() {
            this.$nextTick(() => {
                this.HeaderItem = Object.assign({}, this.defaulHead)
                this.DetailItem = Object.assign({}, this.defaultDetail)
                this.HeaderItemcopy = Object.assign({}, this.defaulHead)
                this.DetailItemcopy = Object.assign({}, this.defaultDetail)
                // this.editedIndex = -1
            })
        },
        editItem(item) {

            if (this.currentTitle == 'Header') {
                this.HeaderItem.id = item.id
                this.HeaderItem.code = item.code
                this.HeaderItem.description = item.description
                this.HeaderItem.user = item.user
                this.HeaderItemcopy = {...this.HeaderItem}
                this.HeaderDialog = true
                this.NewDialog = true;
                // console.log(item);
            } else if (this.currentTitle == 'Detail') {
                this.DetailItem.id = item.id
                this.DetailItem.headerId = item.headerId
                this.DetailItem.code = item.code
                this.DetailItem.description = item.description
                this.DetailItem.attribute = item.attribute
this.DetailItemcopy = {...this.DetailItem}
                this.DetailDialog = true
                this.NewDialog = true;
            }
        },
        FetchDetail(event,row) {
            const item = row.item;
            if (this.activeTab == 0) {
                this.FetchLookUp(item.id)
            }

        },
        FetchLookUp(HeaderId){
            this.$globalData.overlay = true;
            EventServices.FetchLookUpDetails(HeaderId)
                    .then((response) => {
                        if (response.data.status == "S") {
                            if (response.data.details != [] && response.data.details != null) {
                                this.details = response.data.details;
                                this.DetailItem.headerId = HeaderId
                                this.activeTab = 1;
                            } else {
                                this.details = []
                                this.activeTab = 1
                                this.DetailItem.headerId = HeaderId
                            }
                            this.$globalData.overlay = false;
                        } else {
                            this.MessageBar("E", response.data.errMsg)
                            this.$globalData.overlay = false;

                        }
                    })
                    .catch((error) => {
                        this.MessageBar("E", error)
                        this.$globalData.overlay = false;

                    });
        },
    },

};

</script>
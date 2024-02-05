<template>
               <div>

                   <!-- <v-toolbar-title class="black--text subtitle-1 text--primary ml-5">{{ currentTitle }}</v-toolbar-title> -->
                   <!-- <v-layout row wrap class="mb-5" style="padding-left: 12px;">
                       <v-flex xs12 lg9 class="mt-6 d-flex justify-left">
                           <v-icon left color="blue darken-4" medium>
                               mdi-eye
                            </v-icon>
                            <span class="text-subtitle-1 text-capitalize">
                                Ipo Registrar Details
                            </span>
                        </v-flex>
                    </v-layout> -->
                    <v-card elevation="0">
                        <v-toolbar flat>
                            <v-text-field v-model="search" append-icon="mdi-magnify" label="Search" single-line hide-details
                            class="rounded px-2"></v-text-field>
                            <v-divider class="mx-4" inset vertical></v-divider>
                            <v-btn color="primary" text small class="mb-2 text-capitalize" @click="openNewRegistrar()">
                                + Add
                            </v-btn>
                        </v-toolbar>
                        <v-data-table :search="search" :headers="RegistrarHeaders" :items="RegistrarArr" :items-per-page="10" fixed-header
                        :loading="loading" >
                        <!-- <template v-slot:item.id="{ index}">
                          {{ index+1 }}
                    </template> -->
                    <template v-slot:item.registrarLink="{ item}">
                        <span class="d-inline-block text-truncate" style="max-width: 150px">{{
            item.registrarLink
          }}</span>
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
            <IpoRegistrarDialog :RegistrarDialog="RegistrarDialog" :Registrar="Registrar" :Registrarcopy="Registrarcopy"
            @closeDialog="closeDialog" @closeOnly="closeOnly"></IpoRegistrarDialog>
        <!-- </v-card> -->
    </div>
    </template>
<script>
import EventServices from '@/services/EventServices';
import IpoRegistrarDialog from './IpoRegistrarDialog.vue';

export default {
    components: {
        IpoRegistrarDialog
},
    data() {
        return {
            RegistrarDialog:false,
            loading: false,
            Registrar: {
               id:0,
               registrarName:"",
               registrarLink:""
            },
            DefaultRegistrar:{
                id:0,
               registrarName:"",
               registrarLink:""
            },
            Registrarcopy:{},
            search: "",
            RegistrarArr: [],
            RegistrarHeaders: [
                // { text: 'S.No', align: 'start', sortable: false, value: 'id' },
                { text: 'RegistrarName', value: 'registrarName' },
                { text: 'RegistrarLink', value: 'registrarLink' },
                { text: 'CreatedBy', value: 'createdBy' },
                { text: 'CreatedDate', value: 'createdDate' },
                { text: 'UpdatedBy', value: 'updatedBy' },
                { text: 'UpdatedDate', value: 'updatedDate' },
                { text: "Actions", value: "actions", sortable: false },
            ],

        }
    },

    methods: {
        closeDialog() {
            this.close()
            this.RegistrarDialog = false
            this.GetRegistrarDetails()
        },
        closeOnly() {
            this.close()
            this.RegistrarDialog = false
        },
        GetRegistrarDetails() {
            this.loading = true
            EventServices.GetRegistrarDetails()
                .then((response) => {
                    if (response.data.status == 'S') {
                        this.RegistrarArr = response.data.registrarList == null ? [] : response.data.registrarList
                        this.loading = false
                    } else {
                        this.MessageBar('E', response.data.errMsg)
                    }
                })
                .catch((error) => {
                    console.log("Error :", error)
                })
        },
        openNewRegistrar() {
            this.RegistrarDialog = true           
        },
        close() {
            this.$nextTick(() => {
                this.Registrar = Object.assign({}, this.DefaultRegistrar)
            })
        },
        editItem(item) {
            this.Registrar.id = item.id
            this.Registrar.registrarLink = item.registrarLink
            this.Registrar.registrarName = item.registrarName
            this.Registrarcopy ={...this.Registrar}
          this.RegistrarDialog = true
        },
    },
    mounted() {
        this.GetRegistrarDetails()
    }

};

</script>
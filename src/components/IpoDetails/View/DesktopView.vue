<template>
  <!-- Use this page as a reference to design both mobile and desktop view-->
  <div>
    <!-- <v-layout row wrap class="mb-5" style="padding-left: 12px;">
      <v-flex xs12 lg9 class="mt-6 d-flex justify-left">
        <v-icon left color="blue darken-4" medium>
          mdi-link-variant
        </v-icon>
        <span class="text-subtitle-1">IPO
          Link Details</span>
      </v-flex>
    </v-layout> -->

    <v-slide-x-transition mode="out-in" appear>
      <v-data-table :headers="headers1" :items="linksArr" sort-by="calories" class="elevation-0" :loading="loading"
        :search="search" :items-per-page="10">
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
        <template v-slot:item.blogLink="{ item }">
          <span class="d-inline-block text-truncate" style="max-width: 170px">{{
            item.blogLink
          }}</span>
        </template>
        <template v-slot:item.drhpLink="{ item }">
          <span class="d-inline-block text-truncate" style="max-width: 170px">{{
            item.drhpLink
          }}</span>
        </template>
        <template v-slot:item.sme="{ item }">
          <span v-if="item.sme == true">SME</span>
        </template>
        <template v-slot:item.actions="{ item }">
          <v-hover v-slot="{ hover }">
            <v-btn small icon :class="hover ? 'secondary' : 'blue lighten-4'">
              <v-icon small @click="editItem(item)" :class="hover ? 'white--text' : 'primary--text'"> mdi-pencil
              </v-icon>
            </v-btn>
          </v-hover>
        </template>
      </v-data-table>
    </v-slide-x-transition>

    <v-dialog v-model="dialog" max-width="800px" persistent>
      <v-card>
        <v-card-title>
          <span class="text-h5">{{ formTitle }}</span>
        </v-card-title>
        <v-card-text>
          <v-slide-x-transition mode="out-in" appear>
            <v-container>
              <v-form ref="form" lazy-validation>
                <v-row>
                  <v-col cols="12" sm="6" xs="6" md="4">
                    <v-text-field v-model="editedItem.symbol" label="Symbol"></v-text-field>
                  </v-col>
                  <v-col cols="12" sm="6" xs="6" md="4">
                    <v-text-field v-model="editedItem.isin" label="ISIN" :rules="rules"></v-text-field>
                  </v-col>
                  <v-col cols="12" sm="6" xs="6" md="4">
                    <v-text-field v-model="editedItem.blogLink" label="BLOG Link"></v-text-field>
                  </v-col>
                  <v-col cols="12" sm="6" xs="6" md="4">
                    <v-text-field v-model="editedItem.drhpLink" label="DHRP Link" :rules="rules"></v-text-field>
                  </v-col>
                  <v-col cols="12" sm="6" xs="6" md="4">
                    <v-menu v-model="menu1" :close-on-content-click="false" :nudge-right="40">
                      <template v-slot:activator="{ on, attrs }" style="left:40px;top: 343px;">
                        <v-text-field v-model="editedItem.allotmentFinal" label="Allotment" readonly v-bind="attrs"
                          v-on="on"></v-text-field>
                      </template>
                      <v-date-picker v-model="editedItem.allotmentFinal" @input="menu1 = false" color="primary"
                        no-title></v-date-picker>
                    </v-menu>
                  </v-col>
                  <v-col cols="12" sm="6" xs="6" md="4">
                    <v-menu v-model="menu2" :close-on-content-click="false" :nudge-right="40">
                      <template v-slot:activator="{ on, attrs }" style="left:40px;top: 343px;">
                        <v-text-field v-model="editedItem.refundInitiate" label="Refunds" readonly v-bind="attrs"
                          v-on="on"></v-text-field>
                      </template>
                      <v-date-picker v-model="editedItem.refundInitiate" @input="menu2 = false" color="primary"
                        no-title></v-date-picker>
                    </v-menu>
                  </v-col>
                  <v-col cols="12" sm="6" xs="6" md="4">
                    <v-menu v-model="menu3" :close-on-content-click="false" :nudge-right="40">
                      <template v-slot:activator="{ on, attrs }" style="left:40px;top: 343px;">
                        <v-text-field v-model="editedItem.dematTransfer" label="Demat" readonly v-bind="attrs"
                          v-on="on"></v-text-field>
                      </template>
                      <v-date-picker v-model="editedItem.dematTransfer" @input="menu3 = false" color="primary"
                        no-title></v-date-picker>
                    </v-menu>
                  </v-col>
                  <v-col cols="12" sm="6" xs="6" md="4">
                    <v-menu v-model="menu4" :close-on-content-click="false" :nudge-right="40">
                      <template v-slot:activator="{ on, attrs }" style="left:40px;top: 343px;">
                        <v-text-field v-model="editedItem.listingDate" label="Listing" readonly v-bind="attrs"
                          v-on="on"></v-text-field>
                      </template>
                      <v-date-picker v-model="editedItem.listingDate" @input="menu4 = false" color="primary"
                        no-title></v-date-picker>
                    </v-menu>
                  </v-col>
                  <v-col cols="1" sm="1" xs="1" md="1">
                    <v-checkbox v-model="editedItem.sme" label="SME"></v-checkbox>
                  </v-col>
                </v-row>
              </v-form>
            </v-container>
          </v-slide-x-transition>
        </v-card-text>

        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="error" text @click="close"> Cancel </v-btn>
          <v-btn color="blue darken-1" text @click="save" :disabled="issave"> Save </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </div>
</template>
<script>
import Info from "../../Script/IpoInfo/Info"
export default {
  computed: {
    issave() {
      return  this.editedItem.allotmentFinal == ""  || this.editedItem.dematTransfer == "" || this.editedItem.drhpLink == "" || this.editedItem.isin == "" || this.editedItem.refundInitiate == "" || this.editedItem.listingDate == ""
    }
  },
  mixins: [Info],
}
</script>

<style scoped>
.row {
  margin-top: 2px;
  margin-bottom: 2px;
}

.text {
  font-size: 10px;
}

::v-deep .v-data-table__mobile-row__header {
  display: none !important;
}

::v-deep .v-data-table__mobile-row__cell {
  width: 100%;
}

::v-deep .v-data-footer {
  justify-content: end;
}

/* ::v-deep .v-input__control .v-input__slot .v-messages .v-messages__wrapper .v-messages__message {
    font-size: 9px;
} */

::v-deep .v-messages__message {
  /* width: 0%; */
  font-size: 9px;
}
</style>
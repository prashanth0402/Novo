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
      <v-data-table :headers="headers2" :items="linksArr" sort-by="calories" :search="search" :loading="loading"
        :footer-props="{ 'items-per-page-options': [5] }">
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
        <template v-slot:item.symbol="{ item }">
          <div style="width: 100%">
            <v-layout class="mb-1 mt-2">
              <v-flex class="d-flex justify-space-between text">
                <span>
                  <span><b>{{ item.symbol }} </b></span>
                  <!-- <span v-if="item.sme == true"> ( SME )</span> -->
                  <v-chip v-show="item.sme == true" x-small color="blue lighten-5" label
                    class="smechip--text">SME</v-chip>
                </span>
                <span>
                  <v-btn icon small class="blue lighten-4 primary--text">
                    <v-icon x-small @click="editItem(item)">
                      mdi-pencil
                    </v-icon>
                  </v-btn>
                </span>
              </v-flex>
            </v-layout>
            <v-layout class="mb-1">
              <v-flex class="d-flex justify-space-between text">
                <span class="text-large">ISIN</span>
                <span> {{ item.isin }} </span>

              </v-flex>
            </v-layout>
            <v-layout class="d-flex flex-column">
              <v-flex class="d-flex justify-space-between text mb-1">
                <span class="text-large">Blog Link</span>
                <span class="d-inline-block text-truncate" style="max-width: 100px">
                  {{ item.blogLink }}</span>
              </v-flex>
              <v-flex class="d-flex justify-space-between text mb-2">
                <span class="text-large">DRHP</span>
                <span class="d-inline-block text-truncate" style="max-width: 100px">
                  {{ item.drhpLink }}</span>
              </v-flex>
            </v-layout>
          </div>
        </template>
      </v-data-table>
    </v-slide-x-transition>

    <v-dialog v-model="dialog" max-width="800px" persistent :fullscreen="dialog" transition="dialog-bottom-transition">
      <v-card>
        <v-card-title>
          <span class="text-h5">{{ formTitle }}</span>
        </v-card-title>
        <v-card-text>
          <v-slide-x-transition mode="out-in" appear>
            <v-container>
              <v-form ref="form" lazy-validation>
                <v-row>
                  <v-col cols="12" xs="6">
                    <v-text-field v-model="editedItem.symbol" label="Symbol"></v-text-field>
                  </v-col>
                  <v-col cols="12" xs="6">
                    <v-text-field v-model="editedItem.isin" label="ISIN" :rules="rules"></v-text-field>
                  </v-col>
                  <v-col cols="12" xs="6">
                    <v-text-field v-model="editedItem.blogLink" label="BLOG Link" :rules="rules"></v-text-field>
                  </v-col>
                  <v-col cols="12" xs="6">
                    <v-text-field v-model="editedItem.drhpLink" label="DHRP Link" :rules="rules"></v-text-field>
                  </v-col>
                  <v-col cols="12" xs="6">
                    <!-- <v-text-field v-model="editedItem.allotmentFinal" label="Allotment" :rules="rules"></v-text-field> -->
                    <v-menu v-model="menu1" :close-on-content-click="false" :nudge-right="40"
                      :menu-props="{ minWidth: '0px' }">
                      <template v-slot:activator="{ on, attrs }">
                        <v-text-field v-model="editedItem.allotmentFinal" label="Allotment" readonly v-bind="attrs"
                          v-on="on" :menu-props="{ minWidth: '0px' }"></v-text-field>
                      </template>
                      <v-date-picker v-model="editedItem.allotmentFinal" @input="menu1 = false" color="primary" no-title
                        :menu-props="{ minWidth: '0px' }"></v-date-picker>
                    </v-menu>
                  </v-col>
                  <v-col cols="12" xs="6">
                    <!-- <v-text-field v-model="editedItem.refundInitiate" label="Refunds"></v-text-field> -->
                    <v-menu v-model="menu2" :close-on-content-click="false" :nudge-right="40" class="menusize">
                      <template v-slot:activator="{ on, attrs }">
                        <v-text-field v-model="editedItem.refundInitiate" label="Refunds" readonly v-bind="attrs"
                          v-on="on"></v-text-field>
                      </template>
                      <v-date-picker v-model="editedItem.refundInitiate" @input="menu2 = false" color="primary"
                        no-title></v-date-picker>
                    </v-menu>
                  </v-col>
                  <v-col cols="12" xs="6">
                    <!-- <v-text-field v-model="editedItem.dematTransfer" label="Demat"></v-text-field> -->
                    <v-menu v-model="menu3" :close-on-content-click="false" :nudge-right="40">
                      <template v-slot:activator="{ on, attrs }" style="left:40px;top: 343px;">
                        <v-text-field v-model="editedItem.dematTransfer" label="Demat" readonly v-bind="attrs"
                          v-on="on"></v-text-field>
                      </template>
                      <v-date-picker v-model="editedItem.dematTransfer" @input="menu3 = false" color="primary"
                        no-title></v-date-picker>
                    </v-menu>
                  </v-col>
                  <v-col cols="12" xs="6">
                    <!-- <v-text-field v-model="editedItem.listingDate" label="Listing"></v-text-field> -->
                    <v-menu v-model="menu4" :close-on-content-click="false" :nudge-right="40">
                      <template v-slot:activator="{ on, attrs }" style="left:40px;top: 343px;">
                        <v-text-field v-model="editedItem.listingDate" label="Listing" readonly v-bind="attrs"
                          v-on="on"></v-text-field>
                      </template>
                      <v-date-picker v-model="editedItem.listingDate" @input="menu4 = false" color="primary"
                        no-title></v-date-picker>
                    </v-menu>
                  </v-col>
                  <v-col cols="1" xs="1">
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
  mixins: [Info],
  computed: {
    issave() {
      return this.editedItem.allotmentFinal == "" || this.editedItem.dematTransfer == "" || this.editedItem.drhpLink == "" || this.editedItem.isin == "" || this.editedItem.refundInitiate == "" || this.editedItem.listingDate == ""
    }
  }
}
</script>

<style scoped>
.row {
  margin-top: 2px;
  /* margin-bottom: 2px; */
}

.text-large {
  font-size: 13px;
}

.text {
  font-size: 12px;
}

::v-deep div.v-menu__content {
  width: 0px !important;
  min-width: 0px !important;
}

/* ::v-deep .v-menu__content .theme--light .v-menu__content--fixed .menuable__content__active { */
::v-deep div.v-menu__content.theme--light.v-menu__content--fixed.menuable__content__active {
  top: 343px !important;
  left: 40px !important;
  min-width: 500px !important;
  width: 0px !important;
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

.check {
  width: 0;
}
</style>
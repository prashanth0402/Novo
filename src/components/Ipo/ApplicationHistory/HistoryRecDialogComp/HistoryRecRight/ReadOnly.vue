<template>
  <div>
    <v-alert dense class="pt-4 px-4 mb-3 outline" elevation="0" colored-border :color="bid.color">
      <div>
        <v-layout class=" text d-flex ">
          <v-flex lg4 md4 sm3 xs3 class="mr-sm-1">
            <span><b>No of Lot</b></span>
            <p class="ml-6">{{ bid.quantity }}</p>
          </v-flex>
          <v-flex lg4 md4 sm4 xs4 >
            <span><b>Cut off-price</b></span>

            <p>
              <v-chip label :color="bid.cutOff ? 'green lighten-1' : 'orange lighten-2'" text-color="white" x-small>
                {{ changeText(bid.cutOff) }}
              </v-chip>
            </p>
          </v-flex>
          <v-flex lg2 md2 sm2 xs2 class="text-center">
            <span><b>Price</b></span>
            <p>{{ bid.price }}</p>
          </v-flex>
          <v-flex class="d-flex align-center justify-center" lg2 md3 sm3 xs3>
            <span class="mr-2 caption text-capitalize green--text" dark v-if="bid.activityType != 'cancel'">Placed
              Bid</span>
            <span v-else-if="bid.activityType == 'cancel'" class="caption red--text">Bid deleted</span>
          </v-flex>
        </v-layout>
        <!-- <v-divider class="mb-4" v-show="idx != bids.length - 1"></v-divider> -->
      </div>
    </v-alert>
  </div>
</template>
<script>
export default {
  props: {
    bid: {},
  },
  methods: {
    changeText(text) {
      if (text == true) {
        this.text = 'Yes'
        return this.text
      } else {
        this.text = 'No'
        return this.text
      }
    }
  }
};
</script>
<style scoped>
.text {
  font-size: 10px;
}
.outline{
  border: 1px solid #F0F0F0;
}
</style>
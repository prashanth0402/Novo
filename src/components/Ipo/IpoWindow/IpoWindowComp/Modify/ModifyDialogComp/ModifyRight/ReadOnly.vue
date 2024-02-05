<template>
  <div>
    <v-alert :border="bid.activationType == 'cancel' ? 'left' : undefined" dense class="pt-4 px-4 mb-3 outline" elevation="0"
      colored-border :color="bid.color">
      <div>
        <v-layout class="caption ">
          <v-flex lg4 md4 sm4 xs4>
            <span><b>No of Lot</b></span>
            <p class="ml-6">{{ bid.quantity }}</p>
          </v-flex>
          <v-flex lg4 md4 sm4 xs4>
            <span><b>Cut off-price</b></span>

            <p>
              <v-chip label :color="bid.cutOff ? 'green lighten-1' : 'orange lighten-1'" text-color="white" x-small
                class="ml-5">
                {{ changeText(bid.cutOff) }}
              </v-chip>
            </p>
          </v-flex>
          <v-flex lg2 md2 sm2 xs2 class="ml-1">
            <span><b>Price</b></span>
            <p>{{ bid.price }}</p>
          </v-flex>
          <v-flex class="d-flex justify-center mt-3" lg2 md2 sm2 xs2 v-if="!Disable">
            <a class="mr-2" dark @click="ModifyBid(Idx)" v-if="bid.activationType != 'cancel'">
              <span class="caption text-capitalize"> Edit </span>
            </a>
            <span v-else-if="bid.activationType === 'cancel'" class="caption red--text">Bid deleted</span>
          </v-flex>
        </v-layout>
        <!-- <v-divider class="mb-4" v-show="idx != bids.length - 1"></v-divider> -->
      </div>
    </v-alert>
  </div>
</template>
<script>
export default {
  data() {
    return {
      text: ""
    }
  },
  props: {
    bid: {},
    detail: {},
    Idx: Number,
    Disable: Boolean
  },
  methods: {
    ModifyBid(ind) {
      this.$emit("modify", ind);
    },
    changeText(text) {
      if (text == true) {
        this.text = 'Yes'
        return this.text
      } else {
        this.text = 'No'
        return this.text
      }
    }
  },
};
</script>
<style scoped>
.outline {
  border: 1px solid #F0F0F0;
}
</style>
<template>
  <div>
    <v-alert border="left" dense class="pa-4 mb-3 outline" elevation="0" height="80" colored-border :color="bid.color">
      <v-form ref="form" v-model="valid" lazy-validation>
        <v-layout>
          <v-flex lg3 md4 sm4 xs4>
            <v-text-field v-model.number="bid.quantity" label="No of Lot" background-color="white" outlined type="number"
              min="1" dense @keypress="onlyForNumber" :rules="bid.price == 0 ? refRule : refRule"></v-text-field>
          </v-flex>
          <v-flex class="d-flex justify-center" lg4 md4 sm4 xs5>
            <v-checkbox v-model="bid.cutOff" @change="disablePrice(bid)" dense
              :disabled="Disable || issueDetails.cutOffFlag == 'N' ? true : false"></v-checkbox>
            <span
              :class="this.$vuetify.breakpoint.name != 'xs' ? 'mt-2 black--text text-sm-caption' : 'mt-2 black--text text-medium'">Cut
              off-price</span>
          </v-flex>
          <!-- <v-flex lg4 class="d-flex justify-center mt-2"> <b>Cut off-price</b></v-flex> -->
          <v-flex lg3 md4 sm4 xs4>
            <v-text-field v-model.number="bid.price" label="Price" background-color="white" outlined type="number"
              :min="issueDetails.minPrice" :max="issueDetails.cutOffPrice" dense
              :rules="bid.quantity == 0 ? refRule : refRule" :error-messages="ErrText" :error="ErrFeild"
              :disabled="bid.cutOff" @keypress="onlyForNumber" @input="checkPrice(Idx)"></v-text-field>
          </v-flex>

          <v-flex lg2 md2 sm2 xs2 class="pt-1 pl-3" v-if="!Disable">
            <a @click="closeSlot" class="caption text-capitalize error--text text">
              undo
            </a>
          </v-flex>
        </v-layout>
      </v-form>
    </v-alert>
  </div>
</template>

<script>
export default {
  name: "modifyBids",

  props: {
    bid: {},
    issueDetails: {},
    totalAmt: Number,
    Idx: Number,
    Disable: Boolean
  },

  data() {
    return {
      ErrFeild: false,
      ErrText: "",
      refRule: [(v) => !!v || "required"],
      valid: false,
      isError: false
    };
  },

  methods: {
    disablePrice(bid) {
      if (bid.cutOff == true) {
        bid.price = parseInt(this.issueDetails.cutOffPrice);
        this.ErrText = ''
      } else {
        bid.price =  parseInt(this.issueDetails.minPrice);
      }
    },

    // this method is to check the enter values are number
    onlyForNumber($event) {
      let keyCode = $event.keyCode ? $event.keyCode : $event.which;

      if ((keyCode < 48 || keyCode > 57) && (keyCode !== 46 || keyCode == 46)) {
        $event.preventDefault();
      }
    },

    // this method is to check wheather the value entered in the price feild lies between the range
    checkPrice(indicator) {
      if (indicator == 0) {
        if (
          (parseInt(this.bid.price) <
            parseInt(this.issueDetails.minPrice) ||
            parseInt(this.bid.price) >
            parseInt(this.issueDetails.cutOffPrice)) &&
          this.bid.price != 0 || this.bid.quantity == ''
        ) {
          this.isError = true
          this.ErrFeild = true;
          this.ErrText = this.issueDetails.priceRange;
          this.$emit('hideUpdate', this.isError)
        } else {
          this.isError = false
          this.ErrFeild = false;
          this.ErrText = "";
          this.$emit('hideUpdate', this.isError)
        }
      } else if (indicator == 1) {
        if (
          (parseInt(this.bid.price) <
            parseInt(this.issueDetails.minPrice) ||
            parseInt(this.bid.price) >
            parseInt(this.issueDetails.cutOffPrice)) &&
          this.bid.price != 0 || this.bid.quantity == ''
        ) {
          this.isError = true
          this.ErrFeild = true;
          this.ErrText = this.issueDetails.priceRange;
          this.$emit('hideUpdate', this.isError)
        } else {
          this.isError = false
          this.ErrFeild = false;
          this.ErrText = "";
          this.$emit('hideUpdate', this.isError)
        }
      } else if (indicator == 2) {
        if (
          (parseInt(this.bid.price) <
            parseInt(this.issueDetails.minPrice) ||
            parseInt(this.bid.price) >
            parseInt(this.issueDetails.cutOffPrice)) &&
          this.bid.price != 0 || this.bid.quantity == ''
        ) {
          this.isError = true
          this.ErrFeild = true;
          this.ErrText = this.issueDetails.priceRange;
          this.$emit('hideUpdate', this.isError)
        } else {
          this.isError = false
          this.ErrFeild = false;
          this.ErrText = "";
          this.$emit('hideUpdate', this.isError)
        }
      }
    },
    // this method is used to close the input card 
    closeSlot() {
      this.$refs.form.validate()
      if (parseInt(this.bid.price) <= parseInt(this.issueDetails.cutOffPrice)) {
        this.$emit('closeSlot')
        this.ErrText = ''
      } else if (this.bid.price == 0) {
        this.ErrFeild = true;
        this.ErrText = 'Price should not be empty'
      }
    }
  },
  watch: {
    bid: {
      handler(bid) {
        if (bid.price == 0 && bid.quantity == 0) {
          bid.quantity = 1
          bid.price = this.issueDetails.minPrice
        }
      }, immediate: true
    }
  }
};
</script>

<style scoped>
.text {
  font-size: 5px;
}

.text-medium {
  font-size: 10px;
}
.outline{
  border: 1px solid #F0F0F0;
}
</style>
